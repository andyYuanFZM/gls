# 基于 chain33 区块链开发 框架 开发的 POS公链 系统（v6.8.0）

开发框架：https://chain.33.cn

```
注意: master 分支不是 发布版本, 不要用于生产环境
```

## 安装

#### golang 1.16+


#### 支持make file的平台

```
git clone https://github.com/andyYuanFZM/gls $GOPATH/src/github.com/andyYuanFZM/gls

//开启mod功能
export GO111MODULE=on

//国内用户需要导入阿里云代理，用于下载依赖包
export GOPROXY=https://mirrors.aliyun.com/goproxy

cd $GOPATH/src/github.com/andyYuanFZM/gls

make
```

就可以完成编译安装

### 更新go.mod

```
make update
```



