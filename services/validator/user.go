package validator

type User struct {
	Username string `binding:"required,min=2,max=20"`
	Password string `binding:"required,min=6,max=32"`
	AvatarId string `binding:"required,numeric"`
}
