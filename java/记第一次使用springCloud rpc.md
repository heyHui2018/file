###0、string类相等用.equals()，而不是"=="，因为"=="比较的是两者是否一致，指针的话比较的是地址而不是指向的值
###1、在application.properties文件中，当有配置有黄色下划线时，不代表此配置无效
###2、用feign进行rgc微服务访问时，需配置超时时间，默认的比较短；还需配置线程池并发数量
ribbon.ConnectTimeout=1000
ribbon.ReadTimeout=9000
hystrix.command.default.execution.timeout.enabled=true
hystrix.command.default.execution.isolation.thread.timeoutInMilliseconds=10000

hystrix.threadpool.default.coreSize=500
hystrix.threadpool.default.queueSizeRejectionThreshold=500
hystrix.command.default.circuitBreaker.forceClosed=true
###3、在使用eureka集群时，不可在application.java文件中设置负载均衡策略，否则会造成找不到节点的问题
###4、在打包时需将配置文件剔除，以免测试环境或现网环境运行时仍然读取的是开发环境的配置，同时需修改pom文件最后的build标签中的配置，还需在工程的src下
新增assembly文件夹及同名xml文件
###5、反射获取不到父类中的私有变量
###6、ResultVO<T>中的T即泛型，在不同的方法中可设置成不同的类型
###7、通过feign进行微服务调用设置fallback时，需implements FallbackFactory<T>（此处的T即为通过feign进行微服务调用的类），并重写create方法来输出相关日志
###8、hscan命令中参数count指的是扫描的数量，而非想要匹配的数量
###9、当某接口在极短时间内通过相同的参数被连续请求时，第一次应设置锁，在结束时解锁，后续请求应循环尝试锁定，锁定后查看第一次的执行结果，成功时返回已成功，
否则再次执行（因这类情况应是幂等，故在第一次成功时返回已成功，在第一次失败时再次执行）
###10、redis set通过hash表实现，添加/删除/查找的复杂度都为O(1)