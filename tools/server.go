package tools

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chai2010/gettext-go"
	"github.com/liuyongshuai/goUtils"
)

//ServerOptions is the options of simple http file server
type ServerOptions struct {
	Host      string
	Port      int
	RootPath  string // root directory
	UploadDir string // default is `RootPath`
	User      string
	Password  string // if user or password specific, use basic authentication
}

var serverOptions *ServerOptions

type HTMLStructure struct {
	Title                string
	UploadUrl            string
	UploadFileButtonName string
	IndexUploadName      string
	IndexListFileName    string
}

var htmlStructure HTMLStructure

var uploadPrefix = "/upload"

var uploadHTML = `<!DOCTYPE html>
<head>
    <title>{{ .Title }}</title>
</head>
<body>
    <form enctype="multipart/form-data" action="{{ .UploadUrl }}" method="post">
        <input type="file" name="uploadFile" />
        <input type="submit" value="{{ .UploadFileButtonName }}" />
</body>`

var indexHTML = `<!DOCTYPE html>
<head>
	<title>{{ .Title }}</title>
</head>
<body>
	<a href="{{ .UploadUrl }}">{{ .IndexUploadName }}</a>
	<br />
	<a href="/files/">{{ .IndexListFileName }}</a>
<body>
`

var hss = make(map[string]hs)

type hs []http.Handler

func next(r *http.Request) {
	r.Header.Set("next", "true")
}

func nonext(r *http.Request) {
	r.Header.Set("next", "false")
}

func (h hs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range h {
		handler.ServeHTTP(w, r)
		if nt, ok := r.Header["next"]; ok && nt[0] == "false" {
			fmt.Println("break")
			break
		}
	}
	fmt.Printf("%s %s %s %s\n", goUtils.Yellow(r.Method), goUtils.Green(r.URL.String()), r.Proto, r.RemoteAddr)
}

func init() {
	localZipBytes := MustAsset("local.zip")
	gettext.BindLocale(gettext.New("filetools", "local.zip", localZipBytes))
	gettext.SetDomain("filetools")
	htmlStructure = HTMLStructure{
		Title:                gettext.Gettext("filetools simple http file server"),
		UploadUrl:            uploadPrefix,
		UploadFileButtonName: gettext.Gettext("Submit"),
		IndexUploadName:      gettext.Gettext("Upload file"),
		IndexListFileName:    gettext.Gettext("View files"),
	}
}

func checkAuth(w http.ResponseWriter, r *http.Request) {
	if username, password, ok := r.BasicAuth(); !ok || username != serverOptions.User || password != serverOptions.Password {
		w.Header().Set("WWW-Authenticate", `Basic realm="filetools"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		nonext(r)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("index.html").Parse(indexHTML)
	t.Execute(w, htmlStructure)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.New("upload.html").Parse(uploadHTML)
		t.Execute(w, htmlStructure)
		return
	}
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		fmt.Printf("Could not parse multipart form: %v\n", err)
		renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
		return
	}
	file, fileHeader, err := r.FormFile("uploadFile")
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileName := fileHeader.Filename
	filePath := filepath.Join(serverOptions.UploadDir, fileName)
	fmt.Println(goUtils.Green(fmt.Sprintf("Recieved file %s from %s, save to %s, size %s", fileName, r.RemoteAddr, filePath, transSize(fileHeader.Size))))

	newFile, err := os.Create(filePath)
	if err != nil {
		renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)

	if err != nil {
		renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("SUCCESS"))
}

func static(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/files/", http.FileServer(http.Dir(serverOptions.RootPath))).ServeHTTP(w, r)
	// http.FileServer(http.Dir(serverOptions.RootPath)).ServeHTTP(w, r)
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", "/", gettext.Gettext("back to home page"))
}

func Server(options *ServerOptions) error {
	serverOptions = options
	_, err := os.Stat(serverOptions.RootPath)
	if err != nil {
		return fmt.Errorf("root path error: %s", err)
	}

	if serverOptions.User != "" || serverOptions.Password != "" {
		hss["/"] = append(hss["/"], http.HandlerFunc(checkAuth))
		hss[uploadPrefix] = append(hss[uploadPrefix], http.HandlerFunc(checkAuth))
		hss["/files/"] = append(hss["/files/"], http.HandlerFunc(checkAuth))
	}

	hss["/"] = append(hss["/"], http.HandlerFunc(index))
	hss[uploadPrefix] = append(hss[uploadPrefix], http.HandlerFunc(upload))
	hss["/files/"] = append(hss["/files/"], http.HandlerFunc(static))

	addr := fmt.Sprintf("%s:%d", serverOptions.Host, serverOptions.Port)

	http.Handle(uploadPrefix, hss[uploadPrefix])
	http.Handle("/files/", hss["/files/"])
	http.Handle("/", hss["/"])

	fmt.Printf("The file server is running at %s, root path is %s\nThe username is %s\nThe password is %s\n...\n", addr, serverOptions.RootPath, serverOptions.User, serverOptions.Password)

	http.ListenAndServe(addr, nil)
	return nil
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
