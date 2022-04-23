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

var _templateSqlobjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x4f\x6f\xda\x40\x10\xc5\xef\x96\xfc\x1d\x46\x9c\x40\xaa\x36\x87\xde\x92\xaa\x4a\x50\x42\x84\x42\x0a\x89\xe9\x07\x58\xec\x07\x9d\x74\xff\x38\xbb\x43\x2a\x17\xf9\xbb\x57\xb6\x49\x45\x55\x10\x50\xf5\xe4\xb5\xf7\xbd\xfd\xbd\x9d\x19\xb3\x2d\x7d\x10\xf2\x61\xa5\x5e\x8a\x05\xab\xb7\x8f\x2a\xf7\x01\xca\xea\xb2\x44\x50\xcf\xfe\xc7\x63\xbb\xba\x4a\x93\x3d\xd2\xf8\x6a\xfc\xe2\x05\xb9\xa8\xdc\xbb\x25\xaf\xd4\x33\x56\x1c\x05\xe1\x74\xdf\x3a\x8a\xb7\xfc\x13\x41\x0d\xd9\x15\x67\xca\x87\xd0\xee\x0c\xcb\x9c\x2d\xa2\x68\x5b\xe2\x28\x28\x8a\x16\x58\x38\x51\xf7\x90\x7b\x38\x04\x2d\x28\x1e\x50\xc5\xd3\x9d\xd9\xab\x79\x5a\x23\x54\x67\x39\xbe\x96\x85\x16\x5c\xa5\x49\x9a\x5c\x5c\xd0\x66\x43\x2a\x93\xb0\xce\x45\x7d\xd1\x16\x54\xd7\x77\x4e\x58\x2a\x0a\x28\x03\x22\x9c\xc4\x5d\xcd\x5c\x2f\x0c\x54\x96\x7f\x83\xd5\x54\xd7\xea\xaf\xad\xed\x21\x69\x52\xae\x17\x86\x73\x62\x27\x08\x4b\x9d\x63\x0f\xe9\xf6\x66\x4a\x9b\x34\x21\x22\xba\x7e\xbf\x4a\xbf\xb7\xd9\x50\x1e\xa0\x05\x19\x0c\x72\xc9\x9e\x26\xef\x3e\xaa\xeb\xde\xa0\xd3\x4f\x38\xca\xa7\x43\xd9\x3f\xd3\x92\x5d\x71\x63\x4c\x7f\xd0\x5e\xf3\x1c\x00\x4d\x47\xa3\xec\x6e\x4e\x97\x7e\xb9\x8c\x10\x9a\x8c\x1f\xc7\x73\xba\x34\x6c\x59\xce\x83\xcf\xf4\x8a\x5d\xd3\xd3\xfe\x75\x33\x49\xfd\xde\xf6\x8c\xa6\x24\xd4\xae\x3f\xd0\x76\xa7\x63\x6d\xb7\xba\x97\x53\x92\x0f\xab\xd9\xf7\xfd\xe5\x99\x96\xc2\xde\x1d\xc9\x38\xac\xc6\xbf\xb3\x71\xd1\x1b\x34\x1d\x5a\x41\x66\x81\xad\x0e\xd5\x03\xaa\x11\xc3\x14\xf3\xaa\xc4\x6e\x7d\xb8\xf8\x33\x5a\x37\x4e\x3b\xd9\x9a\x5f\x66\xec\x22\xc2\x81\xd6\xbd\x79\x2e\x88\x5b\x41\x47\x6f\x0c\xfd\x1e\xba\x00\xfb\x87\x11\xed\xe3\x08\xb8\xfb\x70\xb8\x28\x2d\x78\xdd\xb9\xfe\x2b\xf8\x16\x06\x47\xc1\x45\x2b\xfa\xd7\x7a\xd7\x69\xf2\x2b\x00\x00\xff\xff\xeb\xae\x7a\x44\x49\x05\x00\x00")

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

	info := bindataFileInfo{name: "template/sqlobject.tmpl", size: 1353, mode: os.FileMode(438), modTime: time.Unix(1650747524, 0)}
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
