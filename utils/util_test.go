package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestMarshalStruct struct {
	Name string           `json:"name"`
	Kv   map[int64]string `json:"kv"`
}

func TestMarshal(t *testing.T) {

	obj := TestMarshalStruct{
		Name: "hello",
		Kv:   map[int64]string{1: "abc"},
	}

	// utils.PrettyPrint(obj)
	fmt.Printf("TestMarshal: %+v \n", obj)

	body, err := json.Marshal(obj)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TestMarshal: %s \n", string(body))
}
