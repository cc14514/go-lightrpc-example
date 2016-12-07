# go-lightrpc-example
### 协议描述
<a href="https://github.com/cc14514/go-lightrpc">https://github.com/cc14514/go-lightrpc</a>


### 安装example

* 假设 GOPATH=/usr/local/gopath
* cd /usr/local/gopath/src 
* git clone https://github.com/cc14514/go-lightrpc-example.git
* cd /usr/local/gopath
* go install go-lightrpc-example
* bin/go-lightrpc-example -h

### api 说明
<pre><code>
Rpcserver 定义如下：
type Rpcserver struct {
	Pattern        string // url , 默认 /api/
	Port           int //端口	
	ServiceMap     map[string]ServiceReg //service服务映射	
	CheckToken     func(token TOKEN) bool //校验Token的回调函数	
	AllowedMethods []string //接受的方法，默认 [GET,POST] 
}

main.go 中初始化服务如下：
...
rs := &rpcserver.Rpcserver{
	
	//端口	
	Port:       ctx.GlobalInt("rpcport"),
	//服务映射
	ServiceMap: service.ServiceRegMap,
	// 校验请求中的 TOKEN 是否正确，根据不同的业务需求，会有不同实现
	CheckToken: func(token rpcserver.TOKEN) bool {
		log4go.Debug("Auth token = %s", token)
		if token == "123456" {
			return true
		} else {
			return false
		}
	},
}
rs.StartServer()
...

其中 service.ServiceRegMap 代码在 service/registry.go 中，
每次编写一个 service 时需要为其注册一个名字，




可以访问 http://localhost:8080/api/?body={} 测试接口是否成功启动；
</code></pre>


