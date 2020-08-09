package model

type Post struct {
	Tid uint
	Pid uint
	Uid uint
	Isfirst uint
	Create_date uint
	Userip uint
	Images int
	Files int
	Doctype int
	Quotepid int
	Message string
	Message_fmt string
}

type Attach struct {
	Aid uint
	Tid uint
	Pid uint
	Uid uint
	Filesize uint
	Width uint
	Height uint
	Filename string
	Orgfilename string
	Filetype string
	Create_date uint
	Comment string
	Downloads int
	Credits int
	Golds int
	Rmbs int
	Isimage int
}



type Mypost struct {
	Uid uint
	Tid uint
	Pid uint
}

type Session struct {
	Sid       uint
	Uid       uint
	Fid       uint
	Url       string
	Ip        uint
	Useragent string
	Data      string
	Bigdata   int
	Last_date uint
}

type Session_data struct {
	Sid       uint
	Last_date       uint
	Data string
}

type Modlog struct {
	Logid       uint
	Uid       uint
	Tid uint
	Pid uint
	Subject string
	Comment string
	Rmbs int
	Create_date uint
	Action string
}

type Kv struct {
	K       string
	V        string
	Expiry uint
}
type Cache struct {
	K       string
	V        string
	Expiry uint
}
type Queue struct {
	Queueid       uint
	V        int
	Expiry uint
}

type Table_day struct {
	Year uint
	Month uint
	Day uint
	Create_date uint
	Table string
	Maxid uint
	Count uint
}