# context

https://qcrao.com/2019/06/12/dive-into-go-context/

Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.

Failing to call the CancelFunc leaks the child and its children until the parent is canceled or the timer fires. The go vet tool checks that CancelFuncs are used on all control-flow paths.

The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.

https://golang.google.cn/pkg/context/

goroutine的上下文，包含goroutine的运行状态、环境、现场等信息；用来在goroutine之间传递上下文信息，包括：取消信号、超时时间、截止时间、k-v等。

context成为并发控制和超时机制的标准做法。
context.Context类型的值可以协调多个goroutine中的代码执行“取消操作”，可以存储键值对，并发安全。

go语言的server是一个协程模型，一个协程处理一个请求。

go中不能直接杀死协程，协程的关闭一般通过channel+select方式来控制。


幂等，连续多次调用同一个方法，得到的结果相同。
