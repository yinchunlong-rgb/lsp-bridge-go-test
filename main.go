package main

import (
	"fmt"
	"time"
)

type TTTTT struct {
	a int32
}

func testfunc(s *TTTTT) {}

func main() {
	//todo 这里做测试
	go func() {
		for {
			fmt.Println(time.Now().Local())
			testfunc(&TTTTT{})
			testfunc(&TTTTT{})
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Hour)
	time.Sleep(time.Hour)
}
