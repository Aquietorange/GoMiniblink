package GoMiniblink

import (
	"io/ioutil"
	url2 "net/url"
	"os"
	"strings"
)

type LoadResource interface {
	Domain() string
	ByUri(uri *url2.URL) []byte
}

type FileLoader struct {
	domain string
	dir    string
}

func (_this *FileLoader) Init(dir, domain string) *FileLoader {
	_this.dir = strings.TrimRight(dir, string(os.PathSeparator))
	_this.domain = strings.ToLower(strings.TrimRight(domain, "/"))
	return _this
}

func (_this *FileLoader) Domain() string {
	return _this.domain
}

func (_this *FileLoader) ByUri(uri *url2.URL) []byte {
	path := strings.Join([]string{_this.dir, uri.Path}, "")
	path = strings.ReplaceAll(path, "/", string(os.PathSeparator))
	if data, err := ioutil.ReadFile(path); err == nil {
		return data
	}
	return nil
}
