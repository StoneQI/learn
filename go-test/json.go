package main

import (
	"database/sql"
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
func setupServerPostHandler(t *testing.T) *gin.Engine {
	
	
	
	engine := gin.New()
	//engine.Use(middleware.Logger())
	engine.Use(gin.Recovery())
	PostDaoMock,issql:= GetPostDao(t)
	if !issql {
		... mock 促使化
		
	}
	server := NewServer()
	// 唯一依赖
	server.SetPostService(PostDaoMock,RedisClient)
	
	engine.POST("/Post", server.GetPosts)
	return engine
}

func TestPostHandler(t *testing.T) {
	router := setupServerPostHandler(t)
	Convey("Post Handler接口测试",t, func() {
		req_content := &Post{
			... 内容
		}
		type Data_resp struct {
			Post_id int `json:Post_id`
		}
		type resp_json struct {
			Data  Data_resp `json:data`
			Err_msg  string `json:err_msg`
			Err_no   int `json:err_no`
		}

		req_content.Type = "AddPost"
		Convey("AddPost 测试", func() {

			Convey("AddPost 测试1", func() {

				req_new := req_content
				req_string, _ := json.Marshal(req_new)
				req := httptest.NewRequest(http.MethodPost, "/AddPost", strings.NewReader(string(req_string)))
				req.Header[global.HEADER_TRACEID] = []string{"testTrace"}
				req.Header[global.HEADER_SPANID] = []string{"testSpan"}
				req.Header[global.HEADER_USER] = []string{"testUser"}
				req.Header.Set("Content-Type","application/json")

				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				resp := w.Result()
				resp_json1 := &resp_json{}
				_ = json.Unmarshal(w.Body.Bytes(), resp_json1)
				So(resp_json1.Data.Post_id,ShouldHaveSameTypeAs,1)
				So(resp.StatusCode,ShouldEqual,http.StatusOK)
			})
		})
	})
}


