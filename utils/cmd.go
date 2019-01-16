package utils

import (
	"io/ioutil"
	"os/exec"
)

func ExcuteCommand(command string, params []string) (err error) {

	err = exec.Command(command, params...).Run()

	if err != nil {
		checkErr(err)
		return
	}

	return
}

func ExcuteCommandResult(command string, params []string) (result string, err error) {

	cmd := exec.Command(command, params...)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		checkErr(err)
		return
	}

	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		checkErr(err)
	}

	output , err := ioutil.ReadAll(stdout)
	result = string(output)

	return

}