# Mutex实现原理

借助CAS指令 + 自旋 + 信号量

我们使用 Mutex 是为了不同 goroutine 之间共享某个变量, 所以需要让这个变量做到能够互斥, 不然该变量就会被互相被覆盖. Mutex 底层是由 state sema 控制的, 当 Mutex 变量被复制时, Mutex 的 state, sema 当时的状态也被复制走了, 但是由于不同 goroutine 之间的 Mutex 已经不是同一个变量了, 这样就会造成要么某个 goroutine 死锁或者不同 goroutine 共享的变量达不到互斥；
