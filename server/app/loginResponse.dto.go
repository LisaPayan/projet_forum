package dto 

type loginresponsedto struct {
	Type string 'json:type'
	AcessToken string 'json:access_token'
	Expiresin int 'json:expires_in'
}