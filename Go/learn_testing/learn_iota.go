package learntesting

type Role int

const (
	RoleGuest Role = iota
	RoleUser
	RoleModerator
	RoleAdmin
)

type Priority int

const (
	Unknown Priority = iota
	Low
	Medium
	High
	Critical
)

func ReturnIntOfRole(role Role) int {
	return int(role)
}

func ReturnIsAdmin(role Role) bool {
	return role == RoleAdmin
}

func RequiresImmediateAttention(p Priority) bool {
	return p == Critical || p == High
}
