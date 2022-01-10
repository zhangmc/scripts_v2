# scripts_v2

## 目录说明
```text
├── Dockerfile  //  docker build
├── LICENSE
├── README.md
├── com  // 通用包
│   └── file.go
├── config 
│   ├── config.go    // 根配置
│   ├── config.json  // 配置文件
│   └── jd  // 子配置项
├── constract  // interface 定义
│   ├── jd.go
│   └── notify.go
├── database  // database 相关
│   └── database.go
├── global  // 全局变量
│   ├── const.go
│   ├── logger.go
│   └── user_agent.go
├── go.mod
├── go.sum
├── manage.go
├── models  // 模型定义
├── repositories  // 模型操作
├── scripts // 脚本存放路径
│   └── jd  // 存放某东脚本
├── shell  // 存放shell脚本
│   └── docker-entrypoint.sh
├── storage  // 存放资源，日志文件
│   └── logs  // 日志目录
├── structs // 结构体定义
│   ├── jd.go
│   └── notify.go
└── tools // 工具函数
    ├── build_scripts.go
    ├── update_config.go
    ├── update_crontab.go
    └── update_readme.go
```

## 安装

- 拉取镜像: `docker pull classmatelin/scripts:v2`
- 创建容器: `docker run -itd --name scripts classmatelin/scripts:v2`
- 进入容器: `docker exec -it scripts bash`

## 使用

### 配置文件

- 配置文件: `config/config.yaml`:
```json
{
  "logger": {
    "directory": "storage/logs", // 日志文件夹
    "filename": "logger.log" // 日志名称
  },
  "jd": {
    "cookies": [
      "pt_pin=jd1;pt_key=sssss;",  // 最简配置
      "pt_pin=jd2;ws_key=sfsfsf;",  // 使用ws_key;
      "pt_pin=jd3:ws_key=sfsfsfsf;remark=账号3;", // remark为备注账号名
      "pt_pin=jd4;pt_key=sfasfafaf;remark=账号4;push_plus=fasfasfasfafaf;" // 备注账号4, 并且单独配置push_plus通知.
    ]
  },
  "notify": {
    "server_j": "", // server 酱通知send key
    "push_plus": "",  // push+通知 token
    "tg": { 
      "bot_token": "",  // tg机器人token
      "user_id": ""   // tg chat id.
    }
  }
}
```


### 定时任务

- 定时任务配置文件: `config/crontab.json`
- 键为脚本名称(不带后缀.go), 值为定时任务值, 设为`false`则关闭脚本定时任务。
```json
{
  "jd_check_cookies": "0 */4 * * *",  
  "jd_joy_park": false 
}
```