package app

type User struct {
	id        int
	age       int
	firstname string `json:"first_name"`
	lastName  string `json:"last_name"`
	email     string
}
