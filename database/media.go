package database

/*
	name: media -- tableName
*/

type Media struct {
	// tableName struct{} `sql:"media"` --> to alter tablename in database
	ID       uint64 `json:"id" pg:"id,pk"`
	Name     string `json:"name" pg:"name,notnull"`
	Url      string `json:"url" pg:"url,unique,notnull"`
	Category string `json:"category" pg:"category,notnull"`
}
