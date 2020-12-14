package common

import (
	"encoding/json"
	"fmt"
)

func Unmarshal(txt string, obj interface{}) {
	errs := json.Unmarshal([]byte(txt), &obj)

	if errs != nil {
		fmt.Println(errs)
		//fmt.Println(errs.Error())
		return
	}

}

func Marshal(obj interface{}) []byte {
	body, err := json.Marshal(obj)
	if err != nil {
		panic(err.Error())
	}

	return body
}

func ToJson(obj interface{}) string {
	return string(Marshal(obj))
}

func ToBeautifyJson(obj interface{}) string {
	body, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(body)
}
