《孔乙己》粉板与账本；
先写日志，再写磁盘；InnoDB先把记录写到redo log(粉板)里面，并更新内存。

crash-safe
Redo log, InnoDB保证即使数据库发生异常重启

Binlog归档日志，属于server层的日志；所有引擎都可以使用；

Redo log是物理日志，记录的是‘在某个数据页上做了什么修改’；
Bin log是逻辑日志，记录的是这个语句的原始逻辑；
Redo log是循环写的，空间固定会用完；
Binlog是可以追加写入；
追加写是指binlog文件写到一定大小后会切换到下一个，不会覆盖以前的日志；

Redo log写完了，要flush脏页，这种情况下，InnoDB整个系统都不能接受更新。

正确告诉InnoDB所在主机的IO能力，这样InnoDB才能知道需要全力刷脏页的时候，可以刷多快。

innodb_io_capacity参数

innodb_flush_neighbors参数，连坐；机械硬盘开启，SSD设置为0；因SSD中IOPS不是瓶颈，更快地执行完必要的刷脏页操作，减少SQL语句响应时间。
