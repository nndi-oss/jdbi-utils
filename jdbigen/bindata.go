// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// template/class.tmpl
// template/sqlobject.tmpl
// template/usehandle.tmpl
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

var _templateClassTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\xc1\x8a\x83\x30\x18\x04\xe0\x7b\x20\xef\x30\x2f\xb0\xe6\xbe\x7b\x59\x58\xb7\xde\xea\x41\x5f\xe0\x57\x7f\x4a\x68\x8c\x92\xfc\x16\x24\xe4\xdd\x8b\xda\xda\xdb\xcc\x30\x7c\x5a\xfd\xba\x69\xec\xa6\x7b\x51\x92\x90\x56\xf3\xd2\x39\xdb\xa3\x77\x14\x23\x52\x42\xd1\x48\x58\x7a\x29\xae\x34\x32\x72\xfe\xf7\x62\x65\x45\xd2\x0a\x00\x52\xfa\x42\x20\x7f\xe3\xf3\x76\xb1\xec\x86\x88\x9c\x8f\xc3\x1c\xec\x83\x84\x77\xa8\x5d\xe7\x4d\xd8\xf3\x4b\xfb\x01\x8c\xd9\x87\xbf\xc9\x2d\xa3\x7f\xef\x1f\x9d\xfd\x70\x76\x63\xd0\xd6\x65\xfd\x8d\x8a\x3d\x87\x8d\xad\x58\x84\x43\x04\xf9\x01\xcd\x91\xb5\xca\x5a\x3d\x03\x00\x00\xff\xff\x4c\x9e\x5a\xfb\xd5\x00\x00\x00")

func templateClassTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateClassTmpl,
		"template/class.tmpl",
	)
}

func templateClassTmpl() (*asset, error) {
	bytes, err := templateClassTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/class.tmpl", size: 213, mode: os.FileMode(438), modTime: time.Unix(1650745597, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateSqlobjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x4f\x6f\xda\x40\x10\xc5\xef\x96\xfc\x1d\x46\x9c\x88\x54\x6d\x0e\xbd\x25\x55\x95\xa0\x84\x08\x85\x14\x12\xd3\x0f\xb0\xd8\x0f\x3a\xe9\xfe\x71\x76\x87\x54\x2e\xf2\x77\xaf\x6c\x13\x89\xaa\x20\xb0\xd4\x93\xd7\xde\xf7\xf6\xf7\x76\x66\xcc\xb6\xf4\x41\xc8\x87\xb5\x7a\x2d\x96\xac\xde\x3f\xab\xdc\x07\x28\xab\xcb\x12\x41\xbd\xf8\x5f\x4f\xed\xea\x3a\x4d\x0e\x48\xe3\x9b\xf1\xcb\x57\xe4\xa2\x72\xef\x56\xbc\x56\x2f\x58\x73\x14\x84\xf3\x7d\x9b\x28\xde\xf2\x6f\x04\x35\x62\x57\xf4\x94\x8f\xa0\x5d\x0f\xcb\x82\x2d\xa2\x68\x5b\xe2\x24\x28\x8a\x16\x58\x38\x51\x0f\x90\x07\x38\x04\x2d\x28\x1e\x51\xc5\xf3\x9d\xd9\x9b\x79\xde\x20\x54\xbd\x1c\xdf\xcb\x42\x0b\xae\xd3\x24\x4d\x2e\x2f\x69\xbb\x25\x95\x49\xd8\xe4\xa2\xbe\x69\x0b\xaa\xeb\x7b\x27\x2c\x15\x05\x94\x01\x11\x4e\xe2\xbe\x66\xa1\x97\x06\x2a\xcb\x7f\xc0\x6a\xaa\x6b\xf5\xcf\xd6\xee\x90\x34\x29\x37\x4b\xc3\x39\xb1\x13\x84\x95\xce\x71\x80\x74\x77\x3b\xa3\x6d\x9a\x10\x11\xdd\x7c\x5c\x65\x38\xd8\x6e\x29\x0f\xd0\x82\x0c\x06\xb9\x64\xcf\xd3\x0f\x1f\xd5\xf5\xe0\xa2\xd3\x4f\x39\xca\x97\x63\xd9\xbf\xd2\x8a\x5d\x71\x6b\xcc\xf0\xa2\xbd\x66\x1f\x00\xcd\xc6\xe3\xec\x7e\x41\x57\x7e\xb5\x8a\x10\x9a\x4e\x9e\x26\x0b\xba\x32\x6c\x59\xfa\xc1\xe7\x7a\xcd\xae\xe9\xe9\xf0\xa6\x99\xa4\xe1\x60\x77\x46\x53\x12\x6a\xd7\x9f\x68\xb7\xd3\xb1\x76\x5b\xdd\xcb\x39\xc9\x47\xd5\xfc\xe7\xe1\xf2\xcc\x4a\x61\xef\x4e\x64\x1c\x55\x93\x5d\xb6\xa6\x37\x6b\xc8\x3c\xb0\xd5\xa1\x7a\x44\x35\x66\x98\x62\x51\x95\xd8\xaf\x0c\x17\x7f\x87\xea\x06\x69\x2f\x55\xf3\xb3\x4c\x5c\x44\x38\xd2\xb4\x77\xcf\x05\x71\x2b\xe8\xb8\x8d\x61\x38\xc0\xe0\xe2\xf8\x18\xa2\x7d\x9c\x00\x77\x1f\x8e\x97\xa3\x05\x6f\x3a\xd7\x7f\x05\xdf\xc1\xe0\x24\xb8\x68\x45\xfd\x2b\x5d\xa7\xc9\x9f\x00\x00\x00\xff\xff\x9f\x3b\xc7\x3a\x3d\x05\x00\x00")

func templateSqlobjectTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateSqlobjectTmpl,
		"template/sqlobject.tmpl",
	)
}

func templateSqlobjectTmpl() (*asset, error) {
	bytes, err := templateSqlobjectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/sqlobject.tmpl", size: 1341, mode: os.FileMode(438), modTime: time.Unix(1650748253, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateUsehandleTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x95\xdd\x6e\xda\x4c\x10\x86\xcf\x91\xb8\x87\x91\x8f\x8c\x44\x36\x9f\xf4\x9d\x25\x6d\xa4\x44\x49\x54\xda\xb4\x10\x41\x2f\x60\xf1\x0e\x66\xe8\x7a\x6d\xad\xc7\xb4\x96\xe5\x7b\xaf\xbc\xc6\x94\x1f\x1b\x12\x35\xad\x54\x9f\x78\xd9\x9f\x77\xde\x79\x66\x16\x53\x94\xc4\x96\x21\xb6\xa1\x58\xa9\x39\x89\xf5\xff\x22\x88\x2d\x8a\x8f\x6a\x4e\xd7\xfd\x5e\xbf\xb7\xd9\xb0\x92\x6b\x29\x32\x26\x2d\x9e\x28\xe5\xeb\x96\xf9\x71\xc2\x14\x1b\xa9\xdd\xa9\xcb\x4b\x28\x0a\x10\x53\xb6\x59\xc0\xe2\x8b\x8c\x10\xca\x12\x2c\x26\x16\x53\x34\x9c\xee\xae\xce\xe4\x5c\xa3\x98\x06\x4b\x8c\x24\x94\xa5\x38\x5a\xda\x1c\xef\xf7\x92\x6c\xae\x29\x80\x40\xcb\x34\x6d\xd1\xbf\xbf\x1d\x43\xd1\xef\x01\x00\x24\x96\xd6\x92\x11\x16\x64\xa4\x86\x2a\x19\x58\x35\x19\xb9\xf5\x5a\xa9\x55\xc3\xdf\x6e\x1f\x34\x72\xd5\xc3\x4b\x4a\x1d\x23\x78\xdf\x68\x55\xd3\xe5\x81\x64\x85\xe7\xdd\xb1\xee\x83\x61\xe2\xfc\xa6\x32\xa4\x6e\xb5\xf6\xf7\xa4\x2d\x72\x66\x8d\x53\x15\xdf\x89\x97\x1f\xa4\x51\x1a\xfd\xa5\x7b\xc1\xc5\x0d\xd4\x23\x11\x58\x94\x8c\xcf\x19\xda\xdc\xf7\x8a\x02\xea\xdf\x53\xd4\x18\xf0\xf4\xf9\xa9\x89\x09\x65\xe9\x0d\x7e\xa9\x57\x8f\x88\x64\x32\x8b\xef\x50\x1a\xff\xd8\x9b\x70\x40\x0f\x4f\x68\x4a\xd9\x1f\x0c\x7e\x2b\xcd\x89\x0c\xc9\x48\x46\xe5\x93\x61\xd0\x14\x11\x0f\xa1\x1a\xc6\x8b\x45\x8a\xfc\x87\x21\xc0\xf8\xf1\x71\xfa\x30\x83\xab\x3a\x1a\x3c\x8d\x3e\x8f\x66\x70\xe5\x6c\x1c\x02\x72\x29\xcf\xc9\x28\xdf\xab\xd7\x87\xb5\xdd\xee\x6d\xb5\xa8\x37\x6c\x72\x69\xd9\x78\x0a\x7a\x4d\xaa\x15\xfd\x0b\xf0\x37\x97\xed\x4c\x09\xee\xf2\x91\xaa\x62\x87\xc8\x13\x4b\x91\xb4\xf9\x27\xcc\x1f\x09\xb5\x9a\xe5\x09\xee\xa2\x22\xf5\xe6\xc5\xb8\xcb\x27\xdf\xce\x74\xa5\x23\xf9\xdf\xb0\x8a\xfe\x8a\x7e\x3d\x81\x4e\x54\x59\x8f\x0d\x76\x93\x5b\xc7\xa4\x80\x4c\x8a\x96\x3b\xa5\x01\xdd\x6b\x8f\x48\x37\x8a\x62\xdf\xc2\x06\xdd\x1e\xa0\xaf\x89\x92\x8c\xbe\xd7\xd6\x75\x45\x71\x01\x56\x9a\x70\x5b\x0e\xe1\x0a\x94\xba\x7f\xbc\x8e\xde\xab\x8c\x6f\x1c\x7b\xc3\x8d\x5b\x11\x22\xef\xcc\xfb\x83\x8e\x58\x68\x54\xbb\x34\xfe\xc0\x20\x63\xf4\x1b\x72\x8e\xde\x49\x8c\x59\x9d\xd6\x5f\xc7\xb8\x6d\xb4\x7a\xe2\x7c\xa3\xfd\xe3\x9c\x15\x6a\xac\x39\xbf\xfa\x22\xbf\x2d\xef\x7b\x67\xe4\x65\xbc\xb7\x97\xbb\xe5\x83\x13\x22\x8f\x54\x2b\xb9\xb3\x74\xca\x7e\xef\x67\x00\x00\x00\xff\xff\x7c\x68\xc1\xe9\xb4\x08\x00\x00")

func templateUsehandleTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateUsehandleTmpl,
		"template/usehandle.tmpl",
	)
}

func templateUsehandleTmpl() (*asset, error) {
	bytes, err := templateUsehandleTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/usehandle.tmpl", size: 2228, mode: os.FileMode(438), modTime: time.Unix(1650747494, 0)}
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
	"template/class.tmpl":     templateClassTmpl,
	"template/sqlobject.tmpl": templateSqlobjectTmpl,
	"template/usehandle.tmpl": templateUsehandleTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"class.tmpl":     &bintree{templateClassTmpl, map[string]*bintree{}},
		"sqlobject.tmpl": &bintree{templateSqlobjectTmpl, map[string]*bintree{}},
		"usehandle.tmpl": &bintree{templateUsehandleTmpl, map[string]*bintree{}},
	}},
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
