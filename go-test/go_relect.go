package main

import (
	"unsafe"
)

func a() int { return 1 }
func b() int { return 2 }

func rawMemoryAccess(b uintptr) []byte { // 获取b地址的内存块
	return (*(*[0xFF]byte)(unsafe.Pointer(b)))[:] // 获取b对应的内存块，以便改写内容，从而达到替换目的
}

func assembleJump(f func() int) []byte { // 生成跳转到f函数所在位置的机器码
	funcVal := *(*uintptr)(unsafe.Pointer(&f))
	return []byte{
		0x48, 0xC7, 0xC2,
		byte(funcVal >> 0),
		byte(funcVal >> 8),
		byte(funcVal >> 16),
		byte(funcVal >> 24), // MOV rdx, funcVal
		0xFF, 0x22,          // JMP [rdx]
	}
}

func replace(orig, replacement func() int) {
	bytes := assembleJump(replacement)
	functionLocation := **(**uintptr)(unsafe.Pointer(&orig))
	window := rawMemoryAccess(functionLocation)

	copy(window, bytes)
}

func main() {
	replace(a, b)
	print(a())
}
