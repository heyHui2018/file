###1.查看内存使用情况
```
    free -m
```
```
             total       used       free     shared    buffers     cached
Mem:          3769       2210       1559          1        251        421
-/+ buffers/cache:       1536       2232 
Swap:          511          0        511 
```
Mem行（单位均为M）： 
* total：内存总数 
* used：已使用内存数 
* free：空闲内存数 
* shared：当前废弃不用 
* buffers：缓存内存数（Buffer） 
* cached：缓存内舒数（Page）

(-/+ buffers/cache)行： 
* （-buffers/cache）: 真正使用的内存数，指的是第一部分的 used - buffers - cached 
* （+buffers/cache）: 可用的内存数，指的是第一部分的 free + buffers + cached

Swap行： 指交换分区

实际上不要看free少就觉得内存不足，buffers和cached都是可以在使用内存时拿来用的，应该以(-/+ buffers/cache)行的free和used来看。
只要没发现swap的使用，就不用太担心，如果swap用了很多，那就要考虑增加物理内存了。

###2.查看CPU使用情况
```
    top
```

###3.查看当前监听的端口
```
    netstat -lntp
```