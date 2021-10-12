package typeDef

import "encoding/json"

type JsonObject struct {
	Object map[string]interface{}
}

type JsonArray struct {
	Object []map[string]interface{}
}

func EmptyJsonObject() *JsonObject {
	return &JsonObject{
		make(map[string]interface{}, 1),
	}
}

func NewObject(jsonStr string) (*JsonObject, error) {
	var err error
	data := make(map[string]interface{}, 8)
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	}
	return &JsonObject{data}, nil
}

func EmptyJsonArray() *JsonArray {
	return &JsonArray{
		make([]map[string]interface{}, 8),
	}
}
func NewJsonArray(jsonStr string) (*JsonArray, error) {
	var err error
	data := make(map[string]interface{}, 8)
	err = json.Unmarshal([]byte(jsonStr), &data)

	if err != nil {
		return nil, err
	}

	return &JsonArray{
		make([]map[string]interface{}, 8),
	}, nil
}
