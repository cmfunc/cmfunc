String类型提供“一个键对应一个值的数据”，String支持保存二进制字节流。String保存数据时消耗的内存空间较多。

Redis底层数据结构，压缩列表(ziplist)，非常省内存；
Redis基于压缩列表实现了List、Hash和Sorted Set集合类型；
