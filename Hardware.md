# 硬件

## 计算机同步原语

### CAS

```go

func transferCas(balance *int32, amount int, done chan bool) {
  for {
    old := atomic.LoadInt32(balance)
    new := old + int32(amount)
    if atomic.CompareAndSwapInt32(balance, old, new) {
      break
    }
  }
  done <- true
}

```

### FAA

```go

func transferFaa(balance *int32, amount int, done chan bool) {
  atomic.AddInt32(balance, int32(amount))
  done <- true
}

```

## CPU

自旋对应于CPU的"PAUSE"指令，CPU对该指令什么都不做，相当于CPU空转，对程序而言相当于sleep了一小段时间，时间非常短，当前实现是30个时钟周期。