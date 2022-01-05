package utils

import "encoding/json"

func Map2struct(data map[string]interface{}, st interface{}) error {
	// map 转 json
	jstr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// json 转 struct
	err = json.Unmarshal(jstr, st)
	if err != nil {
		return err
	}

	return err
}
