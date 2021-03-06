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
	ViewObject    *ViewObject

	I18nKey string
}

//DeepClone make a copy of itself
func (m *AlbumMenu) DeepClone() *AlbumMenu {
	c := &AlbumMenu{
		Title:         m.Title,
		IconClassName: m.IconClassName,
		URL:           m.URL,
		SubMenu:       make([]*AlbumMenu, 0, len(m.SubMenu)),
		AcceptRoles:   m.AcceptRoles,
		ViewObject:    m.ViewObject,
		I18nKey:       m.I18nKey,
	}

	for _, sub := range m.SubMenu {
		c.SubMenu = append(c.SubMenu, sub.DeepClone())
	}

	return c
}

//IsCurrent return if menu is current
func (m *AlbumMenu) IsCurrent(vo *ViewObject) bool {
	cvo := m.ViewObject
	for cvo != nil {
		if cvo == vo {
			return true
		}
		cvo = cvo.Parent
	}
	for _, sm := range m.SubMenu {
		if sm.IsCurrent(vo) {
			return true
		}
	}
	return false
}

//HasSub returns menu has sub
func (m *AlbumMenu) HasSub() bool {
	return len(m.SubMenu) > 0
}

//Authorized returns is current user is authorized to view this menu
func (m *AlbumMenu) Authorized() bool {
	return true
}

//RegisterRootMenu register vo as root menu
func RegisterRootMenu(vo *ViewObject, icon string, roles []int) *AlbumMenu {
	return RegisterRootMenuI18n(vo, icon, roles, "")
}

//RegisterRootMenuI18n register vo as root menu with i18n feature
func RegisterRootMenuI18n(vo *ViewObject, icon string, roles []int, i18nKey string) *AlbumMenu {
	m := &AlbumMenu{
		Title:         vo.Title,
		IconClassName: icon,
		URL:           vo.Link,
		SubMenu:       make([]*AlbumMenu, 0, 5),
		AcceptRoles:   roles,
		ViewObject:    vo,
		I18nKey:       i18nKey,
	}
	menus = append(menus, m)
	return m
}

//RegisterSubMenu register sub menu
func (m *AlbumMenu) RegisterSubMenu(vo *ViewObject, icon string, roles []int) *AlbumMenu {
	return m.RegisterSubMenuI18n(vo, icon, roles, "")
}

//RegisterSubMenuI18n register sub menu with i18n feature
func (m *AlbumMenu) RegisterSubMenuI18n(vo *ViewObject, icon string, roles []int, i18nKey string) *AlbumMenu {
	sm := &AlbumMenu{
		Title:         vo.Title,
		IconClassName: icon,
		URL:           vo.Link,
		SubMenu:       make([]*AlbumMenu, 0, 5),
		AcceptRoles:   roles,
		ViewObject:    vo,
		I18nKey:       i18nKey,
	}
	m.SubMenu = append(m.SubMenu, sm)
	return sm
}

//GetMenus return menus
func GetMenus() []*AlbumMenu {
	ms := make([]*AlbumMenu, 0, len(menus))
	for _, menu := range menus {
		ms = append(ms, menu.DeepClone())
	}
	return ms
}

//MenuCover pre-defined cover for
type MenuCover struct {
	ParamKey string
}

//PreProcess implements AlbumCover
func (mc *MenuCover) PreProcess(vr Viewer) {
	ms := GetMenus()
	ms = menuLock(vr, ms)
	vr.SetParam(mc.ParamKey, ms)
}

func menuLock(vr Viewer, ms []*AlbumMenu) []*AlbumMenu {
	r := make([]*AlbumMenu, 0, len(ms))
	for _, m := range ms {
		if TryOpen(vr, m.AcceptRoles) {
			m.SubMenu = menuLock(vr, m.SubMenu)
			r = append(r, m)
		}
	}
	return r
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
