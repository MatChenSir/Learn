package test

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ConvertInterfaceToString() {

	//interface 转string
	//json := JSONData{}
	json2 := &JSONData{
		Code:    200,
		Message: "OK",
		Result: Result{
			AppID:        1,
			AccessToken:  "aB2XvR5wL9yOzQ8",
			ExpireTime:   1609459200,
			RefreshToken: "<PASSWORD>",
		},
	}
	stri := json2.NewJSONData()
	stri2 := convertToString(json2)
	fmt.Println("开始学习自定义 " + stri2)
	fmt.Println("开始学习json " + stri)

}

func (*JSONData) NewJSONData() string {
	json := &JSONData{
		Code:    200,
		Message: "OK",
		Result: Result{
			AppID:        1,
			AccessToken:  "aB2XvR5wL9yOzQ8",
			ExpireTime:   1609459200,
			RefreshToken: "<PASSWORD>",
		},
	}
	return interfaceTOString(json)
}

func interfaceTOString(v interface{}) string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	jsonString := string(jsonData)
	return jsonString
}

func convertToString(data interface{}) string {
	value := reflect.ValueOf(data)
	fmt.Printf("%+v\n", value.Kind())
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", value.Int())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", value.Float())
	case reflect.Ptr:
		str := ""
		// 如果是指针则获取其指向的元素
		elem := value.Elem()
		// 判断指向的元素是否为结构体类型
		if elem.Kind() == reflect.Struct {
			// 使用反射获取结构体字段名称和对应的值进行拼接
			for i := 0; i < elem.NumField(); i++ {
				field := elem.Type().Field(i)
				fieldValue := elem.Field(i)
				str += fmt.Sprintf("%s: %v, ", field.Name, convertToString(fieldValue.Interface()))
			}
			// 去除拼接结果的最后一个逗号和空格
			if len(str) > 2 {
				str = str[:len(str)-2]
			}
		}
		return str
	case reflect.Struct:
		//如果是结构体
		str := ""
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			fieldValue := value.Field(i)
			str += fmt.Sprintf("%s: %v, ", field.Name, convertToString(fieldValue.Interface()))
		}
		if len(str) > 2 {
			str = str[:len(str)-2]
		}
		return str
	case reflect.Slice:
		//如果是切片
		str := ""
		for i := 0; i < value.Len(); i++ {
			elemValue := value.Index(i)
			str += fmt.Sprintf("%v, ", convertToString(elemValue.Interface()))
		}
		if len(str) > 2 {
			str = str[:len(str)-2]
		}
		return str
	case reflect.Array:
		str := "["
		for i := 0; i < value.Len(); i++ {
			elemStr := convertToString(value.Index(i).Interface())
			str += fmt.Sprintf("%s, ", elemStr)
		}
		if len(str) > 1 {
			str = str[:len(str)-2]
		}
		str += "]"
		return str
	case reflect.Map:
		//如果是map
		str := "{"
		keys := value.MapKeys()
		for i := 0; i < len(keys); i++ {
			key := keys[i]
			keyStr := convertToString(key.Interface())
			valueStr := convertToString(value.MapIndex(key).Interface())
			str += fmt.Sprintf("%s: %s, ", keyStr, valueStr)
		}
		if len(str) > 1 {
			str = str[:len(str)-2]
		}
		str += "}"
		return str
	default:
		return fmt.Sprintf("%v", value.Interface())
	}
}
