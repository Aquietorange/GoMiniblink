package Res

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// call_js_html reads file data from disk. It returns an error on failure.
func call_js_html() ([]byte, error) {
	return bindata_read(
		"Res\\call_js.html",
		"call_js.html",
	)
}

// control_html reads file data from disk. It returns an error on failure.
func control_html() ([]byte, error) {
	return bindata_read(
		"Res\\control.html",
		"control.html",
	)
}

// css_test_css_test_css reads file data from disk. It returns an error on failure.
func css_test_css_test_css() ([]byte, error) {
	return bindata_read(
		"Res\\css\\test-css\\test.css",
		"css/test-css/test.css",
	)
}

// css_xspreadsheet_css reads file data from disk. It returns an error on failure.
func css_xspreadsheet_css() ([]byte, error) {
	return bindata_read(
		"Res\\css\\xspreadsheet.css",
		"css/xspreadsheet.css",
	)
}

// drop_file_html reads file data from disk. It returns an error on failure.
func drop_file_html() ([]byte, error) {
	return bindata_read(
		"Res\\drop_file.html",
		"drop_file.html",
	)
}

// embed_loader_html reads file data from disk. It returns an error on failure.
func embed_loader_html() ([]byte, error) {
	return bindata_read(
		"Res\\embed_loader.html",
		"embed_loader.html",
	)
}

// events_html reads file data from disk. It returns an error on failure.
func events_html() ([]byte, error) {
	return bindata_read(
		"Res\\events.html",
		"events.html",
	)
}

// gobindata_html reads file data from disk. It returns an error on failure.
func gobindata_html() ([]byte, error) {
	return bindata_read(
		"Res\\gobindata.html",
		"gobindata.html",
	)
}

// hook_html reads file data from disk. It returns an error on failure.
func hook_html() ([]byte, error) {
	return bindata_read(
		"Res\\hook.html",
		"hook.html",
	)
}

// images_chrome_png reads file data from disk. It returns an error on failure.
func images_chrome_png() ([]byte, error) {
	return bindata_read(
		"Res\\images\\chrome.png",
		"images/chrome.png",
	)
}

// images_close_png reads file data from disk. It returns an error on failure.
func images_close_png() ([]byte, error) {
	return bindata_read(
		"Res\\images\\close.png",
		"images/close.png",
	)
}

// images_runjs_png reads file data from disk. It returns an error on failure.
func images_runjs_png() ([]byte, error) {
	return bindata_read(
		"Res\\images\\runjs.png",
		"images/runjs.png",
	)
}

// images_web_png reads file data from disk. It returns an error on failure.
func images_web_png() ([]byte, error) {
	return bindata_read(
		"Res\\images\\web.png",
		"images/web.png",
	)
}

// images_xspreadsheet_svg reads file data from disk. It returns an error on failure.
func images_xspreadsheet_svg() ([]byte, error) {
	return bindata_read(
		"Res\\images\\xspreadsheet.svg",
		"images/xspreadsheet.svg",
	)
}

// js_hook_js reads file data from disk. It returns an error on failure.
func js_hook_js() ([]byte, error) {
	return bindata_read(
		"Res\\js\\hook.js",
		"js/hook.js",
	)
}

// js_xspreadsheet_js reads file data from disk. It returns an error on failure.
func js_xspreadsheet_js() ([]byte, error) {
	return bindata_read(
		"Res\\js\\xspreadsheet.js",
		"js/xspreadsheet.js",
	)
}

// js_call_html reads file data from disk. It returns an error on failure.
func js_call_html() ([]byte, error) {
	return bindata_read(
		"Res\\js_call.html",
		"js_call.html",
	)
}

// runjs_html reads file data from disk. It returns an error on failure.
func runjs_html() ([]byte, error) {
	return bindata_read(
		"Res\\runjs.html",
		"runjs.html",
	)
}

// test_html reads file data from disk. It returns an error on failure.
func test_html() ([]byte, error) {
	return bindata_read(
		"Res\\test.html",
		"test.html",
	)
}

// transparent_html reads file data from disk. It returns an error on failure.
func transparent_html() ([]byte, error) {
	return bindata_read(
		"Res\\transparent.html",
		"transparent.html",
	)
}

// web_html reads file data from disk. It returns an error on failure.
func web_html() ([]byte, error) {
	return bindata_read(
		"Res\\web.html",
		"web.html",
	)
}

// window_html reads file data from disk. It returns an error on failure.
func window_html() ([]byte, error) {
	return bindata_read(
		"Res\\window.html",
		"window.html",
	)
}

// zipdemo_zip reads file data from disk. It returns an error on failure.
func zipdemo_zip() ([]byte, error) {
	return bindata_read(
		"Res\\zipdemo.zip",
		"zipdemo.zip",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"call_js.html":            call_js_html,
	"control.html":            control_html,
	"css/test-css/test.css":   css_test_css_test_css,
	"css/xspreadsheet.css":    css_xspreadsheet_css,
	"drop_file.html":          drop_file_html,
	"embed_loader.html":       embed_loader_html,
	"events.html":             events_html,
	"gobindata.html":          gobindata_html,
	"hook.html":               hook_html,
	"images/chrome.png":       images_chrome_png,
	"images/close.png":        images_close_png,
	"images/runjs.png":        images_runjs_png,
	"images/web.png":          images_web_png,
	"images/xspreadsheet.svg": images_xspreadsheet_svg,
	"js/hook.js":              js_hook_js,
	"js/xspreadsheet.js":      js_xspreadsheet_js,
	"js_call.html":            js_call_html,
	"runjs.html":              runjs_html,
	"test.html":               test_html,
	"transparent.html":        transparent_html,
	"web.html":                web_html,
	"window.html":             window_html,
	"zipdemo.zip":             zipdemo_zip,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"call_js.html":            &_bintree_t{call_js_html, map[string]*_bintree_t{}},
	"control.html":            &_bintree_t{control_html, map[string]*_bintree_t{}},
	"css/test-css/test.css":   &_bintree_t{css_test_css_test_css, map[string]*_bintree_t{}},
	"css/xspreadsheet.css":    &_bintree_t{css_xspreadsheet_css, map[string]*_bintree_t{}},
	"drop_file.html":          &_bintree_t{drop_file_html, map[string]*_bintree_t{}},
	"embed_loader.html":       &_bintree_t{embed_loader_html, map[string]*_bintree_t{}},
	"events.html":             &_bintree_t{events_html, map[string]*_bintree_t{}},
	"gobindata.html":          &_bintree_t{gobindata_html, map[string]*_bintree_t{}},
	"hook.html":               &_bintree_t{hook_html, map[string]*_bintree_t{}},
	"images/chrome.png":       &_bintree_t{images_chrome_png, map[string]*_bintree_t{}},
	"images/close.png":        &_bintree_t{images_close_png, map[string]*_bintree_t{}},
	"images/runjs.png":        &_bintree_t{images_runjs_png, map[string]*_bintree_t{}},
	"images/web.png":          &_bintree_t{images_web_png, map[string]*_bintree_t{}},
	"images/xspreadsheet.svg": &_bintree_t{images_xspreadsheet_svg, map[string]*_bintree_t{}},
	"js/hook.js":              &_bintree_t{js_hook_js, map[string]*_bintree_t{}},
	"js/xspreadsheet.js":      &_bintree_t{js_xspreadsheet_js, map[string]*_bintree_t{}},
	"js_call.html":            &_bintree_t{js_call_html, map[string]*_bintree_t{}},
	"runjs.html":              &_bintree_t{runjs_html, map[string]*_bintree_t{}},
	"test.html":               &_bintree_t{test_html, map[string]*_bintree_t{}},
	"transparent.html":        &_bintree_t{transparent_html, map[string]*_bintree_t{}},
	"web.html":                &_bintree_t{web_html, map[string]*_bintree_t{}},
	"window.html":             &_bintree_t{window_html, map[string]*_bintree_t{}},
	"zipdemo.zip":             &_bintree_t{zipdemo_zip, map[string]*_bintree_t{}},
}}
