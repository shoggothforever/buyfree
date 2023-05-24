# 开发问题

- [订单问题](#订单问题）
- [资金问题](#资金问题)
- [服务高可用](#服务高可用)
- [数据大屏](#数据大屏)
- [设备管理界面](#设备管理界面)
- [定位系统](#定位系统)
- [支付系统（购物车、订单）](#支付系统（购物车、订单）)
- [营销额分析](#营销额分析)
- [资源投放](#资源投放)
- [服务迁移](#服务迁移)
- [部署脚本](#部署脚本)
- [TODO](#TODO)

## 订单问题
场站端和平台端的商品详情存在重复
应该是场站使用平台管理，平台端管理所有订单信息
![image.png](https://api.apifox.cn/api/v1/projects/2429864/resources/372272/image-preview)
选用合适的设计模式，设计订单结算的服务接口，实现多态性


## 资金问题
乘客与车主仅通过设备关联，（购物车功能）
车主端与场站存在直接资金往来（补货购物车功能）


## 服务高可用
数据库分库分表，建立隔离事务，redis部署集群

## 数据大屏
以场站的销售额为准
广告统计：需要哪些信息，限制条数
热卖排行：场站供应的货品销量（销售额）榜
司机定位信息：哪些司机需要定位？与场站之间存在订单未取货关系的司机
## 设备管理界面
进入界面分页展示所有订单信息，每页 ？ 行（点击这个页面后就已经获取到了所有需要的信息了）30行
其余选项（激活，在线）前端可以自行过滤筛选

## 定位系统
调用GPS API 或许用户实时信息
用户购物时存储用户购物时的车主位置信息
司机挑选补货场站时获取场站的地理位置信息（绝对地址和相对地址）

## 支付系统（购物车、订单）
结算系统调用支付宝API
订单信息传递
乘客购买商品，司机端设备接收到购物信息，收集资金数据，更新车主端的资金数据
车主在选货列表中挑选场站以及它们拥有的商品，添加进购物车
* 在购物车中选择特定商品，选定后付款信息打包成订单信息，开启线程发送到服务端，接收到返回信息，表示远端信息处理完毕，之后再更新本地数据（更新pgsql里的订单信息，设备商品信息，还要更新车主端的销量信息）
* 服务器端处理过程：车主端向场站发送补货请求，场站端接收到订单信息（状态为1），司机资金流动到场站，场站平台拆分司机订单中的货品信息,向redis数据库发送信号，依次更新
  统计数据
  结算接口的幂等性（如果因为服务延迟导致用户在一次服务响应之前发送了多次请求，最后结果应该与用户只发送一次请求保持一致）
  分布式数据锁的设计:redission+redlock or set nx ex+lua

## 营销额分析
建立redis数据库
应用数据结构 bitmap,list(每天的营销额）,zset,hash，GEORadius
建立特殊时间戳函数：获取每周，月，年的第一天的信息
* 所有数据库都要有垃圾清理机制
  Sort 机制：建立排行榜
  ZRankByLex：
### bitmap存储用户连续登录信息，优化用户登录状态判定
### list 存储每日营销额的连续变化数据，list存储每日营业，每周，每月，每年，总营销额:存储累计营销额
（创建守护进程，时刻维护着每个场站/车主的营销日、周、月、年、总营销额，做垃圾清理)
七日连续营销额，当日营销额以及前6天的营销额做累加即可做即可
Key: Product_ID:date
val:营业额
### hash 存储
### GEOXXX 存储场站经纬度信息
key: factory_id/factory_name/场站地址(xx路xx号）
val: 经度、纬度
利用Georadius获取在以车主为中心一定范围内的场站到车主的距离
### Zset （key，val，score）存储营销额排行榜
Key: owner_id:type(daily,weekly,monthly,yearly):date(2006-01-02:15:04:05)：
Val:object_id
Score:销量
分布式策略，管道通信，广播机制



## 资源投放
在服务器上存放视频资源，图片资源，利用ftp技术传输资源，利用minio搭建本地资源服务器

## 服务迁移
从windows移动到linux（感觉没有问题，问题主要在于数据库迁移）
## 部署脚本
shell脚本，dockerfile ,docker-compose

## TODO:
升级单体架构成为微服务架构
RPC,go-zero,consul/etcd

# 部署事项
### Docker network
要让所有容器能够相互连通，先创建一个公用的网络
具体细节细看help手册
docker network --help
创建一个名为bfnet的网络
`docker network create bfnet`
### Redis
```
docker run --hostname=53bb9473f77a --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=GOSU_VERSION=1.16 --env=REDIS_VERSION=7.0.11 --env=REDIS_DOWNLOAD_URL=http://download.redis.io/releases/redis-7.0.11.tar.gz --env=REDIS_DOWNLOAD_SHA=ce250d1fba042c613de38a15d40889b78f7cb6d5461a27e35017ba39b07221e3 --volume=D:\desktop\pr\buyfree\dal/redis/redis.conf:/etc/redis/redis.conf --volume=/data --network=bfnet --workdir=/data -p 6379:6379 --restart=no --runtime=runc -d redis

```
docker 部署单机redis
[(48条消息) 应用部署到docker容器连接不上redis容器_LinktoDream的博客-CSDN博客](https://blog.csdn.net/qq_38363738/article/details/106785711)
从官网下载redis.conf 文件
注释其中的bind 127.0.0.1，将其中的保护模式设置为no
下面的network-alias比较关键，容器变更ip地址是不固定的，每次都从 /etc/hosts 文件读取IP信息不现实，所以用一个固定的别名代替IP地址是最好的选择
```
docker pull redis:latest
docker run -p 6379:6379 --name bfredis --network bfnet --network-alias redis -v D:\desktop\pr\buyfree\dal/redis/redis.conf:/etc/redis/redis.conf -d redis redis-server /etc/redis/redis.conf --appendonly yes

```
#### docker 部署redis集群（非host模式）
```shell
### 创建一号节点
docker create -p 7001:6379 --name node1 --network bfnet --network-alias redis-node1 -v E:\dockerdbvolume\redis\data\node1:/data redis --cluster-enabled yes --cluster-config-file nodes-node-1.conf 
### 创建二号节点
docker create -p 7002:6379 --name node2 --network bfnet --network-alias redis-node2 -v E:\dockerdbvolume\redis\data\node2:/data redis --cluster-enabled yes --cluster-config-file nodes-node-2.conf 
### 创建三号节点
docker create -p 7003:6379 --name node3 --network bfnet --network-alias redis-node3 -v E:\dockerdbvolume\redis\data\node3:/data redis --cluster-enabled yes --cluster-config-file nodes-node-3.conf 

```

```shell
# 执行命令查看每个节点的网络信息
docker inspect node1
docker inspect node2
docker inspect node3
```
```shell
#创建集群
#1 进入容器
docker exec -it node1 /bin/bash
redis-cli --cluster create 172.20.0.4:6379 172.20.0.5:6379 172.20.0.6:6379 --cluster-replicas 0
```
#### docker部署redis集群(host模式)
构建docker镜像
```
docker create  -p 6380:6380 --name node1 --net host -v E:\dockerdbvolume\redis\data\node1:/data redis --cluster-enabled yes --cluster-config-file nodes-node-1.conf --port 6380

docker create -p 6381:6381 --name node2 --net host -v E:\dockerdbvolume\redis\data\node2:/data redis --cluster-enabled yes --cluster-config-file nodes-node-2.conf --port 6381

docker create -p 6382:6382 --name node3 --net host -v E:\dockerdbvolume\redis\data\node3:/data redis --cluster-enabled yes --cluster-config-file nodes-node-3.conf --port 6382

```
创建集群
```shell

docker exec -it node1 /bin/bash
redis-cli --cluster create 127.0.0.1:6380 127.0.0.1:6381 127.0.0.1:6382 --cluster-replicas 0
```

### PostgreSQL
1. Postgres 部署策略
   宝塔云连接不上postgresql
   原因：自动安装的pgsql不会创建root用户，需要自己手动创建(只能使用root用户了)
   同时需要给数据库所有权限
````
su - postgres
psql
create user root with 'nyarlak';
create database root owner root;
\password root //修改密码
grant all privileges on database root to root;
dsn = "host=localhost port=5432 user=root dbname=root password=nyarlak  sslmode=disable  TimeZone=Asia/Shanghai"
````
2. 升级云服务的部署策略

```bash
#init.sql
su - postgres
psql
create user 'bf' with 'bf123';
create database 'bfdb' owner to 'bf';
grant all privileges on database bfdb to bf;
dsn= "host=postgres port =5432 user=bf dbname=bfdb password=bf123 sslmode=disable  TimeZone=Asia/Shanghai"



```
3.1. dockerfile部署
部署方法 (启用dockerfile)
```Dockerfile
FROM postgres:alpine   
COPY init.sql /docker-entrypoint-init.d
ENV POSTGRES_USER bf  
ENV POSTGRES_PASSWORD bf123  
ENV POSTGRES_DB bfdb
```

```shell
docker build -t bfpsql .
docker run -p 5432:5432 --name bfpsql --network bfnet --network-alias postgres -v E:\dockerdbvolume\pgdata:/var/lib/postgresql/data bfpsql

```

部署成功进入容器
```shell
adduser bf
psql -d bfdb 
#更改时区
%%set timezone='+8'//手动对准%%
set timezone='Asia/Shanghai';
```
打包上传dockerhub,编写docker-compose

```SHELL
docker tag bfpsql:1.0 $DOCKER_USER_NAME:bfpostgres
docker push $DOCKER_USER_NAME:bfpostgres
mkdir ~/postgres && cd ~/postgres
touch docker-compose.yml
```

3.1.2 docker-compose
```docker-compose
version: '3'
services:
	postgres:
		images:版本号，待填
		healthcheck:
			test:[]
			timeout:
			interval:
			retries:
		container_name: bfpostgres
		restart:always
		enviromnent:
			- POSTGRES_USER=root 
			- POSTGRES_PASSWORD=nyarlak
			- APP_DB_USER=bf
			- APP_DB_PASS=bf123
			- APP_DB_NAME=bfdb
		volumes: 
			- ./db:/docker-entrypoint-initdb.d/ 
		ports: 
			- 5432:5432


```
![[Pasted image 20230411202051.png]]

创建数据库创建脚本

```
mkdir db
touch db/01-init.sh
`````
脚本文件
```bash

#!/bin/bash 
set -e 
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

	CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
	 CREATE DATABASE $APP_DB_NAME;
	  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
	   \connect $APP_DB_NAME $APP_DB_USER 
	BEGIN;
		将要生成表的sql文件 carshop.sql
	COMMIT; 
	  EOSQL

  

```


### GO
docker 部署
go 的编译环境
``` Dockerfile
FROM golang:1.20  
WORKDIR usr/src/app  
ENV GOOS=linux  
ENV CGO_ENABLED=0  
ENV GO_PROXY=https://proxy.golang.com.cn,direct  
ENV GO111MODULE=auto  
  
COPY go.mod go.sum ./  
RUN go mod download && go mod verify  
  
COPY . .  
RUN go build -o /bf  
  
CMD ["/bf"]
```

构建镜像: ` docker build -t bfree . `
开启容器:docker run -itd --name bf --network bfnet  -p 6060:6060 -p 9003:9003 -p 9004:9004 -p 9005:9005 bfree `


### Minio
二进制文件部署方式
一定要在官网下载最新版本：[https://dl.min.io/server/minio/release/linux-amd64/minio](https://dl.min.io/server/minio/release/linux-amd64/minio)
```bat
%%windows 环境
start E:\Minio\minio.exe server E:\Minio\SaveFiles

```


一种后台运行方式
![[Pasted image 20230411164636.png]]

docker 部署方案:
````BASH
//先登录minio界面获取对应可以access_key 以及 secret_key
export MINIO_ACCESS_KEY=XXXXXX
export MINIO_SECRET_KEY=XXXXXX
````

```shell
//创建本地文件夹
mkdir -p ~/minio/data

防火墙开启端口号
firewall-cmd --zone=public --add-port=9000/tcp --permanent
firewall-cmd --zone=public --add-port=9090/tcp --permanent
firewall-cmd --reload
success

# linux
docker run -p 9000:9000 -p 9090:9090 --name minio -v ~/minio/data:/data -e "MINIO_ROOT_USER=minioadmin" -e "MINIO_ROOT_PASSWORD=minioadmin" quay.io/minio/minio server /data --console-address ":9090"

命令详解:
1. docker run 运行容器
2. -p指示绑定9000到容器的9000端口，绑定9090到容器的9090端口
3. -name指示创建的容器名为 minio
4. -v 设置本地文件路径作为容器的挂载路径，当MinIO向/data路径写入数据时，会同步到本地的/minio/data
5. -e 设置环境变量 [`MINIO_ROOT_USER`](https://min.io/docs/minio/linux/reference/minio-server/minio-server.html#envvar.MINIO_ROOT_USER "(in MinIO Documentation for Linux)") 和 [`MINIO_ROOT_PASSWORD`](https://min.io/docs/minio/linux/reference/minio-server/minio-server.html#envvar.MINIO_ROOT_PASSWORD "(in MinIO Documentation for Linux)") 默认为"minioadmin"
6. server 服务端启动
7. --console-address ":9090"控制台端口
8. -address: 9000 minio服务api端口
9. 忘记密码直接查看MINIO_ROOT_USER && MINIO_ROOT_PASSWORD
```
### RabbitMQ

