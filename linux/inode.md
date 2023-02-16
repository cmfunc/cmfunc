# 文件存储

文件存储在硬盘上，硬盘的最小存储单位叫“扇区”（Sector）。每个扇区512字节（相当于0.5kb）。
操作系统读取硬盘时，为了提升效率，会一次性连续读取多个扇区（“块”Block），是文件存取的最小单位。块的大小，最常见为4KB，即连续八个sector组成一个block。

文件信息存储在“块”中，文件的元信息（创建者、文件的创建日期、文件的大小等），存储在inode（索引节点）区域。
每个文件都有inode，里面包含了与该文件有关的一些信息。

## Linux文件系统如何通过i节点把文件的逻辑结构和物理结构转换？

inode节点是一个64字节长的表，表中包含了文件的相关信息，其中有文件的大小、文件的所有者、文件的存取许可方式以及文件的类型等重要信息。在inode节点表中最重要的内容是磁盘地址表。磁盘地址表中有13个号块，文件将以号块在磁盘地址表中出现的顺序依次读取相应的块。

Linux文件系统通过把inode节点和文件名进行连接，当需要读取文件时，文件系统在当前目录表中查找该文件名对应的项，由此得到该文件相对应的inode节点号，通过该inode节点的磁盘地址表把分散存放的物理文件块连接成文件的逻辑结构。

## 硬连接与软连接

### 硬链接

Linux下的文件通过索引节点（inode）来识别文件，硬链接可以认为是一个指针，指向文件索引节点的指针，系统并不为它重新分配inode。每添加一个硬链接，文件的链接数就加1。

1. 不可以在不同的文件系统的文件间建立链接
2. 只有超级用户才可以为目录创建硬链接

### 软链接

软链接克服了硬链接的不足，没有任何文件系统的限制，任何用户可以创建指向目录的符号链接，可以跨越不同机器、不同网络文件进行链接；

1. 软链接有原文件的路径信息，所以当原文件从一个目录下移到其他目录中，再访问链接文件，系统就找不到了，硬链接没有这个缺陷；
2. 软链接需要系统分配额外的空间用于建立新的索引节点和保存原文件的路径；

硬链接不可跨分区，软链接可以跨分区；
硬链接指向一个inode节点，而软链接是创建一个新的inode节点；
删除硬链接文件，不会删除原文件，删除软链接文件，会把原文件删除；

## RAID 独立磁盘冗余阵列

基本思想：将多个相对便宜的的硬盘组合起来，成为一个硬盘阵列，使性能达到甚至超过一个价格昂贵、容量巨大的硬盘。RAID通常被用在服务器电脑上，使用完全相同的硬盘组成一个逻辑扇区，因此操作系统只会把它当作一个硬盘。

