package main

import (
	"fmt"
	// "strconv"
	"reflect"
)

var total int

type SumFunc func(...int) SumFunc

/*
注销下面string方法后,fmt.Println(sum(2)(3))和其他关于sum的打印的输出其实都是sum的地址,因为sum的类型是SumFunc,值是SumFunc的地址.
之所以打印的是地址,是因为闭包函数中最后的返回是sum,注销string后调用的是官方包中的string方法.
*/
// func (s SumFunc) String() string { // 因fmt包中的print相关函数会调用string方法,若自己声明了string方法,则会调用该方法
// 	fmt.Println("33333333333333 total = ",total)
// 	tmpTotal := total
// 	total = 0
// 	return strconv.Itoa(tmpTotal)
// }

func main() {

	var sum SumFunc

	sum = func(nums ...int) SumFunc {
		for _, num := range nums {
			total += num
		}
		return sum
	}
	fmt.Println(reflect.TypeOf(sum))
	fmt.Println(reflect.ValueOf(sum))
	// fmt.Println(sum(2, 3))
	fmt.Println(sum(2)(3))
	// 之所以能写成sum(2)(3)的样子,是因为sum(2)返回的还是sum,故可以继续带参数变成sum(2)(3)的样子
}
