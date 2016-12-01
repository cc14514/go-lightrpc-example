package service

import (
	"fmt"
	"github.com/cc14514/go-lightrpc/rpcserver"
	"log"
)

var logger *log.Logger

func init() {
	logger, _ = rpcserver.NewLogger("/tmp", "simple_server.log", "user_service")
	fmt.Println("INIT >>>>>> user_service")
}

type UserService struct{}

type UserVo struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (self *UserService) Login(params interface{}) rpcserver.Success {
	logger.Print("login_params=", params)
	return rpcserver.Success{
		Sn:      "111111",
		Success: true,
	}
}

func (self *UserService) GetUser(params interface{}, token rpcserver.TOKEN) rpcserver.Success {
	logger.Print("get_user=", params, token)
	return rpcserver.Success{
		Sn:      "222222",
		Success: true,
	}
}
