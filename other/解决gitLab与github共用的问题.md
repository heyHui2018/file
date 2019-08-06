####1、生成github秘钥
ssh-keygen -t rsa -C "github对应邮箱地址"
将生成的秘钥放入gitLab秘钥文件夹下,并修改名称以作区分,此时.ssh文件夹如下所示
```
.ssh-
    github_id_rsa
    github_id_rsa.pub
    id_rsa
    id_rsa.pub
```
####2、github账号设置ssh公钥
settings->SSH Keys->New SSH Key
Title处填写任意信息,key处拷贝上面生成的秘钥信息
####3、配置config文件
```
Host github.com
   HostName github.com
   User heyHui2018
   IdentityFile C:\Users\whaley\.ssh\github_id_rsa
   
Host git.moretv.cn
   HostName git.moretv.cn
   User huirenjie
   IdentityFile C:\Users\whaley\.ssh\id_rsa
```
####4、测试连接
ssh -T git@github.com
返回 Hi heyHui2018! You've successfully authenticated, but GitHub does not provide shell access.
ssh -T git@git.moretv.cn
返回 Welcome to GitLab, XXX!
####5、设置全局账户
git config --global user.name "Your Name Here"
git config --global user.email your@email.com
####6、设置局部账户
git config user.name "Your Name Here"
git config user.email your@email.com