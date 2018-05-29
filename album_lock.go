package viewalbum

var (
	currentLocks []AlbumLock
)

func init() {
	currentLocks = make([]AlbumLock, 0, 3)
}

//AlbumLock access control interface
type AlbumLock interface {
	HasAccess([]int) bool
}

//AddLock add lock to album system
func AddLock(lock AlbumLock) {
	currentLocks = append(currentLocks, lock)
}

//TryOpen test accessability
func TryOpen(r []int) bool {
	for _, lock := range currentLocks {
		if !lock.HasAccess(r) {
			return false
		}
	}
	return true
}
