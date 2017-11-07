package main

import (
	"fmt"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
	"github.com/gwuhaolin/lightsocks/server"
	"log"
	"net"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 服务端监听端口随机生成
	port := 7448

	// 默认配置
	config := &cmd.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
		// 密码随机生成
		Password: core.RandPassword().String(),
	}
	config.ReadConfig()
	config.SaveConfig()

	// 解析配置
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// 启动 server 端并监听
	lsServer := server.New(password, listenAddr)
	log.Fatalln(lsServer.Listen(func(listenAddr net.Addr) {
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
密码 password：
%s
	`, listenAddr, password))
		log.Printf("lightsocks-server:%s 启动成功 监听在 %s\n", version, listenAddr.String())
	}))
}
