package dao

import (
	"awesomeProject1/app-web/aputils"
	"awesomeProject1/app-web/framework"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type DaoFetchCallback func(rows *sqlx.Rows) error
type FetchResult struct {
	Kind reflect.Kind
	Data interface{}
}

func MkFetchResultSt(st interface{}) *FetchResult {
	return &FetchResult{reflect.Struct,st}
}
func MkFetchResultSl(sl []interface{}) *FetchResult {
	return &FetchResult{reflect.Slice,sl}
}
func MkFetchResultMp() *FetchResult {
	return &FetchResult{reflect.Map,nil}
}
var EndOfFetch error = errors.New("End of fetch.")

func flError(s string) error {

	msg := fmt.Sprintf("%s\n%s\n",aputils.File_line(1),s)
	log.Errorln(msg)
	return errors.New(s)
}
func hasErr(err error) bool {
	if err != nil {
		log.Errorf("%s\n%s\n",aputils.File_line(2),err.Error())
	}
	return (err != nil)
}
func Fetch(callback DaoFetchCallback,sql string,param ...interface{}) int {
	ret := 0
	defer func() {
		if ret != 0{
			log.Warnf("Some error occurred from Fetch.<%d>\n",ret)
		}
	}()
	db := framework.GetConn()
	rows,err := db.Queryx(sql,param...)
	log.Debugf("sql:%s\n",sql)
	log.Debugln(param...)
	if hasErr(err){
		ret = -1
		return ret
	}
	defer rows.Close()
	for rows.Next() {
		if err := callback(rows);err!=nil{
			ret ++
			if err == EndOfFetch{
				break
			}else{
				hasErr(err)
			}
		}
	}
	return ret
}
func FetchData(row *sqlx.Rows,data *FetchResult) (err error) {
	//typ := reflect.TypeOf(data)
	//log.Debugln("type :",typ,"Kind :",typ.Kind())
	if data.Kind == reflect.Struct {
		if aputils.IsNil(data.Data) {
			typ := reflect.TypeOf(data.Data)
			newData := reflect.New(typ.Elem())
			if hasErr(row.StructScan(newData)) {
				return row.Err()
			}
			data.Data = newData.Interface()
		}else {
			if hasErr(row.StructScan(data.Data)) {
				return row.Err()
			}
		}
	}else if(data.Kind == reflect.Slice) {
		if ret,e := row.SliceScan();hasErr(e) {
			return e
		}else{
			data.Data = ret
		}
	}else if(data.Kind == reflect.Map) {
		if aputils.IsNil(data.Data) {
			data.Data = make(map[string]interface{})
		}
		if hasErr(row.MapScan(data.Data.(map[string]interface{}))) {
			return row.Err()
		}
	}else{
		return flError("Not support result type")
	}
	return
}

func dereferenceIfPtr(value interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(value)).Interface()
}

func FetchResultDataSingleMap(sql string,param ...interface{}) (result framework.BaseDataMap){
	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultMp()
		if err := FetchData(rows,data);err==nil{
			result = (data.Data).(framework.BaseDataMap)
			return EndOfFetch
		}else{
			return err
		}
	},sql,param)
	return
}
func FetchResultDataMultiMap(sql string,key string,param ...interface{}) (result framework.ResultMaps){
	result = make(map[uint]framework.BaseDataMap)
	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultMp()
		if err := FetchData(rows,data);err==nil{
			gdata := (data.Data).(framework.BaseDataMap)
			result[gdata.UInt(key)] = gdata
			return nil
		}else{
			return err
		}
	},sql,param)
	return
}

func makePageOffset(page,pageSize int) (int,int) {
	nPage := Max(1, page)
	offset := (nPage - 1) * pageSize
	return nPage,offset
}