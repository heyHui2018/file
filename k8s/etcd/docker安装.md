####0、查找镜像
docker search etcd

####1、拉取镜像(此处选择v3镜像)
pull xieyanze/etcd3

####2、运行节点0
docker run -d -p 2380:2380 -p 2379:2379 --name etcd xieyanze/etcd3 -name etcd -advertise-client-urls http://172.16.16.114:2379  
-listen-client-urls http://0.0.0.0:2379 -initial-advertise-peer-urls http://172.16.16.114:2380 -listen-peer-urls http://0.0.0.0:2380  
-initial-cluster-token etcd-cluster-1 -initial-cluster "etcd=http://172.16.16.114:2380" -initial-cluster-state new

####3、验证
curl -L http://172.16.16.114:2479/v2/members