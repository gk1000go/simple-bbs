package dbpool

import (
	"awesomeProject1/app-web/app/trash/dbcore"
	"sync"
)


var dbParamPool *sync.Pool
func init()  {
	dbParamPool = &sync.Pool{
		New: func() interface{} {
			return &dbcore.DbParam{}
		},
	}
}

func Get() *dbcore.DbParam {
	p, ok := dbParamPool.Get().(*dbcore.DbParam)
	if !ok { // intアサーションに失敗する
		panic("なんか変")
	}

	p.Reset()
	return p
}

func Put(p *dbcore.DbParam) {
	p.Clear()
	dbParamPool.Put(p)
}