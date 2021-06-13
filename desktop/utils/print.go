package utils

import (
	"encoding/json"
	"fmt"
)

func InterfaceToString(i interface{}) string {
	structBytes, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}
	return string(structBytes)
}
