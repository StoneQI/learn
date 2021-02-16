/*
 * @Author: your name
 * @Date: 2021-02-08 20:32:50
 * @LastEditTime: 2021-02-17 00:38:55
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /learn/go-test/time_learn.go
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	now := time.Now()
	var a *sync.Mutex
	a.Lock()

	a.Unlock()

	fmt.Println(now)
}
