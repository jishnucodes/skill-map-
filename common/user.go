package common



type UserCreationInput struct {
	Name     string `json:"name"` //this is called field tag (the format of object)
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateInput struct {
	Name     string `json:"name"` //this is called field tag (the format of object)
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSigninInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

func NewUserUpdateInput() *UserUpdateInput {
	return &UserUpdateInput{}
}

func NewUserSigninInput() *UserSigninInput {
	return &UserSigninInput{}
}



