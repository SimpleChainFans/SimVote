package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.dataqin.com/sipc/vote_app/cmd"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// @title vote_app
// @version 0.0.1
// @description 投票助手
// @author znddzxx112@163.com

// @schemes {{scheme}}
// @host {{host}}
// @BasePath /api/

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	setLogLevel()
	setDefaultLocation()

	app := cli.NewApp()
	app.Version = cmd.Version
	app.Name = cmd.VoteName
	configPathFlag := cli.StringFlag{
		Name:   "configPath",
		Usage:  "config file path",
		EnvVar: "configPath",
	}
	portFlag := cli.StringFlag{
		Name:   "port",
		Usage:  "port",
		EnvVar: "port",
		Value:  "7688",
	}
	app.Flags = []cli.Flag{
		configPathFlag,
		portFlag,
	}
	app.Action = Start
	go cmd.CronTask()
	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func Start(ctx *cli.Context) error {
	sigCh := make(chan os.Signal)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	configPath := ctx.GlobalString("configPath")

	server := cmd.NewServer(cmd.VoteName)

	go func() {
		select {
		case <-sigCh:
			if err := server.Close(); err != nil {
				logrus.Println(err.Error())
			}
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}()
	err := server.Run(ctx.GlobalString("port"), configPath)
	if err != nil {
		return err
	}
	return nil
}

func setLogLevel() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "level",
		},
		TimestampFormat: time.RFC3339Nano,
	})
	logrus.SetLevel(logrus.InfoLevel)
}
func setDefaultLocation() {
	time.Local = time.FixedZone("UTC", 8*3600)
}
