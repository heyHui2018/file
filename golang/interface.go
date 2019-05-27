package main

import (
	"fmt"
)

/*
Man/Woman均实现了People下的add方法,故后续若要加新的struct,只要实现了People下的add方法,无需修改其余代码即可完成
*/
type People interface {
	add() int
}

type Man struct {
	a int
}

func (m Man) add() int {
	return m.a
}

type Woman struct {
	a int
	b int
}

func (w Woman) add() int {
	return w.a + w.b
}

func main() {
	p1 := Man{1}
	p2 := Man{2}
	p3 := Woman{3, 4}
	list := []People{p1, p2, p3}
	fmt.Println(total(list))
}

func total(list []People) int {
	total := 0
	for _, v := range list {
		total += v.add()
	}
	return total
}
