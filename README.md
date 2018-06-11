# SURLS (For Demo)

SURLS 使用golang编写,实现短域名服务。  
项目基于[go-kit](https://github.com/go-kit/kit)搭建,并集成常用组件.  

###### *PS:转换方式直接使用md5,仅供演示使用.*
---

- [x] 依赖库管理 [glide](https://github.com/Masterminds/glide)
- [x] 实时编译 [realize](https://github.com/oxequa/realize)
- [x] 命令行支持 [cli](https://github.com/urfave/cli)
- [x] [grpc](https://github.com/grpc/grpc)支持 
- [x] grpc => http [协议自动转换](https://github.com/grpc-ecosystem/grpc-gateway)
- [x] 服务熔断[hystrix](https://github.com/afex/hystrix-go) 
- [x] prometheus采集支持
- [x] docker镜像构建
- [x] 自定义中间件
- [x] 实时debug图表信息 [debugcharts](https://github.com/mkevac/debugcharts)
- [x] pprof分析器，图表化
- [x] 多平台打包[gox](https://github.com/mitchellh/gox)
- [x] 服务优雅退出 graceful
- [x] tests
- [x] benchmark
- [x] [yaml](https://github.com/go-yaml/yaml)配置文件支持
- [x] [env](https://github.com/joho/godotenv)配置文件支持
- [ ] 增加sleep接口用于观察graceful效果
- [ ] zipkin全链路追踪

## 环境依赖 (go1.10.2)
```bash
# docker环境
$ curl -fsSL https://get.docker.com/ | sh

# glide
$ brew install glide

# realize
$ go get -v github.com/oxequa/realize

# gox
$ go get -v github.com/mitchellh/gox

# protobuf
$ brew install protobuf
$ go get -u -v github.com/golang/protobuf/protoc-gen-go

# grpc-gateway
$ go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

# pprof
$ brew install graphviz
$ go get -u -v github.com/google/pprof
```

## 文件目录结构
```
.
├── assets //公共资源
├── bin //go build后bin包
├── cli //命令行参数解析
├── clients //grpc客户端
|   ├── get //get服务客户端
|   └── set //set服务客户端
├── conf //多环境配置文件保存目录
|   ├── container.yaml //容器模式下配置文件
|   └── local.yaml //本地开发模式下配置文件
├── docker //docker相关
|   ├── docker-compose.yaml //服务集成环境启动 docker-compose 配置文件
|   ├── image_build //服务构建独立docker镜像相关文件
|   └── ... //其他docker服务相关配置文件
├── global //全局生效 变量&配置
|   ├── conf.go //实例化配置文件
|   ├── errors.go //错误信息配置
|   ├── global.go //全局变量&配置 入口
|   ├── logger.go //日志实例
|   └── redis.go //redis实例
├── lib //公共库目录
├── utils //工具目录
├── pb //protobuf 相关文件保存目录
|   ├── compile.sh //protobuf 文件编译脚本
|   ├── surls.pb.go //protobuf原始文件生成的go文件
|   ├── surls.pb.gw.go //protobuf原始文件生成的 grpc转http 网关文件
|   └── surls.proto //服务定义文件
├── servers //服务启动相关
|   ├── debug.go //debug服务
|   ├── grpc.go //grpc服务
|   ├── http.go //http服务
|   └── metrics.go //数据采集服务
├── handlers //业务逻辑相关目录
|   ├── endpoints //go-kit endpoints实现
|   ├── svc //go-kit 服务定义&实现
|   └── transports //go-kit transport实现
├── runtime //保存程序运行时数据
|   ├── pid //服务运行pid
|   ├── logs //日志保存目录
|   └── ... //其他运行时数据
├── vendor //依赖库保存目录
├── .realize.yaml //实时自动编译配置文件
├── .env //项目运行环境变量
└── glide.yaml //依赖库配置文件

```

## Install
```bash
$ cd $GOPATH/src
$ git clone git@github.com:GxlZ/surls.git
```

## Run
```bash
# run redis server
$ docker run -d \
-p 6973:6379 \
--name redis-local \
redis
$ cd $GOPATH/src/surls
# 自动编译
$ realize start
# 手动启动
$ go run main.go -h 
```

## Run By Docker
```bash
$ cd $GOPATH/src/surls
$ docker-compose -f docker/docker-compose.yaml up -d
```

## Access Service
```bash
# 访问 grpc server
$ go run clients/get/client.go -s test
$ go run clients/set/client.go -s test
# 访问 http server
$ curl -XPOST -d '{"url":"http://www.baidu.com"}' http://localhost:7071/surls/v1/set
$ curl 'http://localhost:7071/surls/v1/get?url=bfa89e563d9509fbc5c6503dd50faf2e'
```

## Graceful
```bash
# 测试优雅退出可以在代码中增加sleep进行测试
$ cat $GOPATH/src/surls/pid | xargs kill -s SIGINT
```

## Build
```bash
$ gox -verbose
```

## Test
```bash
$ cd $GOPATH/src/surls
$ go test -v -cover=true ./...
-----------------------------------------------------------------------------
=== RUN   TestSurlsSet
--- PASS: TestSurlsSet (0.00s)
=== RUN   TestSurlsGet
--- PASS: TestSurlsGet (0.00s)
PASS
coverage: 90.0% of statements
ok  	surls/handlers/transports	0.006s	coverage: 90.0% of statements

```

## Benchmark

```bash
$ hey -m POST -c 1000 -n 100000 -d '{"url":"http://www.baidu.com"}' http://localhost:7071/surls/v1/set
-----------------------------------------------------------------------------------------------
Summary:
  Total:	3.2813 secs
  Slowest:	0.1637 secs
  Fastest:	0.0006 secs
  Average:	0.0323 secs
  Requests/sec:	30475.4100

  Total data:	4196775 bytes
  Size/request:	41 bytes

Response time histogram:
  0.001 [1]	|
  0.017 [968]	|∎
  0.033 [67971]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.049 [28395]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.066 [1464]	|∎
  0.082 [244]	|
  0.098 [59]	|
  0.115 [11]	|
  0.131 [51]	|
  0.147 [450]	|
  0.164 [386]	|


Latency distribution:
  10% in 0.0239 secs
  25% in 0.0271 secs
  50% in 0.0305 secs
  75% in 0.0343 secs
  90% in 0.0400 secs
  95% in 0.0449 secs
  99% in 0.0801 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0001 secs, 0.0006 secs, 0.1637 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0151 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0172 secs
  resp wait:	0.0314 secs, 0.0005 secs, 0.1306 secs
  resp read:	0.0005 secs, 0.0000 secs, 0.0215 secs

Status code distribution:
  [200]	100000 responses
  
```

```bash
$ hey -c 1000 -n 100000 'http://localhost:7071/surls/v1/get?url=bfa89e563d9509fbc5c6503dd50faf2e'
-----------------------------------------------------------------------------------------------
Summary:
  Total:	5.0311 secs
  Slowest:	0.1544 secs
  Fastest:	0.0005 secs
  Average:	0.0494 secs
  Requests/sec:	19876.3453

  Total data:	5800000 bytes
  Size/request:	58 bytes

Response time histogram:
  0.000 [1]	|
  0.016 [4618]	|∎∎∎∎∎
  0.031 [8613]	|∎∎∎∎∎∎∎∎∎
  0.047 [27651]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.062 [38082]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.077 [16573]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.093 [3012]	|∎∎∎
  0.108 [996]	|∎
  0.124 [434]	|
  0.139 [19]	|
  0.154 [1]	|


Latency distribution:
  10% in 0.0277 secs
  25% in 0.0392 secs
  50% in 0.0502 secs
  75% in 0.0599 secs
  90% in 0.0691 secs
  95% in 0.0760 secs
  99% in 0.0996 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0002 secs, 0.0005 secs, 0.1544 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0111 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0199 secs
  resp wait:	0.0486 secs, 0.0004 secs, 0.1544 secs
  resp read:	0.0003 secs, 0.0000 secs, 0.0312 secs

Status code distribution:
  [200]	100000 responses
```

```bash
$ cd $GOPATH/src/surls
$ go test -v -bench=. -benchtime=2s -benchmem -run=none
-----------------------------------------------------------------------------------------------
goos: darwin
goarch: amd64
pkg: surls/tests
BenchmarkSurlsSet-16    	   30000	     89480 ns/op	    3922 B/op	      77 allocs/op
BenchmarkSurlsGet-16    	  100000	     49373 ns/op	    2769 B/op	      52 allocs/op
PASS
ok  	surls/handlers/transports	8.987s
```

## Docker Build
```bash
$ cd $GOPATH/src/surls
$ docker build -f docker/image_build/Dockerfile ../ -t=surls-demo
$ docker run -d \
--name surls-demo \
-p 7070:7070 \
-p 7071:7071 \
-p 7072:7072 \
-p 7073:7073 \
surls-demo
```

## Debug
```bash
# debug默认端口为7073,
$ debugUrl=http://localhost:7073
```
>实时图表化数据
```bash
$ open $debugUrl/debug/charts/
```
<img src="assets/debug-charts.png" />

>常规debug信息
```bash
$ open $debugUrl/debug/pprof/
$ open $debugUrl/debug/pprof/cmdline
$ open $debugUrl/debug/pprof/profile
$ open $debugUrl/debug/pprof/symbol
$ open $debugUrl/debug/pprof/trace
```

>采集数据,展示分析结果
```bash
$ go test -bench=. -benchtime=3s -benchmem -run=none -v &
$ pprof -web $debugUrl/debug/pprof/profile
# Saved profile in /path/to/pprof.xx.pb.gz`
```

```bash
# 原生pprof，比go tool pprof分析结果更加丰富
$ pprof -http=:7074 /path/to/pprof.xx.pb.gz
$ open http://localhost:7074
```
<img src="assets/debug-pprof.png" />

>Prometheus 数据采集
`PS:使用docker-compose方式启动，略过prometheus服务启动`
```bash
# 启动prometheus
$ docker run -d \
  --rm \
  -p 9090:9090 \
  --name prometheus \
  --network=host \
  -v $GOPATH/src/surls/docker/prometheus/data:/prometheus-data \
  prom/prometheus \
  --config.file=/prometheus-data/prometheus.yml

$ open http://127.0.0.1:9090
```
<img src="assets/prometheus.png" />
