####1、搜索mysql镜像
docker search mysql(此处mysql为关键词)

####2、拉取mysql镜像
docker pull mysql(此处mysql为镜像名)

####3、运行容器
docker run -p 3306:3306 --name HRJ_mysql -v $PWD/conf/my.cnf:/etc/mysql/my.cnf -v $PWD/logs:/logs -v $PWD/data:/mysql_data -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```
-p 3306:3306 将容器的3306端口映射到主机的3306端口
--name 设置别名
-v $PWD/conf/my.cnf:/etc/mysql/my.cnf：将主机当前目录下的conf/my.cnf挂载到容器的/etc/mysql/my.cnf
-v $PWD/logs:/logs：将主机当前目录下的logs目录挂载到容器的/logs
-v $PWD/data:/mysql_data：将主机当前目录下的data目录挂载到容器的/mysql_data
-e MYSQL_ROOT_PASSWORD=123456 设置 账号为root,密码为123456
-d 后台运行
```

####4、连接本地mysql
docker run -it --link HRJ_mysql --rm mysql  mysql -h HRJ_mysql -uroot -p"123456"
docker run -it --rm mysql mysql -h 172.17.0.4 -uroot -p"123456"

####5、连接其余mysql
docker run -it --rm mysql mysql -h 172.16.17.121 -uroot -p"111111