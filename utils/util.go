package utils

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

func PrettyPrint(obj interface{}) {
	body := PrettyJson(obj)
	fmt.Printf("%s\n", body)
}

func PrettyJson(obj interface{}) string {
	indent := "\t"
	body, err := json.MarshalIndent(obj, "", indent)

	if err != nil {
		//panic(err)
		fmt.Println(err)
		return ""
	}

	return string(body)
}
