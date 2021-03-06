package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type SignedUp struct{
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

// Data model
type Video struct {
	Id string
	AuthorId int
	Name string
	DisplayCreateTime string
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}


type SimpleSession struct {
	Username string
	TTL int64
}