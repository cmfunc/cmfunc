https://www.pythonf.cn/read/147330

https://segmentfault.com/a/1190000038241863

https://segmentfault.com/a/1190000021250088

从实现方式上看
python的协程通过事件循环监听任务来实现，而go还是通过调用系统的线程来实现并发。
python的协程与线程是N:1的关系，而go是M:N的关系
从生态上看 go天生支持并发 而python的异步库使用还存在很大限制，数据库，网络请求等库都需要特定的异步库来实现。

