package viewalbum

const (
	//FrameScratch skeleton frame
	FrameScratch = 0x000
	//FrameWithData with data field
	FrameWithData = 0x001
	//FrameWithPagination with pagination fields (total, count)
	FrameWithPagination = 0x002

	defaultErrorCode = 500
)

//Frame frame object for response
type Frame struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
	withData
	withPagination
}

type withData struct {
	Data interface{} `json:"data,omitempty"`
}

type withPagination struct {
	Count interface{} `json:"count,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

//PrepareFrame prepare frame for use
func PrepareFrame(mode int) *Frame {
	f := new(Frame)
	if mode&FrameWithData > 0 {
		f.Data = struct{}{}
	}
	if mode&FrameWithPagination > 0 {
		f.Count = 0
		f.Total = 0
	}
	return f
}

//RecordComplain set error into msg
func (f *Frame) RecordComplain(err error) {
	f.Status = defaultErrorCode
	f.Message = err.Error()
}
