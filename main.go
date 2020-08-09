package main

import (
	"awesomeProject1/app-web/conf"
	"awesomeProject1/app-web/framework"
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	// JSONフォーマット
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})

	// 標準エラー出力でなく標準出力とする
	log.SetOutput(os.Stdout)

	// Warningレベル以上を出力
	log.SetLevel(log.DebugLevel)
}

func initConf() bool {
	fConf := "./conf/config.default.toml"
	err := conf.InitialConfig(fConf)
	if err == nil {
		log.Println("fConf load OK")
		//log.Println(conf.DefaultConfig)
		//log.Println(len(conf.DefaultConfig.BaseDb.Slaves))
		log.Println(conf.SiteConf)
		return true
	}
	return false
}

func main() {
	if(!initConf()){
		log.Println("initConf failed...")
		os.Exit(-2)
	}

	r := &MyRouter{"hello world!"}

	app := framework.NewApp(&conf.DefaultConfig)
	app.Run(r)

	gracefulExitWeb()
}

func gracefulExitWeb() {
	app := framework.GetApp()
	server := app.Server

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	log.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil{
		log.Println("err", err)
	}
	app.CloseApp()

	// 看看实际退出所耗费的时间
	log.Println("------exited--------", time.Since(now))
}

func log_test() {

	// ログ例1（出力されない）
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// ログ例2（出力される）
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// ログ例3（出力される）
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// 共通的に使用する変数の設定
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other": "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}