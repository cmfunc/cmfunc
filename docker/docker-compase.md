```bash
docker-compass up -d
```

# 基础
http://www.dockerinfo.net/document

轻量级操作系统虚拟化解决方案；
Docker基础是Linux容器（LXC）等技术；在LXC基础上，Docker进行了进一步的封装，让用户不需要关心容器的管理，使得操作更简便。

 Docker 和传统虚拟化方式的不同之处：
容器是在操作系统层面上实现虚拟化，直接复用本地主机的操作系统，而传统方式则是在硬件层面实现。

Docker是一款针对程序开发人员和系统管理员来开发、部署、运行应用的一款虚拟化平台。Docker 可以让你像使用集装箱一样快速的组合成应用，并且可以像运输标准集装箱一样，尽可能的屏蔽代码层面的差异。Docker 会尽可能的缩短从代码测试到产品部署的时间。

# Dockerfile运行容器
	1. 通过dockerfile build镜像
https://www.runoob.com/docker/docker-build-command.html
docker build -f ./Dockerfile -t article-api001
	2. 运行本地镜像
https://www.runoob.com/docker/docker-run-command.html
docker run -p 

# 挂载卷

--volume -v



