package xmltojson



func JsonFromXml(path string, base interface{}) string  {

	xmlData := contentFromXmlFile(path)
	
	xmlUnmarshal(xmlData, base)

	jsonData := jsonString(base)

	return jsonData
}