package server

import (
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"time"
	"zertu/pkg/devices"
	"zertu/pkg/util"
)

//server
type RtuServer struct {
	name        string
	rs485Device *devices.RS485Device
}

func NewRtuServer() *RtuServer {
	return &RtuServer{
		name:        "rtuServer",
		rs485Device: devices.NewRS485Device("RS485 Sensor"),
	}
}

// 工作函数
func (rs *RtuServer) work() {
	// 连接设备
	if err := rs.rs485Device.Connect("/dev/tty485_2"); err != nil {
		log.Fatalf("Failed to connect RS485 device: %v", err)
		return
	}

	// 模拟设备数据读取
	go func() {
		for {
			// 读取保持寄存器
			// 模拟读取 Modbus 数据
			address := uint16(0x0001) // Modbus 地址
			quantity := uint16(1)     // 读取数量
			data, err := rs.rs485Device.ReadHoldingRegisters(address, quantity)
			if err != nil {
				//log.Fatalf("Error reading registers: %v", err)
				log.Error("Error reading registers: ", err)
			} else {
				log.Info("Received Modbus data: ", data)
				log.Info("Received Modbus hex data: ", hex.EncodeToString(data))
				fmt.Println(util.CurrentTimeFormat(), "  Received Modbus hex data: ", hex.EncodeToString(data))
			}

			// 在这里可以进一步处理数据
			time.Sleep(1 * time.Second) // 假设设备每 2 秒发送一次数据
		}
	}()
}

func (rs *RtuServer) Start() {
	// 创建一个 Gobot 机器人，将 RS485Device 作为设备添加到机器人中
	robot := gobot.NewRobot(
		"rs485Robot",
		[]gobot.Device{rs.rs485Device},
		rs.work,
	)

	// 启动机器人
	robot.Start()
}
