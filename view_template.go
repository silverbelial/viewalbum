package viewalbum

import (
	"errors"
	"sync"
)

var (
	templates     = make(map[string]ViewTemplate)
	templateMutex sync.RWMutex

	currentTemplate ViewTemplate
)

//ViewTemplate interface type for view template
type ViewTemplate interface {
	AssetsMapping() map[string]string
	TemplateName() string
	SetProfile(map[string]string)
	LayoutPage() string
	LoginPage() string
}

//RegisterTemplate register view template
func RegisterTemplate(template ViewTemplate) error {
	templateMutex.Lock()
	defer templateMutex.Unlock()

	if template == nil {
		return errors.New("mvt: Registering nil template")
	}

	if _, dup := templates[template.TemplateName()]; dup {
		return errors.New("mvt: Register called twice for template " + template.TemplateName())
	}

	templates[template.TemplateName()] = template
	return nil
}

//GetTemplate Get Template By Registered name
func GetTemplate(name string) ViewTemplate {
	if templates == nil {
		return nil
	}
	return templates[name]
}

//UseTemplate Set Current running template
//Usage option
func UseTemplate(name string, mv MemoryViewer) error {
	if templates == nil {
		return errors.New("Templates not initlialized")
	}
	template, has := templates[name]
	if !has {
		return errors.New("Invalid template name")
	}
	currentTemplate = template
	if mv != nil {
		for key, value := range template.AssetsMapping() {
			mv.RememberStaticLoc(key, value)
		}
	}
	return nil
}

//GetAssets get current template assets
func GetAssets() map[string]string {
	if currentTemplate == nil {
		return make(map[string]string)
	}
	return currentTemplate.AssetsMapping()
}

//GetLayoutName get current template layout name
func GetLayoutName() string {
	if currentTemplate == nil {
		return ""
	}
	return currentTemplate.LayoutPage()
}

//GetLoginPage get current template login page
func GetLoginPage() string {
	if currentTemplate == nil {
		return ""
	}
	return currentTemplate.LoginPage()
}

//GetTemplateName get current template name
func GetTemplateName() string {
	if currentTemplate == nil {
		return ""
	}
	return currentTemplate.TemplateName()
}
