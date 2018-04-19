package viewalbum

const (
	//FrameScratch skeleton frame
	FrameScratch = 0x000
	//FrameWithData with data field
	FrameWithData = 0x001

	defaultErrorCode = 500
)

//Frame frame object for response
type Frame struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
	withData
}

type withData struct {
	Data interface{} `json:"data,omitempty"`
}

//PrepareFrame prepare frame for use
func PrepareFrame(mode int) *Frame {
	f := new(Frame)
	if mode&FrameWithData > 0 {
		f.Data = struct{}{}
	}
	return f
}

//RecordComplain set error into msg
func (f *Frame) RecordComplain(err error) {
	f.Status = defaultErrorCode
	f.Message = err.Error()
}