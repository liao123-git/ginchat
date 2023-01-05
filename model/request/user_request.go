package systemReq

// Register User register structure
type Register struct {
	Name          string `json:"name" example:"name" binding:"required"`
	Password      string `json:"password" example:"password" binding:"required"`
	ConfirmedPass string `json:"confirmed password" example:"password" binding:"required,eqfield=Password"`
	Email         string `json:"email" example:"ldl@qq.com" binding:"required,email"`
}
