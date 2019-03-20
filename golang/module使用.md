##本文介绍由gopath切换到modules的具体方法

####1、下载新版golang(1.11+),并设置GOMOD为on
golang升级仅需覆盖安装

####2、在gopath外创建目录并将之前的项目移至此目录中

####3、以moretvMVActivity项目为例,其原保存路径为gopath下的 github.com/moretv/moretvMVActivity
* 1、go mod init github.com/moretv/moretvMVActivity
* 2、go build
此时应可编译成功,若失败,需根据具体错误仅需解决.
编辑器中仍有红字.

####4、将编辑器更新至golang1.11之后发行的版本(之后发行的版本才支持go module)

####5、在file -> settings -> Go -> GOPATH中,取消关于GOPATH的两个勾选.在file -> settings -> Go -> GO Module中,勾选关于Module的选项.