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

var _localZip = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x96\x7d\x3c\xd3\x6b\x1f\xc7\x7f\xec\xc1\x43\xc2\x9a\x30\x59\xa4\xb1\x8e\x87\x51\xb6\xa5\xbb\x19\x46\x87\x62\xc8\x53\xbd\x2a\x93\xc4\x88\xb6\x6a\x24\x3b\x9d\x75\x10\x1d\xa2\x16\x62\x2d\xee\x75\x27\x59\x93\x22\x65\x94\x87\xa4\x3c\xc6\x09\x5b\xc5\x3c\x46\xc8\xce\xf4\xa0\x9d\x1e\x0c\xf7\x4b\xf7\xe9\xb4\xbb\xee\xfb\xbc\xba\xfe\xb9\x7e\xdf\xd7\xeb\x73\x7d\xde\xbf\x7f\x3e\xd7\xe7\xf2\xf5\x04\x81\xf5\x80\xc5\x25\xd7\xb3\xf4\x01\x94\x16\x14\x00\x80\x18\x7a\x58\x68\x8c\xed\x17\x8d\xe2\x2b\x8d\xd6\x5f\x1a\x66\x64\x88\xab\xb7\x92\x52\xca\xfa\x6f\xa5\xe1\x57\x4a\x2f\xd7\x10\xf2\x26\x7f\x7f\x17\xf7\x4d\xfe\x7f\x9d\x52\x07\x9e\x26\x59\xfa\x1c\x4a\x19\x1e\x3a\x02\x02\x00\x53\x28\x00\xa0\xfe\xe6\x54\x44\x54\x4c\x78\x2c\x9d\x1e\xc3\xc0\xec\xa7\xc7\x65\x91\xe9\x70\x77\xc3\x66\xf4\x41\xbf\x35\x36\xcf\x37\x59\x3e\xff\x65\xf2\xa6\x40\x24\xb8\xc1\x76\xd6\xcc\x5b\x59\x22\x70\x34\x58\x85\xf2\x8e\x79\x3c\x7a\xe4\x01\x93\x8a\xab\x68\xc2\xe5\x65\x41\xd1\xde\xd1\x27\xb1\x63\x99\x28\x3d\x5f\xb2\x73\x80\x83\x87\x4d\x5a\x39\xf2\x8d\xc6\xc7\xe5\x0e\xb2\xdf\x53\xfd\x89\x2e\xe6\xf3\x77\x2a\x34\x3c\xd7\x58\x75\xcc\x47\x4c\x24\x9c\x9b\xf8\xe0\x34\xfb\x6c\xf2\xae\x75\x03\x03\x2c\x05\x07\xc1\xa5\xf0\xab\x31\x26\xab\xa8\x77\x77\xf0\x2a\xe3\x1a\xc2\x86\xeb\x87\x0f\x3e\x7a\xd9\xc3\x5a\x48\xe4\x25\x99\xbc\x83\xbf\xb1\x77\x00\xb3\x20\x44\x88\x93\x5a\xc8\x5b\x69\xcd\xcf\x88\x8e\xec\xe1\xcd\xaf\xce\x67\x9c\xba\xc3\xe9\x4a\xb8\x16\x11\x3b\x18\xf5\x08\xb2\xd0\x36\x49\x40\x43\xd5\x3f\x18\x29\xb4\x8f\xea\xc4\x36\x14\xee\x46\x28\x74\x15\xcb\x09\xe0\x4e\xf4\x90\x8e\x33\x48\xa1\xa6\x50\x25\xe0\x39\xfd\xeb\xf9\x32\x03\xb5\x5b\x0b\x1e\x6f\x82\xbd\x7f\x7f\xd2\xf0\xae\xc1\x09\x4e\x84\xfe\x96\x2c\x4d\xfe\x58\x97\x5c\xa5\x30\xed\xf2\xb9\xa4\xd6\x0d\xcf\xf1\xd0\x46\x51\x0f\xac\x9e\xcd\x9a\x73\xd7\xd0\x09\x93\xa3\x67\xf9\x7f\xb8\xb6\xb2\x9c\x1e\xab\x9d\xb1\x70\x1b\x5a\x2e\xde\x8a\xb5\x9e\x08\x78\x6b\xf4\x3e\xc8\xae\x31\x0c\x8b\x76\xe4\x56\x80\xed\x1d\x31\x96\xa2\x36\xff\x25\x43\xf9\x9b\x27\x3c\x4f\xdf\xe0\x59\xbf\x28\x22\xd5\x51\xb5\xe1\xa9\x05\xff\xf8\x90\xd8\x8b\x2f\x36\xa8\x2a\x63\xf0\x8a\x4e\x1d\xdb\x1b\xb1\xd9\x43\xaa\x59\x68\x11\x6f\x45\x37\x34\x2e\x6c\xe1\xbb\xcc\xe4\x76\xf8\x4d\x4c\x68\xd6\x18\xf3\xb6\xb6\xb7\x32\xb6\x62\x09\x27\x37\xfe\x30\x9b\xa2\x3a\xa3\xe9\xe5\x8b\xf8\x50\x9d\x74\x5a\xdd\xc2\x1c\x89\x77\xc3\x5b\xf6\xc2\xcd\x73\x33\x35\x6d\x77\x56\x92\xc9\x8d\x78\xcf\xc3\xa2\x15\x24\xfe\xcd\x6b\xeb\xda\x66\xae\xff\x94\x81\x24\x37\xfe\xab\xbf\xdc\x2a\x96\xd3\x95\xb5\x77\x83\x1b\x13\x32\x75\x80\x99\x92\x16\x8a\x41\x74\x48\x9a\x3c\xa6\x58\xd9\x06\xc1\xae\xa6\x61\xe7\x9a\xc3\xb7\xb4\x3d\x90\xe6\xd0\x1d\x6e\x9f\xbd\x54\x1b\x5f\x3d\xc6\x79\x6b\x8d\x9e\x66\x53\xd2\x8d\x6d\x8b\x32\xf2\x10\x0c\x58\x64\xe6\xdd\xc6\xc1\x92\x2c\x2b\xeb\x66\x89\xcd\x99\x3c\x82\x30\x67\xe5\xe9\x31\xf6\xe3\xd1\xfe\x5c\x91\x61\x6a\x1c\xbd\x6a\xcc\x3c\xf4\x86\xf0\x65\xc6\xde\x74\x9d\xc2\xde\x4b\x4b\x7d\x53\xb9\x17\x9f\x31\x8d\xa3\xde\x4d\xee\x78\xb8\x34\x7c\xad\xf6\x28\x16\x97\xbf\xab\x66\xa3\x7d\x2c\x7e\x49\xb6\x79\xd8\xeb\x9b\x35\xe1\x18\x8c\xe7\x0e\x8c\x1d\x5a\xbb\x1b\x9e\xab\xdb\x37\x1b\x12\xd5\xd6\xca\x89\x1f\xc3\x57\x8f\xae\xbd\xa9\x65\x0e\x23\x89\xda\x6c\x1b\x4c\xda\x47\xfc\xac\xaa\x36\xc2\x1e\x92\xdf\x0f\xe0\xc4\x47\xc7\x85\x2b\xa6\x5e\x27\xed\x60\x55\x0d\x94\x94\xf4\x65\x4d\x1d\x0e\x88\xce\xa8\xa9\xa9\x34\xec\x24\xa0\xe6\x9b\x82\xe3\x8e\x81\xf4\xeb\x1d\x0d\x3a\x1f\x0b\x7f\x2a\x2c\xf3\x2a\xc9\x8c\x12\xc7\xf0\xfb\x8c\x12\x36\x74\xf9\xd6\xd3\xf9\x57\x52\xbd\xdd\xb1\x95\x41\xdb\x92\xca\x2d\x35\xf4\x83\xc4\xa2\x2d\x65\xaf\x08\x31\x87\x2b\xf1\x79\x01\x54\xf3\xaa\x9f\x6d\x6f\x9b\x63\x1a\xe9\xf1\xf7\x9e\x6f\x7c\x5d\x24\x54\x0c\x3e\x0b\x92\x52\x9e\x72\x87\x4a\xb7\xe0\x4c\xf8\xe8\xf5\x66\xd2\x4c\x77\xbe\xcd\x39\xc4\x1e\xd1\xa4\x9b\x5d\xbd\x7e\x67\xbf\x46\x44\x34\x6a\x7b\x41\x2f\x93\x38\x2e\x3c\xcf\x9a\x3d\xa3\x3d\x30\x2d\x4e\xea\xeb\xa7\x71\x65\x83\xa8\xae\xa2\xba\x9e\xfc\xdc\xe4\xf7\xe7\x85\xc9\x1f\x23\x1f\x9d\xeb\x7b\x41\x7d\x73\x79\x1f\xc1\xc5\x29\x34\x3e\xff\xbd\x2c\x9a\x45\xdd\x4c\x7e\x72\x6b\xed\xf6\x7f\x06\x88\xeb\x1b\x1f\x48\xf7\xcd\xd5\x09\xae\x65\x07\xe4\x74\x3a\x89\x7f\x61\x9d\x67\x6f\x5d\x00\x09\x28\xe4\xf5\x69\x26\x80\xac\xf5\xde\x8b\xf7\x48\xb2\xc9\x8a\xa9\xf8\x92\x48\x03\xfb\xc4\x23\x77\xc2\x49\x5e\xb8\x79\x11\x7d\xa1\xb2\xa7\x52\x1f\x4f\xa3\x64\x6e\x0b\x94\x22\x12\x20\x5f\x92\x2b\x4e\xb2\xf4\x91\xfb\x9d\xea\xea\x56\x05\x80\xf6\xef\x4e\xee\x01\x7a\x41\x60\xbb\xb7\xc4\x59\x6f\x60\xdb\x94\x93\x57\x6f\x2c\x3c\xf0\xe2\xb2\xb9\x44\x4c\x1f\x2e\xb8\x2c\x7c\x8f\x4f\x16\xac\xb2\x99\x8b\x80\x77\x2f\xcf\x49\x05\xe3\x60\x39\xa4\x0b\x28\x3d\x36\x29\x9d\x6d\xc4\x16\x90\x5a\x40\xac\xa6\xb0\xe2\x12\x05\xe9\xfa\x7e\xcd\xcb\x37\x0f\x89\x32\x71\xb9\xd4\xce\xda\x89\x7e\x46\xd3\xe5\x16\x3f\x24\xd9\x8b\xc0\xda\x9b\x7e\x95\x4d\xd8\x95\x6a\xcb\xdd\xc5\x8e\xbc\x77\x5f\xc4\xd4\x27\x5c\x44\x36\x0c\x3e\xb9\x6f\x67\xc8\x0d\xed\x7f\xc4\xbb\xdf\x59\x9d\x0c\x8a\xc3\x1a\xca\xad\x68\xc5\xe3\xf4\xc0\x1f\xe7\x8a\xdc\x50\x0d\x5e\x49\xa8\xb0\x54\x67\x87\x0e\xe3\xb4\x75\x22\xd3\x83\x2d\x9d\x17\x25\xa5\x25\x99\xe9\xdb\xac\xaa\x6b\x19\xd7\xcb\x05\xe9\x43\x73\x03\xef\xa6\x9e\x79\x2c\xb7\x5d\x36\x12\x02\x3b\x79\x20\xe6\x3e\x70\x0c\x33\xfd\xf1\x7c\x73\x9a\x1a\x2e\xa2\x8f\xe3\x1a\x9d\xa0\x5f\x89\xd2\x1d\xf5\x28\xce\x8f\xcc\x7a\x4d\xd2\xfa\x6d\xe3\x80\x31\xc5\xe6\x4e\xc1\x65\xc9\x86\x34\xa6\xb1\xbc\x9d\x52\xad\x5f\x1a\x38\x73\x48\x14\x82\x74\x85\x9c\x88\x06\x77\x40\xcf\x46\x8c\x0e\x0f\xc8\xd5\xd1\x30\x17\x34\xc9\x6f\x89\x67\xdb\xd5\x13\x76\xac\x5a\x19\x32\xa7\xc9\xad\x6d\xe2\xd7\x87\xb0\x53\x2c\xab\x0b\x6b\x3d\x97\x79\x42\x15\x36\x16\xa3\x1d\xb2\x7c\x81\x3c\xe8\x7e\x48\xa9\x5e\xe0\xec\x3a\x98\x64\x5c\x8f\x3d\xae\xe1\x63\xef\xa2\x9b\x21\x39\x3c\xbd\xea\x1a\xcb\x4f\x5c\x31\x12\xd5\xfe\xb2\x14\xb5\xa7\xb6\xb5\xfc\xa8\x83\x91\x5b\xa9\xd6\x68\x44\x8b\xa3\xb6\x7e\x37\xc4\x8e\x5e\xcd\x51\x87\x66\xac\xee\xb5\x8e\x19\xe4\x39\x63\x9b\xe4\xfb\xcc\x8e\x83\x32\xb0\x94\x32\x73\x90\xd3\xea\xa5\x2b\xb8\x61\xbb\x35\xb4\x55\x23\xfd\x79\x5b\x56\x9d\x69\xcc\x56\xdf\xc9\xac\x62\x12\xe5\x66\x10\x1e\xdf\x19\x01\xa2\xc2\x2a\x32\x54\xdb\xec\x89\x68\xdd\x49\x7a\xca\x09\x33\x9a\x99\xef\xdc\xf5\x4e\x86\xda\x93\xe1\x31\xfe\x25\x81\xde\x09\x99\x6b\x7d\x2d\x52\xbe\x44\x6e\x9a\x62\xb6\x1e\x41\x4c\xde\x53\xe6\x4c\xbb\xce\x37\x6a\xa6\x15\x6b\x66\xdf\xb9\x85\xe6\x17\x06\xb1\xaa\x0a\x74\xe6\xc7\xeb\x84\x81\x2b\x57\x1b\x1e\x56\x19\xe1\x57\x1f\x73\x23\xe1\x4a\x04\x47\xe7\x87\x42\x39\x16\xdc\x46\x48\x81\xce\x55\xaa\xc1\xf4\xa6\x78\xd7\xa5\x78\xb3\xb7\x6e\x39\x79\xf9\xe0\x29\xa9\x99\xc5\xca\x8b\x08\xb7\x91\x02\x58\xa8\xe3\xd9\x5f\xad\x7d\x61\x42\xde\xbe\xda\x99\x3a\xec\xb3\xd0\xc4\x53\x2e\x18\xf9\x7e\xb5\x58\x10\x24\xb0\xfa\x1a\x33\x58\x4c\xf1\xbe\x4a\x64\xc0\xcd\x60\x1c\xd1\xb8\x75\x7b\x9a\xee\x9a\xcc\x04\xcb\x49\x97\x43\x27\x2e\x25\x79\xbf\x11\xd9\xcd\x97\xa5\x5f\xa0\xbb\x3c\x7d\x18\xdc\x33\xf2\xe3\xd2\x9d\x67\xdc\xbb\xc1\x89\x5d\xb1\x0e\x3f\x30\x86\x33\x24\xd6\x10\x0c\x93\xd3\x51\x60\x55\xdf\x8b\xa4\x3c\x09\x14\x3e\x97\x21\xa1\x5a\xa4\x74\x5a\xb1\x74\xfd\x1f\xae\x24\xde\xc0\x3c\xad\x7f\xc3\xdd\xed\xb8\x7a\x41\xc5\xad\x61\x69\x9f\xc6\xab\xba\x2b\x00\x55\xe4\xd7\x2a\x21\x9a\xf4\xa0\x7b\x42\x73\x68\xf9\xcb\x4c\xf6\x5c\x11\x10\xb5\x4d\x7d\xc3\x28\xee\x8c\xe3\x4f\x77\x63\xb9\x29\x92\x8e\x23\xaf\xe2\xab\x06\x67\x2c\x3d\xe3\x85\x63\xb6\xbe\x9e\x2a\xaa\x7a\xc0\xff\xeb\x54\xd4\x9f\x5f\xba\x7f\xee\xff\xc9\x88\x26\x60\xfa\x69\x52\x01\x0c\x81\xdb\x21\xaf\x5f\x39\x52\xc5\x2a\x69\x2a\x7e\x2b\x4b\xa9\x62\x15\x1e\xad\x7d\x72\x71\x56\xf6\xfd\xb6\x87\x95\x7d\xbf\xce\x9e\xb2\x7b\xdc\x26\xe8\x07\x65\xf7\xcf\x34\x65\xf7\x6f\xbb\x5b\xd9\xdd\xfb\x6f\x92\xad\x4c\xd2\xcf\xba\x60\x58\xaa\x44\xfa\x4c\xfe\x42\xfa\x5f\x7d\xff\x99\xb4\xe8\x73\xfc\x3b\xdb\x5f\x99\x5a\x4e\xa3\xc8\xca\xa8\x62\x15\x15\x23\xf1\xf4\xe2\xfe\xf9\x2f\x94\xa9\xdf\xde\x55\xca\x54\x22\xf8\xfb\x6e\x2e\x65\xaa\xbf\x16\x7e\x6a\x91\x36\xd1\x1d\xfc\x89\xde\x82\x3a\x02\x25\x7e\xa2\x42\x16\x9f\x51\x00\x04\x80\x00\x5a\xaa\x00\xe0\xa7\xb6\x38\xfd\x3b\x00\x00\xff\xff\xa1\x12\x91\x1f\x74\x09\x00\x00")

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

	info := bindataFileInfo{name: "local.zip", size: 2420, mode: os.FileMode(438), modTime: time.Unix(1568103772, 0)}
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
