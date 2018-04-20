package viewalbum

//Viewer interface type for context/controller type that handles requests and response
type Viewer interface {
	ServeJSONObject(interface{})
	SetParam(string, interface{})
	ServeHTMLFile(string)
}

//StubbornViewer interface type for context/controller type that handles viewtemplate requests
type StubbornViewer interface {
	Viewer
	ServeReplacable(string, string)
}

//MemoryViewer a viewer remembers static files location
type MemoryViewer interface {
	RememberStaticLoc(string, string)
}
