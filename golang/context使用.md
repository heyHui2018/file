###1.简介
当面临一些复杂多变的网络并发场景时， channel 和 WaitGroup 会显得有些力不从心。比如网络请求 request，每个 request 都需要开启一个 goroutine，这些 goroutine 又可能会开启新的 goroutine，如数据库或RPC服务。所以我们需要一种可以跟踪 goroutine 的方案来实现控制他们的目的。Go为我们提供了 Context，称之为上下文很贴切，它就是 goroutine 的上下文。它包括一个程序的运行环境、现场和快照等。每个程序要运行时，都需要知道当前程序的运行状态，通常 Go 将这些封装在一个 Context 里，再将它传给要执行的 goroutine 。context 包主要用来处理多个 goroutine 之间的共享数据及多个 goroutine 的管理。
###2.Context包
核心是 struct Context，其中：
* Done()：返回一个只能接受数据的channel类型，当该Context关闭或者超时时，该channel就会有一个取消信号。
* Err()：在Done() 之后，返回context取消的原因。
* Deadline()：设置该context cancel的时间点
* Value()：允许 Context 对象携带request作用域的数据，该数据必须是线程安全的。
Context 对象是线程安全的，你可以把一个 Context 对象传递给任意个数的 gorotuine，对它执行取消操作时，所有 goroutine 都会接收到取消信号。一个 Context 不能拥有 Cancel 方法，同时我们也只能 Done channel 接收数据。背后的原因是一致的：接收取消信号的函数和发送信号的函数通常不是一个。一个典型的场景是：父操作为子操作启动 goroutine，子操作也就不能取消父操作。
###3.继承Context
Context 包提供了一些协助用户从现有的 Context 对象创建新的 Context 对象的函数。这些 Context 对象形成一棵树：当一个 Context 对象被取消时，继承它的所有 Context 都会被取消。
###4.例子
```
func childFunc(ctx context.Context, num *int) {
    ctxWithCancel, _ := context.WithCancel(ctx)
    for {
        select {
        case <-ctxWithCancel.Done():
        fmt.Println("child Done : ", ctxWithCancel.Err())
        return
        }
    }
}

func main() {
    gen := func(ctx context.Context) <-chan int {
        dst := make(chan int)
        n := 1
        go func() {
            for {
                select {
                    case <-ctx.Done():
                        fmt.Println("parent Done : ", ctx.Err())
                        return // returning not to leak the goroutine
                    case dst <- n:
                        n++
                        go childFunc(ctx, &n)
                    }
                }
            }()
            return dst
        }

    ctx, cancel := context.WithCancel(context.Background())
    for n := range gen(ctx) {
        fmt.Println(n)
        if n >= 5 {
            break
        }
    }
    cancel()
    time.Sleep(5 * time.Second)
}
```
上面的例子主要描述的是通过一个channel实现一个循环次数为5的循环。
在每个循环中产生goroutine，每个goroutine中都传入context，通过传入ctx创建一个子Context并通过select监控该Context的运行情况。当在父Context退出的时候，代码中并没有调用子Context的Cancel函数，但是通过分析结果可以看出，子Context还是被正确合理的关闭了，这是因为所有基于这个Context或者衍生的子Context都会收到通知，从而进行清理操作，最终释放goroutine ，这就优雅的解决了goroutine启动后不可控的问题。
###5.使用原则
* 不要把Context放在结构体中，要以参数的方式传递
* 以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位
* 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
* Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
* Context是线程安全的，可以放心地在多个goroutine中传递

