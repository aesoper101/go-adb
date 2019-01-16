package Adb

import (
	"testing"
)

func TestAdb_ReadDevices(t *testing.T) {
	adb := NewAdb().ReadDevices()

	println(adb)
	for _,v := range adb{
		println(v.DeviceName, v.DeviceState)
	}
}

func TestAdb_LatestScreenShot(t *testing.T) {
	picName := "sc2.png"
	NewAdb().LatestScreenShot(picName)
}

func TestAdb_Pull(t *testing.T) {
	NewAdb().Pull("/sdcard/sc.png", "/Users/weilanzhuan/go/src/Adb")
}

func TestAdb_ScreenShot(t *testing.T) {
	NewAdb().ScreenShot("/sdcard/sc1.png")
}

func TestAdb_RemoveFile(t *testing.T) {
	NewAdb().RemoveFile("/sdcard/sc1.png")
}

func TestAdb_OldScreenShot(t *testing.T) {
	NewAdb().OldScreenShot("/sdcard/sc2.png", "/Users/weilanzhuan/go/src/Adb", false)
}