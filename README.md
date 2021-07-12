# Simple

[![Build](https://github.com/fregie/simple/actions/workflows/gobuild.yml/badge.svg?branch=main)](https://github.com/fregie/simple/actions/workflows/gobuild.yml)


Simple是一款网络代理/VPN服务端管理平台。使用simple来轻松管理服务器上多种不同协议的服务端。

## 支持的协议
- [x] trojan
- [ ] openvpn
- [ ] ikev2
- [ ] WireGuard
- [ ] shadowsocks
- [ ] V2ray

## 功能
- [x] 创建、删除、查询用户配置
- [x] 配置数据持久化
- [x] 限速（需要协议实现支持）
- [x] 使用grpc api管理
- [x] 使用命令行工具(spctl)管理
- [x] 自定义服务端对接

## 安装
Simple使用多服务模块化设计，需要数个服务同时运行，传统部署方式会略麻烦，推荐[使用docker部署](#使用Docker部署)。

### 使用Docker部署
首先安装docker以及docker-compose。可以在docker目录下修改相关配置。
```shell
cd docker
docker-compose up -d
```
安装命令行管理工具spctl
```shell
go install github.com/fregie/simple/spctl@latest
```

## 使用
### 配置spctl
创建配置文件`~/.spctl.yaml`.  
默认使用`~/.spctl.yaml`,也可以使用参数`-config path`指定配置文件。
```yaml
# simple服务的grpc地址
grpcAddr: 127.0.0.1:4433
```
### session
session是simple的基本单位，含义为一个vpn的客户端会话，可以理解为一个客户端可用的配置。
#### 创建
```shell
$ spctl create session
SUCCESS  Create success!
ID:            trojan-83d33d33-4u8hyxFHu2922D3pD6
Proto:         trojan
Config type:   JSON
Config:
{"run_type":"client","remote_addr":"simple.fregie.cn","remote_port":2443,"password":["4u8hyxFHu2922D3p"],"ssl":{"verify":false,"sni":""},"mux":{"enabled":false,"concurrency":0,"idle_timeout":0},"websocket":{"enabled":false,"path":"","host":""}}
```
#### 查看全部
```shell
$ spctl get sessions
ID                                 | proto  | config type
trojan-8e1a48ee-6KKx8a5BvFOzJpMFD9 | trojan | JSON       
trojan-8cb8b7e3-a01n63djgk4m1egTI5 | trojan | JSON       
trojan-a6534f4a-79gac1JwT668HkFa1S | trojan | JSON       
trojan-21f00bb5-dB55kAKwCAh7oEHLZM | trojan | JSON       
trojan-92c88a90-Z8Pn687n7STxwIwSd3 | trojan | JSON       
trojan-7465a545-rr7jX04KBBY7sEeW0T | trojan | JSON       
```
#### 查看特定session
```shell
$ spctl get session trojan-8e1a48ee-6KKx8a5BvFOzJpMFD9(session ID) -conf
ID:            trojan-8e1a48ee-6KKx8a5BvFOzJpMFD9
Proto:         trojan
Config type:   JSON
Option:
    Upload rate limit:   0 mbps
    Download rate limit: 0 mbps
Config:
{"run_type":"client","remote_addr":"","remote_port":2443,"password":["wSd3rr7jX04KBBY7"],"ssl":{"verify":false,"sni":""},"mux":{"enabled":false,"concurrency":0,"idle_timeout":0},"websocket":{"enabled":false,"path":"","host":""}}
```

#### 删除
```shell
$ spctl delete session trojan-8e1a48ee-6KKx8a5BvFOzJpMFD9(session ID)
SUCCESS  Delete trojan-8e1a48ee-6KKx8a5BvFOzJpMFD9
```