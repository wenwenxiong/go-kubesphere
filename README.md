### 下载编译安装

下载

```
go get github.com/wenwenxiong/go-kubesphere
```

编译

```
go build main.go
```

(go包依赖自己下载)

### 使用

1、增删改查```kubesphere```上的应用商店中的分类；

2、该查```kubesphere```上的应用商店中的预置应用。

例如：

把```kubesphere```上的应用商店中```mongodb```移到自建的```database```分类。

```
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ ./main category create -n database -l {}
Accesstoken: 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGt1YmVzcGhlcmUuaW8iLCJpYXQiOjE1ODkzNjI2MzQsInVzZXJuYW1lIjoiYWRtaW4ifQ.ETKKRThEDxvBeKLj7wzn19XNG1A1xexZG7LHgKooDBc
create app category id: ctg-0V4A71ANOO5R
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ ./main category get 
Accesstoken: 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGt1YmVzcGhlcmUuaW8iLCJpYXQiOjE1ODkzNjQyNTYsInVzZXJuYW1lIjoiYWRtaW4ifQ.HKgvuCTpKj8LZvz0sFT9ridyo1OUJB6DNORKqQTl8fs
get app category result total: 
3
	get app category at 1 result:
		app category name: uncategorized
		app category id: ctg-uncategorized
	get app category at 2 result:
		app category name: bigdata
		app category id: ctg-7yBqy1EQj7Mq
	get app category at 3 result:
		app category name: database
		app category id: ctg-0V4A71ANOO5R
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ 
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ ./main app get 
Accesstoken: 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGt1YmVzcGhlcmUuaW8iLCJpYXQiOjE1ODkzNjQyNjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.gW2zALPQ_zsV-vj2yIx8dE-v69eGzVPeV9u4dO_8UpQ
get app result total: 
9
	get app at 1 result:
		app name: harbor
		app id: app-Yj7xzjOzvrNn
	get app at 2 result:
		app name: memcached
		app id: app-2o4X9B49NRY4
	get app at 3 result:
		app name: mongodb
		app id: app-ywxoOZBnoqOj
	get app at 4 result:
		app name: mysql
		app id: app-7Og2AxymKrJy
	get app at 5 result:
		app name: nginx
		app id: app-QzVMGQBZAy5O
	get app at 6 result:
		app name: postgresql
		app id: app-PYOkwlDnoqOj
	get app at 7 result:
		app name: rabbitmq
		app id: app-YZ3D9n2zvrNn
	get app at 8 result:
		app name: redis
		app id: app-ED7jnGkEzRMr
	get app at 9 result:
		app name: tomcat
		app id: app-W6oKKnooR6B8
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ ./main app update --help
update app in kubesphere store

Usage:
  kubespherectl app update [OPTIONS]  [flags]

Flags:
  -i, --appId string               app id
  -c, --appNewCategoryId string    app new category Id
  -d, --appNewDescription string   app new description
  -n, --appNewName string          app new name
  -h, --help                       help for update

Global Flags:
  -a, --apiGateway string   ks-apigateway url (default "http://192.168.122.162:30881/")
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ 
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ 
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ ./main app update -i app-ywxoOZBnoqOj -c ctg-0V4A71ANOO5R
Accesstoken: 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGt1YmVzcGhlcmUuaW8iLCJpYXQiOjE1ODkzNjQzMzcsInVzZXJuYW1lIjoiYWRtaW4ifQ.8Wi5hcPJfcMOJD1N3XkI08kvPhyYvRNtNiXibCDFWUU
update app id: app-ywxoOZBnoqOj, result success
xww@xww-NUC8i5BEH:~/gowork/src/github.com/wenwenxiong/go-kubesphere$ 

```



### 效果图

![Screenshot from 2020-05-13 18-14-53](/home/xww/Pictures/Screenshot from 2020-05-13 18-14-53.png)