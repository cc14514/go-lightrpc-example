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
<pre><code golang>
		rs := &rpcserver.Rpcserver{
			Port:       ctx.GlobalInt("rpcport"),
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
</code></pre>
