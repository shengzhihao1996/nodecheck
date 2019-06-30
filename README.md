## if you want to build image, please install docker, you must have access to an external network, then use: 

```
make install
```

###### 间隔10s探测节点存活，使用icmp报文、http请求和tcp链接三种方法测试，四次轮询，满足条件驱逐节点，四次之内，只要有一次连接正常，计数器清零。
###### 计数器为节点标签test

