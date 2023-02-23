# 环境配置
**golang 1.19**  
**mysql 8.0**  
**ffmpeg**

# 使用说明
修改config/config.go里面配置信息
![QQ截图20230213155511](https://user-images.githubusercontent.com/57628827/218401417-b5361576-774e-4fb6-ad22-8df456bc8060.png)

视频文件和图片会分别存储在 public/video和public/pic下

## 不同模块介绍：

 **common**: 中间件验证，响应结构体  
 **config**: 配置ip地址和数据库信息  
 **handler**: 处理器，决定下一步如何处理，并响应请求  
 **model**: 数据库模块，数据库交互，返回给service层  
 **public**: 上传视频的目录  
 **response**: 响应请求的封装  
 **router**: 路由处理  
 **service**: 业务逻辑，响应数据的封装，返回给handler层  
 **utils**:  工具类  
 **main.go**:  程序入口  

# tiktok
第五届字节青训营项目-迷你版抖音
## group info
队名: **芜湖起飞**
队伍代号: **10086666**

郭天旗
朴景麟
王浩
解鸿基
王文杰
陈德平
雷天乐

## 项目开发资料
Go 框架三件套详解
https://bytedance.feishu.cn/file/boxcnKHOoYmud2SuUGmhFaGbjVb

Database/sql 及 GORM 相关解读
https://juejin.cn/course/bytetech/7140987981803814919/section/7142752406692954143
![QQ截图20230131171326](https://user-images.githubusercontent.com/57628827/215717992-279ec244-7f8f-4108-9b7e-c540a8752ad2.png)

Gorm指南
https://gorm.io/zh_CN/docs/index.html

api接口文档
https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523

极简抖音App使用说明
https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7

抖音项目方案说明
https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#

实战项目-Go 语言笔记服务
https://bytedance.feishu.cn/docx/Wwa4dfwScogfjLxclXKcStGEncd

【开营直播】第五届字节跳动青训营后端专场 
https://bytedance.feishu.cn/docx/WZDddh2Lqoyfu6x93u1c8km9nug

Apk下载地址
https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#
