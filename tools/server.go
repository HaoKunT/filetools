package tools

import (
	"fmt"
	"net/http"
	"os"
)

//ServerOptions is the options of simple http file server
type ServerOptions struct {
	Host     string
	Port     int
	RootPath string // root directory
	Prefix   string // default is /files
	User     string
	Password string // if user or password specific, use basic authentication
}

var serverOptions *ServerOptions

func Server(options *ServerOptions) error {
	serverOptions = options
	_, err := os.Stat(serverOptions.RootPath)
	if err != nil {
		return fmt.Errorf("root path error: %s", err)
	}

	http.Handle(serverOptions.Prefix, http.StripPrefix(serverOptions.Prefix, http.FileServer(http.Dir(serverOptions.RootPath))))
	addr := fmt.Sprintf("%s:%d", serverOptions.Host, serverOptions.Port)

	fmt.Printf("The file server is running at %s, root path is %s ...\n", addr, serverOptions.RootPath)

	http.ListenAndServe(addr, nil)
	return nil
}
