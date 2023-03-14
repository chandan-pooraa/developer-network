// Contains the DB structure of comment table

package database

type Comment struct {
	Id		int		`json:"id" pg:"id,pk"`
	Content	string	`json:"content"`
	UserId	int		`json:"uid"`
	PostId  int		`json:"pid"`
}
