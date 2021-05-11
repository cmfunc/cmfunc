https://dev.mysql.com/doc/refman/5.7/en/optimize-overview.html

	• 数据库层面优化
	• 硬件层面优化
	• 可以执行与性能间的平衡

数据库层面优化
	1. 表结构的设计，字段类型选择
	2. 为查询建索引
	3. 为每一张表选择合适的存储引擎

硬件优化
	1. 磁盘搜索。现代磁盘，从磁盘中查找一条数据的耗时低于10ms，理论上1秒钟可以执行100条查找。优化寻道时间是将数据分发到多个磁盘上。
	2. 磁盘读写。现代磁盘，至少可提供10-20MB/s的吞吐量。优化，从多个磁盘上并行读取。
	3. CPU周期。
	4. 内存带宽。当CPU需要的数据超出CPU缓存的容量时，主内存带宽将成为瓶颈。

https://www.docs4dev.com/docs/zh/mysql/5.7/reference/optimization.html

CPU周期、IO操作

典型的用户旨在从其现有的软件和硬件配置中获得最佳的数据库性能。高级用户会寻找机会改进 MySQL 软件本身，或者开发自己的存储引擎和硬件设备来扩展 MySQL 生态系统。

	• 合适的行格式
	• 合适的锁机制
	• InnoDB缓冲池的大小设置是否合理
	
# 索引优化
https://www.docs4dev.com/docs/zh/mysql/5.7/reference/select-optimization.html




