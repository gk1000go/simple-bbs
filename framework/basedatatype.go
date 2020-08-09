package framework

type BaseDataMap map[string]interface{}
type ResultMaps map[uint]BaseDataMap
type ResultSlices []BaseDataMap

func Conv2DataMap(org map[string]interface{}) BaseDataMap {
	return BaseDataMap(org)
}

func (b BaseDataMap)Int(key string) int {
	if v,ok := b[key];ok{
		return v.(int)
	}
	return 0
}
func (b BaseDataMap)UInt(key string) uint {
	if v,ok := b[key];ok{
		return v.(uint)
	}
	return 0
}
func (b BaseDataMap)Str(key string) string {
	if v,ok := b[key];ok{
		return v.(string)
	}
	return ""
}