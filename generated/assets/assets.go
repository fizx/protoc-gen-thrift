// Code generated by vfsgen; DO NOT EDIT.

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 4, 4, 21, 50, 20, 231127992, time.UTC),
		},
		"/baseplate.thrift": &vfsgen۰CompressedFileInfo{
			name:             "baseplate.thrift",
			modTime:          time.Date(2020, 4, 4, 21, 35, 5, 339332972, time.UTC),
			uncompressedSize: 2033,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x55\xcb\x6e\x1c\x37\x10\xbc\xf3\x2b\x0a\xbe\xc4\x5e\xc8\x23\xc4\x09\x72\x90\x4f\x4e\x9c\x00\x02\x72\x92\x75\xf1\x49\xe0\x92\xbd\x4b\x46\x33\xec\x09\xd9\xa3\xd9\x85\x91\x7f\x0f\x48\x0e\xf7\x25\x29\xb9\x06\x39\x8d\xc4\x47\x57\x75\x55\x35\x37\xe8\x81\xd2\xa8\x0d\x61\xdc\x63\xad\x13\x8d\xbd\x16\xea\xc4\x45\xbf\x11\x75\xdc\xdd\x32\x22\x59\xeb\xa5\x3b\x1c\x3a\xd9\xfd\x43\x3f\x69\x18\x1e\xba\x67\x67\xd4\xf5\x6a\x85\x7b\x47\xa5\x36\x36\x1c\xa1\xc3\x09\xd0\xfb\xfc\x97\x45\xa2\xf8\xe4\x0d\x75\x4a\x7d\xe5\x29\xb6\x7f\x91\x1c\x4f\xbd\x85\x0f\x8e\xa2\x17\x6c\x22\x0f\x10\xe7\x13\x38\x10\x12\x43\x9c\x96\x0c\x3b\x70\x80\x30\xf7\x09\x46\x07\xf8\x20\x14\xb5\x11\x35\x7b\x71\x05\x8e\x76\x23\x19\x21\x5b\xb7\x36\xda\x50\xea\x94\x5a\x5d\xab\x06\xf4\x73\xe3\xf3\x65\x59\xf8\xa6\x00\x20\x73\xbf\x23\x99\x62\xc0\xec\x48\x1c\x45\x70\x44\x60\x81\x38\x3a\xb0\xf4\x09\x8e\x74\x2f\x6e\xdf\xa9\x72\x2d\xb7\x5b\x57\x8c\x23\xf3\x48\x11\x6f\x8f\xca\xe6\x5b\x14\xbb\x93\xfd\x77\x0b\xbf\x54\x7b\xa3\x60\x47\xf6\x41\x20\x5c\xaa\xd1\xce\x27\xc9\xdd\x7a\x29\xed\x59\x12\x8a\x83\x0f\x84\xfd\x89\x56\xdf\x35\x12\x07\x0e\x3e\x35\xfd\x62\x6d\xe1\x3e\x4e\x04\xbf\x79\x8d\x3b\x6e\x2f\xb7\x4a\xa1\x29\x2c\x07\xae\x1a\x81\xa5\xdc\x6f\xba\x4f\x94\xf5\x88\xda\x27\x82\x0e\xa0\x9d\xa1\x51\x3c\x87\x85\xc3\xea\xba\x7c\xd6\xcc\x3d\x7c\x7a\x58\xea\xbc\x7d\x77\xa5\xfe\x52\xc7\x60\x18\x1e\x46\x0e\x14\x24\x81\x2b\x83\xbb\x12\x22\xfc\xce\xb7\x9f\x61\x98\x1f\x3d\x55\xa7\x67\xc2\xac\x8b\x30\x18\x23\x8f\xbc\xd5\x42\x58\x93\xcc\x44\xa1\x59\x99\x8d\x2d\xbd\x0f\x6c\x29\xc3\x42\x1f\x11\x1a\xc0\x9b\x5f\xed\x96\xde\xdf\xd1\x9f\x13\x25\x79\x93\x15\xb0\x14\x3b\xe0\x2b\x4f\x4d\xb3\x6c\x72\x20\xb2\xd9\x85\x16\x28\x94\x40\xc9\xb1\xba\xf5\x91\x8c\xf4\xfb\x2b\xac\x27\x41\xd4\x25\x21\xe2\x22\x4f\x5b\x57\x80\x32\xce\x02\xf3\x0b\x07\xa1\x9d\xa8\x43\x04\x73\x0f\x4f\xde\x92\xc5\xfa\x74\xf2\x6a\x2e\x25\x4e\x26\x2b\xe0\xed\x49\x14\xb3\x5a\xb7\x9f\x5b\x13\x27\xf2\x9c\xeb\xfd\xfd\x0d\x92\x44\x1f\xb6\xf0\xf6\xe3\xd9\x65\xf1\x03\xe5\x24\x87\xcb\x02\x98\x75\x82\x89\xa4\xeb\x90\x80\x46\x36\x0e\x83\xef\x7b\x9f\xc8\x70\xb0\xe9\x1c\xe2\xc3\x0d\xfc\x4f\x3f\xb6\x1b\x0f\x43\xfa\x98\x2d\xfd\x57\x47\xbf\x50\x4a\x3e\xcf\x6a\xd4\x65\x2c\x5e\x36\x57\x3d\x33\x17\xff\x47\x73\x9b\x18\xaf\xf9\xfb\xb2\x58\xff\x64\x75\xb5\xe0\x13\xa2\x9e\xa1\x27\x71\x14\xc4\x1b\x2d\xa5\x08\x3f\x52\x80\x4e\xcb\xf0\x56\x5e\x19\xe5\xe2\xdc\xf1\x21\x5e\x5d\x2b\xd9\x8f\x64\x69\xd3\x20\x3e\x9d\x1d\xbd\xcf\x15\xab\xe7\xb9\x7b\xed\x03\xc5\x45\xb9\xfc\xca\x37\x81\x9a\x11\x30\x55\xa2\xe6\x87\x52\x87\x37\x17\xb3\xef\xfb\x4c\x84\x07\x9d\xab\xf7\xfd\x1e\xa3\x8e\x89\xaa\x1b\xcb\xab\xff\x8a\xb1\xd0\xc1\xaa\x45\xec\x84\xf6\xfa\x17\x0b\x6a\xac\xa2\x1e\x97\x87\xf5\xbe\xfc\xa6\x55\x8e\xff\xc1\x3c\x34\xa1\xbe\x35\x67\xcb\xf4\xf7\xdc\x86\xf8\xc3\xcd\x21\x12\xa9\x7e\xeb\xfa\x0f\x37\x2f\x39\x73\x61\xec\x43\x09\x40\x8e\xc8\xdf\x01\x00\x00\xff\xff\xe6\x75\x5e\x65\xf1\x07\x00\x00"),
		},
		"/gen.go": &vfsgen۰CompressedFileInfo{
			name:             "gen.go",
			modTime:          time.Date(2020, 4, 5, 4, 0, 16, 847502973, time.UTC),
			uncompressedSize: 3833,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x51\x8f\x9b\x38\x10\x7e\x2e\xbf\x62\xb4\x95\x56\xb8\xa2\x20\xb5\x55\x1f\xf6\xb4\x0f\xa7\xd3\xe9\xae\x6f\xa7\xeb\xbe\x55\x5b\xe4\x65\x07\x62\x05\x6c\x6a\xcc\x2a\x39\x8b\xff\x7e\xb2\xb1\x1d\x20\x24\x69\xd3\xbe\x84\x78\x3c\xf3\xcd\x37\x5f\x3e\x4c\x68\x69\xb1\xa5\x15\x02\x7d\xa6\xad\x42\x19\x45\xac\x69\x85\x54\x10\x47\x00\x37\x85\xe0\x0a\x77\xea\xc6\x7c\xaf\x98\xda\xf4\x4f\x69\x21\x9a\xac\x12\x35\xe5\x55\xd6\x4a\xa1\xc4\x53\x5f\x66\xad\xda\xb7\xd8\x65\xd8\xb4\x6a\xbf\xcc\x2d\xd9\x7f\xbb\x31\xb3\x78\x5b\x21\x7f\xab\x36\x92\x95\x2a\xc3\x1d\x6d\xda\x1a\xb3\x0a\x39\x4a\xaa\xf0\xd9\x47\x6e\x22\x50\xd7\x01\x2c\x36\xfc\xd5\x96\xde\x44\x24\x8a\xca\x9e\x17\xd0\x29\xc9\x78\xf5\x20\xfe\x31\xe1\xb8\x73\x6b\xe2\xae\xa0\x23\x00\x89\xaa\x97\x1c\xba\x68\x98\x97\x3c\xd8\x06\xdf\x57\xc3\x3e\x7e\x38\xf4\x60\x5c\x7d\xfc\x40\xc6\xcb\xe9\xec\x00\x7f\x39\xfd\xfd\xbb\x19\xf8\xfb\x77\x64\xbc\x9c\xce\x9e\x82\x9f\x4f\x7f\x12\xa2\x3e\xa0\x9b\x15\xb1\x9f\x67\x92\x03\xf8\xe9\x6c\xa7\xfe\xb3\xe8\x9f\x6a\x3c\xc0\x97\xb5\xa0\x76\x58\xf7\x65\xb5\x89\x2f\x0a\x6d\x2e\x55\x45\x5a\xbf\x6e\x68\xdb\x0d\x0e\x40\xeb\x8e\x96\x98\x73\xda\xe0\x30\xf8\xe6\x8c\x43\x43\xdb\x2f\x5a\xd7\x58\xaa\x7c\x34\xcf\x30\x3c\x6a\x2d\x59\xb5\x39\x04\xc8\x34\xcb\x7a\x69\x92\xe4\xd6\xb6\x3f\x2b\x81\x71\xb8\xbf\x07\xce\xc6\xe9\x03\x25\xce\xea\x08\x60\x88\x00\x44\xaf\xe0\xee\x1e\x1a\xba\xc5\xf8\x32\x2c\x89\x00\x4a\x21\x61\x9b\xc0\x8b\x29\x93\x94\x57\x68\x9a\x8c\xe8\xa2\x57\x81\xfd\xbe\x9d\x4c\xb6\x25\x8f\x70\x0f\x61\x90\xd9\xde\x0b\x71\x54\x1c\x37\xd1\x2b\x23\xd8\x9a\x4c\x4e\xee\xb9\x4e\xa7\xa8\x7e\x97\x96\x3f\xab\xd3\x69\xe0\x1f\x97\xca\x4d\xb7\xae\x95\xdb\x3c\x21\x96\xd6\x99\x73\x97\x31\x5a\xcd\x3a\x15\x9c\x66\x16\xb9\xd3\x71\x24\xe7\x30\xf3\x89\xed\xbe\x3c\x6a\x3d\xdb\x24\x36\x64\xa5\x74\x91\xab\x94\x5a\x82\x24\x50\x23\x8f\x19\x27\x5e\x1d\x96\x00\x9e\x50\x87\x8d\x32\x1c\x13\xf7\xbc\xf1\x8c\x71\xce\x4d\x7d\x70\xd1\x92\x1e\x39\x16\xe2\xea\xb1\x67\x28\xbf\x68\x6e\xc7\x7c\x7d\xf0\x2c\xd3\x3a\x1d\x06\xe3\x05\x6f\x00\x7b\xea\x60\xd7\xd1\x0a\x1f\x2c\x84\xe9\x3e\xde\x57\x47\x27\xcf\x1b\x95\xfa\x28\x81\x37\xee\x51\x15\x42\x8e\xdf\x8a\x10\x0b\x29\x46\x62\x41\x8e\xdb\x25\x90\x1e\xb7\xb5\x7e\x5d\x32\xac\x9f\x87\xc1\x81\x68\xfd\x95\x75\xb9\xe0\x98\x8b\x32\x04\x2d\x4c\xaa\x75\x41\x1b\xac\xdd\x29\x70\xc1\x14\x8c\xa7\x5a\x77\x2d\x16\x8c\xd6\xf9\xb4\x8e\x84\x46\xd9\xb2\x91\xd6\x99\x27\xe3\xc9\x09\x8e\xa2\xec\x7c\xc0\x86\x5e\x68\xdd\xe3\x84\x9a\xd5\x22\xd5\x7a\x8b\xfb\x61\xf8\x8c\x35\x16\x8a\x09\x2b\xce\x41\x49\xb7\x9b\x6b\xcd\x38\x47\x99\xcf\xc4\x9c\x4e\x68\xd3\x60\x45\xaf\x73\xa5\x86\xd6\x74\xf3\x2e\xac\x57\x34\x99\x26\x92\x64\x82\x72\x98\xe8\xf0\x5b\x64\xd3\x61\xb5\xce\xe6\x72\x4c\x6c\x37\x16\x2d\x3c\x75\xb8\xc1\x8e\x7c\x44\xa6\x3e\xbb\xda\x54\xea\xe7\xec\xb4\x66\x8f\x4b\xf7\x9c\xd5\xb0\x95\x58\xb2\x9d\xf9\x59\x7f\x81\xb3\xce\x19\x6b\x97\x80\xd8\x9a\x51\x19\x4f\xff\x42\xe5\xfc\x11\x93\x34\x3e\x52\x74\x61\x10\xf2\x9b\xa9\x3c\x61\xb0\x89\x4b\x2f\x9a\x34\x7a\xf5\x6a\x5a\x3f\x33\xe1\xfd\x91\xd1\x9c\x4a\xbb\xa5\xd1\x22\x07\x32\x8c\x97\x6b\x0d\x66\x1f\x70\xb3\x83\xcc\x9e\x6d\x1d\xca\x17\x56\xe0\x58\x66\xa8\x04\x13\xfe\x3e\xbe\x3b\x7c\x6a\xda\xda\xfc\x27\xee\x0b\xe5\xcd\x66\xd8\xc1\x52\xc4\xcf\x28\x5f\x50\xda\x66\xde\xce\x31\x33\xb5\x2b\x78\x04\x3e\x75\x7f\x23\xad\xd5\x66\x1f\x17\x6a\x07\xee\x8d\x24\xfd\x63\xbc\x12\x88\xcd\xbf\xcd\x04\x50\x4a\x21\x89\x6b\x9b\xdb\xb5\xfd\x45\x9b\xb6\x4e\x2d\x8b\x74\x86\x93\xc0\xad\x7d\x5b\x49\xff\x34\x9f\xda\xb9\xca\x29\x61\x6a\xc7\x1b\xc4\xe2\x78\x55\xac\x7d\x1a\x54\x1b\xf1\x1c\xa4\xfc\xba\xb1\x98\x79\xb1\xc1\x62\xeb\xa2\xb3\x81\x9c\x68\xf9\xca\x60\x7e\xda\xb5\xb9\x12\x90\xf8\x6d\xbc\x7b\x19\x6f\x7b\xf5\xe0\x9e\x98\xb1\x0d\x89\x5e\x85\xd8\x62\x74\xeb\xa0\xb5\xf1\xa7\xdd\x12\x98\xc1\xfa\x93\x4b\xe2\x37\x42\xe6\x27\xc2\xbc\x57\x70\x9e\xe8\x15\xf1\xda\xf8\xe3\x42\xeb\xec\x48\x0c\x6b\x24\x27\xd8\xf2\xe4\x72\x5a\xc4\x9d\x75\xc3\x09\x93\x10\x50\xe9\xbf\x94\x6f\xfd\x6b\x56\x20\x76\xbb\xe6\x3d\x2f\x81\x9d\xf8\x0e\x46\xe4\x24\x70\xb4\x4f\xeb\xe0\xe2\xff\x03\x00\x00\xff\xff\x4b\x30\xd6\x65\xf9\x0e\x00\x00"),
		},
		"/gen.thrift": &vfsgen۰CompressedFileInfo{
			name:             "gen.thrift",
			modTime:          time.Date(2020, 4, 5, 3, 43, 23, 451133674, time.UTC),
			uncompressedSize: 677,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x3f\x6f\xc2\x30\x10\xc5\x77\x7f\x8a\x13\x91\x2a\x2a\xb5\x89\x3a\x42\xa7\x56\x74\xe8\x52\x06\x98\x1b\x99\xe4\x42\x2c\x1c\x9b\xfa\x0f\x2a\x3a\xf9\xbb\x57\x49\x00\x9b\xc2\x10\xe5\x74\xbe\xf7\x7b\xef\xae\x28\xe0\xcd\x3b\xfd\xbc\x45\x85\x86\x3b\xac\xa1\x31\xba\x03\x22\xc5\x3b\x0c\x01\xf2\x3c\x87\xc5\x12\xbe\x96\x6b\xf8\x58\x7c\xae\x19\x13\xaa\x92\xbe\x46\x98\x6c\xb8\xc5\xbd\xe4\x0e\x73\xd7\x1a\xd1\xb8\x09\x63\x44\x99\x45\x73\x10\x15\x86\xc0\x00\x4e\x75\x84\xe1\xaf\x43\x55\x5b\x88\xd2\xf7\x73\xb5\x3a\xcf\x32\x00\x80\x8d\xd6\x12\x84\x2d\x5b\xe4\xd2\xb5\xc7\xe9\xe3\xeb\xd0\x26\xca\x3a\x74\xad\xae\x07\xfc\xd8\xf9\x1e\x67\xca\xaa\xc5\x6a\x97\xf4\xb5\x77\x7b\xef\xd6\xc7\x7d\x6f\x7c\x49\x30\x7d\x99\x03\x91\x50\xf1\xc9\xe0\x8f\x47\xeb\x4e\x0e\xbd\xb2\xb8\x43\x24\x2a\x12\xe7\xc0\x88\x8a\xb8\x29\x1b\x72\x59\xcb\xb7\x38\x42\x59\x2f\x18\xaf\x52\xba\x7f\x09\xe8\xe2\x93\x35\x02\x65\x5c\xa5\x6f\x29\xdf\x6d\xd0\x84\xd0\x87\x7c\xb8\x0f\x48\x72\x5e\xeb\x89\x32\xad\x50\x37\x36\x21\xce\x66\xb3\xf9\x45\x49\xb4\xc3\xe3\x40\x1a\xfe\x2b\x94\x58\x39\xa1\x55\x42\x4c\x00\xfd\xc7\x6e\xa9\xa8\x7c\x77\x4b\x8c\x11\x0e\x5c\x7a\xbc\xda\x49\x28\x85\xa6\x1c\x05\x4f\xd1\x2a\x1d\x1c\x0f\x16\xdd\x87\x6b\x27\xf7\xfc\x0b\x00\x00\xff\xff\x86\xc1\xd6\x8a\xa5\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/baseplate.thrift"].(os.FileInfo),
		fs["/gen.go"].(os.FileInfo),
		fs["/gen.thrift"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}