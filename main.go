package main

/* TODO
https://github.com/tidwall/gjson 处理json
https://github.com/alecthomas/log4go 日志
*/
import (
	"fmt"
	"github.com/cc14514/go-lightrpc-example/service"
	"os"

	"github.com/alecthomas/log4go"
	"github.com/cc14514/go-lightrpc/rpcserver"
	"github.com/urfave/cli"
)

var logLevel []log4go.Level = []log4go.Level{log4go.ERROR, log4go.WARNING, log4go.INFO, log4go.DEBUG}

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "JSON-RPC 接口框架"
	app.Version = "0.0.1"
	app.Author = "liangc"
	app.Email = "cc14514@icloud.com"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "rpcport,p",
			Usage: "HTTP-RPC server listening `PORT`",
			Value: 8080,
		},
		cli.StringFlag{
			Name:  "logfile",
			Usage: "log file path",
		},
		cli.IntFlag{
			Name:  "loglevel",
			Usage: "0=errr, 1=warn, 2=info, 3=debug",
			Value: 3,
		},
	}
	app.Before = func(ctx *cli.Context) error {
		filepath := ctx.GlobalString("logfile")
		idx := ctx.GlobalInt("loglevel")
		level := logLevel[idx]
		if filepath != "" {
			fmt.Println("logfile =", filepath, "level =", level)
			log4go.AddFilter("file", log4go.Level(level), log4go.NewFileLogWriter(filepath, false))
		}
		log4go.AddFilter("stdout", log4go.Level(level), log4go.NewConsoleLogWriter())
		return nil
	}
	app.Action = func(ctx *cli.Context) error {
		log4go.Debug(">> Action on port = %s", ctx.GlobalInt("p"))
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
		return nil
	}
	app.Run(os.Args)
}
