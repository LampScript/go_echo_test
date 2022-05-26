package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

//怎么实现
//go build ./pkg/echo.go
//go run process.go

func main() {
	i := 0
	for i < 1 {
		i++
		tmp := exec.Command("./echo")
		tmp.Stdout = os.Stdout
		stdin := bytes.Buffer{}
		tmp.Stdin = &stdin
		_, err := stdin.WriteString(fmt.Sprintf("this is %d", i))
		if err != nil {
			panic(err.Error())
		}
		tmp.Start()
		if err != nil {
			panic(err.Error())
		}
		err = tmp.Process.Kill()
		if err != nil {
			panic(err.Error())
		}
		time.Sleep(time.Second)
	}
	fmt.Println("end")
}
