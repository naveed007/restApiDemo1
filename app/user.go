package app

type User struct {
	id        int    `json:"id"`
	age       int    `json:"age"`
	firstname string `json:"first_name"`
	lastName  string `json:"last_name"`
	email     string
}
