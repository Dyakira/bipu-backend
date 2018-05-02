# bipu-backend

## 启动
### 本地
     `go run main.go`

### 配置

conf/config.json 格式如下

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
