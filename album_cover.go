package viewalbum

import (
	"sync"
)

var (
	covers     []AlbumCover
	coverMutex sync.RWMutex
)

func init() {
	covers = make([]AlbumCover, 0, 10)
}

//AlbumCover Pre processer interface
type AlbumCover interface {
	PreProcess(Viewer)
}

//AddCover add AlbumCover
func AddCover(ac AlbumCover) {
	if ac != nil {
		coverMutex.Lock()
		defer coverMutex.Unlock()
		covers = append(covers, ac)
	}
}
