package user

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"last_name"`
	Secret    string `json:"-"`
}
