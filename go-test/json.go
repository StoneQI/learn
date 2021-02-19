package main

import "encoding/json"

func main() {

	str := `
	`
	result := map[string]interface{}
	json.Unmarshal(str,result)
}
