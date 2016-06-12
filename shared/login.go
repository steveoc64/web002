package shared

type LoginCredentials struct {
	Username   string
	Password   string
	RememberMe bool
	Channel    int
}

type LoginReply struct {
	Result string
	Token  string
	Role   string
	Site   string
	ID     int
	// Menu   []UserMenu
	Routes []UserRoute
}

type UserRoute struct {
	Route string
	Func  string
}
