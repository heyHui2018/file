package main

import (
	"unsafe"
	"fmt"
)

func main() {
	//在 cap 小于1024的情况下是每次扩大到 2 * cap ，当大于1024之后就每次扩大到 1.25 * cap

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr before: %v , the len: %v , the cap: %v \n",arr, len(arr), cap(arr))
	slice := arr[:3]//ptr:0xc04200c240
	fmt.Printf("slice before: %v , the len: %v , the cap: %v , ptr: %v \n",slice, len(slice), cap(slice),*(*unsafe.Pointer)(unsafe.Pointer(&slice)))
	slice01:= arr[1:3]//ptr:0xc04200c248 与slice相差8，即一个int，故，slice的cap是5，slice01是4，也可理解成slice01从第二位开始复制，少了一位，少了一个cap
	fmt.Printf("slice01 before: %v , the len: %v , the cap: %v , ptr: %v \n",slice01, len(slice01), cap(slice01),*(*unsafe.Pointer)(unsafe.Pointer(&slice01)))
	slice02:= arr[0:1]
	fmt.Printf("slice02 before: %v , the len: %v , the cap: %v , ptr: %v  \n",slice02, len(slice02), cap(slice02),*(*unsafe.Pointer)(unsafe.Pointer(&slice02)))

	slice[1]=100
	fmt.Printf("arr : %v \n",arr)//slice是引用，故slice修改了其第二个元素，它引用的arr也修改了
	fmt.Printf("slice01 : %v , the len: %v , the cap: %v \n",slice01, len(slice01), cap(slice01))
	fmt.Printf("slice02 : %v , the len: %v , the cap: %v \n",slice02, len(slice02), cap(slice02))

	slice=append(slice, 4,5,6)//cap翻倍,此时ptr：0xc0420780a0，与原来不同，故下面改slice时，arr不会被修改
	fmt.Printf("slice after: %v , the len: %v , the cap: %v , ptr: %v \n",slice, len(slice), cap(slice),*(*unsafe.Pointer)(unsafe.Pointer(&slice)))
	slice[1]=1000
	fmt.Printf("arr[1] : %v \n",arr[1])
}