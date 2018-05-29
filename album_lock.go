package viewalbum

var (
	currentLocks []AlbumLock
)

func init() {
	currentLocks = make([]AlbumLock, 0, 3)
}

//AlbumLock access control interface
type AlbumLock interface {
	HasAccess(Viewer, []int) bool
}

//AddLock add lock to album system
func AddLock(lock AlbumLock) {
	currentLocks = append(currentLocks, lock)
}

//TryOpen test accessability
func TryOpen(v Viewer, r []int) bool {
	for _, lock := range currentLocks {
		if !lock.HasAccess(v, r) {
			return false
		}
	}
	return true
}
