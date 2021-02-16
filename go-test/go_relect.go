/*
 * @Author: your name
 * @Date: 2021-02-17 00:34:35
 * @LastEditTime: 2021-02-17 00:43:00
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /learn/go-test/go_test.go
 */
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type nonEmptyInterface struct {
	// see ../runtime/iface.go:/Itab
	itab *struct {
		ityp *rtype // static interface type
		typ  *rtype // dynamic concrete type
		hash uint32 // copy of typ.hash
		_    [4]byte
		fun  [100000]unsafe.Pointer // method table
	}
	word unsafe.Pointer
}
type rtype struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      uint8   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte // garbage collection data
	str       int32 // string form
	ptrToThis int32 // type for pointer to this type, may be zero
}

type a interface {
	say()
}
type b struct {
	aa int
}

func (bb *b) say() {
	fmt.Println(bb.aa)
}

type c struct {
	aa string
}

func (bb *c) say() {
	fmt.Println(bb.aa)
}

func main() {

	bbb := &b{aa: 12}
	ccc := &c{aa: "123"}

	templateDaoMock1 := reflect.ValueOf(bbb).Elem()

	cccc := templateDaoMock1.Pointer()

	_ = cccc
	templateDaoMock2 := (*nonEmptyInterface)(unsafe.Pointer(templateDaoMock1.UnsafeAddr()))

	fpsTemplateDaoPtr := reflect.ValueOf(ccc).Elem()
	fpsTemplateDaoUnPtr := (*nonEmptyInterface)(unsafe.Pointer(fpsTemplateDaoPtr.UnsafeAddr()))
	fpsTemplateDaoUnPtr.itab.fun = templateDaoMock2.itab.fun

	ccc.say()
}
