package utils

import (
	"fmt"
	"testing"
)

func TestExcuteCommand(t *testing.T) {
	str := fmt.Sprintf("adb")
	ExcuteCommand(str, []string{"devices"})
}

func TestExcuteCommandResult(t *testing.T) {
	str := fmt.Sprintf("adb")
	str1,_ := ExcuteCommandResult(str, []string{"devices"})
	fmt.Println(str1)

	str2,_ := ExcuteCommandResult("echo", []string{"devices"})
	fmt.Println(str2)
}

