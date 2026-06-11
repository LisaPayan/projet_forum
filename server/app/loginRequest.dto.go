package dto

type loginrequestdto struct {
	username string 'json:username'
	password string 'json:password'
}