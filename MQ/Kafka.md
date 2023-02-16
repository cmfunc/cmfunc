# 消息

Kafka的消息是字节数组

## 主题与分区

消息通过主题分类，

## 处理批量消息

Kafka内部，消息都是以“批”为单位处理的，一批消息从发送端到接收端；
Kafka的客户端SDK在实现消息发送逻辑时，采用了异步批量发送的机制；

当你send()一条消息之后，无论是同步发送还是异步发送，Kafka都不会立即将这条消息发送出去。这条消息会先被存放在内存中缓存起来，然后选择合适的时机把缓存中的所有消息组成一批，一次性发给Broker。简单说，就是攒一波一起发。

在Kafka的服务端，也不会把一批消息再还原成多条消息，再一条一条处理。Kafka将每批消息都当作一个“批消息”来处理，在Broker整个处理流程中，无论是写入磁盘、从磁盘读出来、还是复制到其他这些流程中，批消息都不会被解开，一直作为一条“批消息”来处理。

在消费时，消息同样时以批为单位进行传递，Consumer从Broker拉到一批消息后，在客户端把消息解开，再一条一条交给用户代码处理。

构建批消息和解开批消息分别在发送端和消费端的客户端完成，不仅减轻了Broker的压力，最重要的是减少了Broker处理请求的次数，提升总体的处理能力。

## 使用顺序读写提升磁盘IO性能

对于磁盘来说，顺序读写的邢恩那个要远远好于随机读写。SSD上，顺序读写的性能要比随机读写快几倍，如果是机械硬盘，差距会达到几十倍。

操作系统每次从磁盘读写数据时，需要先寻址，找到数据在磁盘的物理位置，然后进行数据读写。机械硬盘的寻址时间较长，因为它要移动磁头，是一个机械运动，机械硬盘工作的时候会发出咔咔声，就是移动磁头的声音。

顺序读写比随机读写省去了大部分的寻址时间，只要寻址一次，就可以连续地读写下去。Kafka充分利用磁盘的顺序读写特性，对于每个分区，把从Producer收到的消息，顺序地写入log文件中，一个文件写满了，就开启一个新的文件顺序写下去。消费的时候，也是从某个全局的位置开始，某个log文件中的某个位置开始，顺序地把消息读出来。

## 利用PageCache加速消息读写

PageCache时现代操作系统具有的一项基本特性，操作系统在内存中给磁盘上的文件建立的缓存。无论使用什么语言，在调用系统的API去读写文件的时候，并不会直接去读写磁盘上的文件，应用程序实际操作的都是PageCache，也就是文件在内存中缓存的副本。

### LRU算法

优先清除最近最少使用的数据。

## ZeroCopy：零拷贝技术

可以不将PageCache的数据拷贝到应用程序的内存空间中。因为在从文件读取数据后再通过网络发送出去的过程中，不涉及到要对文件数据进行处理，那一定要用到零拷贝的方法，有效提升性能。

## Kafka高性能设计关键技术点

- 使用批量处理的方式来提升系统吞吐能力
- 基于磁盘文件高性能顺序读写的特性来设计的存储结构
- 利用操作系统的PageCache来缓存数据，减少IO并提升读性能
- 使用零拷贝技术加速消费流程