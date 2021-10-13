package typeDef

import (
	"encoding/json"
)

type JsonObject struct {
	Object map[string]interface{}
}

func EmptyJsonObject() *JsonObject {
	return &JsonObject{
		make(map[string]interface{}, LEN),
	}
}

func NewObject(jsonStr string) (*JsonObject, error) {
	var err error
	data := make(map[string]interface{}, LEN)
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	}
	return &JsonObject{data}, nil
}

func getVal(jsonObject *JsonObject, key string) interface{} {
	if v, present := jsonObject.Object[key]; present {
		return v
	}
	return nil
}

func (jsonObject *JsonObject) GetString(key string) string {
	val := getVal(jsonObject, key)
	if val != nil {
		return val.(string)
	}
	return ""
}

func (jsonObject *JsonObject) GetFloat(key string) float64 {
	val := getVal(jsonObject, key)
	if val != nil {
		return val.(float64)
	}
	return 0
}

func (jsonObject *JsonObject) GetBool(key string) bool {
	val := getVal(jsonObject, key)
	if val != nil {
		return val.(bool)
	}
	return false
}

func (jsonObject *JsonObject) GetObject(key string) *JsonObject {
	val := getVal(jsonObject, key)
	if val != nil {
		return &JsonObject{val.(map[string]interface{})}
	}
	return nil
}

func (jsonObject *JsonObject) GetArray(key string) []interface{} {
	val := getVal(jsonObject, key)
	if val != nil {
		return val.([]interface{})
	}
	return nil
}
