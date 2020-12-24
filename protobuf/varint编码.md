Varint是一种使用一个或多个字节序列化整数的方法，会把整数编码为变长字节。对于32位整型数据经过Varint编码后需要1~5个字节，小的数字使用1个byte，大的数字使用5个bytes。64位整型数据编码后占用1~10个字节。在实际场景中小数字的使用率远远多于大数字，因此通过Varint编码对于大部分场景都可以起到很好的压缩效果。


https://developers.google.cn/protocol-buffers/docs/encoding
https://zhuanlan.zhihu.com/p/84250836

最高有效位（most significant bit - msb）
