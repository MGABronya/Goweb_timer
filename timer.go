// @Title  timer
// @Description  利用通道和time包制作一个定时器
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-10 00:39
package main

import (
	"fmt"
	"time"
)

// @title    Timer
// @description  创建一个定时器
// @auth      MGAronya（张健）             2022-10-10 00:39
// @param     duration time.Duration				定时时间长度
// @return    chan bool			被传入的通道
func Timer(duration time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(duration)
		// TODO 到达时间
		ch <- true
	}()
	return ch
}

// @title    main
// @description		使用一个五秒的定时器
// @auth      MGAronya（张健）             2022-10-10 00:39
// @param     duration time.Duration				定时时间长度
// @return    chan bool			被传入的通道
func main() {
	// TODO 新建一个5s的定时器
	timeout := Timer(5 * time.Second)
	for {
		select {
		case <-timeout:
			// TODO 五秒到达
			fmt.Println("already 5s!")
			return
		}
	}
}
