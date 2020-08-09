package framework

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func (ap *AppBase)Run(r IAppRouter) {
	if r != nil{
		r.InitRouter()
	}else {
		log.Warn("No AppRouter...")
		ap.InitRouter()
	}

	ap.Server = &http.Server{
		Addr:           ":8080",
		Handler:        siteApp.GinEngine,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if(ap.Server == nil){
		log.Println("initServer failed...")
		os.Exit(-1)
	}

	//go r.Run(":8080")
	go ap.Server.ListenAndServe()
}

func (ap *AppBase)CloseApp()  {
	closeAllConn()
}