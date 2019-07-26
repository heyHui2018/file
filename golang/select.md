###1.为请求设置超时时间
```
func main() {
    ch := make(chan int)
    timeout := time.After(5 * time.Second)
    for {
        select {
        case s := <-ch:
            fmt.Println(s)
        case <-timeout:
            fmt.Println("timeout!")
            return
        }
    }
}
```
###2.quit channel
```
select {
case c <- fmt.Sprintf("%s: %d", msg, i):
    // do something
case <-quit:
    cleanup()
    return
}
```
###3.生成随机串
```
for {
    select {
    case c <- 0:
    case c <- 1:
    case c <- a:
    case c <- b:
    }
}
```
###4.永久阻塞
```
select {}
```