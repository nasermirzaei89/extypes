package extypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSONObject struct {
	v interface{}
}

func NewJSONObject(v interface{}) JSONObject {
	return JSONObject{v: v}
}

func (o *JSONObject) Scan(src interface{}) error {
	var b []byte
	switch v := src.(type) {
	case []byte:
		b = v
	case string:
		b = []byte(v)
	case nil:
		b = make([]byte, 0)
	default:
		return errors.New("unsupported type")
	}

	if len(b) == 0 {
		// this is empty json field. simply create an empty map
		o.v = nil

		return nil
	}

	return json.Unmarshal(b, &o.v)
}

func (o JSONObject) Value() (driver.Value, error) {
	return json.Marshal(o.v)
}

func (o *JSONObject) Set(v interface{}) {
	o.v = v
}

func (o JSONObject) Get() interface{} {
	return o.v
}

func (o JSONObject) GetStringSlice() []string {
	if o.v == nil {
		return nil
	}

	vv, ok := o.v.([]interface{})
	if !ok {
		return nil
	}

	res := make([]string, len(vv))

	for i := range vv {
		v, ok := vv[i].(string)
		if !ok {
			return nil
		}

		res[i] = v
	}

	return res
}

func (o JSONObject) GetStringInterfaceMap() map[string]interface{} {
	if o.v == nil {
		return nil
	}

	return o.v.(map[string]interface{})
}

func (o JSONObject) String() (string, error) {
	res, err := json.Marshal(o.v)
	if err != nil {
		return "", fmt.Errorf("error on marshal: %w", err)
	}

	return string(res), nil
}

func (o JSONObject) Decode(v interface{}) error {
	b, err := json.Marshal(o.v)
	if err != nil {
		return fmt.Errorf("error on marshal: %w", err)
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return fmt.Errorf("error on unmarshal: %w", err)
	}

	return nil
}
