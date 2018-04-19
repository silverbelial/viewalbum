package viewalbum

//Viewer interface type for context/controller type that handles requests and response
type Viewer interface {
	ServeObject(interface{})
	SetParam(string, interface{})
	ServeHTMLFile(string)
}

//MemoryViewer a viewer remembers static files location
type MemoryViewer interface {
	RememberStaticLoc(string, string)
}
