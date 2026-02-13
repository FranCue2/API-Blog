package models

type Role string

const(
	RoleAdmin	Role = "admin"
	RoleUser	Role = "user"
)

func (r Role) IsValid() bool{
	switch r {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}