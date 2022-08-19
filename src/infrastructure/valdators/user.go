package valdators

type UserModelValidator struct {
	User struct {
		ID string `binding:"alphanum,len=7`
	}
}

type SignupUserModelValidator struct {
	User struct {
		name string `binding:"alphanum,min=4,max=13"`
	}
}
