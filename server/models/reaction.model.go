package models

type Reaction struct {
	Type_reac string  `json:"type_reac"`
	User_c    User    `json:"user_c"`
	Message_c Message `json:"message_c"`
}
