package model

func u(d uint) *uint {
	a := new(uint)
	*a = d
	return a
}

func i(d int) *int {
	a := new(int)
	*a = d
	return a
}

var All []interface{}=[]interface{}{
	&User{},
	&Group{},
	&Forum{},
	&Forum_access{},
	&Thread{},
	&Thread_top{},
	&Post{},
	&Attach{},
	&Mythread{},
	&Mypost{},
	&Session{},
	&Session_data{},
	&Modlog{},
	&Kv{},
	&Cache{},
	&Queue{},
	&Table_day{},
}