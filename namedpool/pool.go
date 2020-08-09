package namedpool

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

//var dataPool *sync.Pool

const NAME0 = `package<namedpool>`


var mapDataPool sync.Map
type NewPoolFunc func() interface{}

//func init()  {
//	mapDataPool = make(map[string]*sync.Pool)
//}

func New(named string,funcNew NewPoolFunc) {
	if _, ok := mapDataPool.Load(named); !ok {
		dataPool := &sync.Pool{
			New: funcNew,
		}
		mapDataPool.Store(named,dataPool)
	}
}

func Get(named string) interface{} {
	if pool := checkPool(named); pool !=nil {

		p := pool.Get()
		if (p == nil) { // intアサーションに失敗する
			log.Fatalf("%s name(%s) pool(Get) error.\n",NAME0,named)
		}
		if d,ok := p.(BasePoolFunc);ok{
			d.Reset()
		}else{
			log.Warnf("%s pool(Get) :name(%s) is not exists(Reset).\n",NAME0,named)
			return nil
		}
		return p
	}
	return nil
}

func Put(named string,p interface{}) {
	if pool := checkPool(named); pool !=nil {
		if d,ok := p.(BasePoolFunc);ok{
			d.Clear()
		}else{
			log.Warnf("%s pool(Put) :name(%s) is not exists(Clear).\n",NAME0,named)
			return
		}
		pool.Put(p)
	}else{
		log.Warnf("%s pool(Put) :name(%s) is not exists.\n",NAME0,named)
	}
}
func checkPool(named string) *sync.Pool{
	pool, ok := mapDataPool.Load(named)
	if !ok {
		log.Warnf("%s pool :name(%s) is not exists.\n",NAME0,named)
		return nil
	}
	return pool.(*sync.Pool)
}