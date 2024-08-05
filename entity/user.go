package entity

type User struct {
	Name string `json:"name"`
}

func (u *User) GetName() string {
	return u.Name
}
