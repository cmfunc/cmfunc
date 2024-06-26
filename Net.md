# NET

## 通过网络传输结构化的数据

想要使用网络框架的API来传输结构化的数据，必须得先实现结构化的数据与字节流之间的双向转换。

TCP连接上，传输数据的基本形式是二进制流（一段一段的1和0）；一般编程语言或网络框架提供的API中，传输数据的基本形式是字节Byte，一个字节是8个二进制位，8个Bit。

**序列化：**将结构化数据转换成字节流的过程；

**反序列化：**将字节流转换成结构化数据的过程；

序列化的用途：

1. 用于在网络上传输数据
2. 将结构化数据保存在文件中

好的序列化实现具备的特点：

1. 序列化后的数据最好是易于人类阅读
1. 实现的复杂度足够低
1. 序列化和反序列化的速度越快越好
1. 序列化后的信息密度越大越好，也就是序列化后占用的存储空间越小越好

实现高性能的序列化与反序列化:将结构化数据转化成字节流或二进制流;
