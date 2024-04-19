package entity

type UserTable struct {
	UserID         string   `json:"user_id" dynamodbav:"user_id"`
	UserName       string   `json:"user_name" dynamodbav:"user_name"`
	Email          string   `json:"email" dynamodbav:"email"`
	CreateTime     int64    `json:"create_time" dynamodbav:"create_time"`
	UpdateTime     int64    `json:"update_time" dynamodbav:"update_time"`
	PreferredGenre []string `json:"preferred_genre,omitempty" dynamodbav:"preferred_genre,omitempty"`
}
