###1、概述
可解决数据持久化、共享两个问题
Kubernetes支持多种类型存储卷,pod可以同时使用任意类型、数量的存储卷,pod通过制定字段来使用:
* spec.volumns // 指定存储卷
* spec.containers.volumeMounts // 将存储卷挂载到容器中
###2、分类
####A、EmptyDir
此类存储卷是为pod对应的目录创建一个空的文件夹,此文件夹会随着pod的删除而删除
####A、HostPath
此类存储卷是将宿主机文件系统的文件/目录直接挂载到pod中,除了需要指定path字段外,还可以设置type:
* -空字符串(默认)用于向后兼容,即挂载到宿主机之前不执行任何检查
* DirectoryOrCreate-若path指定的目录不存在,则会根据path在宿主机上创建一个新的,权限设置为0755,此目录将和kubelet拥有一样的组及拥有者
* Directory-path指定的目录必须存在
* FileOrCreate-若path指定的文件不存在,则会根据path在宿主机上创建一个新的,权限设置为0644,此文件将和kubelet拥有一样的组及拥有者
* File-path指定的文件必须存在
* Socket-path指定的unix socket必须存在
* CharDevice-path指定的字符设备必须存在
* BlockDevice-path指定的块设备必须存在

在使用hostPath时,还需注意：
* 具有相同配置的pod可能会因为node的文件不同而行为不同
* 在宿主机上创建的目录或文件,只有root用户有写入权限,故须以root权限运行进程,或修改这些目录或文件的权限以便写入

示例：YAML文件,定义了一个名为test-pd的Pod,选择hostPath类型存储卷将宿主机上的/data目录挂载到容器的/test-pd目录
```
apiVersion: v1
kind: Pod
metadata:
  name: test-pd
spec:
  containers:
  - image: k8s.gcr.io/test-webserver
    name: test-container
    # 指定在容器中挂接路径
    volumeMounts:
    - mountPath: /test-pd
      name: test-volume
  # 指定所提供的存储卷
  volumes:
  - name: test-volume
    hostPath:
      # 宿主机上的目录
      path: /data
      # this field is optional
      type: Directory
```
####B、NFS
此类存储卷是将现有的NFS(网络文件系统)直接挂载到pod中,在pod被移除时,NFS存储卷上的内容会被保留,故此方式可以在pod间共享数据.NFS可被同时挂载到多个pod中,
且可同时进行写入.使用此方式仅需确保在使用前NFS服务器已正确部署且设置了共享目录

示例：YAML文件,使用NFS存储卷,NFS服务器地址为192.168.8.150,路径为/k8s-nfs/redis/data,容器通过volumeMounts.name的值确定确定所使用的存储卷
```
apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
kind: Deployment
metadata:
  name: redis
spec:
  selector:
    matchLabels:
      app: redis
  revisionHistoryLimit: 2
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
      # 应用的镜像
      - image: redis
        name: redis
        imagePullPolicy: IfNotPresent
        # 应用的内部端口
        ports:
        - containerPort: 6379
          name: redis6379
        env:
        - name: ALLOW_EMPTY_PASSWORD
          value: "yes"
        - name: REDIS_PASSWORD
          value: "redis"   
        # 持久化挂接位置，在docker中 
        volumeMounts:
        - name: redis-persistent-storage
          mountPath: /data
      volumes:
      # 宿主机上的目录
      - name: redis-persistent-storage
        nfs:
          path: /k8s-nfs/redis/data
          server: 192.168.8.150
```
```
apiVersion: v1
kind: Pod
metadata:
  name: nfs-web
spec:
  containers:
    - name: web
      image: nginx
      imagePullPolicy: Never  #如果已经有镜像，就不需要再拉取镜像
      ports:
        - name: web
          containerPort: 80
          hostPort: 80        #将容器的80端口映射到宿主机的80端口
      volumeMounts:
        - name : nfs          #指定名称必须与下面一致
          mountPath: "/usr/share/nginx/html"        #容器内的挂载点
  volumes:
    - name: nfs               #指定名称必须与上面一致
      nfs:            #nfs存储
        server: 192.168.66.50        #nfs服务器ip或是域名
        path: "/test"                #nfs服务器共享的目录
```
####C、PersistentVolumeClaim(PVC)
此类存储卷是将PersistentVolume直接挂载到pod中,用户并不知道存储卷的详细信息

PersistentVolume(PV)是一种待分配的存储资源,而PVC表示用户对存储资源的申请,这种概念与pod类似,PVC消耗了持久卷资源,pod小号了节点上cpu和内存等资源

示例：YAML文件,使用镜像为busybox,基于此镜像的容器需要对/mnt目录下的数据进行持久化,YAML文件中指定使用名称为nfs的PersistenVolumeClaim对容器的数据进行持久化
```
# This mounts the nfs volume claim into /mnt and continuously
# overwrites /mnt/index.html with the time and hostname of the pod. 
apiVersion: v1
kind: Deployment
metadata:  
  name: busybox-deployment
spec:  
  replicas: 2  
  selector:    
    name: busybox-deployment
  template:    
    metadata:      
      labels:        
        name: busybox-deployment    
    spec:      
      containers:      
      - image: busybox        
        command:          
        - sh          
        - -c          
        - 'while true; do date > /mnt/index.html; hostname >> /mnt/index.html; sleep $(($RANDOM % 5 + 5)); done'        
        imagePullPolicy: IfNotPresent        
        name: busybox
        volumeMounts:
        # name must match the volume name below          
        - name: nfs            
          mountPath: "/mnt"     
     volumes:      
     - name: nfs        
       persistentVolumeClaim:          
         claimName: nfs-pvc
```
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nfs-pv
  labels:
    type: nfs         #指定类型是NFS
spec:
  capacity:           #指定访问空间是15G
    storage: 15Gi
  accessModes:        #指定访问模式是能在多节点上挂载，并且访问权限是读写执行
    - ReadWriteMany
  persistentVolumeReclaimPolicy: Recycle        #指定回收模式是自动回收，当空间被释放时，K8S自动清理，然后可以继续绑定使用
  nfs:
    server: 192.168.66.50
    path: /test
```
```
apiVersion: v1
kind: Pod
metadata:
  name: redis111
  labels:
    app: redis111
spec:
  containers:
  - name: redis
    image: redis
    imagePullPolicy: Never
    volumeMounts:
    - mountPath: "/data"
      name: data
    ports:
    - containerPort: 6379
  volumes:
  - name: data
    persistentVolumeClaim:        #指定使用的PVC
      claimName: test-pvc         #名字一定要正确
```
因PVC允许用户小号抽象的存储资源,故提供不同类型、属性、性能的PV就是一个比较常见的需求,此时就能通过StorageClass来提供不同种类的PV资源

PV提供3种访问模式：
* ReadWriteOnce 当前卷可以被一个节点使用读写模式挂载
* ReadOnlyMany 当前卷可以被多个节点使用只读模式挂载
* ReadWriteMany 当前卷可以被多个节点使用读写模式挂载

PV还有3种回收方式：
* Retain(保留)数据,若需此PV被重新使用,需要删除被使用PV并手动清除数据
* Delete(删除)数据,若当前卷支持此回收策略,则PV及相关数据会被删除
* Dynamic Provisioning(动态配置)