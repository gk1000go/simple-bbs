package controllers

import (
	"awesomeProject1/app-web/framework"
	log "github.com/sirupsen/logrus"
	"reflect"
)


type initInfo struct {
	Count  int
	Object framework.IController
}

var definedControllers = map[string]*initInfo{
	"HomeController":&initInfo{0,&HomeController{}},
	"UserController":&initInfo{0,&UserController{}},
}

func Ctrl(name string,new bool) framework.IController {
	info,ok := definedControllers[name]
	if new {
		ok = true
		typ := reflect.TypeOf(info.Object).Elem() // *HomeController
		newCtrl := reflect.New(typ)
		//return unsafe.Pointer(&newCtrl)
		//log.Debugf("reflect new before :%p\n",info.Object)
		info.Object = newCtrl.Interface().(framework.IController)
		//log.Debugf("reflect new after :%p\n",info.Object)
		//return newCtrl.Interface().(framework.IController)
	}
	if !ok {
		log.Warnf("Get %s Failed.\n",name)
		return nil
	}
	if info.Count > 0 {
		log.Warnf(">>> Use %s Plural.\n",name)
	}
	info.Count ++
	return info.Object
}