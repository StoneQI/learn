package main

import "encoding/json"

func main() {

	str := `123
	123
	`
	result := map[string]interface{}
	json.Unmarshal(str,result)
}
