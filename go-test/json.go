package main

import (
	"encoding/json"
)

func main() {

	str := `{
		"base_configer": {
			"kind": 5,
			"by_file": false,
			"data_string": {
				"basicInfo": {
					"PostId": 5,
					"tag": "tag_traffic_last_number_limit",
					"testForbid": "1",
					"beginDate": "2021-03-16 00:00:00",
					"endDate": "2021-03-17 23:59:00",
					"dayType": [
						"1",
						"2",
						"3"
					],
					"daysOfWeek": [
						"1",
						"2",
						"3",
						"4",
						"5",
						"6",
						"7"
					],
					"area": "1",
					"toArea": "1",
					"orderType": "3",
					"orderProductTemp": [
						"2",
						"1"
					],
					"driverProductTemp": [
						"2",
						"1"
					],
					"name": "测试测试-ziqi",
					"warnLevel": 200,
					"showTxt": "123",
					"desc": "213"
				},
				"matchInfo": {
					"driverValueId": "",
					"orderValueId": "",
					"matchCondition": 2,
					"city_mode": 2
				},
				"orderInfo": [
					[
						[
							{
								"key": "dest_region",
								"type": "region_struct",
								"group": 1,
								"operator": "sub",
								"value": [
									"{\"user_type\":2,\"big_regions\":[\"281691\"],\"regions\":[\"281697\"]}"
								]
							}
						]
					]
				],
				"driverInfo": [
					[
						[
							{
								"key": "limit_plate_no_type",
								"type": "plate_struct",
								"group": 4,
								"operator": "week_limit",
								"value": [
									"{\"plate_no_type\":\"prefix\",\"detail\":[{\"not_prefix\":[],\"prefix\":[\"1\"],\"index\":1},{\"not_prefix\":[],\"prefix\":[\"2\"],\"index\":2},{\"not_prefix\":[],\"prefix\":[\"3\"],\"index\":3},{\"not_prefix\":[],\"prefix\":[\"4\"],\"index\":4},{\"not_prefix\":[],\"prefix\":[\"4\"],\"index\":5},{\"not_prefix\":[],\"prefix\":[\"5\"],\"index\":6},{\"not_prefix\":[],\"prefix\":[\"4\"],\"index\":7}]}"
								]
							}
						]
					]
				],
				"extend": 1
			}
		}
	}
	`
	result := make(map[string]interface{})
	json.Unmarshal([]byte(str), &result)
	// aa, err := sql.Open("", 2)
	aa, b := result["data_string"].(map[string]interface{})
	print(aa, b)
}

func a(a int, b int) int {
	if a > 10 {
		return b(a+10, b)
	}
}

func b(a int, b int) int {
	if b < 10 {
		return c(a, b+10)
	}
}
func c(a int, b int) int {
	return a + b
}
