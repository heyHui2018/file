###1、整个map的直接赋值会导致两个变量指向同一个底层数据,修改其中一个时,另一个也会被改变
```
a := make(map[string]int)
b := a
当修改b时,a也会被改变
```