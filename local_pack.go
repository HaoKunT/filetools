// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// local.zip
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _localZip = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x97\x67\x54\x93\x59\xbf\xc5\x13\x08\x55\x8a\x18\x4a\x68\xa1\x85\x32\xf4\x48\x47\x02\x84\xaa\x14\x03\x02\x0a\xd2\x9b\x10\xba\x42\xa8\x0a\x88\x44\x21\xf4\x1a\xe9\x5d\x01\x45\x8a\x80\x54\x69\x82\x42\x10\x44\x3a\x28\x28\x4d\x0c\x3d\x18\x91\xa2\xdc\xe5\xdc\xd7\x19\xee\xcc\x7d\x67\xcd\xf9\x72\x9e\xff\x5a\xfb\xec\xdf\xf3\x69\x9f\x7d\x4c\x0c\x29\x41\xac\x80\x9f\x8b\xcc\x2a\x8e\x02\x9c\x58\xd4\x00\x00\xc0\xcb\xd7\xd9\xd1\x4b\xe6\x4f\xcd\xd1\x5f\x34\x0c\x7f\x68\x42\xdd\xed\xb5\x2f\x9e\x50\xae\x85\xff\x5f\x25\xe4\x2f\x4a\x23\x6d\x7b\x63\x5d\x33\x33\xa4\xbe\xae\xd9\x1f\xa7\x68\x01\x86\x18\x49\xd4\xcd\x92\x48\x30\x8a\x0a\x00\x18\xa3\x03\x00\x60\xff\x70\xea\x1a\xda\xcb\x35\xc0\xd7\xd7\xcb\x5f\xda\xdb\x17\x63\x66\xe8\x05\xd6\x87\x40\xb3\x35\xcf\xcf\x97\xa1\x2a\xaa\xd1\x68\x36\xf3\xac\xf3\xfc\x5d\x40\xf2\x04\x65\x7b\x23\x9b\x16\x1b\xfe\xf1\xc6\x85\xb6\x81\x50\xb7\xd6\xd6\x3e\x05\xbc\x08\x63\x22\x2f\x4b\x9a\x3c\x2c\x5a\x4d\x8c\x1f\x1b\x99\xec\x57\x7c\x19\xd6\x17\x0b\x16\x20\x5f\x22\x4f\x68\x27\x59\x30\x9f\xf3\x7c\x88\xe7\x09\x5d\x6d\x98\x1f\x94\x4c\x4c\xeb\xf6\x70\xeb\xc8\xdb\xe7\xdb\xd5\x78\x7e\x3c\xac\xbd\x61\x06\x1a\x5e\xb3\xdc\x72\xf8\x1e\x45\x96\xce\x6c\x48\x4a\x5e\xec\xec\x88\xac\x06\xd8\xdf\xba\x9b\xac\x14\xc9\x47\x39\x4f\x11\x4e\x7f\x13\x1c\x4c\xfd\x4d\x5d\x96\xf1\x88\xc9\x92\x02\x0a\x7c\xfe\x46\xef\xd5\xbd\xce\x34\xef\x99\xad\x49\x96\x7c\x1a\x28\x95\x1d\x30\x10\xc8\x07\x9c\xa7\xf4\x69\x40\x6b\xd0\xcc\xd2\xed\x53\x07\x41\xd1\x93\x3f\xdc\xf6\xd6\x0e\xc3\x67\xaa\xc9\xe1\x94\xaa\x9d\x1d\x7b\x6b\x61\x43\xbc\xfc\x4b\xe0\x0f\xba\xfc\x76\x75\xd1\xc3\x02\x1f\xe8\xb7\x43\x3e\x54\x77\xe2\x23\x23\x74\x0e\x4f\xbb\x03\x8e\x81\xe1\xb0\xc5\xad\x37\xe1\x32\x01\xee\xfb\x46\xf7\x26\xec\x41\xcf\xa5\xbe\x07\xb8\xb7\xcf\x0b\xed\x77\xa4\x84\x1f\x22\xb5\x98\x84\x86\x7a\x3a\x2f\xae\xa3\x82\x3c\x8e\x8a\xcd\x66\x0b\x28\xf3\x81\x5f\x83\x88\xc1\xd0\x0b\x02\xd0\x68\x63\x3b\xa1\x19\xfc\xcd\x64\x11\x74\x1a\x59\xb0\x65\xa8\x8a\xd9\x46\xc2\x7e\xd8\x11\xc1\x4d\x5e\xd5\x84\x69\x67\x65\xcc\x33\x4d\x73\x3f\xe6\x28\xd9\xc5\xc7\x26\x6a\xab\x1f\xba\xd8\x33\x8c\x6e\x66\x16\x7c\x1d\xb6\xb9\xf6\x2d\x5e\x04\x51\x20\x62\x8c\x17\xbe\xb9\xb9\x52\x3e\x6d\x68\xbc\xae\xb2\xbb\x2e\x5d\x50\xe7\x6b\x3c\xec\x0b\x57\x79\xed\xb5\xc2\x75\x99\xd0\xff\x69\xb0\x4e\x19\x41\xb0\x95\xd8\x47\x12\x76\xf0\x92\x29\xd7\x11\xbb\x6a\x25\x74\x9b\x4b\x81\x74\x5e\x42\xaf\xcb\x91\x61\xc6\x71\x0d\xb9\x11\x19\x0b\x56\x18\x3d\xe3\x56\x7c\x2b\xba\x47\xa9\xf9\xbd\xa0\x46\x8e\x53\x38\xae\x79\x2c\x2d\x34\x91\x70\x68\xb3\x12\x22\xef\xd3\x04\xe7\xac\xdb\xba\x66\x69\x44\x9d\xaf\x20\x15\x27\xc2\x20\x91\x52\xeb\xe2\x23\x5b\x94\x1f\x72\xdb\x74\xae\x35\xad\xe4\x92\x11\x44\x71\xb9\x31\xa3\xe4\xba\x9e\x1c\xd4\xb0\xfb\x2d\xca\x34\x67\x74\x77\x30\xc1\x3f\x69\xc4\x8e\xb0\xea\xd9\xac\xd1\x35\x9a\xf8\xb1\xdb\x39\xf3\x91\x2c\x62\xe0\x3d\x2f\xbe\x95\x35\xcf\x26\x7d\x6e\xd4\x3c\x75\x70\xd7\xd8\xda\xf5\x7c\x58\xa9\x13\x2a\xae\xa6\x00\x4e\x73\xf7\x1a\x9f\xfc\xe3\x76\x8b\xba\x41\x79\x27\xbc\x15\x79\x2a\x2b\x20\xb7\x19\x71\x56\x8e\x37\x69\xd8\xe2\x7a\x20\xba\xa4\xe1\xae\xab\x2e\xfd\x83\x84\xa6\x7c\xd6\xa4\xc5\xec\xc6\x79\xa3\x5c\xb7\x9b\xa2\xa3\x33\x01\x49\x25\x8e\xea\x9f\x32\xf5\x13\x94\x4b\xc3\xee\x90\x9a\x1e\x39\xbf\x9d\x7c\xe1\x6f\x6e\xf5\x25\x7b\xd0\x20\x14\x3e\x26\x6f\xd8\x1e\xbf\xa8\x19\x27\x8c\x32\x9d\xfe\x71\xdf\xc3\x20\x7e\x25\x58\x4f\x81\x82\x31\x19\x51\x3b\x18\x3e\x7d\x16\xac\xa7\x9f\x64\x50\xed\x27\x3c\x8e\x2a\x72\x88\xcd\x13\x42\x87\x0d\x6c\x9f\x81\x57\x0e\x8d\x0e\xd6\x20\xc8\x84\xf5\x52\xf4\x5e\x6f\xf9\x88\x9a\x42\x20\x53\xc6\x13\x47\xd1\x0c\x73\x3b\x5c\x9f\x9b\xed\x28\x3e\x3a\x0d\xab\xcd\x84\xcb\xb2\xd5\xc3\x5f\x5a\x33\x08\xc4\x48\x1c\x9c\x8f\x1d\xc0\x72\x25\xf6\x30\x4c\xf6\x0c\x91\x05\x2b\x7b\x15\x5e\xba\xf0\xd6\x5f\x32\xe2\x0c\xe3\x59\x22\x66\x0e\xa6\xaf\xc0\xcb\x28\xa9\xcb\xca\xbc\x98\x75\x33\x9e\xde\xaf\x4f\x36\x60\xf2\x67\x3b\xdc\x00\x79\xda\x7d\xb0\x52\x23\xbc\xdb\xb4\x65\x1b\x6e\x32\xf9\xc6\x17\xc4\x0a\x5f\x2a\x1a\x74\xa9\x76\x70\xeb\x0b\xe3\xe9\x41\x35\xe0\x5c\xbf\x83\x7d\x86\x80\xd8\xed\x4d\x68\x41\x00\x09\x54\x6c\xc8\xf9\xd2\xc5\x89\xb8\xf5\x2d\x36\x0d\x3e\x2e\xa2\x20\x14\x8a\xab\x7b\xb7\xb1\xe7\xae\xd8\xbb\x2e\x02\x16\x63\xf5\xda\xd4\x48\xb9\x3e\xed\x9e\xcd\x87\xb2\x88\x9a\x55\xa7\xa4\x81\xcb\xca\x0a\x38\xed\x1c\x6e\xbf\xdc\x3f\xb0\x51\x69\xb3\xa2\xad\x66\x8f\x6a\x60\xab\xdd\x96\x2d\x4f\x7f\xf1\xdd\xd3\xe9\xc7\xe6\xd7\x83\x72\xb1\xa1\xd8\x86\x4a\x31\xe6\x94\xaa\xdb\xd3\x07\x4b\xb6\x67\xc6\x7c\x70\xb3\x0a\xe7\xa3\x06\xc6\x0f\xc6\x56\xbf\x1e\xf0\xc3\xc0\xc5\x8a\x02\xec\x06\xc5\x4d\x5a\x8d\xf4\xc2\x3b\xda\xc6\xf1\xf8\xfa\x73\xc5\x45\xb5\xf5\x10\xa6\xe5\x0b\xc1\xde\x17\xe2\x65\xee\x4e\x2a\x12\x4f\x8d\xf7\xad\xe4\x96\x4b\xdf\x9d\xf4\xed\xb7\x36\x87\x20\xcb\xae\x16\x74\x14\x49\xe7\xd1\x4d\xb8\x71\x2f\x5b\x93\x49\x01\x72\xe2\xb5\x7e\xfa\x92\x49\x17\x3e\x4f\x21\x5e\xb0\x33\xb2\x78\xbd\x70\x10\xc8\xb5\xab\x40\xcd\xd1\x3e\xc4\xc1\x6c\xb2\xbd\xbc\x18\xa8\x32\xb0\x85\x37\x50\x39\x3e\xd5\x1c\x3d\x0b\xe3\x0f\x72\x36\xc8\x2c\xc5\xda\xb7\x2a\x9b\x1a\x78\x60\x53\x6d\xd0\x44\xa6\xbb\xe3\x3b\xee\x1d\x0a\x72\xd2\xde\x16\xf4\xe4\xe9\xb8\x03\xf9\xdc\xab\x73\xe2\xf1\x2c\xa6\xb9\xec\x58\x0f\x13\xce\xce\x65\x3f\x41\x0f\x33\xe4\x72\xaa\x39\x17\x27\xed\x27\x7d\x28\x87\xcb\x12\x4e\x4c\xaa\x84\x9f\xdc\xe0\x28\x88\x79\xe4\xef\xe4\x24\x4f\x92\xd3\x10\x24\x69\xf4\x68\x3d\x9e\x63\xdf\xb7\xb0\xda\xed\xbb\x26\x70\xa9\x45\x83\x37\x93\x2f\x1a\x2f\xaa\x63\x71\xb9\x28\x3c\x84\xd3\x82\x56\x14\x7f\xda\x76\xad\x94\xf0\x23\x5e\x8a\xca\xf3\x59\xef\x1b\x4b\x67\x7f\x91\xb6\x66\xf1\xbb\x6e\x85\xc7\xd1\xc8\x24\x84\xf7\xdb\x91\x07\x7e\xf5\x35\x0a\x66\x6f\xaf\xf4\x23\x33\xf3\x24\xd8\xbe\xd8\x1d\x56\xe3\x69\x3e\x70\x92\x5f\xf2\x08\x01\x8d\x6a\x59\x97\xe3\xca\x01\x4e\x69\x2a\xe5\x91\xbd\x24\x3f\x93\xb4\x2a\x46\x71\xcf\x84\xde\x2e\x84\xd5\x6a\xdf\x9a\xb7\x09\x6f\x66\xc8\xfb\xf6\xe8\x08\xcb\xae\x8c\x6b\x91\x4f\x5c\x9e\x7e\xff\xd8\x99\xd7\x67\xc2\x78\xa4\xe9\xcd\x22\x38\x8d\x36\x1e\x7b\xb1\x73\x66\x44\x24\xa6\xa8\xb6\x05\xa6\x28\x1d\x55\x9f\xf3\xe1\x09\x54\x0d\x9e\xed\xeb\xca\xfe\x80\xe9\x62\xfb\x18\x9c\x9c\x21\x24\x2a\xc7\x34\x41\x2a\x97\xa6\x46\x66\x2e\xc0\x95\x3a\x1d\xe0\x47\x9a\x8d\xf2\x0b\xd9\xd7\x9e\x3a\xe5\x42\x86\xe2\x48\xda\x24\xbe\x3f\x03\x59\x12\x23\x89\xc2\xd7\xa5\x97\x44\x81\x00\x80\x0a\xfa\x7f\x1b\xc8\x7e\xbe\x79\x97\x8d\x51\xdc\x2a\x90\xe7\xd5\xf5\xc7\x5d\x04\x2c\x81\xb5\x42\x02\xc3\x6c\xb3\x57\x13\xb6\x09\x0d\x4d\x90\x5b\x84\xb1\xcb\x2d\xa6\x20\xd0\xd9\x88\x2e\x2d\x75\x93\xb4\xee\x4b\xcf\x8b\xac\xc5\x91\xb8\xe4\x6c\xfe\x2a\x7a\xb4\xd6\x87\xe1\x78\xb7\xee\xd8\x8e\x9a\xea\x0d\x56\x6a\xf7\x0c\x77\x99\x92\xac\xda\xef\xdf\x6e\x15\xec\x85\xd8\xcd\x82\x43\x88\x85\x15\x15\x7c\x11\xdb\x3a\x76\xd1\x36\x1d\x44\xe6\x73\x37\xd9\x8a\x7f\x63\x36\x8b\xd0\x43\x0c\x32\x28\xec\x05\x2d\xb4\x30\xb4\x3b\xbc\xc3\x84\x1b\x54\xc5\x51\xd3\x90\xdc\xfd\xfc\x47\x62\x5a\x72\xfa\xf5\x16\x57\x9c\x1c\x25\x53\x1d\xa0\x48\x01\xe7\x27\x9a\xca\x04\x9c\xc0\x80\x29\xbe\x54\xa0\x75\xd7\xd4\x17\xf5\x44\x61\xe1\xf5\xd5\x31\x8d\xb8\x3d\xbc\xeb\xf2\xbc\x4b\xd0\x85\x7c\xbe\xfe\x81\x7d\xcb\x6f\xf3\x31\x28\xe5\x40\xb3\xc8\xc9\xdd\xd6\xfe\x83\x9d\x29\xac\xe2\x2e\xc5\xa0\x40\xfd\x99\x51\x01\xba\x17\xca\x37\xc4\x44\x03\x5f\x05\xab\x60\x39\x82\x26\xee\xe1\xe3\x4c\xf7\x64\x89\xf9\xf7\x21\xbe\xd2\xc4\xeb\xb7\x46\xdb\x08\x3d\x32\x4a\x1c\x04\x70\x8d\xc9\x92\x22\x92\xb3\xc1\x4c\x84\x71\xce\x5a\x41\xd9\x99\x85\x47\x50\x8b\xb8\x98\x7c\xf6\x18\x6e\x54\x37\x99\xc8\xec\x01\x16\xcc\xa0\x87\x49\x9e\x4e\x7c\x7e\xd6\x41\x4e\x03\x16\x5e\x4e\xda\xb3\xba\xcf\xe4\xd3\x2a\x74\xc5\x5a\x24\xc2\xaa\xea\xd5\x57\x7c\xed\x14\xf0\x61\xe4\x97\xd4\x6d\x49\x30\x08\xd9\x7c\x67\x01\xf1\x62\xf9\x59\x5d\x88\xea\x7a\x7d\x6a\x38\x46\x72\x82\x8f\xea\xf2\xdb\x32\xe4\xd8\x30\x7c\xf3\xf3\xbb\x51\x07\x55\x5b\x40\x1c\x87\x84\x73\x17\xbf\xb3\x76\x4e\x52\xfb\x21\x9b\xac\xab\x8d\x77\x22\x56\x4c\x4b\x4f\x51\x86\x46\x80\x12\x21\xc4\xdd\x1a\x38\xd2\x2b\x20\x46\x5b\x88\xf8\x94\x85\xd5\x01\x6a\xe8\x4e\x05\x7d\x99\x1d\x1e\x0e\xaf\x75\x88\xa1\x4f\xf7\xbf\x0e\x85\xcf\x4b\x95\x9c\x59\xf4\xa7\x00\x1e\x3f\xb7\x73\xb9\x01\xc6\x91\x7d\xc5\x65\xb8\x6c\x6e\xa7\xbe\xce\xac\x33\xe5\x6c\xd4\x80\x6d\x48\x81\x65\xd6\x69\xa4\x59\xb1\x0f\x0f\x2a\x14\xef\x70\xb0\x92\x25\x6e\x1c\x2b\x13\x15\x8b\x6e\x3c\x23\x81\x02\x5d\xc8\xfe\x63\xb7\xf2\x1e\x77\x24\x29\x11\xc1\x10\x21\x31\x21\xaf\xa8\xc3\x44\xe1\x25\x47\xb3\x89\xe1\x1f\xcf\x38\xf3\xc1\xad\xaf\xa8\x6e\x55\x94\x72\xa4\x6f\x59\x05\x19\x33\x4d\x30\x1c\x58\xa6\xdb\xab\x98\xae\x9f\x13\xd4\x85\x16\xab\xe8\x7c\xcc\x61\xf9\x88\xba\x17\xd3\x56\x2c\x28\xa5\x31\x7d\x20\xf1\xc3\x4d\xc3\x9b\x3d\x06\xb2\x9a\x5a\xd0\x78\x5b\xc9\xc1\x64\x62\xcf\xa5\x43\x74\x36\xce\x39\x28\x02\x7e\x36\xc5\x44\x90\x51\x99\xe0\x2f\x08\x1a\x70\xf2\xa9\x80\x68\xd9\xc1\x32\x71\xa5\x3a\xb6\x7b\xb7\xcd\x41\xdf\xc1\x31\xaf\xdf\x53\x55\x89\xce\x4c\x29\x71\xf5\x9f\x01\x69\xf9\x3b\xb8\x5a\x02\x5c\xdf\x85\x5e\xa8\x58\xad\x55\x7a\x26\x93\xd6\x33\xb7\xed\x81\xfa\x7a\x75\xa7\x95\x59\x9c\x32\xcf\x69\x6c\x43\xf8\x70\x0b\xa3\x75\x6e\xa4\xd7\xe7\xaa\x4a\x72\xad\xda\x91\x18\xea\xca\xd8\xea\x00\xdd\x12\x62\x06\x38\xf1\x0a\x7b\xb5\x0d\xa3\x58\x74\x31\xdf\x53\x22\x21\x3d\x55\xda\xa9\xea\xb7\x63\x68\x64\x91\x88\x97\xbe\x7f\x34\x88\x5f\x3e\xfb\xde\xe6\x8a\xba\x2a\xef\x93\xb9\xde\x04\x6c\x7f\x6b\xe9\x29\x2a\xfd\x29\x8c\xbe\x54\xde\xc1\xe9\x79\x25\x36\x9b\xf9\xd4\x86\x7c\xd5\xb0\x16\xd1\x42\x1a\x95\x29\x13\xea\x69\x5b\x1c\xfe\xd3\x47\x50\xb6\x9f\x5c\x57\x70\xb0\xa7\x8b\x58\x38\xe7\xe7\x1e\x69\xfd\xd8\x35\xda\x53\xb8\x2b\x5a\xc5\x10\x90\x6c\x9c\xf2\xd3\xf3\xb5\xae\x9f\x89\xd6\xe5\x83\xb8\x96\x85\xca\x32\xcd\x1d\x5b\x55\xfd\x80\xea\x9a\x37\x69\xa2\xf4\x6c\x85\x63\x5f\xac\xdc\x0f\xbe\xbe\x97\x69\x08\x64\x3f\x35\xab\x75\xd0\xa1\x5f\x23\xf3\x45\x62\x87\x97\xf8\x49\x2a\xda\xf7\x78\xa1\xb5\x20\x54\x3a\x40\x18\x51\xa9\x48\x85\xcd\x50\x66\xb4\xb0\xb6\x20\x61\x79\x02\x04\x1e\xee\xd4\xe7\x2b\x6e\x18\x0d\xbb\xe7\x55\xb2\xe4\x9d\xea\x46\x6a\xca\xde\xbb\xa8\xd7\x0c\x41\xab\x9e\x23\x86\x69\xc6\x64\xb9\x7c\xee\x8d\xba\x8d\x73\x4b\xa7\x1a\xbb\x61\x38\x7c\xdb\xbf\xcf\xcc\x99\xb9\xad\xb0\x4b\xb2\x41\xcb\x98\x9e\x76\x12\x31\x77\xcf\x8c\xb4\x46\xdd\xed\x2a\xdc\xed\xb6\x59\xb0\xd1\xf3\x68\x66\x75\x60\x1b\x2d\xe1\x13\x91\x38\xff\xb8\x3f\xff\xed\x0d\xd2\x0a\x6b\x3d\x4f\x75\x90\xa6\x0c\xba\x72\x3c\xc3\xb5\x0c\x4c\x4e\x5a\x5a\x18\xd8\xf6\x10\x55\x57\xc6\x8f\xf3\x8e\x4b\xf2\xc9\x79\xdf\x4a\xf2\x18\xbf\xc8\xdd\xb0\x96\xf7\x60\xda\x00\x5e\x71\x34\xf9\x38\xcb\xf7\x4e\x41\xfd\x27\x20\xa7\x0d\x48\xa7\x1b\x43\xed\x6e\xae\x9d\xa1\xba\xca\xe0\x64\xa6\xc4\xc5\x5b\xa3\x56\xa5\x47\xe9\xfa\xa6\x7c\x96\xaf\x91\x1d\xea\xe2\xb8\xee\xa9\x5e\x62\x67\xb1\x8c\x39\xe2\x9a\x73\x64\xd3\xe4\xe4\xe3\x6e\x9f\xcb\x56\xce\x3f\x56\xb6\xe1\x89\x24\x20\xbb\x4e\x2f\x80\xd8\x77\xca\x53\xbc\xfb\x64\x42\xad\xba\xe5\x2a\xb9\x39\xe9\x7a\x9a\x1f\x70\x5d\xcd\xde\xc9\xb3\x18\x27\x62\x0a\x4c\x7f\xc8\xe2\xeb\x13\x46\xd7\x66\x06\x74\xdc\x36\x6d\x3d\x87\x40\xca\x92\x6a\x4d\xf5\x9a\x50\x80\xaa\xda\xd6\x2b\x82\x30\x64\x9f\xd9\xc4\x10\x48\xc1\x0a\xf8\x6f\xd5\x16\xf6\x9f\xaf\xd3\xff\xd9\xff\x37\xd3\xe8\x01\xfc\xbf\x4f\x40\x00\x04\xd0\x62\xbf\xb3\x8d\x70\x1b\x07\xc6\x02\x2f\x41\xab\xdc\xc6\x81\x05\x3e\x84\xd5\x9f\xf3\x49\xdf\xbf\xd7\xe1\x93\xbe\x7f\xcd\xca\x93\xee\x18\x5d\xea\xfd\x93\xee\xbf\x68\x27\xdd\xff\x5e\xa1\x4f\xba\x5f\xfc\x87\x24\x3e\x49\x62\x4f\x2d\x84\x54\x9d\x20\xfd\x22\xff\x49\xfa\xff\x6a\xf7\x2f\xd2\x4f\x9f\xe8\x7f\x59\xc2\x4f\x52\x45\x73\x7e\x53\x55\x40\x8f\x03\x9b\x46\x0a\xce\xfd\xdc\x7f\xfd\xc5\x49\xea\xdf\xef\x96\x93\x54\x36\xea\x7f\x77\xd3\x9c\xa4\x3e\x58\xa5\x80\xfc\xa4\xd9\x4e\x89\xfd\x4e\x7f\x05\x0b\xa6\x56\xff\x9d\x4a\xf5\xf3\x35\x03\xa0\x02\x50\x01\x18\x28\x00\x80\x69\xfa\x9f\xd3\xff\x04\x00\x00\xff\xff\x6b\x49\x95\xac\xfb\x0c\x00\x00")

func localZipBytes() ([]byte, error) {
	return bindataRead(
		_localZip,
		"local.zip",
	)
}

func localZip() (*asset, error) {
	bytes, err := localZipBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "local.zip", size: 3323, mode: os.FileMode(438), modTime: time.Unix(1568270560, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"local.zip": localZip,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"local.zip": &bintree{localZip, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
