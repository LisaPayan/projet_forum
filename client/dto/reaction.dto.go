package dto

type Reaction struct {
	Type_reac string     `json:"type_reac"`
	User_c    UserDto    `json:"user_c"`
	Message_c MessageDto `json:"message_c"`
}
