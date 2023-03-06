# Golang

## 数据结构

### slice

golang slice data[:6:8] 两个冒号的理解

常规slice , data[6:8]，从第6位到第8位（返回6， 7），长度len为2， 最大可扩充长度cap为4（6-9）

另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8

a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x

### map

The comparison operators == and != must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice.

<https://golang.google.cn/ref/spec#Comparison_operators>

### interface

interface内部构造:

golang的多态特点，从语法上并不明显；

发生多态的几个要素：

- 有interface接口，且有定义接口的方法；
- 有子类去重写interface的接口；
- 有父类指针指向子类的具体对象；

多态效果：父类指针可以调用子类的具体方法；

#### interface 内部结构

##### 空接口``interface{}``用``eface`` ``struct``表示

```go
type eface struct{ //空接口
    _type *_type  //类型信息
    data unsafe.Pointer //指向数据的指针（类似C语言的void*）
}
```

##### 非空接口``interface{function()}``用``iface`` ``struct``表示

```go
type iface struct{
    tab *itab // 包含了interface的一些关键信息
    data unsafe.Pointer
}
```

##### *interface{}是什么

因为interface底层是struct，所以interface{}也是可以被取指针的；

### error

Go语言中，错误被认为是一种可以预期的结果；异常是一种非预期的结果，发生异常表示程序存在bug或发生其它不可控的问题；
Go语言推荐使用recover函数将内部异常转为错误处理；

Go语言库的实现习惯: 即使在包内部使用了panic，但是在导出函数时会被转化为明确的错误值。

Go语言错误是一种接口类型，接口类型包含原始类型和原始的值。只有当接口的类型和原始的值都为空的时候，接口值才为nil；接口中类型为空时，原始值必然为空；反之，不成立；

必须要和有异常的栈帧只隔一个栈帧，recover函数才能正常捕获异常。换言之，recover函数捕获的是祖父一级调用函数栈帧的异常（刚好可以跨越一层defer函数）！

## 内置方法

### copy

copy(dst, src []T) int
copy(dst []byte, src string) int

var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")

### panic

Calling a nil function value causes a run-time panic.

如果运行时移位计数为负数，则会发生运行时死机。

For an operand x of pointer type *T, the pointer indirection*x denotes the variable of type T pointed to by x. If x is nil, an attempt to evaluate *x will cause a run-time panic.

A receive operation on a closed channel can always proceed immediately, yielding the element type's zero value after any previously sent values have been received.

### select

<https://lessisbetter.site/2018/12/17/golang-selete-advance/>

If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed.

### make与new

slice使用make，引用类型使用make；
array使用new，值类型使用new；

The built-in function make takes a type T, which must be a slice, map or channel type, optionally followed by a type-specific list of expressions.

make()：对应map、slice、channel都是引用类型；在函数内部修改值后，外部的值也会跟着改变。

### append

If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, append re-uses the underlying array.

### recover

```go
func protect(g func()) {
 defer func() {
  log.Println("done")  // Println executes normally even if there is a panic
  if x := recover(); x != nil {
   log.Printf("run time panic: %v", x)
  }
 }()
 log.Println("start")
 g()
}
```

### delete

If the map m is nil or the element m[k] does not exist, delete is a no-op.

## 范型

泛型需求： 创建一组操作相同或类似的算法，这些算法与数据类型无关，不管什么数据类型只要符合要求就可以操作。

静态类型：在编译期就能确定的变量、表达式的数据类型，编译期就能确定某个类型的内存布局。

动态类型：编译期间无法确定某个变量、表达式的具体类型。

动态类型语言不关心数据的具体类型是什么，即使没有泛型也可以写出类似泛型的代码。

泛型可以减少重复的代码，确保代码的类型安全。

获取go接口类型的值数据，需要类型断言，类型断言是在运行时进行的，如果断言类型错误会导致panic。真正的泛型可以在编译期发现这类错误，而不是在运行时。

## GC

垃圾回收，一种自动内存管理的机制。当程序向操作系统申请的内存不再需要时，垃圾回收主动将其回收供其他代码进行内存申请时候复用，或者将其归还给操作系统，这种针对内存级别资源的自动回收过程，称为垃圾回收。

负责垃圾回收的程序组件，即为垃圾回收器。

程序需要进行特殊优化时，通过提供可调控的API，对gc的运行时机、运行开销进行把控。

垃圾回收器的执行过程被划分为两个半独立的组件：

 1. 赋值器：用户态的代码，只是在修改对象之间的引用关系，也就是在对象图（对象之间引用关系的一个有向图）进行操作。
 2. 回收器：负责执行垃圾回收的代码。

根对象：在垃圾回收术语中叫做根集合，是垃圾回收器在标记过程时最先检查的对象，包括：

 1. 全局变量：程序在编译期就能明确的那些存在于程序整个生命周期的变量。
 2. 执行栈：每个goroutine都包含自己的执行栈，包含栈上的变量及指向分配的堆内存区块的指针。
 3. 寄存器：寄存器的值可能表示一个指针，参与计算的这些指针可能指向某些赋值器分配的堆内存区块。

常见gc实现方式：

 1. 追踪（tracing）：从根对象出发，根据对象之间的引用信息，一步步推进直到扫描完毕整个堆并确定需要保留的对象，从而回收所有可回收的对象。Go、Java、V8的实现均为追踪式GC。
 2. 引用计数（reference counting）：每个对象自身包含一个被引用的计数器，当计数器归零时自动得到回收。该方法缺陷较多，追求高性能时通常不被采用。Python、Objective-C为引用计数式GC。

追踪式：
 • 标记清扫：从根对象出发，将确定存活的对象进行标记，并清扫可以回收的对象。
 • 标记整理：为了解决内存碎片问题提出，在标记过程中，将对象尽可能整理到一块连续的内存上。
 • 增量式：将标记与清扫的过程分批执行，每次执行很小的部分，从而增量的推进垃圾回收，达到近似实时、几乎无停顿的目的。
 • 增量整理：在增量式的基础上，增加对对象的整理过程。
 • 分代式：将对象根据存活时间的长短进行分类，存活时间小于某个值的年轻代，存活时间大于某个值的为老年代，永远不会参与回收的对象为永久代。并根据分代假设（如果一个对象存活时间不长泽倾向于被回收，如果一个对象已经存活很长时间则倾向于存活更长时间）对对象进行回收。

引用计数：
根据对象自身的引用计数来回收，当引用计数归零时立即回收。

Go的GC目前使用的是无分代（对象没有代际之分）、不整理（回收过程中不对对象进行移动与整理）、并发（与用户代码并发执行）的三色标记清扫算法。

选择三色标记的原因：

 1. 对象整理的优势是解决内存碎片问题以及“允许”使用顺序内存分配器。Go运行时的分配算法基于tcmalloc，基本上没有碎片问题。并且顺序内存分配器在多线程的场景下并不使用。go使用基于tcmalloc的现代内存分配算法，对对象进行整理不会带来实质性的性能提升。
 2. 分代GC依赖分代假设，即GC将主要的回收目标放在新创建的对象上（存活时间短，更倾向于被回收），而非频繁检查所有对象。Go的编译器会通过逃逸分析将大部分新生对象存储在栈上（栈直接被回收），只有那些需要长期存在的对象才会被分配到需要进行垃圾回收的堆中。分代GC回收的那些存活时间短的对象在Go中是直接被分配到栈上，当goroutine死亡后，栈会被直接回收，不需要GC的参与，所以分代假设并没有带来直接优势。  Go的垃圾回收器与用户代码并发执行，使得STW的时间与对象的代际、对象的size没有关系。Go团队更关注如何更好地让GC与用户代码并发执行（使用适当的CPU来执行垃圾回收），而非减少停顿时间上。

三色标记

核心：对象的三色抽象以及波面（wavefront）推进。
三色抽象是一种描述追踪式回收器的方法，标记清扫的垃圾回收。

三色抽象规定三种不同类型的对象：
 • 白色对象（可能死亡）：未被回收器访问到的对象。在回收开始阶段，所有对象均为白色，当回收结束后，白色对象不可达。
 • 灰色对象（波面）：已被回收器访问到的对象，但回收器需要对其中的一个或多个指针进行扫描，因为他们可能还指向白色对象。
 • 黑色对象（确定存活）：已被回收器访问到的对象，其中所有字段都已被扫描，黑色对象中任何一个指针都不可能直接指向白色对象。

垃圾回收开始时，只有白色对象。随着标记过程开始进行，灰色对象开始出现（着色），这时候波面便开始扩大。当一个对象所有子节点均完成扫描时，会被着色为黑色。当整个堆遍历完成时，只剩下黑色和白色对象，这时的黑色对象为可达对象，即存活；白色对象为不可达对象，即死亡。过程中，将灰色对象视为波面，将黑色对象和白色对象分离，使波面不断向前推进，直到所有可达的灰色对象都变为黑色对象为止的过程。

STW

Stw在垃圾回收过程中为了保证实现的正确性、防止无止境的内存增长等问题而不可避免的需要停止赋值器进一步操作对象图的一段过程。

Stw已被优化到半毫秒级别以下。

Runtime.GC()，需要通知并让所有用户态代码停止，但for{}所在goroutine永远不会被中断，从而始终无法进入STW阶段。Go1.14后，这些goroutine会被异步的抢占。

内存泄漏
预期的能很快被释放的内存由于附着在长期存活的内存上、或生命周期意外地被延长，导致预计能够立即回收的内存而长时间得不到回收。
 • 预期能被快速释放的内存因被根对象引用而没有得到迅速释放；
 • Goroutine泄漏：goroutine需要维护执行用户代码的上下文信息，如果程序不断产生新的goroutine、且不结束已经创建的goroutine并复用这部分内存，会造成内存泄漏。
 • Channel 作为一种同步原语，会连接两个不同的 goroutine，如果一个 goroutine 尝试向一个没有接收方的无缓冲 channel 发送消息，则该 goroutine 会被永久的休眠，整个 goroutine 及其执行栈都得不到释放

如何保证标记与清除过程的正确性？

写屏障：
三色标记算法的强弱不变形和赋值器的颜色；写屏障是在并发垃圾回收器中出现的概念，垃圾回收器的正确性体现，不应出现对象的丢失，也不应错误的回收还不需要回收的对象。

为了保证强弱三色不变性的并发指针更新操作，需要通过赋值器屏障技术来保证指针的读写操作一致性。
Go中的写屏障、混合写屏障，其实是指赋值器的写屏障，赋值器的写屏障作为一种同步机制，使赋值器在进行指针写操作时，能够“通知”回收器，进而不破坏弱三色不变性。
Dijkstra插入屏障、Yuasa删除屏障。

触发GC的时机：

 1. 主动触发，通过调用 runtime.GC 来触发 GC，此调用阻塞式地等待当前 GC 运行完毕。
 2. 被动触发，分为两种方式：
  ○ 使用系统监控，当超过两分钟没有产生任何 GC 时，强制触发 GC。
  ○ 使用步调（Pacing）算法，其核心思想是控制内存增长的比例。
GC 调优时，通常是指减少用户代码对 GC 产生的压力，这一方面包含了减少用户代码分配内存的数量（即对程序的代码行为进行调优），另一方面包含了最小化 Go 的 GC 对 CPU 的使用率（即调整 GOGC）。

## 并发

并发体系的理论：CSP（通讯顺序进程）；
Go并发编程核心CSP理论的核心概念是：同步通讯；

Go语言中常见的并发模式：
并发不是并行。并发更关注程序的设计层面，并发的程序完全可以顺序执行，只有在真正的多核CPU上才能真正地同时执行。
并行执行更关注的是程序的运行层面，并行一般是简单的大量重复。

并发编程中，对资源的正确访问需要精确的控制，go将共享的值通过channel传递（多个独立执行的线程很少主动共享资源）；go并发编程哲学：不要通过共享内存来通讯，而应通过通讯来共享内存；通过channel来传值是go语言推荐的做法。

并发编程的核心是同步通信，同步方式有多种；
对一个未加锁的sync.Mutex进行解锁，会导致panic；

### CSP 含义

不要通过共享内存来通信，而是通过通信来共享内存。

Go并发，依赖CSP模型，基于channel实现。

输入驱动，并产生输出，供其他processes消费，processes可以是进程、线程、或代码块

大部分语言并发模型是基于线程和内存同步访问控制，Go的并发编程的模型则用goroutine和channel来替代。

### goroutine

<https://www.pythonf.cn/read/147330>

<https://segmentfault.com/a/1190000038241863>

<https://segmentfault.com/a/1190000021250088>

从实现方式上看
python的协程通过事件循环监听任务来实现，而go还是通过调用系统的线程来实现并发。
python的协程与线程是N:1的关系，而go是M:N的关系
从生态上看 go天生支持并发 而python的异步库使用还存在很大限制，数据库，网络请求等库都需要特定的异步库来实现。

### channel

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

### GPM 调度器

[刘丹冰博客](<https://www.jianshu.com/p/fa696563c38a>)

#### 调度器的由来

CPU调度器、时间片、进程；

CPU调度切换的是进程和线程，多线程开发设计要考虑同步竞争、锁等问题；

进程虚拟内存会占用4GB[32位操作系统],线程大约4MB；

线程分为“内核态”线程和“用户态”线程；一个“用户态线程”必须绑定一个“内核态线程”，CPU并不知道“用户态线程”的存在，CPU只知道它运行的是一个“内核态线程”（Linux的PCB进程控制块）。

多个goroutine绑定一个或多个线程（thread）上。

goroutine与系统线程是M:N的关系。线程由CPU调度是抢占式的，协程由用户态调度是协作式的，一个协程让出CPU后，才执行下一个协程。一个goroutine只占几kb，goroutine栈的大小是可以伸缩的，runtime会自动为goroutine分配，可以在有限的空间内支持大量的goroutine，支持更多的并发。

#### GMP模型及设计思想

##### 2012年前的goroutine调度器

G：goroutine协程

M：thread线程

全局goroutine队列；锁；系统线程；

老调度器的缺点：

1. 创建、销毁、调度G都需要每个M获取锁，这就形成了激烈的锁竞争；
1. M转移G会造成延迟和额外的系统负载；
1. 系统调用（CPU在M之间的切换）导致频繁的线程阻塞和取消阻塞操作增加了系统开销；

##### 2012年的GMP

引入了P（Processor），它包含了运行goroutine的资源，如果线程想运行goroutine，必须先获取P，P中还包含了可运行的G队列。

在Go中，线程是运行goroutine的实体，调度器的功能是把可运行的goroutine分配到工作线程上。

```go
GOMAXPROCS
```

1. 全局队列：存放等待运行的G。
1. P的本地队列：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建G‘时，G’优先加入到P的本地队列，如果队列满了，则会把本地队列中的一半移动到全局队列。
1. P列表：所有的P都在程序启动时创建，并保存在数组中，最多有GOMAXPROCS（可配置）个。
1. M：线程想运行任务就得获取P，从P的本地队列获取G，P队列为空时，M也会尝试从全局队列拿一批G放在P的本地队列，或从其他P的本地队列偷一半放到自己P的本地队列。M运行G，G执行之后，M会从P获取下一个G，不断重复下去。

Goroutine调度器和OS调度器是通过M结合起来的，每个M都代表了1个内核线程，OS调度器负责把内核线程分配到CPU的核上运行。

P的数量：由启动时环境变量$GOMAXPROCS或者是由runtime的方法GOMAXPROCS决定。这意味着程序执行的任意时刻都只有$GOMAXPROCS个goroutine在同时运行。

M的数量：go语言本身的限制，go程序启动时，会设置M的最大数量，默认10000.但是内核很难支持这么多线程数，所以这个限制可以忽略。runtime/debug中的SetMaxthread函数，设置M的最大数量。一个M阻塞了，会创建新的M。

P何时创建：在确定了P的最大数量n后，运行时系统会根据这个数量创建n个P。

M何时创建：没有足够的M来关联P并运行其中的可运行的G。比如所有的M此时都阻塞住了，而P中还有很多就绪任务，就会去寻找空闲的M，而没有空闲的，就会去创建新的M。

##### 调度器的设计策略

复用线程：避免频繁的创建、销毁线程，而是对线程的复用。

1. work stealing机制：当本线程无可运行的G时，尝试从其他线程绑定的P偷取G，而不是销毁线程。
1. hand off机制：当M因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的闲扯给

#### 调度场景过程

### 内存模型

<https://golang.google.cn/ref/mem>

go内存模型，规定了保证一个goroutine读取被其他goroutine写的变量的条件；

程序中，被多个goroutine同时修改的数据，必须保证多个goroutine串行访问这个数据；保证串行化访问数据的方式有：channel、sync和sync\atomic。

多线程、消息传递

Go语言基于消息并发模型，基于CSP模型；
goroutine之间共享内存；

goroutine和系统线程
Goroutine是一种轻量级的线程，goroutine和系统线程不等价；两者之间的区别实际上只是一个量的区别；
每个系统线程会有固定大小的栈(默认是2MB)，用来保存函数递归调用时参数和局部变量；固定大小的栈会导致对于只需要很小的栈空间的线程内存空间的巨大浪费，对于少数需要巨大栈空间的线程来说面临栈溢出的风险。
Goroutine以很小的栈启动(2kb或4kb)，当遇到深度递归导致当前栈空间不足时，goroutine会根据需要动态地伸缩栈的大小（主流实现中栈的最大值可达1GB）。启动代价小，所以可以轻易启动成千上万个goroutine。

Go运行时包含了自己的调度器，调度器可以在n个操作系统线程上多工调度m个goroutine。go调度器的工作和内核的调度是相识的，但go调度器只关注单独的Go程序中的goroutine。
Goroutine采用半抢占式（？？？）的协程调度，只在当前goroutine发生阻塞时才会导致调度；同时发生在用户态，调度器会根据具体函数只保存必要的寄存器，切换的代价比系统线程低很多。
runtime.GOMAXPROCS变量，用于控制当前正常非阻塞goroutine的系统线程数目。

原子操作
原子操作是并发编程中“最小的且不可并行化”的操作。
原子操作对于多线程并发编程模型来说，不会发生有别于单线程的意外情况，共享资源的完整性得到保证。
原子操作一般通过“互斥”访问来保证，通常由特殊的CPU指令提供保护。

用互斥锁保护一个数值型的共享资源，效率低下。
sync/atomic包对原子操作提供了丰富的支持。

顺序一致性模型
同一个goroutine线程内部，顺序一致性内存模型是得到保证的；在不同goroutine之间，并不满足顺序一致性内存模型，需要通过明确定义的同步事件来作为同步的参考。
如果两个事件不可排序，那么这两个事件是并发的。
为了最大化并行，Go语言的编译器和处理器在不影响上述规定的前提下，会对执行语句重新排序。
如果两个并发程序无法确定事件的顺序关系，那么程序的运行结果往往会得到不确定的结果。
Go语言规范（？？？），main函数退出时程序结束，不会等待任何后台线程；goroutine的执行和main函数的返回事件是并发的；

通过同步原语给两个事件明确排序；goroutine 中向chan中发送数据，main中接收chan；不带缓冲的chan会阻塞main；

基于channel的通讯
在无缓存的channel上每一次发送操作都有与其对应的接收操作相配对，发送和接收操作通常发生在不同的goroutine上，在同一个goroutine上执行2个操作很容易导致死锁。
无缓存的channel上的发送操作总在对应的接收操作完成前发生；

在关闭channel后继续从中接收数据，接收者会收到该channel返回的零值；
根据控制channel的缓存大小来控制并发执行的goroutine的最大数目；
select{}是一个空的管道选择语句，会导致goroutine阻塞；
调用os.Exit(0)可以让程序正常退出；

#### 内存分配

<https://deepu.tech/memory-management-in-golang/>

<https://deepu.tech/memory-management-in-programming/>

<https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html>

Go程序进程由操作系统分配一些虚拟内存，这部分虚拟内存是Go程序可以访问的全部内存，虚拟内存中实际正在使用的内存称为Resident Set（驻留内存）。

<https://tonybai.com/2020/03/10/visualizing-memory-management-in-golang/>

#### happens before

内存重排会打乱预期的代码执行顺序；

To specify the requirements of reads and writes, we define happens before, a partial order on the execution of memory operations in a Go program. If event e1 happens before event e2, then we say that e2 happens after e1. Also, if e1 does not happen before e2 and does not happen after e2, then we say that e1 and e2 happen concurrently.

When multiple goroutines access a shared variable v, they must use synchronization events to establish happens-before conditions that ensure reads observe the desired writes.

Program initialization runs in a single goroutine, but that goroutine may create other goroutines, which run concurrently.

If a package p imports package q, the completion of q's init functions happens before the start of any of p's.

The start of the function main.main happens after all init functions have finished.

The go statement that starts a new goroutine happens before the goroutine's execution begins.
goroutine的创建发生在goroutine的执行之前。

channel communication
Channel communication is the main method of synchronization between goroutines. Each send on a particular channel is matched to a corresponding receive from that channel, usually in a different goroutine.

The closing of a channel happens before a receive that returns a zero value because the channel is closed.

The kth receive on a channel with capacity C happens before the k+Cth send from that channel completes.

This program starts a goroutine for every entry in the work list, but the goroutines coordinate using the limit channel to ensure that at most three are running work functions at a time.

```golang
//协程池
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

mutex

A single call of f() from once.Do(f) happens (returns) before any call of once.Do(f) returns.
