
不要通过共享内存来通信，而是通过通信来共享内存。

Go并发，依赖CSP模型，基于channel实现。

输入驱动，并产生输出，供其他processes消费，processes可以是进程、线程、或代码块

大部分语言并发模型是基于线程和内存同步访问控制，Go的并发编程的模型则用goroutine和channel来替代。
