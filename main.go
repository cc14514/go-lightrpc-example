package main

import (
	"fmt"
	"github.com/cc14514/go-lightrpc/rpcserver"
	"github.com/urfave/cli"
	"go-lightrpc-example/service"
	"os"
)

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
	}
	app.Action = func(ctx *cli.Context) error {
		fmt.Println("Action ====", ctx.GlobalInt("rpcport"), ctx.GlobalInt("p"))
		rs := &rpcserver.Rpcserver{
			Port:       ctx.GlobalInt("rpcport"),
			ServiceMap: service.ServiceRegMap,
			CheckToken: func(token rpcserver.TOKEN) bool {
				fmt.Println("Auth token =", token)
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
