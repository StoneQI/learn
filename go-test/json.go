package main

import (
	"encoding/json"
)

func main() {

	str := `123
	123
	`
	result := new(map[string]interface{})
	json.Unmarshal([]byte(str), result)

	print(result)
}
