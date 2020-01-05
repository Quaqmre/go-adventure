package wrappingService

//User is a our custemers branchs
type User struct {
	id   int
	name string
}

type UserService interface {
	UserName(int) string
	ChangeUserName(int) bool
}


