package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		s1, _ := reader.ReadString('\n')
		if len(s1) == 0 {
			continue
		}
		s1 = strings.Replace(s1, "\n", "", -1)
		if s1 == "end" {
			break
		}
		s2 := fmt.Sprintf("%v: 读到的数据：%s\n", time.Now(), s1)
		fmt.Println(s2)
		f, err := os.OpenFile("./echo.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			panic(err.Error())
		}
		_, err = f.WriteString(s2)
		if err != nil {
			panic(err.Error())
		}
	}
}
