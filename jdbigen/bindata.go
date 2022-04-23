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

var _templateSqlobjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x4f\x6f\xda\x40\x10\xc5\xef\x96\xfc\x1d\x46\x9c\x40\xaa\x36\x87\xde\x92\xaa\x4a\x50\x42\x84\x42\x0a\x89\xe9\x07\x58\xec\x07\x9d\x74\xff\x38\xbb\x43\x2a\x17\xf9\xbb\x57\xb6\x49\x45\x55\x10\x50\xf5\xc4\xa2\x79\x6f\x7f\x6f\x67\xc6\x6c\x4b\x1f\x84\x7c\x58\xa9\x97\x62\xc1\xea\xed\xa3\xca\x7d\x80\xb2\xba\x2c\x11\xd4\xb3\xff\xf1\xd8\x9e\xae\xd2\x64\x8f\x34\xbe\x1a\xbf\x78\x41\x2e\x2a\xf7\x6e\xc9\x2b\xf5\x8c\x15\x47\x41\x38\xdd\xb7\x8e\xe2\x2d\xff\x44\x50\x43\x76\xc5\x99\xf2\x21\xb4\x3b\xc3\x32\x67\x8b\x28\xda\x96\x38\x0a\x8a\xa2\x05\x16\x4e\xd4\x3d\xe4\x1e\x0e\x41\x0b\x8a\x07\x54\xf1\x74\x67\xf6\x6a\x9e\xd6\x08\xd5\x59\x8e\xaf\x65\xa1\x05\x57\x69\x92\x26\x17\x17\xb4\xd9\x90\xca\x24\xac\x73\x51\x5f\xb4\x05\xd5\xf5\x9d\x13\x96\x8a\x02\xca\x80\x08\x27\x71\x57\x33\xd7\x0b\x03\x95\xe5\xdf\x60\x35\xd5\xb5\xfa\xab\xb4\xbd\x24\x4d\xca\xf5\xc2\x70\x4e\xec\x04\x61\xa9\x73\xec\x21\xdd\xde\x4c\x69\x93\x26\x44\x44\xd7\xef\x4f\xe9\xf7\x36\x1b\xca\x03\xb4\x20\x83\x41\x2e\xd9\xd3\xe4\xdd\x47\x75\xdd\x1b\x74\xfa\x09\x47\xf9\x74\x28\xfb\x67\x5a\xb2\x2b\x6e\x8c\xe9\x0f\xda\x67\x9e\x03\xa0\xe9\x68\x94\xdd\xcd\xe9\xd2\x2f\x97\x11\x42\x93\xf1\xe3\x78\x4e\x97\x86\x2d\xcb\x79\xf0\x99\x5e\xb1\x6b\x66\xda\xbf\x6e\x36\xa9\xdf\xdb\xde\xd1\xb4\x84\xda\xf3\x07\xda\x56\x3a\xd6\xb6\xd4\xfd\x39\x25\xf9\xb0\x9a\x7d\xdf\xdf\x9e\x69\x29\xec\xdd\x91\x8c\xc3\x6a\xfc\x3b\x1b\x17\xbd\x41\x33\xa1\x15\x64\x16\xd8\xea\x50\x3d\xa0\x1a\x31\x4c\x31\xaf\x4a\xec\xf6\x87\x8b\x3f\xa3\x75\xeb\xb4\x93\xad\xf9\x64\xc6\x2e\x22\x1c\x18\xdd\x9b\xe7\x82\xb8\x15\x74\xf4\xc6\xd0\xef\xa1\x0b\xb0\x7f\x19\xd1\xfe\x1c\x01\x1f\x85\xae\x3b\xc7\x7f\x85\xde\xc2\x40\x70\x78\x12\x2d\xb8\x68\x45\xff\xda\xeb\x3a\x4d\x7e\x05\x00\x00\xff\xff\x09\xd7\xea\xc9\x45\x05\x00\x00")

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

	info := bindataFileInfo{name: "template/sqlobject.tmpl", size: 1349, mode: os.FileMode(438), modTime: time.Unix(1650746184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateUsehandleTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\xcd\x6e\xda\x40\x10\xc7\xef\x48\xbc\xc3\xc8\x27\x23\x91\x4d\xa5\xde\x92\x36\x52\xa2\x24\x2a\x6d\x5a\x88\xa0\x0f\xb0\x78\x07\x33\x74\xbd\xb6\xd6\x63\x5a\xcb\xf2\xbb\x57\x5e\x63\xca\x87\x0d\x8d\x92\xde\xea\x8b\x97\xfd\xf8\xcf\x7f\x7e\x3b\x83\x29\x4a\x62\xcb\x10\xdb\x50\xac\xd4\x9c\xc4\xfa\xbd\x08\x62\x8b\xe2\xb3\x9a\xd3\x75\xbf\xd7\xef\x6d\x36\xac\xe4\x5a\x8a\x8c\x49\x8b\x27\x4a\xf9\xba\x65\x7e\x9c\x30\xc5\x46\x6a\x77\xea\xf2\x12\x8a\x02\xc4\x94\x6d\x16\xb0\xf8\x26\x23\x84\xb2\x04\x8b\x89\xc5\x14\x0d\xa7\xbb\xab\x33\x39\xd7\x28\xa6\xc1\x12\x23\x09\x65\x29\x8e\x96\x36\xc7\xfb\xbd\x24\x9b\x6b\x0a\x20\xd0\x32\x4d\x5b\xf4\xef\x6f\xc7\x50\xf4\x7b\x00\x00\x89\xa5\xb5\x64\x84\x05\x19\xa9\xa1\x4a\x06\x56\x4d\x46\x6e\xbd\x56\x6a\xd5\xf0\xb7\xdb\x07\x8d\x5c\xf5\xf0\x92\x52\xc7\x08\x3e\x36\x5a\xd5\x74\x79\x20\x59\xe1\xf9\x70\xac\xfb\x60\x98\x38\xbf\xa9\x0c\xa9\x5b\xad\xfd\x3d\x69\x8b\x9c\x59\xe3\x54\xc5\x4f\xe2\xe5\x27\x69\x94\x46\x7f\xe9\x5e\x70\x71\x03\xf5\x48\x04\x16\x25\xe3\x73\x86\x36\xf7\xbd\xa2\x80\xfa\xf7\x14\x35\x06\x3c\x7d\x7e\x6a\x62\x42\x59\x7a\x83\x3f\xea\xd5\x23\x22\x99\xcc\xe2\x3b\x94\xc6\x3f\xf6\x26\x1c\xd0\xc3\x13\x9a\x52\xf6\x07\x83\x57\xa5\x39\x91\x21\x19\xc9\xa8\x7c\x32\x0c\x9a\x22\xe2\x21\x54\xc3\x78\xb1\x48\x91\xff\x31\x04\x18\x3f\x3e\x4e\x1f\x66\x70\x55\x47\x83\xa7\xd1\xd7\xd1\x0c\xae\x9c\x8d\x43\x40\x2e\xe5\x39\x19\xe5\x7b\xf5\xfa\xb0\xb6\xdb\xbd\xad\x16\xf5\x86\x4d\x2e\x2d\x1b\x4f\x41\xaf\x49\xb5\xa2\xff\x0b\xfc\x4d\xb3\x9d\xb9\x82\xbb\x7c\xa4\xaa\xd8\x21\xf2\xc4\x52\x24\x6d\xfe\x05\xf3\x47\x42\xad\x66\x79\x82\xbb\xa8\x48\xbd\xf9\x65\xdc\xe5\x93\x1f\x67\xaa\xd2\x91\x7c\x37\xac\xa2\xbf\xa0\x5e\x4f\xa0\x13\x55\xd6\x63\x83\xdd\xe4\xd6\x31\x29\x20\x93\xa2\xe5\x4e\x69\x40\xf7\xda\x23\xd2\x8d\xa2\xd8\xb7\xb0\x41\xb7\x07\xe8\x7b\xa2\x24\xa3\xef\xb5\x55\x5d\x51\x5c\x80\x95\x26\xdc\x5e\x87\x70\x17\x94\xba\x7f\xbc\x8e\xda\xab\x8c\x6f\x1c\x7b\xc3\x8d\x5b\x11\x22\xef\xcc\xfb\x83\x8e\x58\x68\x54\xbb\x34\xfe\xc2\x20\x63\xf4\x1b\x72\x8e\xde\x49\x8c\x59\x9d\xd6\x7f\x8c\xaf\xc3\xa8\x50\x63\x8d\xf1\xc5\x7d\xfa\x26\x38\xb7\x7d\x7b\xef\x8c\x9c\xef\x5b\xd8\xed\xdd\x96\xef\x49\x88\x3c\x52\xad\xe4\xce\xd2\x29\xfb\xbd\xdf\x01\x00\x00\xff\xff\x6f\x07\x9f\x9b\x93\x08\x00\x00")

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

	info := bindataFileInfo{name: "template/usehandle.tmpl", size: 2195, mode: os.FileMode(438), modTime: time.Unix(1650746150, 0)}
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
