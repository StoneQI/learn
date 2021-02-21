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
					"templateId": 5,
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

type PostMysql struct{
    db *sql.DB
}

func NewPostMysql(db *sql.DB){
	return &PostMysql{
		db: db,
	}
}


func (s *PostMysql) FindPostByID(ctx context.Context, id int) (*Post, error){
	...
}

func (s *PostMysql)     FindPosts(ctx context.Context, filter PostFilter) ([]*Post, int, error){
	s.db.Exec("...",...)
}

func (s *PostMysql)     CreatePost(ctx context.Context, post *Post) error{
	...
}

func (s *PostMysql)     UpdatePost(ctx context.Context, id int, upd PostUpdate) (*Post, error){
	...
}

func (s *PostMysql)     DeletePost(ctx context.Context, id int) error
{
	...
}

type PostService struct{
	postDAO *PostDAO
}

func NewPostService(postDAO *PostDAO)  {
	return &PostService{
		postDAO: postDAO,
	}	
} 

func (post *PostService)GetAllPost()  {
	post.postDAO.FindPosts()
	....
} 



type Service struct{
	postDAO *PostDAO
}

func NewPostService(postDAO *PostDAO)  {
	return &PostService{
		postDAO: postDAO,
	}	
} 

func (post *PostService)GetAllPost()  {
	post.postDAO.FindPosts()
	....
} 


