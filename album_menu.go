package viewalbum

const (
	defaultMenuParamKey = `Menus`
)

var (
	menus []*AlbumMenu
)

func init() {
	menus = make([]*AlbumMenu, 0, 10)
}

//AlbumMenu menu object
type AlbumMenu struct {
	Title         string
	IconClassName string
	URL           string
	SubMenu       []*AlbumMenu
	AcceptRoles   []int
}

//RegsiterRootMenu register vo as root menu
func RegsiterRootMenu(vo *ViewObject, icon string, roles []int) *AlbumMenu {
	m := &AlbumMenu{
		Title:         vo.Title,
		IconClassName: icon,
		URL:           vo.Link,
		SubMenu:       make([]*AlbumMenu, 0, 5),
		AcceptRoles:   roles,
	}
	menus = append(menus, m)
	return m
}

//RegisterSubMenu register sub menu
func (m *AlbumMenu) RegisterSubMenu(vo *ViewObject, icon string, roles []int) *AlbumMenu {
	sm := &AlbumMenu{
		Title:         vo.Title,
		IconClassName: icon,
		URL:           vo.Link,
		SubMenu:       make([]*AlbumMenu, 0, 5),
		AcceptRoles:   roles,
	}
	m.SubMenu = append(m.SubMenu, sm)
	return sm
}

//GetMenus return menus
func GetMenus() []*AlbumMenu {
	return menus
}

//MenuCover pre-defined cover for
type MenuCover struct {
	ParamKey string
}

//PreProcess implements AlbumCover
func (mc *MenuCover) PreProcess(vr Viewer) {
	//TODO check role
	vr.SetParam(mc.ParamKey, GetMenus())
}

//EnableMenuCover enable MenuCover
func EnableMenuCover(tag string) {
	if tag == "" {
		tag = defaultMenuParamKey
	}
	AddCover(&MenuCover{
		ParamKey: tag,
	})
}
