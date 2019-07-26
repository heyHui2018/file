####1、搜索nsq镜像
docker search nsq(此处nsq为关键词)

####2、拉取nsq镜像
docker pull nsqio/nsq

####3、运行lookupd
docker run -d --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd

####4、查询lookupd ip
docker inspect lookupd

####5、运行nsq
docker run -d --name nsqd -p 4150:4150 -p 4151:4151 nsqio/nsq /nsqd --broadcast-address=172.16.16.114 --lookupd-tcp-address=172.16.16.114:4160
(此处的ip为所在服务器ip)

####6、运行nsqadmin
docker run -d --name nsqadmin -p 4171:4171 nsqio/nsq /nsqadmin  --lookupd-http-address=172.16.16.114:4161
(此处的ip为所在服务器ip)

####7、访问nsqadmin
http://ip:4171/nodes
(此处的ip为nsqadmin所在服务器ip)