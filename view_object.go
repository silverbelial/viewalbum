package viewalbum

import (
	"fmt"
	"os"
	"sync"
)

var (
	tagVo      map[string]*ViewObject
	tagVoMutex sync.RWMutex
)

func init() {
	tagVo = make(map[string]*ViewObject)
}

//ViewObject object stores view object info
type ViewObject struct {
	Title    string
	HTMLFile string
	JsFile   *string
	CSSFile  *string
	Location string
	Parent   *ViewObject
	Link     string
}

func (vo *ViewObject) js() (string, bool) {
	if vo.JsFile == nil {
		return "", false
	}
	return *vo.JsFile, true
}

func (vo *ViewObject) css() (string, bool) {
	if vo.CSSFile == nil {
		return "", false
	}
	return *vo.CSSFile, true
}

//RegisterViewObject register child view object
func (vo *ViewObject) RegisterViewObject(tag, htmlFile, jsFile, cssFile, title, uri string) *ViewObject {
	cvo := buildVo(htmlFile, jsFile, cssFile, title, uri, nil)
	tagVoMutex.Lock()
	defer tagVoMutex.Unlock()
	tagVo[tag] = cvo
	return cvo
}

//RegisterRootViewObject register root vo into view album
func RegisterRootViewObject(tag, htmlFile, jsFile, cssFile, title, uri string) *ViewObject {
	vo := buildVo(htmlFile, jsFile, cssFile, title, uri, nil)
	tagVoMutex.Lock()
	defer tagVoMutex.Unlock()
	tagVo[tag] = vo
	return vo
}

func buildVo(htmlFile, jsFile, cssFile, title, uri string, parent *ViewObject) *ViewObject {
	vo := &ViewObject{
		Title:    title,
		HTMLFile: htmlFile,
		Link:     uri,
		Parent:   parent,
	}
	if jsFile != "" {
		vo.JsFile = &jsFile
	}
	if cssFile != "" {
		vo.CSSFile = &cssFile
	}
	vo.prepareLocation()
	return vo
}

func (vo *ViewObject) prepareLocation() {
	//current view template name
	templateName := GetTemplateName()
	if templateName != "" {
		path := fmt.Sprintf("views/%s/%s", templateName, vo.HTMLFile)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			vo.Location = fmt.Sprintf("%s/", templateName)
		}
	} else {
		path := fmt.Sprintf("views/%s", vo.HTMLFile)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			vo.Location = ""
		}
	}

	if vo.JsFile != nil {
		path := fmt.Sprintf("views/scripts/%s/%s", templateName, *vo.JsFile)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			*vo.JsFile = templateName + "/" + *vo.JsFile
		}
	}
}

//DisplayView main entrance for display registered view
func DisplayView(v Viewer, tag string) {
	tagVoMutex.RLock()
	vo, has := tagVo[tag]
	tagVoMutex.RUnlock()
	if !has {
		return
	}
	v.ServeHTMLFile(vo.HTMLFile)
	jsf, has := vo.js()
	if has {
		v.SetParam("JsFile", jsf)
	}
	cssf, has := vo.css()
	if has {
		v.SetParam("CSSFile", cssf)
	}
}
