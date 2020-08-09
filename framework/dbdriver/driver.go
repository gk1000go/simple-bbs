package dbdriver

type SqlErrInfo struct {
	Errno uint16
	ErrMsg string
	ErrStr string
}

func GetSqlError(err error) (*SqlErrInfo) {
	return getMysqlError(err)
}

