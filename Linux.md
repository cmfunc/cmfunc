# Linux

## 1. shell 命令

## 1.1 进程有关

通过程序监听的端口号查找程序

```bash
lsof -i :2379
```

杀死指定pid的进程

```bash
kill -9 16067
```

## 1.2 文件有关

在指定目录下查找文件

```sh
find . -iname "Channel*"
```

在当前目录下，查找名称以Channel开头的文件
