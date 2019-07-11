####1、搜索redis镜像
docker search redis(此处redis为关键词)

####2、拉取redis镜像
docker pull redis(此处redis为镜像名)

####3、运行容器
docker run -p 6379:6379 -d --name redis-2 -v /root/docker/redis/:/data redis redis-server
```
-p 6379:6379 将容器的6379端口映射到主机的6379端口
-d 后台运行
-name 设置别名
-v /root/docker/redis/:/data 将/root/docker/redis/挂载到容器的/data(数据默认存储在VOLUMN /data目录下)
redis 镜像名
redis-server redis命令
```

####4、运行redis-cli
docker run -it --link redis-2 --rm redis redis-cli -h redis-2 -p 6379
docker run -it --rm redis redis-cli -h ip -p 6379(此处ip为容器自有的ip,可用docker inspect redis-2或容器id查看)
```
-it 交互模式
--link 连接另一个容器,这样可以使用容器名作为host
redis-2 上面设置的容器别名
--rm 一旦进程退出就删除容器
redis 镜像名
redis-cli -h redis-2 -p 6379 redis命令,此处用redis-2是因为上面设置了容器别名(因前面link了)
```

####5、解决‘Cannot connect to the Docker daemon’问题
* 切换到root权限
* service docker start