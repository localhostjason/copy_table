# copy_table

工具：
支持多数据库，多表格 复制

fix: 地形与勇士 拍卖行和金币寄售 无法开启



在 dist 目录下 运行：

```shell
main.exe -x
```

注：
dist 下面 config/agent.json 请填写正确的 mysql 配置
```json
{
  "db": {
    "enable": true,
    "mysql": [
      {
        "key": "TestMysql1",
        "user": "root",
        "password": "123456",
        "host": "127.0.0.1",
        "port": 3306,
        "db": "taiwan_cain_auction_gold",
        "charset": "utf8mb4",
        "timeout": 5,
        "multi_statements": false,
        "debug": false
      },
      {
        "key": "TestMysql2",
        "user": "root",
        "password": "123456",
        "host": "127.0.0.1",
        "port": 3306,
        "db": "taiwan_cain_auction_cera",
        "charset": "utf8mb4",
        "timeout": 5,
        "multi_statements": false,
        "debug": false
      }
    ]
  }
}

```


若 邮件显示乱码，数据中 配置文件中设置
/etc/my.cnf 中：
```shell
[mysqld]
default-character-set=latin1

[client]
default-character-set=latin1
```