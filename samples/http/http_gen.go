package http

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"time"
)

type fileSystem struct {
	files map[string]file
}

func (fs *fileSystem) Open(name string) (http.File, error) {
	name = strings.Replace(name, "//", "/", -1)
	f, ok := fs.files[name]
	if ok {
		return newHTTPFile(f, false), nil
	}
	index := strings.Replace(name+"/index.html", "//", "/", -1)
	f, ok = fs.files[index]
	if !ok {
		return nil, os.ErrNotExist
	}
	return newHTTPFile(f, true), nil
}

type file struct {
	os.FileInfo
	data []byte
}

type fileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool

	files []os.FileInfo
}

func (f *fileInfo) Name() string {
	return f.name
}

func (f *fileInfo) Size() int64 {
	return f.size
}

func (f *fileInfo) Mode() os.FileMode {
	return f.mode
}

func (f *fileInfo) ModTime() time.Time {
	return f.modTime
}

func (f *fileInfo) IsDir() bool {
	return f.isDir
}

func (f *fileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return make([]os.FileInfo, 0), nil
}

func (f *fileInfo) Sys() interface{} {
	return nil
}

func newHTTPFile(file file, isDir bool) *httpFile {
	return &httpFile{
		file:   file,
		reader: bytes.NewReader(file.data),
		isDir:  isDir,
	}
}

type httpFile struct {
	file

	reader *bytes.Reader
	isDir  bool
}

func (f *httpFile) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

func (f *httpFile) Seek(offset int64, whence int) (ret int64, err error) {
	return f.reader.Seek(offset, whence)
}

func (f *httpFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *httpFile) IsDir() bool {
	return f.isDir
}

func (f *httpFile) Readdir(count int) ([]os.FileInfo, error) {
	return make([]os.FileInfo, 0), nil
}

func (f *httpFile) Close() error {
	return nil
}

// New returns an embedded http.FileSystem
func New() http.FileSystem {
	return &fileSystem{
		files: files,
	}
}

// Lookup returns the file at the specified path
func Lookup(path string) ([]byte, error) {
	f, ok := files[path]
	if !ok {
		return nil, os.ErrNotExist
	}
	return f.data, nil
}

// MustLookup returns the file at the specified path
// and panics if the file is not found.
func MustLookup(path string) []byte {
	d, err := Lookup(path)
	if err != nil {
		panic(err)
	}
	return d
}

// Index of all files
var files = map[string]file{
	"/js/padstart.js": {
		data: file0,
		FileInfo: &fileInfo{
			name:    "padstart.js",
			size:    577,
			modTime: time.Unix(1571842233, 0),
		},
	},
	"/leftpad.js": {
		data: file1,
		FileInfo: &fileInfo{
			name:    "leftpad.js",
			size:    288,
			modTime: time.Unix(1571842233, 0),
		},
	},
}

//
// embedded files.
//

// /js/padstart.js
var file0 = []byte(`String.prototype.padStart =
    function (maxLength, fillString=' ') {
        let str = String(this);
        if (str.length >= maxLength) {
            return str;
        }

        fillString = String(fillString);
        if (fillString.length === 0) {
            fillString = ' ';
        }

        let fillLen = maxLength - str.length;
        let timesToRepeat = Math.ceil(fillLen / fillString.length);
        let truncatedStringFiller = fillString
            .repeat(timesToRepeat)
            .slice(0, fillLen);
        return truncatedStringFiller + str;
    };
`)

// /leftpad.js
var file1 = []byte(`if (!String.prototype.leftPad) {
	String.prototype.leftPad = function (length, str) {
		if (this.length >= length) {
			return this;
		}
		str = str || ' ';
		return (new Array(Math.ceil((length - this.length) / str.length) + 1).join(str)).substr(0, (length - this.length)) + this;
	};
}
`)
