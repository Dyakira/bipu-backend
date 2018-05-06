# bipu-backend

## 启动
### 本地
     `go run main.go`

### 配置

conf/config.json 格式如下
（除github键之外，其它所有字段均需填写，否则会导致数据库连不上，用户session不能用，登录注册不了等情况）

```
{

  "Github":{
    "clientid": "",
    "secret": ""
    "redirectUrl": ""
  },
  "Db": {
    "User": "",
    "Pass": "",
    "Host": "",
    "Port": "",
    "Database": ""
  },
  "Session":{
    "sessionName": "",
    "sessionSecret": ""
  },
    "Others":{
      "md5Salt": ""
    }
}
```


### 正式  
 + 编译  
    
        `sudo CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go`
 + 运行   
        
         `$> nohup ./main > /dev/null 2>stderr.log &`
 + 结束
            
         kill `cat bipu_server.pid`

### etc

接口文档：etc/interface.md