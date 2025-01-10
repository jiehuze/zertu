package main

import (
	"fmt"
	"os"
	"zertu/pkg/logger"
	"zertu/pkg/server"
)

var (
	VersionInfo string
	BuildInfo   string
)

func main() {
	var runMode string
	if len(os.Args) > 1 {
		runMode = os.Args[1]
	} else {
		runMode = "dev"
	}
	logger.Init()
	// 打印运行信息，logger未初始化，使用fmt
	fmt.Println("------------------------------------------------")
	fmt.Printf("     build at: %s\n     version: %s\n     runMode: %s\n", BuildInfo, VersionInfo, runMode)
	fmt.Println("------------------------------------------------")

	rtuServer := server.NewRtuServer()
	rtuServer.Start()
}
