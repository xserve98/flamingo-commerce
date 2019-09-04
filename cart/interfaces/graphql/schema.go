// Package graphql Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// schema.graphql
package graphql

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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\xdd\x6e\xdb\x38\x13\xbd\x6e\x9e\x42\x41\x6e\x52\xa0\xf8\x1e\xc0\x77\x8e\xf5\xa1\x30\xb6\x71\xd3\xc4\xdd\xbd\x28\x82\xc5\x44\x1c\xdb\xdc\x50\xa4\x96\x1c\x25\x11\x16\x7d\xf7\x05\x29\xca\x26\x29\x4a\xce\x62\xb7\x17\x45\xcd\x39\x1c\xce\xcf\xe1\x19\xaa\xd4\x35\x58\xac\x54\x5d\xa3\xae\xf0\xf7\x12\x2b\xa5\x81\x90\xad\x40\x53\xf1\xd7\x45\x51\x14\x45\x05\x9a\x16\x27\x88\xb5\x5c\x3a\x03\x1b\xc0\x25\x0a\xfe\x82\x9a\xa3\x59\x14\x3f\x22\x60\x99\x40\xba\xcb\xc7\x8b\x9f\x17\x17\xf1\xa1\xc1\x59\x9c\x2d\x8a\x75\xd9\xbb\x47\x49\x9c\xba\x75\xb9\x28\x1e\x48\x73\xb9\xef\x57\x9f\xb8\x10\x5c\xee\x97\x8c\x69\x34\x26\x89\xcb\xaf\x3a\x60\xd3\xea\xea\x00\x06\x75\x82\xb9\x43\x6d\x94\xf4\x09\x4c\xc7\x7d\x0c\xd7\x02\xaf\x3e\x00\x63\x9c\xb8\x92\x20\x4a\x20\x18\x1f\x1b\x18\x2f\xfd\x96\x06\xba\x1a\x25\x3d\xa0\xc0\xca\x5a\xd3\x38\x12\xb3\x4f\x0f\x85\x92\x7b\xb3\x55\xcb\x96\x0e\xb6\x02\x95\xad\xde\x77\x97\xc6\x8d\x52\x02\xc1\x03\x21\xb5\xa7\x85\xba\xfa\x00\x4d\x23\x38\xb2\x95\x6a\x1b\x25\x57\x8a\x8d\xd3\x3c\x99\x7c\xa2\x0c\x77\xd0\x0a\x5a\xb5\x5a\xa3\xac\xba\x93\xc7\x2b\x6b\x25\x45\x20\x38\x61\x3d\x72\xb4\x1d\x2c\xde\x8f\xfd\xe7\x4a\xb5\x92\x16\xc5\x5a\x7a\xba\x34\x5a\xb1\xb6\xa2\x74\x99\x9b\xa8\x12\xc8\x92\x44\xf7\x1a\x24\x73\xfe\x83\x02\xde\x69\x5e\x61\x92\xe7\x67\xbe\xa3\x15\x68\x36\x0a\x6e\x19\xdb\xa7\x28\x38\xe2\xaa\xe7\xa4\x67\x49\xb7\x48\xe1\x9e\x22\xf1\x5d\x58\xe7\xca\x53\x86\xd6\xe9\xf3\xb3\xc7\xae\xe5\x4e\x4d\x1c\x6d\x4d\xc7\x2b\x9a\xed\xcb\xfa\xd4\x12\x73\xe0\x4d\xc3\xe5\xde\x2e\x25\xfe\x1e\x02\x53\x8f\x6d\x9f\x5c\xc5\x3f\x6b\x15\xdd\x31\x57\xf6\x33\x5d\xf1\x0e\x6a\x67\xdc\xc2\xdb\xb2\xee\x1b\x9e\x45\xf5\xc7\x6c\x70\xca\xde\x7b\x29\xb9\xa9\xac\x93\x59\x57\xf5\x46\x49\x9b\xc0\x3d\x0a\xd7\xc4\x77\xed\xf9\xa7\x1b\x82\xb2\xfc\xc6\xe9\x30\xec\xc9\xd7\x28\x48\xef\x3c\xf8\x00\xc6\x53\xe7\xc8\xfe\x59\x96\xd8\xde\x0f\x02\xad\x18\xc6\x77\xff\x55\xe9\xe7\x9d\x50\xaf\xf1\x6a\x8d\x74\x50\x2c\x5e\xab\x40\x6b\x6e\xc5\x25\x5c\x1c\xa8\xf7\x45\x55\x90\x51\xae\x32\x31\xfb\x3d\x86\x6b\x64\x5b\x5e\xe3\xa2\xb0\x7f\xf7\x34\xc1\x44\x1c\xaf\x9f\xf1\xa4\x2a\x1f\xe3\x63\x63\x8d\xfd\x05\x3b\x4b\x67\x0f\x78\xec\x15\x28\x80\x04\x75\x30\x8b\xa2\x86\xe6\x87\x71\xd0\xc7\x3f\x8c\x92\xff\xbb\x87\xd7\x5b\x34\x06\xf6\x38\x5f\xc6\x21\x87\xc2\xd7\xd2\x22\x47\x51\x9d\x99\x34\xad\xc1\x9b\x64\x2a\x45\x0a\x16\xf7\x27\x1b\x4e\x78\x03\x87\x48\x38\x89\x24\x94\xc6\xb2\x65\xea\xb2\xd0\xec\x55\xf3\x1a\x99\xe3\x60\x20\x91\x47\xf3\x44\x98\x29\x6c\x98\xda\x39\xe9\x49\xb0\x53\xaa\x97\xc0\x8e\x6f\x8e\xba\x01\xbe\x77\xa3\x69\x51\xc4\x8c\x3d\xce\xac\x85\xfd\x19\x9a\x04\x3c\xa1\x70\xab\x45\x6a\xf2\xf9\x0f\xc6\xdc\x20\xe9\x5b\x9f\xdd\xcd\x4d\x20\x14\x49\x77\x8d\xd2\xf4\x55\x33\x7b\x87\xec\x1f\x37\xd7\xe6\xe7\x4b\xd0\x64\x3e\x16\xe3\xa3\x08\xfb\x79\x19\xb5\xd3\xad\xe4\xdd\x87\x5e\xc3\x57\xd4\x1b\xa1\x96\x20\xee\x71\x87\x76\xa6\x27\x94\xaa\x41\x3f\x23\x35\x02\x2a\x5c\x8d\x64\xe4\x05\x34\x07\x49\xb7\x0e\x73\x97\xc7\xf8\x28\x37\x50\x27\x06\xa3\x5a\x5d\x61\xfa\x2a\xf9\x93\xba\x60\xf6\xcf\xdf\xf8\x31\xe2\x57\x10\x2d\x8e\x30\xef\x14\x99\x41\x66\x97\xe9\xa1\x29\xdc\xb7\xb7\xcf\x82\xcb\xbd\x40\xc7\x92\xec\x30\xbc\x4c\x51\x99\xeb\xd9\x63\xb4\x7a\x3d\xe7\x66\x80\x64\x7d\x5c\x79\xc4\x16\xde\x30\xdc\xef\x7e\x4f\xdd\xd7\x5e\x8f\x3c\x2d\x5e\x80\x02\x7e\xe7\x99\xbe\xe3\xda\x90\x74\xcd\x9c\xc4\x08\xc8\x42\x62\x5e\x71\xc6\x04\x6e\x46\xa8\x10\xe3\x15\x6e\x36\x1e\x03\xa2\x25\x3f\x84\x26\x31\xa4\x11\x33\xa9\x8d\x31\x1b\x3d\x17\xf3\x89\x6b\xbe\x6e\x5f\xb8\x1c\xb3\xad\x52\x75\x03\xb2\x1b\x1d\x17\x69\x14\xa7\x31\x20\xc1\x34\xca\xd0\x51\xc5\x26\xa3\x06\x3a\x57\x21\x8d\x7b\x1e\xe8\x61\x3e\x1e\x2b\xad\xfa\x4c\xcc\x3d\x66\xe4\x28\xea\x18\x0a\x6c\x0e\x4a\xce\xb1\x03\x6b\xe0\x62\x26\xe6\x2c\x51\xfb\x8f\x32\xcf\xd3\xf3\x13\xb7\x71\x70\xfb\x08\x20\xe0\x22\x45\xde\xc5\xd6\x41\x07\xb9\x21\x2e\xf7\xab\xd6\x90\xaa\x51\x67\x3e\xe3\xfe\x9f\x81\xe4\xc3\xcd\x21\x13\xed\x9d\x49\xf3\x18\xd9\xf0\xd2\x07\xc2\xaf\xbb\x1b\xae\xe9\x90\x68\x2b\x18\xd3\x28\xdd\x7f\x32\xe9\x2e\x6f\xdc\xb4\xf5\x53\xfa\x80\x93\xd0\xf3\xd8\xd1\x30\x28\xbc\x1d\x06\x92\xb9\x49\x57\x7c\x6b\x4f\xdf\x1a\x55\x18\xe4\x62\xe2\xbf\x03\x46\x1e\x6e\xfd\xcd\x4c\x9d\x2c\x19\xdb\x2a\xbb\xe3\x7a\x34\x5d\xd6\xe5\xe5\xa7\xd3\x0c\xf8\x74\x7c\x69\x46\x73\xe5\xe3\x74\x04\xd1\x39\x25\x0a\x24\x0c\x5f\x73\xd7\xff\x81\x3f\x3b\x47\xaf\xed\x5c\xb6\xa3\xcb\xc5\xfb\xaf\x9c\x7e\x6f\x6c\x7b\xad\xd3\x6f\xd4\xbd\xc3\x6f\x50\x9e\x99\x23\x7e\x5e\xfc\x1d\x00\x00\xff\xff\xbd\xce\x68\xd9\xb8\x11\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"schema.graphql": schemaGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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