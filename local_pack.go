// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var _localZip = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x97\x79\x3c\xd4\x7b\xdf\xc6\xc7\x32\xb2\x8c\x65\xe4\xd8\xb7\x41\x14\x42\xc6\xde\xa0\x2c\x29\x43\x06\x43\xd9\xb7\x10\xc6\x56\xb6\x91\x35\xb2\xa6\x94\x35\x21\x3b\x8d\x83\xca\x52\xd6\xc9\xbe\x8c\x31\x0c\x93\x7d\x29\x8d\x30\x94\xb1\x84\x70\x3c\xaf\x9e\xe7\x3e\xf7\x71\x9f\xf3\xdc\xe7\x75\xbe\xff\x7c\x7f\x9f\xd7\xeb\xfa\x5e\xef\xff\x3e\xd7\xef\x42\xc0\xe9\xe8\xb9\x00\x3f\xcf\x0e\x97\xb4\x09\xe0\xc4\x61\x00\x00\x00\x28\x5f\x17\x27\x94\xfc\x1f\x9a\xc3\x3f\x69\x40\xff\xd6\x84\xde\x76\xd0\xbd\x7e\x42\x49\x89\xf8\x4f\x25\xef\x9f\x94\x46\xba\x0e\xc6\xfa\xe6\xe6\x97\x0d\xf4\xcd\xff\xfd\x8a\x11\xf0\x32\x53\xd1\xe4\x70\xcd\x4b\x68\x09\x08\x00\xd8\xb0\x00\x00\x12\x7f\xf3\xca\xcd\x03\xe5\x1a\xe0\xeb\x8b\xba\x2b\xe7\xed\x1b\x88\x34\x36\x12\xb4\x10\x8c\xf4\x06\xb3\x58\x54\x5a\x54\x26\x3a\xd4\x44\xd0\xbb\x70\xd7\x19\xaf\x16\x52\x97\xda\x43\x91\x2e\xfc\xec\xca\xa0\x66\xf6\x20\xe2\xe4\x33\x7f\xb7\x06\x19\x65\xd3\xd4\x0e\x9d\x01\x6e\xbd\xb3\x2f\xea\xc7\x00\xdd\x3d\x72\x12\xf5\xae\xcd\xb8\x0e\x52\xf1\x05\x61\xd8\xde\x93\x6c\x74\x2d\x2e\x51\x4f\x33\x74\x63\x77\x4e\x20\x88\x87\x39\x2d\xef\x08\x7f\xef\xf8\xdb\xfc\xc1\xf1\xbc\x88\xc1\x6a\x0f\x63\x7e\x39\x55\x2b\x79\x57\xb5\xa5\x66\x21\x6f\xd2\xa0\x8f\x69\x41\xe0\x80\x09\x0d\xe2\x05\x0a\x03\x25\x77\x65\x66\x2f\x2f\x98\xbf\x47\x2f\xb0\xbc\x17\xdb\x60\x59\xf8\xf6\xfe\x46\xb4\x3b\x3b\xc7\xa6\xd8\x2b\xf6\x3d\xf0\x46\x33\xdd\xc0\x71\xb4\x86\xce\x1c\x2d\x51\xa3\x5e\x8b\x05\xc5\x60\x0f\x50\xa1\xe1\x89\x26\xd0\x45\x62\xcc\x23\x54\xa3\x5f\x88\x66\x75\xb2\x57\x84\xcc\xee\xc4\x0a\xd2\x5d\xff\x50\x09\x63\xde\x31\x3c\xe6\x29\x60\x3d\x2c\xb7\xef\x7e\xac\x79\xb5\xcf\xeb\x3d\x88\xc3\x46\xe2\x14\x8c\xd6\xf6\x94\x16\xfd\x21\x5f\x08\xef\x9e\xfb\xbb\x39\x7a\xcf\xf7\xc7\x51\xf2\x74\xda\x5a\x15\x87\xdf\x01\xf3\x74\x11\xf7\x10\x2b\x47\x2c\x61\x0c\xef\x6a\x2e\x4f\x6a\x8f\x17\x7c\xc0\xb3\x0e\x7c\x9d\x65\x28\x75\x57\x39\x7b\x40\x88\x92\x87\x1c\x9f\x13\x8d\xb8\x55\x65\x3a\x79\x34\x51\xf1\x01\x1b\x47\xc1\x46\x89\xd0\x1e\xfb\x20\xd1\xf9\x5a\x7a\x9f\xb7\xe8\xed\x11\x86\xfc\xd7\x76\xd4\xad\xb7\xce\x87\x31\x1e\x85\xdf\xb7\x29\x52\xe3\x24\x2c\x42\x98\xe7\x4d\x3e\x85\x4b\xd8\x8a\x88\x4d\xbf\x88\x81\x24\xe5\x02\x4d\x02\x3f\x5a\xe7\xee\x77\x3a\x39\xc0\x92\x78\xb6\x4f\x23\x74\x39\x73\x7c\x1a\x12\xb0\x4d\x6f\x85\xca\xd0\x0f\x7b\x6c\x1f\xf7\x04\xc4\xeb\x6b\xe2\x56\x4a\x42\x2f\x92\x5e\x37\x64\x7b\xa9\xa0\xd3\xc3\x31\xcd\xf6\x29\xfd\x17\x07\x6f\x3d\x72\x1b\x5b\xcb\x75\x8b\x73\x1a\x81\x66\xba\x74\x3d\x22\xa1\xb0\x84\xa2\x6b\xbd\x24\x02\x94\xbf\x7e\x6d\xc8\x9b\xd6\x1c\xae\x80\x19\x09\x9a\x75\x54\x3e\x10\xe9\xaa\x4a\x9f\x29\x6e\x9a\xf4\x51\x67\xc2\x8f\x9e\xaf\xbd\x12\x21\x39\xd4\xa2\x76\xb7\x42\x66\x1b\xd5\xbb\xea\x97\x4d\x21\xaa\xe0\x13\x13\xd2\x35\xa5\xfa\xc5\xc4\xa8\xe9\x3e\xbc\x4a\x18\xbe\x91\x12\x88\x67\x0c\x53\xe0\x0e\xcb\x43\xeb\x75\x12\x7f\xa9\x33\x12\xfa\xd2\x95\x3f\xaf\x2e\x65\xa4\x04\xe7\xfc\x6a\xfb\xd3\x58\xdd\x32\x7a\xa9\xce\x4a\xae\x78\x3d\xee\x83\xe9\x69\x6d\xcf\x33\x28\x7c\x74\x6f\xfa\x99\xd1\x80\x97\x64\xa4\x9d\x8f\x79\xdd\xc5\xce\xcd\xef\x66\x89\x17\x62\xd0\x28\xd5\x52\x4f\x56\x26\x96\x99\x6b\x8d\xee\xfd\x8a\x5a\x91\xad\xfa\xb2\x9a\x38\xc5\x19\x37\xc2\xc8\x8c\x5a\xca\xeb\xa0\xf3\x1e\x21\xd3\xb0\xe6\x7c\x1f\xc3\xaf\xb7\x2b\x11\xbe\x55\xe0\x1e\x65\xa9\x2b\x16\x6b\x9e\xb9\x0a\x3c\xe0\xc2\x29\xfa\x55\x5c\x05\x7e\x5f\xd0\x45\xc5\x8c\x74\x71\x61\x6e\x02\x1d\x12\x97\x1f\xa5\xfd\xb1\x98\x64\x24\x39\x95\x3d\x2a\xcd\x78\x43\x09\x12\xee\xef\x36\x5d\x37\xe2\x9f\xe3\x14\xcb\xdc\xc4\x24\x19\x54\xdc\x1f\x14\x30\x5a\x99\x1b\x9f\xb1\xbd\x76\xcb\xbb\xaa\x3a\x83\xe3\x73\x14\x1b\xa6\xfe\x72\x13\xf4\xdc\xf8\x1d\xe4\x75\x69\x23\x1b\xa9\x0c\xee\x75\x0d\xa8\xb3\x0f\xbd\x41\x5e\x91\x0f\xb1\xa1\xb0\xdb\xb9\xbc\x62\x79\x20\x61\x2d\x83\x71\xbd\xee\xe9\x1b\xed\x61\x9e\xfb\x52\xba\xda\x52\x8c\x44\xa7\xb1\x19\xd9\x96\x8e\x0c\x84\x36\xb0\xa1\x98\x74\xde\xe9\x9e\xe1\xdd\x62\xb0\xf4\xee\x69\x23\xa8\xdd\x7e\xff\x60\xb6\x4f\x90\x5e\x85\xe5\x24\x09\xf3\xdc\xec\x06\xd4\xc5\x4c\xfa\x7e\x91\x88\xf9\xa3\xaf\x89\xea\x1c\x5d\x85\x0c\xea\x0a\xc6\xd6\x69\x48\xaf\x24\xe5\xfa\xbe\x6a\x6e\x4e\x2b\x16\x12\xea\xb9\xc7\xa3\xe4\x8c\xc5\xb8\xf2\xa9\xb2\xeb\xc8\x8d\x82\x74\x31\xe2\x44\xfd\xf0\xd6\x79\x67\xc7\x58\xb4\xde\x6a\xc1\xeb\xd3\x3a\x0d\x2e\x77\x4a\x6a\x13\x5e\x71\x18\xaa\xdc\xfd\x8c\xc8\x31\xb5\x64\x5a\xe0\xb7\x4c\x52\x3f\x3f\xcd\x6b\xa2\x81\x78\x57\x43\xbd\xe5\xd2\x87\x6f\x95\x2c\x2a\x82\x78\xce\x7a\x58\xaf\xe3\x2c\x6d\xb6\x3e\xd5\x3b\xfb\xeb\x98\x96\x8f\x72\x43\x9d\xd7\xb3\x0f\xf4\xd6\x24\x36\x1f\x7b\xf5\xea\xaf\x6c\x52\xba\xfa\x55\xc5\x37\x45\xc9\x7e\x95\x54\x13\xab\x21\xf5\xac\xb2\x4f\x83\xbf\x6e\x33\x3f\xb8\xef\x81\x6e\x96\xcc\x61\xae\x29\xac\xa2\x88\x7a\x3e\x13\xd6\xcd\x2a\x92\xcb\xe6\x68\xdb\x70\xdd\xc0\x1b\x79\x68\x08\x0c\x96\xc5\xc8\x78\x60\xfd\xb2\x08\xb3\x5b\x77\xf9\x22\xbc\xee\x13\x64\x7d\xaa\xd9\x36\x7f\xf3\xb0\xef\xdb\x42\xb1\x4f\x2e\xe7\x48\xa7\xc8\x71\x62\xd5\x70\x5f\x92\x53\xf7\x1a\x65\x4a\x5f\xb7\xe3\x49\x3e\x57\xaf\x84\xf8\xc7\x12\x5c\x43\x8e\xed\xde\x3c\x19\xda\xa8\x49\x0a\x53\x2b\x9e\x01\x03\xc9\xa5\x0f\x0a\xca\xb3\xe8\x4f\xc3\xc3\x42\x20\xd9\x37\x65\xf0\x64\x1c\x65\x14\x86\xe5\xb0\x6d\x95\x2b\xea\x53\x3a\xb4\x6d\x53\xba\x6b\xcc\x90\xb7\x7d\xa3\x7c\x60\x92\x52\xba\x16\x72\x46\x46\x2b\x38\x46\x9e\x9d\x69\x65\xd8\xf8\x33\xf7\x23\x3e\x54\x48\xb0\x06\x2f\xd6\x86\xbf\x95\xfb\xe1\x4e\x9a\xa3\x47\x5d\xcb\x87\xc5\x2f\x4b\x87\x33\xc6\xae\x4d\x8d\x53\x22\xdb\x0e\x2f\xdf\x4d\xda\x0a\x89\x4a\x93\x5a\xbb\x8f\x6a\x05\xd2\xfc\x86\x2d\xa8\x34\xac\x7e\x5d\xc9\x2e\x8c\x58\xd3\xec\xb3\xee\x5f\x10\xac\xdf\x9a\x17\xbf\xc8\x41\x46\x1c\x03\xaa\x72\xdb\xc2\x8f\xc3\xcb\x8d\x0a\x74\x3d\x99\x24\x4b\x29\xf3\x95\x8c\x45\x99\xe4\x36\xe2\x36\x7b\x3b\x19\x5d\xaf\x7b\xb8\x53\xd9\xee\x18\x9c\x2a\xc2\x16\xa8\x2a\x19\xaf\x5f\xf6\xf4\x15\x71\xac\x5b\xe1\xf9\x6b\x72\xca\x36\x07\x51\xf0\xcb\x4a\x8b\x61\x2d\xa7\xe1\x59\xd6\x00\x20\x94\xce\x66\x9b\x3d\x0c\x0f\xc5\x15\x60\x91\xcb\xe2\x3a\xaa\xda\x78\x03\x06\x5d\x6d\x28\x5d\xcb\xae\x8a\xec\xde\x00\xec\x9b\x1d\x58\xca\xd6\xde\x24\xde\xf0\x6d\x4f\xcc\x5a\xde\xc2\x26\x23\x81\x65\x08\x9e\x95\xd5\x8e\x56\x68\x8a\x5d\xf3\xf4\x15\x1e\x61\xd5\x85\x47\x38\xe5\xef\x82\x32\xc6\xa5\x62\x26\x7f\x99\x53\x65\x3a\xb2\x93\xcb\xef\xb3\x57\x65\xcf\x11\xe3\x35\xf2\x35\x72\x43\x9e\xf5\x7f\x9d\xa0\xd5\x5f\x42\xb2\xde\x97\x77\x6d\x5e\xe0\x2b\x4f\xb9\x07\x04\x2e\x58\x82\x8b\xe6\x22\x14\xa2\x66\xf4\x58\x15\x9d\x36\x56\x78\xf4\xc5\xf7\x1b\x75\xd0\x85\xf7\x0c\x7e\xc1\xe0\xc7\x88\x84\xdd\xc4\x26\x6e\xf9\x8f\x02\x96\x7e\xa5\x3f\x8c\x2f\x58\x79\x34\x14\x50\xae\x34\x96\x27\xe8\x32\xbc\xbd\xfa\x04\xf9\xd8\xcd\xc0\x3d\xd6\x4b\x49\xcb\x54\xa0\xef\x65\xcb\x9b\xae\x71\x39\xe7\x5b\x1f\x2e\xa8\x7b\xae\xed\x3b\xfa\x7a\x3b\xb8\x2c\x81\xbb\x52\x44\xb2\x8e\x9d\x17\x80\x74\x11\x85\x3c\x28\x17\x10\x50\xd9\x26\xcf\x6a\xe5\x5d\xd5\xf5\x1b\xb5\x12\x07\x79\xbb\x3b\x9b\x6b\x9a\x9b\xda\x7f\xe4\x84\x69\xa6\xa2\x89\x73\x5b\x94\xe1\x01\x3d\x00\xa0\x03\xfa\xa7\x39\xe1\xe7\x9b\x67\x09\x37\x11\x37\xe0\x68\x33\x18\x88\x5c\xf6\xfb\xa8\xc8\xe1\x5d\x82\x8f\xf6\x62\xeb\x47\xdf\x9d\x90\xc1\x1b\x4d\x0b\x5f\x80\x6a\xae\x0a\x2f\x06\xb2\x81\xe9\x0f\x2f\x79\xd1\x41\x5a\x40\x28\x6e\x89\x4e\xb1\x74\xe5\xab\xcf\x81\x35\x29\x1f\x09\x46\x84\xec\xd3\x91\x10\xd2\xac\x3d\x74\x09\x62\xd3\xf8\x30\xd4\xad\x09\x8b\xf9\xb1\x5a\x89\x09\x42\x4f\x35\x0c\x4f\x4c\xbc\x9d\x7f\x07\xc3\xd3\x95\x65\xae\x76\x83\xe6\xe8\x6f\x0f\x76\x37\xa8\x9e\x13\x09\x8d\x9f\x39\x60\xf1\x17\x4c\xb2\x37\x5f\xba\x39\x7f\x66\xe8\x41\x1e\x9d\xff\x58\x63\x33\x26\x16\x5f\x15\x52\x7b\x27\xb8\x5e\x81\x2b\xd1\x3c\xab\x20\xca\x2c\x5e\x1c\xa6\x5b\x4b\x63\x38\xc9\x76\x4d\x47\xe5\x4b\xe7\x87\xfc\x9b\x35\x7a\x36\xa5\xaf\xc2\xc1\xeb\xf2\xc9\x1d\xe5\x21\xc1\xd7\x6c\x85\x13\x3b\xc8\x88\xbd\x17\x89\x06\xbc\x41\x9c\x51\x6f\xb7\x9a\xa9\x3d\xfd\xc3\x31\x92\x81\xe1\x83\xef\xeb\x05\x2d\xe4\xfc\xf5\xf8\xde\xa0\xae\x0c\x46\xbd\x09\xa3\x95\xd7\x98\x5e\xb6\x97\x8f\xc3\x87\x4e\xc0\x48\x5a\x73\x76\x19\xc5\x9b\x99\xf6\xc9\x3f\xb2\x9a\x2b\xc9\xe0\x9b\x7d\x8a\x69\x05\x1e\x1e\x1e\x1a\x4c\x73\xd6\xa3\x6a\xba\x60\xd9\x0e\xba\xe9\xce\x68\xe3\xd6\x5e\x05\xaa\xbd\xb8\xc6\xd9\x94\x84\x74\x5c\xaa\xa8\x2e\xef\x8f\x81\xa8\x6b\xad\x51\x21\xd1\x0f\xe7\x8b\x2f\x13\x1b\xf3\xa0\xaf\x83\x7e\x5d\xe4\xe9\xbe\x7f\x3b\x8c\x2f\x69\x1d\x2a\x90\xba\xc4\xf8\xbc\x30\x57\xf4\xea\xdb\x17\x42\x97\x6c\x7f\x83\x23\x09\xcd\xee\xbd\xfe\xcd\x6a\x5c\x94\x71\x10\xfe\xe5\x98\xe8\x34\xdc\x74\x09\x33\x91\xe6\x28\x6d\x0b\x48\xe6\x66\xd1\x6d\x17\x75\x36\x7e\x1e\x3f\xdb\x0e\x5a\x56\xaa\xb6\x88\xd7\xb9\x05\x11\x4b\x23\x1f\x45\xdd\xcf\x7f\x92\x28\x04\xa5\x2c\x00\xb8\xe8\x87\xab\x77\x88\x97\x60\x31\x0e\xa2\x6b\x56\x9e\x3b\x9d\x38\x11\xd6\x3d\xd7\xbb\xe9\x77\xfc\xf8\x4d\x09\xfd\xc5\x82\xae\xb5\x82\x34\xc7\xf9\x1a\x25\x2e\x8b\x4f\x77\xa4\x32\x78\x82\x7c\x2e\xf5\x4f\x95\x54\x70\x35\xa5\x74\xbe\x7b\x5d\x75\x04\xbc\x87\x78\x00\xdb\xd2\xd1\x2c\xbd\x70\xb3\xe0\x95\x1d\xc7\x81\x94\x53\x2d\xf3\x33\x0a\xee\xf4\x04\xb8\x9b\xe0\xb4\xe5\xdd\x42\xd5\x9a\x92\x7d\x8c\xe0\x8d\x10\x76\x5a\x9f\x81\x9e\x81\xac\x5a\x98\xb6\xed\xcd\x77\x6f\x24\x65\x5b\xf5\xac\xab\x9f\xd3\xd4\xd6\x46\x34\x51\x8a\x6a\xab\x47\xe2\x8b\x0e\xbb\x12\x6e\xcd\x7f\x62\xda\x4c\xcd\xd3\x51\x69\x31\xef\x68\xa5\x02\x9e\xa8\xaa\x98\x0c\x98\x08\x1e\x83\x8f\x5a\x5f\xc8\x72\x6f\xf0\x2e\x73\x8a\x04\x47\xa9\x3a\x22\xd2\x77\xf5\x9a\xcf\xf0\x27\xeb\x56\x61\xbf\xfe\x02\xce\x12\xdb\xe1\x33\x86\x76\xa8\xe7\xb4\x0f\xa6\x6a\xd1\xcb\xc2\x41\x2b\xd7\x86\x82\x72\x4f\x99\x7d\xfc\xee\x7b\x09\xb5\x3a\xe7\x51\x3f\xd2\x95\xff\xa1\x0f\x74\x59\x54\x89\xe5\x3a\x03\x1f\xe2\x0a\x36\xb3\x24\x49\x78\x7c\x69\x28\xd0\xd6\x37\xd0\xf7\xbb\x1e\x45\x2d\x34\xf8\xac\x38\x02\x12\x87\xb7\x6c\x39\xb6\x07\x22\x59\x6a\x82\xd2\x06\xa2\x13\xf4\x8f\x62\xd4\x1c\x32\x22\x6f\x98\x40\xbe\xe6\x03\xd5\xc4\xb9\x0d\x47\xbe\x77\x27\x68\xa0\x4b\x0d\x5f\x96\xa7\xdf\xe2\x84\x5b\x50\xdd\x4f\x31\x7b\x76\x05\x29\xc1\x13\xd8\x75\x6a\x47\x87\x67\x73\x0d\x46\x3f\x2e\xd9\x21\xad\x78\x5b\xe1\x1a\xa9\x7e\xb1\x42\xe5\xe3\x44\x6f\xf6\xc3\x7b\x66\xed\x64\x6f\x53\xcf\xfe\xc9\x51\x18\x50\x33\x61\xa8\x4c\x36\x75\xb6\xf4\x36\x5a\x58\x72\x38\xa1\x47\xcd\xf5\x32\x36\x15\xb2\x8f\x2e\x89\xb3\x1b\x90\x7c\xd4\xcd\x7a\xd5\x9e\x8d\xfa\xeb\xb7\x4e\x39\x67\x63\x48\x7f\x66\xe1\x1b\xdf\x02\x18\x1e\xf3\x10\x97\x49\xdb\x1f\xbc\x66\x55\x14\xfa\x20\x5a\x52\x52\x42\xfa\x52\xff\xba\xf9\xf0\x36\xfe\xc0\x64\x73\x8d\x6b\xa9\x2c\x7a\x2c\x38\xcb\xf1\x71\xe5\x19\xaa\xec\x34\x5c\x80\x71\x17\xfb\xe9\xe0\xb1\xa7\x7c\xc0\x99\x9d\x1b\xfc\xc0\x07\xe4\x6c\xa1\xbe\x06\x69\x32\xcb\x85\x0a\xbd\x47\x8b\xfe\xf5\xde\x33\xfc\x2f\x90\x1a\x19\x91\x48\x59\xc6\x03\xad\x76\x90\xc6\x33\x7c\x88\xc9\x28\xb4\xb5\xed\x54\xec\xb2\x71\x9c\xc6\xe0\xb9\x8a\x10\xb3\xcc\x28\x98\x4f\x85\xa6\x82\xbb\xed\xfe\xb7\xcb\xe7\xcf\x2a\xe8\xe6\xa7\x3e\x0f\xeb\x40\x3e\xde\x75\x34\x8c\x50\x7e\x2f\x94\x23\x8e\xba\xe3\xd7\x10\xec\x6b\xd8\x2a\xd3\xd9\xbb\x5e\x41\x55\xb7\xb6\xe2\x39\xff\x92\x1c\x42\x16\xac\x13\xf4\xce\x8d\xb6\x7f\x8a\xc4\xc9\xe0\x56\xad\x26\xb4\x38\x93\x2b\x67\x9f\x5b\x07\xf8\x9d\x4b\x52\x9e\x94\x15\x81\x7a\xc3\x0c\xea\x51\xe5\x19\x44\xef\x2a\x8c\xf1\x35\x49\xd7\x23\x0b\xf8\x0f\xbb\xc5\xde\x14\xf4\xa5\x38\x69\x5a\xb1\x8f\x1a\x74\x4a\x03\x0e\x0f\x14\x97\x41\x61\x69\xaa\xfc\xe4\x57\x6a\xf2\x71\x42\x56\x05\x4e\xb5\xb3\x9a\xf6\x21\x1c\x0f\xe4\x1c\x8e\xf8\x1b\x92\x14\xab\xd3\xde\xbe\x8a\x33\x07\x94\x52\x6e\x52\xea\x1d\xba\x46\x0e\xe3\xbf\x3a\x9f\xc2\x08\x72\x81\xe4\x1e\x21\xc4\xaa\xab\x5c\x2f\xae\xbb\xce\xf2\x75\x94\xb9\xaa\x12\x39\xe3\x1a\x25\x05\xac\x9f\x52\xa7\x72\x50\x3b\x18\x42\xdf\x8f\xe1\xac\x3a\xfc\xd6\x34\x71\xf5\x2a\x74\xb7\x4d\xd9\xe4\x40\xad\x0f\xf6\xae\xee\x92\x20\x60\xaa\x15\x76\x7a\x1c\x2e\x3e\x13\x25\xc8\x99\x5a\x7e\x65\x34\xbc\x32\x70\xc5\xd6\xb7\x9a\xc1\xfa\x06\x36\x52\xfb\x9c\x03\x09\x18\x20\xe6\x98\xed\x2e\x5a\x7c\x2e\x37\x90\x61\x49\x67\x7f\xa1\xc1\x72\x7c\xb6\x80\xc7\x09\x0d\x6f\x3e\x45\x3f\xdf\x10\xdc\x94\x12\x1a\xf9\x90\x76\x74\xa4\x27\xf0\xed\xf6\xca\x85\x00\x6c\xda\xc3\x08\xa1\x67\x11\xdf\xa7\xd4\x97\x7b\x11\x5d\x20\x37\x6e\x22\xd7\xec\xe8\x48\x31\x4a\x01\x1c\x4a\x96\x48\xe9\x8e\x3f\xbe\xcd\x56\x42\x4f\x48\xa6\xce\x61\x47\x06\x9e\xfc\x56\x04\x3a\x3e\x85\x80\xd3\xd0\x72\x01\xfe\xdb\xaf\xbe\xc4\xbf\xbe\x38\xfe\x75\xff\xdf\x32\x65\x06\x88\xfe\xef\x44\x03\xe0\x05\x34\x39\x50\x37\x34\xdd\x49\x34\x89\x34\x66\xc2\x55\xee\x24\x9a\x17\x3e\xb8\xe5\x9f\xf3\x49\xdf\xbf\xd6\x83\x93\xbe\x7f\x5e\xd2\x27\xdd\x03\xf5\x19\xf6\x4f\xba\xff\x4e\x3b\xe9\xfe\xd7\x4a\x71\xd2\xfd\xfa\xdf\x44\xc0\x49\x12\xf7\xd3\x02\xde\xaa\x13\xa4\xdf\xc9\x7f\x90\xfe\xbf\x1a\xf2\x3b\xe9\xa7\x4f\xec\x3f\x2c\x25\x27\xa9\x68\x9b\xd8\x2b\xec\x3e\x24\x1a\xef\xdd\xef\x26\x3f\xef\x88\x8b\x18\x9c\xaa\xc7\x7f\x52\xff\x1a\x6a\x27\xa9\xbf\x32\xfc\xb3\x88\x3b\x49\x45\x16\x50\x3e\xb3\xf9\x90\x68\xec\x84\x77\xff\x44\x05\xfe\x6c\x77\x00\x20\x00\x08\x00\xd1\x02\x00\x2b\x2c\x3f\xa7\xff\x09\x00\x00\xff\xff\x57\x66\xe9\x39\x0b\x0e\x00\x00")

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

	info := bindataFileInfo{name: "local.zip", size: 3595, mode: os.FileMode(438), modTime: time.Unix(1568804028, 0)}
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
