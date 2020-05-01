package tools

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type WgetOptions struct {
	Url        string
	OutputName string
	Timeout    time.Duration
}

type tempFile struct {
}

func Download(options *WgetOptions) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	verbose("Test support Partial Content or not")

	req, e := http.NewRequest(http.MethodHead, options.Url, nil)
	if e != nil {
		return e
	}
	req.Header.Set("Range", "bytes=0-")

	client := &http.Client{
		Timeout:       options.Timeout,
		CheckRedirect: checkRedirect,
	}

	resp, e := client.Do(req)
	if e != nil {
		return e
	}

	defer resp.Body.Close()

	var rangeSupported bool

	switch resp.StatusCode {
	case http.StatusPartialContent:
		verbose("Partial Content is supported")
		rangeSupported = true
	case http.StatusOK, http.StatusRequestedRangeNotSatisfiable:
		verbose("Partial Content is not supported")
	default:
		panic(fmt.Errorf("Unexpected status code: %d", resp.StatusCode))
	}

	var saveFileName string
	if options.OutputName != "" {
		saveFileName = options.OutputName
	} else {
		// parse the Content-Disposition (filename)
		contentDisposition := resp.Header.Get("Content-Disposition")
		if contentDisposition != "" {
			reg := regexp.MustCompile(`filename="(.*?)"`)
			for _, v := range strings.Split(contentDisposition, ";") {
				v = strings.TrimSpace(v)
				if reg.MatchString(v) {
					saveFileName = reg.ReplaceAllString(v, "$1")
				}
			}
		} else {
			hash := md5.Sum([]byte(options.Url))
			verbose("Unkown filename, use hash value: %x", hash)
			saveFileName = fmt.Sprintf("%x", hash)
		}
	}

	verbose("file will be saved in %s", saveFileName)

	if rangeSupported {
		verbose("File content length: %d bytes (%s)", resp.ContentLength, transSize(resp.ContentLength))
		// may be we need a download manager

	}

	// if rangeSupported {
	// 	req, e := http.NewRequest(http.MethodGet, options.Url, nil)
	// 	if e != nil {
	// 		return e
	// 	}
	// 	req.Header.Set("Range")
	// }
	return nil
}

func checkRedirect(req *http.Request, via []*http.Request) error {
	if len(via) > 10 {
		return errors.New("too many redirect")
	}
	verbose("Redirect %s to %s", via[len(via)-1].URL, req.URL)
	return nil
}
