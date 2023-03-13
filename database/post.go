// Contains the DB structure of post table
package database

type Post struct {
	Id      int    `json:"id" pg:"id,pk"`
	Title   string `json:"title" pg:"title,notnull"`
	Content string `json:"content" pg:"content,notnull"`
	UserId  int    `json:"uid" `
}

type Likes struct {
	PostId int   `josn:"pid"`
	Users  []int `json:"users"`
}
