package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var i int64
	var s string
	for {
		cmd := exec.Command(os.Args[1], os.Args[2:]...)
		b, err := cmd.CombinedOutput()
		if err != nil {
			s = fmt.Sprintln(err.Error())
		} else {
			s = string(b)
		}
		fmt.Printf("[%d]%s", i, s)
		i++
	}
}
