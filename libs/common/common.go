package common

import (
	"bytes"
	"os/exec"
)

func ExecShell(cmd string) (ret string, err error) {
	r := exec.Command("/bin/bash", "-c", cmd)
	var out bytes.Buffer
	r.Stdout = &out
	err = r.Run()
	ret = out.String()
	return
}
