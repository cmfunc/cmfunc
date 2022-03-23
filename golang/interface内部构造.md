# interface内部构造

golang的多态特点，从语法上并不明显；

发生多态的几个要素：

- 有interface接口，且有定义接口的方法；
- 有子类去重写interface的接口；
- 有父类指针指向子类的具体对象；

多态效果：父类指针可以调用子类的具体方法；

## interface 内部结构

### 空接口``interface{}``用``eface`` ``struct``表示

```go
type eface struct{ //空接口
    _type *_type  //类型信息
    data unsafe.Pointer //指向数据的指针（类似C语言的void*）
}
```

### 非空接口``interface{function()}``用``iface`` ``struct``表示

```go
type iface struct{
    tab *itab // 包含了interface的一些关键信息
    data unsafe.Pointer
}
```

### *interface{}是什么

因为interface底层是struct，所以interface{}也是可以被取指针的；
