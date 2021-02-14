/*
 * @Author: your name
 * @Date: 2021-02-08 20:32:50
 * @LastEditTime: 2021-02-14 20:41:20
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

func main() {
	now := time.Now()
	var a *sync.Mutex
	a.Lock()

	a.Unlock()

	fmt.Println(now)
}
