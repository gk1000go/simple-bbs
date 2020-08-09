package cache

var errLast error
func LastError() error {
	return errLast
}

func CacheDelete(key string) bool {

	return true
}
