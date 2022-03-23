# channel

 1. 在不改变channel自身状态的情况下，无法获知一个channel是否关闭；
 2. 关闭一个closed channel会导致panic。所以，如果关闭channel的一方在不知道channel是否处于关闭状态时就去贸然关闭channel是很危险的。
 3. 向一个closed channel发送数据会导致panic。向channel发送数据的一方不知道channel是否处于关闭状态时就贸然向channel发送数据是很危险的。

不要从receiver侧关闭channel，不要在有多个sender时，关闭channel；

关闭channel的方式：

 1. Defer + recover
 2. sync.Once保证只关闭一次

channel引发内存泄漏

原因：goroutine操作channel后，处于发送或接收阻塞状态；而channel处于满或空的状态，一直得不到改变。同时，垃圾回收器不会回收此资源，进而导致goroutine会一直处于等待队列中，不见天日。
程序运行过程中，对于一个channel，没有任何goroutine引用后，gc会对其进行回收操作，不会引起内存泄漏。

Channel发送和接收元素的本质：“值的拷贝”，从sender goroutine的栈到chan buf，还是chan buf到recevier goroutine，或者直接从sender goroutine 到recevier goroutine。

带缓冲的channel，for range带缓冲的channel，程序会一直阻塞在for range循环中，无法退出，必须对channel进行close()，才能退出；

对closed的chan发送数据会panic；
重复关闭chan会panic；

Channels act as first-in-first-out queues.
if one goroutine sends values on a channel and a second goroutine receives them, the values are received in the order sent.

Sending to or closing a closed channel causes a run-time panic. Closing the nil channel also causes a run-time panic.

channel实现协程池

```go
var limit = make(chan int, 3)

func main() {
 for _, w := range work {
  go func(w func()) {
   limit <- 1
   w()
   <-limit
  }(w)
 }
 select{}
}
```

nil的channel永久阻塞
