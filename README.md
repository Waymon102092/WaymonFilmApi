### WaymonFilm 电影票系统
```
  WaymonFilm包含
  1 、WaymonMini         微信小程序
  2 、WaymonTouTiao      抖音小程序
  3 、WaymonAlipay       支付宝小程序
  4 、WaymonApi          用户端小程序api
  5 、WaymonWap          票商端公众号
  6 、WaymonWapApi       票商端api
  7 、WaymonAdmin        管理系统
  8 、WaymonAdminApi     管理系统api
  9 、WaymonMQ           异步程序
  10 、WaymonCrontab      定时器
```

#### 一、介绍
1 、WaymonFilm是一款用户端在小程序(微信、抖音、支付宝)上购买电影票、票商在公众号网页
上出票的电影票系统。
2 、电影票接口是基于网络上提供的免费电影票api接口、如有变动需要及时切换api源
3 、请自主注册微信小程序
4 、请自主注册微信公众号
5 、请自主注册企业支付宝
6 、支付方式 ： 微信支付
7 、提现方式 ： 支付宝提现

#### 二、功能介绍
1 、电影票列表
2 、座位选择
3 、影院列表
4 、个人中心
5 、员工角色
6 、代理商角色
7 、票商角色

#### 三、 用到的技术
1 、Go
2 、Gin
3 、Gorm
4 、Redis
5 、mysql 主从复制
6 、mysql 分库分表
7 、Rabbitmq
8 、jwt
9 、viper
10 、zap
11 、captcha
12 、crontab


#### 四、 go mod tidy
```
  github.com/spf13/viper
  github.com/gin-gonic/gin 
  github.com/dgrijalva/jwt-go
  github.com/go-redis/redis 
  gorm.io/gorm
  gorm.io/driver/mysql
  gorm.io/plugin/dbresolver
  go.uber.org/zap
  github.com/streadway/amqp
  github.com/dchest/captcha
  github.com/robfig/cron
  github.com/go-redsync/redsync/v4
  github.com/satori/go.uuid
  github.com/aliyun/alibaba-cloud-sdk-go
  github.com/aliyun/aliyun-oss-go-sdk
  gorm.io/sharding
  github.com/smartwalle/alipay/v3
  github.com/wechatpay-apiv3/wechatpay-go
```

#### 五、构建镜像发布服务器 (阿里云容器服务)
```
docker build --platform linux/amd64  -t waymon_api:v0.0.1 .  //这里因为本地电脑是mac os 的所以选择打包到 x86的linux服务器
docker login --username=Waymon registry.cn-hangzhou.aliyuncs.com  //这里需要替换成自己的 用户名
docker tag imageId registry.cn-hangzhou.aliyuncs.com/替换成自己的/waymon_api:v0.0.1
docker push registry.cn-hangzhou.aliyuncs.com/替换成自己的/cinema_api:v0.0.1
//登录到自己的服务器上
docker pull registry.cn-hangzhou.aliyuncs.com/替换成自己的/cinema_api:v0.0.1
docker run --name cinema_api56 --restart=always -p 8081:8081 -d  registry.cn-hangzhou.aliyuncs.com/替换成自己的/waymon_api:v0.0.1
```

#### 五、如有疑问可以添加微信咨询
个人主页：https://github.com/Waymon102092/
微信公众号： Waymon
微信：
![](/Users/waymon/Desktop/Waymon/WechatIMG260.jpeg)