package xmltojson

import (
	"encoding/json"
	"fmt"
)

func jsonString(base interface{}) string {

	str, err := json.Marshal(base)

	if err != nil {
		fmt.Print("Erro ao criar json: ", err)
	}

	return string(str)
}