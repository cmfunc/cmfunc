# 多服务

把复杂的大应用，解耦拆分为几个小的应用；

优势：
有利于团队组织架构的拆分；每个应用独立运维，独立扩容，独立上线，各个应用之间互不影响。

弊端：
服务之间调用关系变得更复杂，平台的整体复杂熵升高，出错的概率、debug问题的难度都会变高。服务治理是微服务的技术重点。

服务治理：就是管理微服务，保证平台整体平稳的运行。涉及内容：鉴权、限流、降级、熔断、监控告警。

## 链路追踪

<https://wu-sheng.gitbooks.io/opentracing-io/content/>

<https://opentelemetry.io/docs/>

<https://www.jaegertracing.io/docs/1.14/>

## 监控

<https://prometheus.io/docs/introduction/overview/>

<https://prometheus.fuckcloudnative.io/>

<https://grafana.com/tutorials/grafana-fundamentals/?pg=docs>

Kubernetes集群监控系统；

## 日志收集

<https://docs.fluentd.org/>

<https://www.elastic.co/guide/en/kibana/current/index.html>

<https://www.elastic.co/guide/en/logstash/current/index.html>

<https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html>

Elasticsearch is the distributed search and analytics engine at the heart of the Elastic Stack. Logstash and Beats facilitate collecting, aggregating, and enriching your data and storing it in Elasticsearch. Kibana enables you to interactively explore, visualize, and share insights into your data and manage and monitor the stack. Elasticsearch is where the indexing, search, and analysis magic happens.

<https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html>

## 熔断与降级

<https://github.com/afex/hystrix-go/>

## 限流

• 固定窗口计数法
• 滑动窗口算法
• 漏桶算法
• 令牌桶算法

固定时间窗口限流算法；

滑动时间窗口限流算法；
循环队列

基于时间窗口的限流算法，不管是固定时间窗口还是滑动时间窗口，只能在选定的时间粒度上限流，对选定时间粒度内的更加细粒度的访问频率不做限制

## 权限控制

### 控制策略

#### acl

#### abac

#### rbac

### impl

#### opa

#### casbin

#### ladon

## Docker

```bash
docker-compass up -d
```

### 基础

<http://www.dockerinfo.net/document>

轻量级操作系统虚拟化解决方案；
Docker基础是Linux容器（LXC）等技术；在LXC基础上，Docker进行了进一步的封装，让用户不需要关心容器的管理，使得操作更简便。

 Docker 和传统虚拟化方式的不同之处：
容器是在操作系统层面上实现虚拟化，直接复用本地主机的操作系统，而传统方式则是在硬件层面实现。

Docker是一款针对程序开发人员和系统管理员来开发、部署、运行应用的一款虚拟化平台。Docker 可以让你像使用集装箱一样快速的组合成应用，并且可以像运输标准集装箱一样，尽可能的屏蔽代码层面的差异。Docker 会尽可能的缩短从代码测试到产品部署的时间。

### Dockerfile运行容器

 1. 通过dockerfile build镜像
<https://www.runoob.com/docker/docker-build-command.html>
docker build -f ./Dockerfile -t article-api001
 2. 运行本地镜像
<https://www.runoob.com/docker/docker-run-command.html>
docker run -p

### 挂载卷

--volume -v

## etcd

 • 服务注册：同一service的所有节点注册到相同的目录下，节点启动后将自己的信息注册到所属服务的目录中。
 • 健康检查：服务节点定时发送心跳，注册到服务目录中的设置信息，设置一个较短的TTL，运行正常的服务节点每隔一段时间会去更新信息的TTL。（定时续租，更新TTL）
 • 服务发现：通过名称能查询到服务提供外部访问的IP和端口号。网关代理及时发现服务中新增节点、丢弃不可用的服务节点，同时各个服务之间也能感知对方的存在。

TTL失效、数据改变监视、多值、目录监听、分布式锁原子操作。

对比consul：
Consul Agent、Consul Server（一主多从）

## istio

<https://istio.io/latest/zh/docs/concepts/what-is-istio/>

将限流、熔断、降低、监控、链路追踪，都放在一个代理层；
代理层对服务做转发，有点类似python中wsgi接口协议，开发人员不用重新开发server（gunicron、uWSGI服务器、fcgi、bjoern），只用开发application部分的代码。
uwsgi也是通信协议，由uWSGI服务器独占。
<https://www.fullstackpython.com/wsgi-servers.html#:~:text=%20WSGI%20servers%20learning%20checklist%20%201%20Understand>,requests%20to%20the%20WSGI%20server%20for...%20More%20

## jenkins

<https://jenkins.io/zh/download/>

<https://www.jenkins.io/zh/doc/tutorials/build-a-python-app-with-pyinstaller/>

使用docker安装Jenkins
*************************************************************
*************************************************************
*************************************************************

Jenkins initial setup is required. An admin user has been created and a password generated.
Please use the following password to proceed to installation:

8b6616b79d7048c481a8ad95c62c9a2f

This may also be found at: /var/jenkins_home/secrets/initialAdminPassword

*************************************************************
*************************************************************
*************************************************************

## protobuf

### message中的编号

protobuf在传输二进制shu数据时，不会放松数据的key名称和类型，只会记录key所对应的编号和value的二进制，当接收方收到数据化，再通过本地的proto文件生成的语言解码结构或算法，通过编号对应相应的key的名称和类型。

### varint编码

Varint是一种使用一个或多个字节序列化整数的方法，会把整数编码为变长字节。对于32位整型数据经过Varint编码后需要1~5个字节，小的数字使用1个byte，大的数字使用5个bytes。64位整型数据编码后占用1~10个字节。在实际场景中小数字的使用率远远多于大数字，因此通过Varint编码对于大部分场景都可以起到很好的压缩效果。

<https://developers.google.cn/protocol-buffers/docs/encoding>
<https://zhuanlan.zhihu.com/p/84250836>

最高有效位（most significant bit - msb）

### 原码、反码、补码

原码、反码、补码是机器存储一个具体数字的编码方式；

原码：符号加上真值的绝对值；8位二进制
反码：正数的反码是其本身，负数的反码是在其原码基础上，符号位不变，其余各个位取反。（？反码如何识别是负数并对其转换为原码）
补码：正数的补码是其本身；负数的补码是在其原码的基础上，符号位不变，其余各位取反，最后+1。（也即在反码的基础上+1）

反码、补码出现的原因：计算机辨别“符号位”会让计算机的基础电路设计变得十分复杂。利用1+（-1）的原理，将负数和正数直接相加；

### 机器数

一个数在计算机中的二进制表示形式，叫做这个数的机器数；
机器数带符号，计算机中用机器数的最高位存放符号，正数为0，负数为1；
最高位1代表负、0代表正；
带符号的机器数对应的真正数值称为机器数的真值；

## 哪些因素会导致网站访问慢？

1. 服务器出口带宽不够用，本身服务器购买的出口带宽比较小
1. 服务器负载过大，导致响应不过来
1. 数据库瓶颈
1. 网站代码

## MQTT

物联网通信协议，构建于TCP/IP协议上。MQTT最大的优点在于，可以以极少的代码和优先的带宽，为连接远程设备提供实时可靠的消息服务。

<https://github.com/eclipse/paho.mqtt.golang>

<https://www.runoob.com/w3cnote/mqtt-intro.html#:~:text=MQTT%E6%98%AF%E4%B8%80%E4%B8%AA%E5%9F%BA%E4%BA%8E%E5%AE%A2%E6%88%B7>,%E8%AE%BE%E5%A4%87%E4%B8%AD%E5%B7%B2%E5%B9%BF%E6%B3%9B%E4%BD%BF%E7%94%A8%E3%80%82

MQTT中有三种身份：发布者（Publish）、代理（Broker）（服务器）、订阅者（Subscribe）。
消息的发布和订阅都是客户端，消息代理是服务器，消息发布者可以同时是订阅者。
MQTT传输的消息分为：主题（Topic）和负载（payload）两部分。
 • topic，消息的类型，订阅者订阅（Subscribe）后，就会收到该主题的消息内容（payload）。
 • payload，消息的内容，指订阅者具体要使用的内容。

MQTT构建有序、无损、基于字节流的双向传输。
当应用数据通过MQTT网络发送时，MQTT会把与之相关的服务质量（QoS）和主题名（Topic）相关连。

## 布隆过滤器

数据结构：bitMap，一个超大的bit数组；
算法：一些列的哈希函数；

某样东西一定不存在，或可能存在；

布隆过滤器是一个初值都是0的bit数组和N个哈希函数组成。

一定不存在

## 负载均衡算法

 1. 轮询
 2. 随机
 3. 加权轮询

在同一个客户端上，再一次绘画中的所有请求都路由到同一个服务器；

## 缓存淘汰策略

### 先进先出策略FIFO

first in, first out

### 最少使用策略LFU

Least frequently used

### 最近最少使用策略LRU

least recently used
