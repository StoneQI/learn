package main

import "testing"

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		Test1()
	}
}
func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		Test2()
	}
}
