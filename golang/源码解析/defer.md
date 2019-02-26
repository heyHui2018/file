源码位置 runtime/panic.go
defer结构体定义位置 runtime/runtime2.go:_defer
###1、规则
* 1、延迟函数的参数在defer语句出现时就已经确定下来了（若参数是一个地址,后续对该地址指向的变量的修改会影响延迟函数）
* 2、延迟函数执行按后进先出顺序执行，即先出现的defer最后执行
* 3、延迟函数可能操作主函数的具名返回值,可分为3种情况
    * A、主函数拥有匿名返回值，返回字面值
    ```
    func foo() int {    
    var i int
    defer func() {
        i++
    }()    
    return 1}
    ```
    此种情况不会影响
    * B、主函数拥有匿名返回值，返回变量
     ```
     func foo() int {    
     var i int
     defer func() {
         i++
     }()    
     return i}   
     ```
     <big>此种情况不会影响</big>
    * C、主函数拥有具名返回值
     ```
     func foo() (ret int) {    
     defer func() {
         ret++
     }()    
     return 0}
     此种情况会影响
     ```
###2、原理
每次defer都会把参数压入栈中,函数返回前再把参数取出并执行.编译器将defer处理成两个函数调用：
* 1.deferproc定义了一个延迟调用对象
```
func deferproc(siz int32, fn *funcval)
```
* 2.然后在函数结束前即return指令前,更准确地说是汇编的ret指令前,通过deferreturn完成最终调用
```
func deferreturn(arg0 uintptr)
```
###3、结构
type _defer struct {
    sp      uintptr   //函数栈指针
    pc      uintptr   //程序计数器
    fn      *funcval  //函数地址
    link    *_defer   //指向自身结构的指针，用于链接多个defer
}