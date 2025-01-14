package jsoniter

import (
	"fmt"
	"testing"
)

func TestInt64MapUnmarshal(t *testing.T) {
	var str = `{"20100001":"2","20100003":1,"20300003":1,"20500003":1}`
	var m map[int64]int
	if err := jsonUnmarshal(&m, str, ""); err != nil {
		t.Fatal(err)
	}
	t.Log(m)
}
func jsonUnmarshal(data interface{}, jsonString string, tagKey string) error {
	var newJson = Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 tagKey,
		ConvertStringTo64:      true,
	}.Froze()
	err := newJson.UnmarshalFromString(jsonString, data)
	if err != nil {
		return fmt.Errorf("got error:%w while unmarshal data", err)
	}
	return nil
}
