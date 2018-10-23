package viewalbum

import (
	"testing"
)

func TestAlbumMenu(t *testing.T) {
	m := RegsiterRootMenu(&ViewObject{
		Title: "Root",
		Link:  "/root",
	}, "", []int{0})
	m.RegisterSubMenu(&ViewObject{
		Title: "Sub", Link: "/root/sub",
	}, "", []int{0})
	ms := GetMenus()
	if len(ms) == 0 {
		t.Log("root length failed")
		t.Fail()
		return
	}
	if len(ms[0].SubMenu) == 0 {
		t.Log("sub length failed")
		t.Fail()
		return
	}
}
