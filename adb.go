package Adb

import (
	"Adb/utils"
	"regexp"
	"strings"
)

type Adb struct {

}

func NewAdb() *Adb {
	return new(Adb)
}

type Device struct {
	DeviceName string
	DeviceState string
}

// 获取设备列表
func (self *Adb) ReadDevices() []Device{
	DeviceList := make([]Device, 0)

	output, err := utils.ExcuteCommandResult("adb", []string{"devices"})

	if err != nil {
		return DeviceList
	}

	output1 := strings.Split(output, "\n")


	if len(output1) > 1 {
		for i := 1; i < len(output1); i++ {

			if 0 == len(output1[i]) || output1[i] == "\r\t" {
				continue
			}

			str := strings.Replace(output1[i], "\n", "", -1)
			re := regexp.MustCompile(`\s`)
			deviceInfo := re.Split(str, -1)

			var device Device
			device.DeviceName = deviceInfo[0]
			device.DeviceState = deviceInfo[1]
			DeviceList = append(DeviceList, device)

		}
	}

	return DeviceList
}

// 新版截图，直接存到电脑中
func (self *Adb) LatestScreenShot(pcPicPath string) error{

	return utils.ExcuteCommand("adb", []string{"exec-out", "screencap", "-p", ">", pcPicPath})
}

// 截图存到手机卡中
func (self *Adb) ScreenShot(sdPath string) error{

	return utils.ExcuteCommand("adb", []string{"shell", "screencap", "-p", sdPath})
}

// 上传截图到PC
func (self *Adb) Pull(sdFile string, pcPath string) error {
	return utils.ExcuteCommand("adb", []string{"pull", sdFile, pcPath})
}

// PC文件推送到手机
func (self *Adb) Push(pcFile string, sdPath string) error {
	return utils.ExcuteCommand("adb", []string{"push", pcFile, sdPath})
}

// 点击屏幕
func (self *Adb) Tap(x string, y string) error {
	return utils.ExcuteCommand("adb", []string{"shell", "input", "tap", x, y})
}

// 滑动屏幕
func (self *Adb) Swipe(x string, y string) error{
	return utils.ExcuteCommand("adb", []string{"shell", "input", "Swipe", x, y})
}