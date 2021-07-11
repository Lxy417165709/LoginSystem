@[toc]
## 环境说明
`Centos` 版本: `Centos 7.3`

<br>

## 安装 docker
### 步骤
1. 执行安装命令 
	```shell
	$ curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
	```
	也可以使用国内 `daocloud` 一键安装命令：
	```shell
	$ curl -sSL https://get.daocloud.io/docker | sh
	```
2. 测试是否安装成功
	```shell
	$ docker --version
	Docker version 1.13.1, build 64e9980/1.13.1
	```



### 参考资料
- [Docker](https://www.runoob.com/docker/centos-docker-install.html)

<br>

## docker  镜像加速
### 需要加速原因
国内从 `DockerHub` 拉取镜像有时会遇到困难，此时可以配置镜像加速器。
### 步骤
1. 执行命令 
	```c
	$ cd /etc/docker
	$ vim daemon.json
	```
2. 在 `deamon.json` 中写入以下内容
	```c
	{"registry-mirrors":["https://reg-mirror.qiniu.com/"]}
	```
3. 重启 `docker`
	```c
	$ sudo systemctl daemon-reload
	$ sudo systemctl restart docker
	```

### 参考资料
- [Docker 镜像加速](https://www.runoob.com/docker/docker-mirror-acceleration.html)

<br>

## 安装 docker-compose
### 步骤
1. 执行安装命令 
	```c
	$ sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
	```
2. 将可执行权限应用于二进制文件
	```c
	$ sudo chmod +x /usr/local/bin/docker-compose
	```

3. 测试是否安装成功
	```c
	$ docker-compose --version
	docker-compose version 1.24.1, build 4667896b
	```

### 问题
- [ ] `步骤1` 安装时，下载太慢。
	`步骤1` 安装时换个 `docker-compose` 源: 
	```c
	$ curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
	```
	`步骤2`、`步骤3` 一样。

### 参考资料
- [Docker Compose](https://www.runoob.com/docker/docker-compose.html)
- [docker-compose速度太慢解决方式](https://blog.csdn.net/weixin_43299268/article/details/105108738)


## 部署

### 步骤
1. 将 文件夹`LoginSystem/deploy` 的内容拷贝到阿里云服务器。 (可以使用 `File Zilla`)
2. 在阿里云服务器，切换到该文件夹中。
	```shell
	$ cd /root/myApp/deploy
	```
3. 执行命令，运行应用
	```shell
	$ docker-compose up
	```
	如果想要后台运行，则执行:
	```shell
	$ docker-compose up -d
	```
4. 检查是否部署成功
	```shell
	$ docker ps
	CONTAINER ID        IMAGE                                                   COMMAND                  CREATED             STATUS              PORTS                    NAMES
	36084ff2e248        nginx                                                   "/docker-entrypoin..."   12 hours ago        Up 4 seconds        0.0.0.0:80->80/tcp       nginx
	89fc0f1a4194        postgres                                                "docker-entrypoint..."   12 hours ago        Up 4 seconds        0.0.0.0:8989->5432/tcp   pg_db
	281f6376ce96        deploy_golang                                           "./go"                   12 hours ago        Up 3 seconds        0.0.0.0:9000->8080/tcp   golang
	72e300a0f549        registry.cn-hangzhou.aliyuncs.com/xylink/redis:3.2_v1   "docker-entrypoint..."   12 hours ago        Up 4 seconds        0.0.0.0:9999->6379/tcp   redis
	```

## 问题
### `ping` 得通阿里云服务器，但是无法访问服务器端口
#### 原因
这可能是没有配置阿里云服务器安全组的问题。

#### 步骤


![在这里插入图片描述](https://img-blog.csdnimg.cn/20200724100558650.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE5MDE4Mjc3,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200724100708822.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE5MDE4Mjc3,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200724100721112.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE5MDE4Mjc3,size_16,color_FFFFFF,t_70)
	

