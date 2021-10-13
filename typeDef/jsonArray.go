package typeDef

import "encoding/json"

type JsonArray struct {
	// []map[string]interface{} or []interface{}
	Object []interface{}
}

func EmptyJsonArray() *JsonArray {
	return &JsonArray{make([]interface{}, LEN)}
}
func NewJsonArray(jsonStr string) (*JsonArray, error) {
	var err error
	data := make([]interface{}, LEN)
	err = json.Unmarshal([]byte(jsonStr), &data)

	if err != nil {
		return nil, err
	}

	return &JsonArray{
		data,
	}, nil
}
