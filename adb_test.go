package Adb

import "testing"

func TestAdb_ReadDevices(t *testing.T) {
	adb := NewAdb().ReadDevices()

	println(adb)
	for _,v := range adb{
		println(v.DeviceName, v.DeviceState)
	}
}
