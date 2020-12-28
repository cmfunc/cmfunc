# 项目环境配置

## 使用cfssl在本地生成STL证书，并保存至docker-compose.yml中指定的etcd证书地址
```shell script
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=etcd server-csr.json | cfssljson -bare server
```
## 运行etcd、kafka、redis docker-compose
进入docker-compose.yml文件所在目录

运行 docker-compose up -d

# Q&A

## 1. 使用``make proto``生成go文件
如果提示以下错误：

```bash
protoc -I. --proto_path=/Users/huxiu/go/src:./proto/channel --go_out=plugins=grpc:./proto/channel channel.proto
--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
make: *** [proto] Error 1
```

使用以下命令安装``protoc-gen-go``插件,可解决
```bash
go get github.com/golang/protobuf/protoc-gen-go
```

问题原因：
https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable

## 2. 生成``.pb.go``文件，``grpc.ClientConnInterface``未定义问题

```go
// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

```

https://github.com/grpc/grpc-go/issues/3750

https://github.com/grpc/grpc-go

> https://github.com/Eson-Jia
>
> > for google.golang.org/grpc v1.26.0,I usego get github.com/golang/protobuf/protoc-gen-go@v1.3.2to downgrade the version of protoc-gen-go to v.1.3.2,then rebuild proto file, problem has been resolved.