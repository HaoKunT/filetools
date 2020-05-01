// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package tools generated by go-bindata.// sources:
// local.zip
package tools

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

var _localZip = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x97\x67\x3c\xdc\xf9\xbe\xc7\x67\x8c\x11\xa3\x26\x21\xc6\xe8\x84\x28\xd1\x89\x2e\x6a\x58\x82\xd1\x7b\x0b\x51\xc2\xe8\x3d\x4a\x88\x10\xa2\x45\x5f\x44\x89\x1a\x89\x10\x7d\x44\x17\x8c\x92\x31\x88\x16\x75\xf4\xde\x83\x31\xda\x7d\xed\xd9\x7b\xce\x3d\x7b\xf6\x64\xef\x7e\x1f\xfc\xcb\x83\xef\xfb\xf3\xfa\x3d\xfa\x7d\xde\xda\x0f\x41\x84\x24\x80\xdf\x86\xc6\x9f\x17\x0e\xf8\xb7\x21\x02\x30\x00\x10\x6e\xb6\x8f\x10\x02\x06\xfa\x10\x00\x88\x7e\x3b\xc0\x22\x8f\xaa\xd0\xd2\xc7\x9f\x14\x00\x24\x5c\x07\x81\x00\xff\x78\xfc\x7c\x9f\xec\x5f\xfb\x4f\x1d\xad\x94\xb4\xfe\x26\x65\x1b\x96\xa1\xfd\xef\x14\xda\xff\xa0\x68\x28\x59\x69\x3e\xd0\xd3\x53\x50\x7d\xa0\xf7\x3b\x11\x2d\x55\x60\xf9\x13\x22\x35\x00\x00\x20\x06\x5c\x4f\xcd\xd0\x96\xc8\x6e\xef\x68\xba\x06\x00\xb8\x51\x00\x00\xec\x7f\x41\xb4\x7f\x82\xb0\xf3\x76\x73\x43\x78\xf1\xbb\xb8\xfd\x83\x9e\x4b\x55\xf8\xdf\xe8\x3e\x06\x73\x6e\x0c\x46\xb7\x5a\xe4\x6d\x16\x24\xb4\xa6\x5e\x5f\x62\x55\xe7\x8b\x23\xe8\x06\xef\x18\x90\x72\xb3\x11\xbc\x20\xbc\xf9\x40\x46\x7e\xb3\x30\x78\x3d\xdc\xc1\xaf\xa2\xb9\xe0\xeb\x00\x99\x48\xb5\x49\xac\xf6\xe3\xf1\x5a\xbf\x6c\x7f\x17\xf1\x09\x2d\x86\xc6\x52\x3e\x06\x69\x9e\x4e\x67\x7d\x06\xb7\x31\x4b\x7a\x48\x27\xa8\x09\x78\xd3\x6d\x4f\xd8\xcf\xf7\x22\x45\x0e\xc0\x84\x1b\x68\x5e\xb4\x40\xb9\x18\xaa\xb7\x34\x9c\xe4\x78\xed\x9c\x0e\x9c\xee\xec\xb7\xc2\x19\x32\xed\x71\xc1\xe0\xab\x96\xd0\x67\x72\xe7\x84\xb6\x0e\x61\x63\x9c\xb0\x33\x9e\xbd\xb3\x5c\x9b\xad\x6f\xf2\x54\xfe\x64\x38\x2f\xac\x07\xd6\xce\x9a\xf3\x13\x61\xd0\x4d\x89\x6b\x08\xd0\x25\xd3\x04\xb0\x0c\x60\x05\xbc\x84\xf9\x53\xe3\xac\xac\x29\x26\x7c\xa3\x2a\x39\x43\x5b\x09\x42\x6e\x48\x50\x04\x11\xf9\x43\x71\x7a\x58\xab\x50\xce\x50\x38\xa3\xfd\x16\xf5\x9e\x73\x5b\x7a\x98\x9c\x64\xd9\x89\x00\x28\xf7\x5a\x0b\x69\x10\x0d\x92\x42\x03\x70\x05\xf4\x8e\x26\x4e\x92\x50\x0e\x20\x80\x13\x08\xd4\xae\x9f\x97\x13\x9c\xb5\xb9\x54\x6a\xae\xce\x84\xbc\xa8\x2c\x09\x83\x33\x0f\x05\x59\xca\xe4\x6f\xb7\x5e\x3b\x67\xf8\x74\x16\x35\xd0\x4f\x14\x42\x6a\x0c\x38\x26\xf0\x0b\x85\x03\xad\x1c\xc7\x4f\x9f\x41\x71\x3e\x58\xb3\xed\xcc\x67\x6f\x2d\x56\x25\x67\xcb\x37\x6b\x89\x83\xc8\x90\x74\x15\xf5\xb6\x8d\xad\x5d\x9b\xad\xa0\x56\xc8\x39\xbc\xe9\x7e\xb4\x0d\xac\x4d\x4f\xf5\xe8\x29\x31\xce\x0e\xbb\xa4\x13\x3c\x27\xb5\x77\x0f\x2b\xb7\x78\x08\xd5\xde\x5a\x0d\x9d\x03\x84\x30\xe1\x54\x5e\xb6\x60\x62\xc6\xdd\xfc\x1c\xaf\xcc\x18\x33\x51\xb9\xa1\x73\xc5\x21\xde\x6b\x32\x03\x4c\xdf\x3c\x5a\xbc\x34\x50\x57\x42\x7b\x01\x23\xdc\x20\x65\x75\x44\x99\x1e\x2a\x6b\x7d\x64\x73\xd9\xbf\x99\xa6\xe2\x6a\xa2\xf4\xe2\x6a\xb2\x74\xf7\x14\xba\x47\xb0\xf7\x6c\xff\x74\x13\xe5\x23\x3b\x2b\x55\x7e\x76\x20\x71\x19\x1c\xbd\xce\xbc\xbe\x1d\x8d\xcb\x42\xd1\x15\xce\x39\x64\x94\x0b\x5d\x70\x71\x2e\x28\xc8\xca\x31\xec\x26\xf2\xe7\xc4\xb5\x4e\xd9\x9e\x93\x7d\x0b\x56\x0e\xd4\xbf\xfa\xa5\x1a\xd3\xe5\x90\x50\x39\x96\x50\x03\x79\xaa\x9e\x9c\x22\xf0\xb8\xf4\xc7\x20\x4d\xae\x1c\x6c\x51\x44\x28\x67\x7a\x35\x86\xf3\xad\xa8\x17\x9b\x25\xd5\x2e\xb3\x7f\x07\x31\x21\xfc\x0c\x08\xcf\xca\xf9\x05\x1c\x3f\x8e\x41\x5b\xba\xda\x6c\x13\xcf\xed\x90\x20\x71\x11\xeb\xd7\xb6\x54\x3b\x68\x72\x1e\x55\x1a\x45\xd8\x49\x9e\x6b\xa2\xaa\x87\x4b\x26\x36\xd9\x52\x0f\x83\x0e\x93\xb0\xdc\xa8\x6d\xd4\xb0\x58\xd5\x73\xfa\xa6\xf1\x1a\xc2\xcd\x01\x89\x80\xa3\x2e\xcd\xc6\xf1\xfd\x2a\x15\x6b\x29\xb1\xcd\x01\x8e\x6a\xfb\x36\x7e\x64\x9a\x6d\x61\xed\x3c\x78\xae\xe6\x0c\x3b\xea\x62\x6f\x61\xdc\xfa\xf4\xf2\x1c\xe3\x5f\xbf\xc7\x0f\x49\xb8\x46\x31\xdc\x33\xd5\x37\xb6\xef\x5f\xbe\xbc\x5c\x2d\xae\x91\x98\x20\x37\xf7\x35\x8a\x05\xe4\x95\x5b\xee\x2d\x5b\x82\xa1\x6a\xc1\x24\x47\xa6\xb3\xf0\x87\xa5\x53\x41\x6c\x52\xc4\xf9\x62\x4a\xc8\xee\xa6\x56\x47\x89\x9a\x23\xdd\x2e\x75\xdd\x4b\xea\x1f\x58\xc2\x8a\xc9\xb2\xa4\xeb\x46\x31\xc9\xcf\x57\x83\x12\x23\x1d\x5e\xcf\xd7\x7d\x49\xd4\x1b\xde\x07\xef\xd6\xc6\xde\x4f\x13\xd5\x85\x72\x7f\x89\x05\x0d\x97\x0f\x71\x1a\xf7\xf5\x7e\x8e\x7c\xd9\x13\xa0\xea\xa0\xa2\x26\xae\x24\xc2\xa3\x6b\x9e\xfd\x92\x72\xb8\xe3\x66\x18\x0e\x59\x60\x38\x6b\x72\xcd\xce\x70\x15\x99\xa4\xdb\x7a\xbe\x00\x33\x25\x8a\xb7\x3d\x2d\x01\x61\xf6\x44\x64\xf8\x9e\xc6\xbd\xb9\xc5\x92\xc3\xfc\x6b\x1d\x8f\xc4\x2e\x6b\x01\xd3\x98\x28\xd8\xce\x60\x45\xe3\x3a\x52\x2a\xa7\xc2\x09\x2d\x1c\x99\xbe\x22\x95\x3c\xf6\xd1\xc3\xde\xdb\xea\x21\x33\xbe\x7f\xc3\x01\x1a\xe1\x9e\x58\xfe\xfd\x65\x3f\xef\x8c\x46\xbf\x72\x60\x6a\x70\x42\x76\xff\x9d\x88\x71\xa5\xe5\xd1\x2e\xf9\x85\x11\xd6\x09\x79\xd7\x48\xbb\xae\xa9\x62\xf2\x95\x8f\x36\x5c\xf9\x09\xbc\x11\x8b\x16\x3f\x1a\xec\x2c\x95\xcf\xb7\x8a\x98\x56\xe8\x5a\x12\x9a\x45\x2d\xf2\xf1\xa5\xc8\x77\x75\x27\xd2\x92\x35\xbc\xd8\x96\xbb\xba\xf6\x02\xb7\x5a\x9e\x20\x26\x9e\x6c\x7c\x50\xf5\xd4\xdc\x3b\xd4\x94\x43\x67\xbc\x71\x27\xb9\xc0\xd9\xbf\xa2\x97\xec\xac\x05\x89\x14\x1d\x6d\x4a\x59\xb6\x99\x63\x6a\xc2\x8b\x69\xdf\x98\x63\x57\x08\x56\x92\xc9\xbc\xa3\xe3\xd8\x8f\xbd\x65\x83\x26\x3e\x7c\x15\x1d\x7f\x2c\x67\xe4\xdb\x5c\xfc\x69\x38\x84\xc0\x6f\xb9\x21\xd7\x8b\xfb\xfd\x2f\xb9\x8b\x52\xac\x37\xed\x14\x97\xe1\x8a\x85\xfa\x6d\xe6\x1e\x14\xc2\x77\xde\x68\xe5\x16\xbf\x23\x7e\xb0\xb2\x3a\xb8\xa5\xef\x1c\x13\x84\x81\x85\xf0\xb4\x8f\x86\x18\xd2\xd6\x55\xdc\x8f\xc8\xee\x16\x79\xe8\x5e\xd0\xe0\x5c\x23\x96\xca\x67\x10\xa1\xb5\xae\x4b\x2c\xb7\x9b\x01\x6c\xb6\xd8\x80\x49\xbc\x2b\x99\xb1\xf3\xe6\xae\x2e\xf7\x1e\x87\x19\xb8\x4a\xe9\x98\x23\xe7\x3f\x77\xec\x05\x95\xc9\xc8\xd1\xc0\x75\x49\x39\x9d\x59\xa2\x84\xd5\xb9\x1c\xe6\xda\x17\xa5\x94\xf8\xd8\x4e\x5e\x54\xb5\xf7\xf7\x0e\x7f\x0f\xe3\x7c\xff\x9c\xa3\x27\x3c\xd6\x7c\xbe\x54\xc1\xa6\x44\xf9\x90\x60\xce\xf3\xa0\x43\x3a\x5c\x6a\x55\x06\x81\x72\x32\xc2\xeb\xe5\x52\xca\xd5\xac\xdc\xaf\x1c\xb2\xe7\x33\x39\x41\xaa\xbc\x96\x44\x42\xd0\x77\x6c\x03\x0a\x47\x75\xba\x81\xb2\x1b\x9a\xb3\xa1\x44\xee\xb1\x97\x9f\xfc\x05\xb7\xbf\x39\x9d\x53\x37\xac\x3b\x5b\xbc\x6f\x5f\xf3\xc2\x78\x1c\x6b\x9a\xaa\x34\x6b\x06\x5c\xf0\xcb\x3e\x59\x6d\x73\xec\x2b\xc6\x3c\x0f\xe8\xf6\x8f\x7e\x55\xdc\x62\x79\xff\x4c\x5d\x51\x4d\x1e\xcd\x2c\xfb\xf6\x47\x78\x8f\x52\xbe\x66\xea\x50\xf6\x71\xbb\x73\xca\x18\x73\x08\x07\x2f\x73\x02\xba\xb9\x32\xd5\x73\x0b\xa0\x5b\x92\xd4\x9e\xdb\xbb\x5c\x5f\x60\x76\x78\xd4\x30\xea\xf8\xf6\x3a\x9a\xba\x03\x1d\xe8\xb9\x1e\x62\xcc\xdd\xdb\x28\xbe\x74\x9e\x7d\xcd\xcc\xb6\x4a\x82\x53\xfc\x83\xa7\x4f\xd4\x36\xc5\x74\x19\x7c\xe6\xd1\x23\x17\x85\xd9\xf2\x79\xf6\xb2\xaf\x23\x25\x28\x26\x04\xf9\xc2\x52\x66\x99\x45\xf2\x43\x0e\x93\x26\x94\x2f\x29\xc7\xa9\xab\xa6\x56\x9a\xb7\x44\xf1\x5b\x1d\xcb\xf8\xc8\x48\x13\x51\x21\x98\x61\xba\xb1\x71\xc3\x4d\x57\xbe\x4e\xa2\xa3\x51\xa0\x49\x66\xf9\x5c\xd0\xcb\x4b\xae\xb5\x66\xc7\x16\xf6\x0a\xde\x05\x0a\x73\x5a\x43\xd4\x97\xbe\x51\x95\x8f\xc7\xaf\xf5\x6a\xee\x0f\x19\x73\xef\xf0\xb9\xa6\x2c\x2f\x79\x7a\xdb\x36\xb6\xe8\xbb\x6c\x7d\xed\xdd\xf0\xfc\xc1\x14\xe0\xec\xf9\xa8\xc8\x34\xf7\x99\x95\x4e\x06\x79\xbd\xab\x98\x21\x74\xf3\xc3\x78\x96\xf9\xf2\xf0\x2b\xf3\xac\xbc\x6f\xd1\x94\x6d\x12\x21\xbb\x7b\x89\xe1\x11\x6b\xf8\x66\xfc\xc8\xd9\x90\x77\x62\x6f\x83\x59\x15\x92\xc6\x50\xcc\x0f\x3a\x6e\xa2\x9f\xe3\x84\xa0\x8c\x48\x2d\xfa\x74\x10\x16\x7f\x68\xfd\x81\xf6\x96\x43\xa6\x10\x33\x22\x4e\xd8\x22\x4e\x95\x19\x0a\xf5\x75\xb8\x61\x9d\x07\x03\xd1\x38\x86\x97\x69\xde\x4e\x9f\x41\xd3\xc7\x8d\x72\x50\x6b\xc4\x97\x85\xcb\x54\xc5\x17\xd7\xd1\x47\x50\xca\xe1\x57\xe5\x1b\xdb\x22\xb7\x77\x1f\x6d\x0a\x00\xfb\xde\x6f\x06\x7e\x42\x96\xb5\xe0\xa4\xb3\x09\x3d\xac\xe8\x4b\x31\x17\xac\x74\xad\x1f\xfd\x8c\x91\x87\x96\xa2\x62\xb4\x06\x9c\x8b\xee\x8d\xd5\xc9\xca\x2a\xb7\x68\x6f\x2c\x4a\x7c\x73\x6e\x8c\x80\x08\x36\xf0\x93\x1a\xa0\x0b\x33\xc8\x61\x71\x14\xca\x87\x05\x81\xb2\x96\xdf\x53\x8c\x76\x82\x52\xde\x35\x89\xa9\xdc\x66\x7a\x0e\x11\x4f\x5f\xea\xde\xfe\x24\xc6\x8c\x9e\x4e\x37\xa3\x8a\x82\xf2\x21\x04\x4d\x98\x4a\x05\xc8\xc4\xcc\xa2\x43\xb5\x95\xf5\xa7\x43\x9b\xb0\xef\xe1\x43\x86\x0d\xf7\x58\x38\xfb\x33\xb9\x48\xc6\xb9\x28\xab\xf0\x6e\xfd\x9e\x0b\xe2\xbc\x74\x17\xb5\xbc\x1f\x15\x82\x8f\xd0\x37\x54\xd4\x0c\x57\xb1\x48\xa3\x28\x1d\xc9\x73\x56\xdd\xae\x2e\x4f\xff\x70\xc2\xd9\x09\xee\x4e\x15\xbf\x0a\xd5\x0f\x9b\x43\xb8\x7b\x26\x1c\x1e\x23\xa8\x71\x96\x78\xc7\xf7\xb1\x55\x5b\xb9\x73\x3f\x68\xef\xf4\xcd\xaf\xb8\x48\x8a\x0e\x4d\x64\x3c\xe8\xf3\x7c\x4c\x73\xcc\x2c\xd2\x42\x95\x90\x39\xa7\xed\x26\x6a\x38\x91\xcd\xb7\xfe\xc1\x51\x26\x39\xef\x21\xb8\x79\x6c\x3f\x7a\xaa\x87\xf7\x71\xf4\x3a\x19\xc1\xaf\x68\x6f\x51\x73\xe5\xf5\x44\x7a\x0f\x6d\x48\x41\x22\x17\xeb\x57\x96\xd1\xee\x03\x2c\x90\xc6\xe3\xc6\xa8\x70\x59\xfc\xf5\x0a\x69\xa3\x7d\x72\x81\x17\xc3\x94\x4a\xf5\xc5\xa7\x3c\xf9\x61\xa9\xdd\xda\xe4\x04\xca\x7a\x89\x6c\xce\x4f\x16\xe4\x45\x16\xef\xd1\x7d\xf9\x41\xe4\xe9\x84\x48\x59\x3e\x96\xe6\x35\x45\x32\x4f\xb1\x7c\x70\x6c\x21\xc9\xc2\xf7\x99\xca\x10\x52\x38\x32\x4d\x1f\x2b\x17\x12\x5e\x1f\xf3\x1d\xf4\x90\xbd\xbf\x3c\xb1\x77\x89\x8a\xf9\x0e\x11\x97\x3b\x4c\xc2\xe7\x12\x71\x1a\x08\x4c\x33\xa8\x94\x08\xb6\xbf\x2b\x0b\x14\x28\x3f\x4d\x5a\x58\x9d\xa6\x1e\x9a\x5d\xd3\x16\x30\x8d\x27\xeb\x1f\x3d\xdb\xba\xfa\xbf\x66\xf0\x5b\xd7\x48\xfc\xe2\xb4\x69\x41\x04\x00\xc0\xff\x76\x33\x70\x77\xfb\xab\xde\x11\x67\xf4\x10\x4e\x6f\x44\x3d\x6b\xb4\xf1\x0c\x89\x8e\x8c\x53\x4a\x55\x5f\x67\x9b\xe2\xcf\xb4\x4b\x55\xad\xb6\xdb\x6a\xe8\xbe\xdd\x2d\x52\x10\x88\x42\x7a\xeb\x28\x40\xea\xd4\x04\x8b\x53\xc4\xd9\x81\xa3\x32\x29\x62\x75\x45\x29\xd6\xe7\xa1\x72\x22\x7e\x08\xc3\xf3\x88\x1c\xaf\x0a\x21\x3f\x19\x16\x9f\xf0\x6e\x91\xd5\xe0\x96\xc3\x6c\xa5\x34\xb8\xe9\x9d\x80\x3d\xec\x66\x45\x55\x50\xd2\x99\xc6\xfd\x3e\x66\x1c\x9e\xbb\xd6\x25\xbd\xc9\x33\xc3\x25\xa4\x20\x9a\x70\x26\xd8\x63\xa2\x79\x87\x17\x8c\x8e\x9b\xc5\x0f\xfb\x38\x68\x19\x2d\xb8\xa6\xdd\x9c\xf3\x71\xd1\x8a\x2d\xcc\x3a\x54\x9b\xac\x2a\x2d\x9f\x79\x43\x6e\xda\x68\xe3\xc3\x3d\xcd\x45\xb4\x2d\x30\x3f\x9b\x6f\x8c\x9b\x9e\x0a\xe0\x1b\xa0\xe8\xea\xc8\x70\x5d\x57\xdd\xd0\xb6\x9f\xc8\xdc\x19\xdd\xd2\x72\xdb\x56\x4e\x57\x6b\x30\xc1\x58\x63\x05\x77\xd5\x55\xdf\x20\x63\x7f\xa5\x29\xe0\x9e\xb3\x13\xc9\x74\x32\xb6\xda\xb4\xfa\x7a\x0d\xed\x34\x78\xfe\x4e\xc6\xca\x4f\xfd\xd1\x13\x25\xd9\x98\x8d\x53\x1f\x07\xa1\x4c\x95\xf8\x22\xae\x0c\xf9\x02\xaa\xcf\xd9\x63\xaf\x74\x87\x22\x34\xfb\x82\xfb\x85\x82\xfa\x7d\x16\xe8\x9a\x67\x4c\x91\x56\xb1\x77\x97\x45\x46\x30\xda\x2a\xae\xba\xbf\x34\xc4\xc3\x19\x99\xa9\xdc\xcd\xad\xfc\x1a\x0c\xa0\x77\x55\x79\x80\x3a\xc6\xd7\x6b\xbd\x66\x62\x51\xbc\xda\x99\xc5\xbc\xeb\x9e\xea\x79\x58\x4b\xfe\xfc\x26\x5c\x09\x97\x38\xfe\xa8\xc0\xb2\x31\xef\x83\x2c\x06\x19\x60\x22\x64\x23\xf9\x24\xdd\xce\x6e\xef\x73\x4e\x2c\xf4\x47\xc7\x29\xd8\x82\x52\xd1\x90\x5b\x4a\x45\xcc\x94\xd6\xa0\x27\xb5\x8a\x3f\x3b\xfc\xb5\x2e\xa2\xfb\xd0\x22\x94\x50\x0d\xe6\xcb\xe3\x40\x7f\xbf\x86\xb4\xec\x31\x1a\x05\xe9\x10\xe9\xbe\x6d\x23\x51\x1c\x84\xdf\xd9\x3f\x81\xe0\x1e\x7b\x24\xd5\xc5\xa9\x2d\x48\x1d\x93\xf5\x36\x22\x5e\x51\xbd\x4b\xbd\x9d\xaa\xdf\xc9\x70\xd7\xdb\x21\x2b\x7c\x87\x83\x45\xb9\x81\xfa\x21\xe9\xc5\xa4\x42\x45\x5d\xa9\x22\xe8\x59\xa8\x49\x4e\x00\xfb\x55\xbe\x88\x3f\x8b\x42\xa0\xc2\x56\x36\x21\x5a\x58\x3f\x20\x0a\xdc\x63\x42\xad\x15\xbe\x32\x35\x3b\xf9\xd2\xbd\x99\x1a\xfc\x46\x5c\x89\xec\x16\x4a\x56\x9b\x6a\xb4\x8a\x2b\xf8\x54\x8e\x25\x50\x45\x4f\xb2\xfb\xe3\x13\x44\x8b\x38\x15\x10\xff\x7e\xad\x9c\x47\xe9\x73\x07\x65\xac\xcd\xa9\x46\xb9\x1d\xbc\x4c\xb4\x12\x5b\xff\xb0\x41\x64\x85\x0a\x1f\xd3\xe5\x03\x88\x2a\x98\xfb\x72\x12\x28\x45\xcc\x4a\xd1\x63\x14\x34\xf9\xbc\xf3\x2d\x61\x7e\xcd\xdd\x49\x3e\x9d\xac\x71\xc6\x66\xba\xd0\xe4\xdb\x1a\x32\x2b\x5a\xc2\x7d\x92\x6a\x74\xbc\x61\x3b\x41\x69\x05\x10\xef\xcc\x7a\x44\x60\xe1\xc7\x69\xfc\x2d\x9e\xe4\x96\x38\x38\x49\x31\x99\xea\x91\x88\x1a\x24\xb4\x83\x50\x53\x0e\x64\xc1\x40\x00\x95\x66\x42\x9d\x8f\x3d\xa8\x98\xdc\xb0\xda\x79\x35\xa6\x83\xdb\x01\x84\xa7\x7e\x27\x44\x1b\xba\xed\x86\xc7\x06\x4b\xaf\x6d\xad\x59\x44\x2e\x29\x1f\x19\x8b\xf2\xe0\x6a\x0f\x6e\xe8\x32\xf8\x03\x18\x85\x8c\xc8\xc4\x5b\x3b\x03\xf1\x1a\xcf\x03\x06\x0d\x9a\x49\xbb\xaa\x70\x05\xa4\xaf\xe3\x29\x77\xa9\x8b\x98\x7e\xcc\x65\xbe\xa8\xac\xb7\xa5\xf6\x3b\xe5\x37\xcf\x6c\x4c\xe6\x46\x38\x52\x08\x21\x02\xe5\x0d\xd3\x40\x96\xbb\x9d\xda\x4f\xf4\xa8\xcd\x7a\x52\xd1\x09\x83\xd2\x43\x16\xaa\xa2\x3b\x4d\x14\x50\x87\x30\xa1\x45\xdf\x86\xfe\xf3\xf9\x50\xf9\x02\xfb\xa8\xa5\xf9\x9a\x07\x2e\xc9\x9c\x79\x83\x01\x1d\x5e\x39\xd7\x74\x47\x82\x14\x05\x67\x5e\x6c\xe9\x3f\xff\x8a\x94\xfc\xce\xc8\x60\xe0\x98\xe2\xa2\xa9\x23\x19\x0e\x39\xd8\x63\x35\x05\xeb\x71\x5d\x12\x4a\xf0\xfb\x9c\x91\xf8\xe4\xd2\xca\xa4\xb2\x1c\x04\xa7\xd1\x51\xa4\x27\x8c\x08\xf1\x39\x08\x2b\xc4\xce\x9b\x2d\x1e\x17\xa5\x44\x3c\x49\xb7\x1a\xe2\xce\x63\xc5\xf0\x32\xab\x11\xbf\x25\x92\xec\xd1\x8a\x85\xd9\xea\x22\xbe\x32\xd0\xa7\xff\xda\x21\x20\x0c\xb5\xf1\x67\xba\x47\x2b\xf8\xa2\x0b\x07\x78\xd6\x0b\xcc\xa5\xb5\x56\x26\xe1\x28\x14\xf6\x03\x7a\x2c\x12\x1b\xe9\xc3\x7f\x81\x61\x27\x79\x43\xe8\xdb\x3f\x8a\xa6\xc9\x93\x53\x1f\x70\x25\xc7\xae\x01\xae\xe6\x24\xb0\x75\x59\x25\x08\x99\xcf\x31\xf8\xf4\xae\x76\xb1\xb3\x1a\xd8\x98\xd6\x97\x65\x83\x21\xc5\xe9\xa3\x34\xfa\xd0\x46\xf7\x8f\x14\x5c\x83\x0f\xe3\x62\xaa\x75\x29\x27\x5f\xda\xb4\x0c\xb7\x4e\x53\x9c\x60\x94\x1f\xd1\xab\x71\x59\xae\x77\x06\x17\xa9\x64\x9a\xf2\x58\xba\xa4\xa7\xde\x2d\x63\x1f\x50\x6f\x78\x17\x72\xf7\xac\x79\x92\xd7\x98\xa8\x53\x89\xf5\xd1\x37\x45\xd8\xfe\x37\x45\x08\x0b\xb8\x72\xbe\x2a\x15\x18\xa3\xc2\xba\xb6\x20\x08\x46\xfb\x5e\x39\xba\xb2\x93\xde\x73\xda\x70\x18\xe4\xb1\xb6\xe4\x74\x9b\x87\xd5\xb4\xa5\x64\x4a\x82\xd1\xf3\x6d\x11\xd5\xa8\x8b\x53\x7e\x0d\xf6\xd0\x5b\x8a\xb1\x04\x72\xce\xdc\x2f\xb4\x24\x1a\x10\x49\x55\x10\x76\x48\x84\xa1\x31\x79\x8f\xc8\x53\x8f\x7a\xa3\xd7\x8a\xcb\xb8\xdb\x1b\xa3\xa5\x9f\x66\x91\x3c\x86\x2a\xe9\x43\x7c\x94\x5a\xfe\x91\xb9\xb0\xef\x4d\xd7\x19\xd7\x4e\x30\x1b\x0b\x0b\x3e\xde\xb6\x53\x23\x25\x7e\x37\x4e\xf9\x2a\xec\x66\xd6\xdb\x3b\xf6\x6f\x09\x86\x9c\xf9\x32\x3c\xeb\x5b\x5c\x79\x8b\x21\xcb\x09\x4b\x2b\x8c\x7b\xc3\x38\x3f\x35\x7c\x39\xea\x6b\xa0\xb9\xc3\xb2\x33\x1c\xf9\x8a\x90\xc7\x7b\x74\x76\x63\xfc\xfd\x97\x0b\x82\x22\x92\x71\x7e\x7c\x17\xc6\xe3\xdb\x97\xd2\x56\xaf\x33\xad\xe9\xc0\x55\x65\x8b\x8d\x3e\x58\x78\x86\x32\x47\x48\xd6\x52\xde\xf5\x21\xf5\xf7\x9c\x66\xf5\x53\xd2\x38\x55\xbb\xc8\x1e\x47\xea\xa7\x8d\x57\x91\xfd\x2f\x37\xe2\x76\x94\xc6\xb6\x66\xdf\xf7\x59\xf0\xc3\xec\xd9\xa4\x19\xd7\x47\x3c\xcc\x9a\x24\xa3\xd3\x49\xdb\x62\x92\xda\x3a\x14\x1a\x49\xa0\x6c\x6d\xf7\xa6\x25\xa0\xd9\x24\x23\xd0\x6e\xd5\x93\xf3\x24\x50\x57\xe8\xe6\xaa\xe1\x18\x68\x15\x6e\xd5\x84\x8b\x96\x29\x21\x9b\x60\x9e\x72\xf8\x1c\x80\xf2\x25\xa7\x0a\xef\x0b\x38\x50\xb3\x81\x27\x00\x23\x8b\x63\x1d\x15\x46\xb3\x16\x6b\x78\x67\xb1\x93\x61\x59\x6c\x8e\x7c\x4a\x22\xf8\x19\xd7\xc1\x20\x2a\xef\xe3\x45\x6a\x06\x3e\x56\x9c\xf2\x64\x1b\x40\xa3\x59\xf6\x6c\x75\xbf\x3b\xcf\xbe\x1b\x2c\xa5\xc4\x3f\x40\xef\x3e\x10\x26\xb6\xe7\xb1\xe4\xe4\xfa\x32\x90\xcf\xf8\x41\xf3\xa2\xca\x12\x1e\x14\x3f\xe8\x85\x6b\xe4\xe7\x79\x16\x67\x6e\xad\xe1\xc9\xa3\x91\x55\xa8\x6b\x72\xd3\xc2\x6e\x9b\x9d\x37\x3b\x6e\xda\x25\xa7\x5d\xd2\xbb\x3f\xd6\xa9\x77\x26\x41\x68\xe1\xa8\x4b\x36\xb3\x6a\xd8\x9f\x94\x3c\x22\xd1\xc4\xed\x9c\xa3\x27\xe6\x6d\x8e\x22\x1d\xcf\x7c\x92\x1a\xdd\xf4\x73\xbe\xd7\xa0\xc6\x78\xb7\x22\x15\x4a\xb8\x8e\xc0\x8f\xe3\x55\xd7\x53\x07\x37\x0f\x5a\xd6\x3c\x85\xe0\x88\x5a\x01\xa9\x6a\x9b\xa2\xf2\xa7\x44\x54\xa3\x15\xaa\x97\x10\x58\x8d\x71\xf8\x9c\xb5\xf6\xbe\x6f\x0a\xa4\x37\x02\x03\xf1\xda\x95\x66\xb7\x4b\x57\x5d\xa1\x21\x40\xc2\xfa\x75\x70\xe4\x25\x50\x9c\x29\x81\xe0\xe2\x59\x4e\x66\xb8\x71\xef\xcb\x2a\x80\x58\x6a\x3b\xd7\x12\x5b\xfa\x53\x7e\xc8\x48\xa9\xed\xbb\xaa\x9b\x35\x6c\x7a\x27\x07\x34\x43\x95\xd4\x0b\xf2\x17\xdf\xf3\x4f\x69\xf1\x5f\xbe\xa9\x5f\x21\xa9\xa1\x5a\x74\xa5\xdc\x2f\x6c\x89\x30\x10\xcc\xaf\x23\x8c\x51\x33\x27\xf9\x23\x00\x47\x4b\xc6\x37\xea\xd9\x57\x00\xed\x87\x40\x02\x26\xd0\xcf\x74\x9d\xf6\x7f\xbf\xae\x03\xb6\x15\x7e\x7b\xff\x53\xde\xc1\xbf\x6b\xf7\x9f\x04\xf9\xe7\x2c\xb2\x3f\xb0\xe4\xff\xc5\xfa\xa7\xc8\xff\xff\xc4\x3f\x6b\xfc\xbf\x13\x23\xff\x83\xf8\x47\xa9\x07\xff\x7e\xb9\xfe\x57\xfa\xcf\x94\x9e\xf6\x5f\x49\x85\x61\x13\x7f\x41\xff\xa3\xe0\x83\x7f\x17\xfc\xbf\x48\xfa\x73\x45\xf8\x3d\x09\xf8\x8f\xa4\x03\xe2\xbf\x93\xe4\xee\xf6\xf3\x33\x81\x89\x7e\xa3\x81\x01\x60\xc0\x20\x10\x00\xf8\x48\xf9\xdb\xdf\xff\x04\x00\x00\xff\xff\x9a\x33\xd6\xee\x93\x11\x00\x00")

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

	info := bindataFileInfo{name: "local.zip", size: 4499, mode: os.FileMode(438), modTime: time.Unix(1587811232, 0)}
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