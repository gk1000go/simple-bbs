package services

import (
	"awesomeProject1/app-web/framework"
	"awesomeProject1/app-web/lang"
	"encoding/binary"
	"net"
	"strconv"
	"time"
)

//https://github.com/syyongx/php2go

// default format:"0000-00-00"
func IntToTime(inTime int64,f ...string) string {
	//i, err := strconv.ParseInt("1405544146", 10, 64)
	//if err != nil {
	//	panic(err)
	//}
	timeFmt := "0000-00-00"
	if inTime == 0 {return timeFmt}
	if len(f) > 0 {timeFmt = f[0]}
	loc, _ := time.LoadLocation("Asia/Tokyo")
	tm := time.Unix(inTime, 0).In(loc)
	return tm.Format(timeFmt)
}
// default format:"0000-00-00"
func UIntToTime(inTime uint,f ...string) string {
	return IntToTime(int64(inTime),f...)
}

// Long2ip long2ip()
// IPv4
func Long2ip(properAddress uint32) string {
	if properAddress <=0 {return ""}
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}
func Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

func Str2Uint(s string) uint {
	if v,e := strconv.Atoi(s); e== nil{
		return uint(v)
	}
	return 0
}

func GetTimeStamp() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	t := time.Now().In(loc)
	return t.Format("20060102150405")
}
func GetTodaysDate() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	current_time := time.Now().In(loc)
	return current_time.Format("2006-01-02")
}

func GetTodaysDateTime() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	current_time := time.Now().In(loc)
	return current_time.Format("2006-01-02 15:04:05")
}

func GetTodaysDateTimeFormatted() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	current_time := time.Now().In(loc)
	return current_time.Format("Jan 2, 2006 at 3:04 PM")
}

func GetTimeStampFromDate(dtformat string) string {
	form := "Jan 2, 2006 at 3:04 PM"
	t2, _ := time.Parse(form, dtformat)
	return t2.Format("20060102150405")
}

func GetGuestUser() map[string]interface{} {
	return map[string]interface{}{
		"uid" : 0,
		"gid" : 0,
		"groupname" : lang.Lang["guest_group"],
		"username" : lang.Lang["guest"],
		"avatar_url" : framework.GetApp().Cfg.BaseSite.View_url+"img/avatar.png",
		"create_ip_fmt" : "",
		"create_date_fmt" : "",
		"login_date_fmt" : "",
		"email" : "",
		"threads" : 0,
		"posts" : 0,
	}
}