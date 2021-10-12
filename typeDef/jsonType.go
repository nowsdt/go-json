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

func NewJsonObject(jsonStr string) (*JsonObject, error) {
	var err error
	data := make(map[string]interface{}, 8)
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	}
	return &JsonObject{data}, nil
}

func newJsonArray() *JsonArray {
	return &JsonArray{
		make([]map[string]interface{}, 8),
	}
}
