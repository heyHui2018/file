####1、pprof有两种包,net/http/pprof及runtime/pprof,前者对后者进行了封装并在http上暴露端口.
在使用了某些web框架如beego后再使用pprof,无法直接在项目监听的端口直接使用pprof,因这些web框架对端口监听及router进行了封装,不会读取pprof的init中的router,
所以需要重新设置监听端口
####2、使用情况：
* 直接引入包 _"net/http/pprof"
* 引入包 _"net/http/pprof"后开启额外的goroutine来监听
```
go func(){
    log.Pringtln(http.ListenAndServe("localhost:8999",nil))
}()
```
* 引入包 "runtime/pprof"
```
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	...
```
####3、使用方法：
* 先运行程序
* 执行命令：go tool pprof http://localhost:8999/debug/pprof/profile,会进入30秒获取数据状态，需同时进行并发请求
(profile-分析cpu,heap-分析内存)
* boom -n 1000 -c 100 "http://127.0.0.1:8081/kidsEdu/userCenter/courseProfiler/newCourses"
* top、web或其余命令