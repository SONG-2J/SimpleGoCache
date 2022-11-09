## 代码结构
```
SimpleGOCache/
    |--cachepb/
        |--cachepb.pb.go    // cachepb.proto文件转出的go文件
    
    |--consistenthash/
        |--consistenthash.go   // 一致性hash算法
    
    |--lru/
        |--lru.go           // lru 缓存淘汰策略
    
    |--simplegocache/
        |--byteview.go      // 支撑并发读写
        |--cache.go         // 缓存结构体及add和get方法
        |--getter.go        // 接口Getter和Get回调函数（缓存不存在时，调用这个函数，得到源数据）
        |--group.go         // 分布式缓存的主体结构Group
        |--httpgetter.go    // 提供被其他节点访问的能力(基于http)
        |--httppool.go      // http池，self记录自己的地址，basePath是通信地址前缀
        |--peers.go         // 抽象出2个接口，PeerPicker和PeerGetter
    
    |--singleflight/
        |--singleflight.go     // 防止缓存击穿

    |--main.go              // 分布式缓存测试
```

## Group数据结构
```
                           是
接收 key --> 检查是否被缓存 -----> 返回缓存值 ⑴
                |  否                         是
                |-----> 是否应当从远程节点获取 -----> 与远程节点交互 --> 返回缓存值 ⑵
                            |  否
                            |-----> 调用`回调函数`，获取值并添加到缓存 --> 返回缓存值 ⑶
```

## 测试
```shell
    ./run.sh
```
