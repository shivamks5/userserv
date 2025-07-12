package model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type IDNumber struct {
	ID string `json:"id"`
}

type Response struct {
	Data interface{} `json:"data"`
}
