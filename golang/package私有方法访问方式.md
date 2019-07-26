###1、将exported类型变为非本package不可访问
```
定义一个internal文件夹,将这些package放在internal文件夹下。
即internal文件夹下的exported类型只能由internal所在的文件夹以下的文件访问。
这是通过go命令实现的，在1.4版本引入。

如 $GOPATH/src/mypkg/internal/foo只能被$GOPATH/src/mypkg import.
```

###2、访问其他package的私有方法
```
通过go:linkname指令，格式：//go:linkname localname importpath.name
这个指令告诉编译器为函数或者变量localname使用importpath.name作为目标文件的符号名。(因这个指令破坏了类型系统和包的模块化，所以只能
在import "unsafe" 的情况下使用)(importpath.name可以是这种格式:a/b/c/d/apkg.foo,这样在package a/b/c/d/apkg中就可以使用这个函数foo了)
```