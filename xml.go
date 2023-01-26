package xmltojson 

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
)
 

func contentFromXmlFile(path string) []byte {

	data, err := ioutil.ReadFile(path)
 
	if err != nil {
		fmt.Printf("Erro ao carregar arquivo")
	}

	return data
}


func xmlUnmarshal(xmlData []byte, base interface{}) {
	xml.Unmarshal(xmlData, &base)
}

