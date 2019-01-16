// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/home.tpl (1.191kB)

package web

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesHomeTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\xcd\x4e\xe4\x38\x10\x3e\x2f\x4f\x51\xf8\xc0\x09\x77\xba\x57\x82\x85\xee\x24\x17\x16\x69\x91\x58\x0d\x07\x34\xd2\x08\x71\x30\x49\x25\x2e\x70\x6c\x8f\x5d\x1d\xe8\x41\xbc\xfb\x28\x7f\xd3\xdd\xc0\x68\x72\x71\x95\xeb\xef\xab\xf2\x57\x49\x0f\xff\xfd\x72\x71\xfb\xed\xe6\x12\x34\x37\x26\x3f\x48\xbb\x03\x8c\xb2\x75\x26\xd0\x8a\xfc\x00\x20\xd5\xa8\xca\x4e\x00\x48\x1b\x64\x05\x85\x56\x21\x22\x67\x62\xcd\x95\x3c\x13\xa3\x89\x89\x0d\xe6\xb7\x81\xec\xf7\x35\x72\x9a\x0c\xfa\x4e\x98\x55\x0d\x66\xa2\x25\x7c\xf6\x2e\xb0\x80\xc2\x59\x46\xcb\x99\x78\xa6\x92\x75\x56\x62\x4b\x05\xca\x5e\x39\x06\xb2\xc4\xa4\x8c\x8c\x85\x32\x98\x2d\x66\x73\xf1\x31\x55\x89\xb1\x08\xe4\x99\x9c\xdd\xc9\x76\x43\xc6\xb1\x0a\xa4\x9e\x40\xc2\x04\xe7\x93\x68\xb5\x66\xed\xc2\xa7\x81\x93\xb7\x21\xfb\x04\x3a\x60\x95\x89\x24\xb2\x62\x2a\x92\x22\xc6\x24\xa0\x2a\xd5\x83\xc1\x59\x43\x76\x56\xc4\x28\x20\xa0\xc9\x44\xe4\x8d\xc1\xa8\x71\x5b\xed\x37\xf1\x95\xb3\x2c\xd5\x33\x46\xd7\xfc\x39\x47\x7f\x03\xbc\xf1\x98\x09\xc6\x17\xee\x32\x8c\x36\x80\x07\x57\x6e\xe0\x75\x54\x00\xbc\x2a\x4b\xb2\xb5\x64\xe7\x97\xf0\xcf\xdc\xbf\xac\x46\xd3\xdb\x78\xea\x05\xe4\x10\xbd\xb2\x90\x83\x3a\xde\x57\x97\xda\xb5\x18\xde\x5f\xb6\x14\x89\xb1\xdc\x29\x02\x50\x38\xe3\x02\x2c\x81\xac\xc6\x40\xbc\xda\x31\xf5\xad\x45\xfa\x81\x4b\x58\x9c\x6e\xeb\x77\x5f\xa3\x42\x4d\x56\x1a\xac\x78\x09\x8b\x0f\xe0\xd2\xa4\xef\x74\xec\xfa\x50\x4a\xb8\x46\xf8\xef\xf6\xff\xeb\x13\x88\x9a\x9a\x63\xa8\x5c\x80\xab\xcb\x53\x79\x06\x71\xed\x3b\x06\x81\xab\x46\x07\x34\xd8\xa0\xe5\x08\x52\x6e\xe3\xef\xa8\x02\xc3\x70\x75\x09\xe7\xf7\xd3\xbc\xd2\x81\x30\x10\x43\xb1\x7d\x92\xc7\x98\x74\xa4\x3f\x89\x9a\xda\xfe\x39\x1e\xa3\xc8\xd3\x64\x70\x9d\xf2\xdd\xa1\x2d\xa9\xba\xef\x0a\x74\x2b\x91\x0c\x3b\xd1\x89\xdd\x1b\x8c\x5e\x25\xb5\x50\x18\x15\x63\x26\x3a\x56\x29\xb2\x18\x44\xef\xf5\xde\xec\x55\x8d\xb2\xcb\xd1\x3b\x6c\x87\x94\xea\x45\x9e\x52\x53\xef\x23\xa4\xa6\x4e\x78\x64\xf2\xcc\xdb\x5a\xc0\xb0\x32\xe2\x64\xee\x5f\x44\x92\x1f\xd9\x87\xe8\x57\x13\xd7\x3b\x74\x8b\x5f\x1d\x27\x25\xb5\x23\x84\x3d\xb9\x72\x8e\x31\xe4\x9f\x60\xdb\x85\xfe\xd7\x60\xde\xa6\x56\x13\x9f\x45\xde\xbe\xbe\xc2\xec\x2b\x86\x48\xce\xc2\xdb\x5b\x9a\xa8\x1c\xe4\x14\x71\xe1\xfc\x26\x50\xad\x19\x8e\x0a\xe7\x37\x2b\xf8\x7b\xbe\x38\x87\xed\x8a\xed\xc3\x1b\xc4\x09\x52\x3f\xe0\x61\xac\x69\x32\xfc\x96\x7e\x06\x00\x00\xff\xff\x2a\x0f\xf2\x1b\xa7\x04\x00\x00")

func templatesHomeTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHomeTpl,
		"templates/home.tpl",
	)
}

func templatesHomeTpl() (*asset, error) {
	bytes, err := templatesHomeTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/home.tpl", size: 1191, mode: os.FileMode(0644), modTime: time.Unix(1547205467, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0x19, 0xed, 0x8d, 0xde, 0x7a, 0xb7, 0x39, 0x9, 0xad, 0x79, 0x58, 0x3, 0xb9, 0x7, 0x57, 0x0, 0xbc, 0x30, 0xf2, 0x48, 0x12, 0x27, 0xb3, 0xe2, 0xd1, 0x3f, 0xf4, 0x8d, 0x91, 0x0, 0xba}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"templates/home.tpl": templatesHomeTpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"templates": &bintree{nil, map[string]*bintree{
		"home.tpl": &bintree{templatesHomeTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}