package main

import (
	"fmt"
	"log"
	"os"
	"staff-mall-center/pkg/setting"
	"staff-mall-center/src/router"

	cli "gopkg.in/urfave/cli.v1"
)

const APP_VER = "0.0.1"

func init() {
	setting.AppVer = APP_VER
}

func main() {
	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "Hualvmall"
	//程序的用途描述
	app.Usage = "Staff Mall BackEnd"
	//程序的版本号
	app.Version = "1.0.0"

	//预置变量

	//TODO 这里考虑最后放进config配置里面
	//设置启动参数
	app.Flags = []cli.Flag{
		//参数类型string,int,bool
		cli.StringFlag{
			Name:        "host",           //参数名字
			Value:       "127.0.0.1",      //参数默认值
			Usage:       "Server Address", //参数功能描述
			Destination: &setting.Host,    //接收值的变量
		},
		cli.IntFlag{
			Name:        "port,p",
			Value:       8888,
			Usage:       "Server port",
			Destination: &setting.Port, //接收值的
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug mode",
			Destination: &setting.DevMode,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "start",
			Aliases:  []string{"start"},
			Usage:    "start server",
			Category: "arithmetic",
			Action: func(c *cli.Context) error {
				host, port, debug := setting.Host, setting.Port, setting.DevMode
				fmt.Printf("host启动为：%v，端口号启动为：%v，debug模式为 %v \n", host, port, debug)

				router.Start()
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
