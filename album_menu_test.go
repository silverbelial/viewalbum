package viewalbum_test

import (
	"testing"
	"viewalbum"
)

func TestAlbumMenu(t *testing.T) {
	m := viewalbum.RegsiterRootMenu(&viewalbum.ViewObject{
		Title: "Root",
		Link:  "/root",
	}, "", []int{0})
	m.RegisterSubMenu(&viewalbum.ViewObject{
		Title: "Sub", Link: "/root/sub",
	}, "", []int{0})
	ms := viewalbum.GetMenus()
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
