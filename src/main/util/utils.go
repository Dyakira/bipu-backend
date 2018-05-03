package util

import (
	"reflect"
	"strings"
)

//获取某结构体中某字段的某Tag的值，如果没有则返回空字符串
func GetStructTagName(structName interface{}, fieldName string, tagKey string) (tagVal string) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}

	// 遍历字段
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		if 0 == strings.Compare(fieldName, t.Field(i).Name) {
			tagVal = t.Field(i).Tag.Get(tagKey)
			return tagVal
		}
	}
	return ""
}

func GetStructTagName4Gorm(structName interface{}, fieldName string, tagKey string) string {
	// 这里得到的是一个grom的val字符串，里面的各个键又被分号隔开，有的是冒号连接的形式，有的没有，比如：
	// type:varchar(100);unique_index
	gormStr := GetStructTagName(structName, fieldName, "gorm")
	gormStrArr := strings.Split(gormStr, ";")

	for i := 0; i < len(gormStrArr); i++ {
		if strings.HasPrefix(gormStrArr[i], tagKey) {
			return strings.Split(gormStrArr[i], ":")[1]
		}
	}
	return ""
}
