package todogo

type RegisterUser struct {
	Id       int    `json:"-"`
	Username string `json:"username" binding:"required" valid:"-"`
	Email    string `json:"email" binding:"required" valid:"email"`
	Password string `json:"password" binding:"required" valid:"-"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required" valid:"-"`
	Password string `json:"password" binding:"required" valid:"-"`
}

type JWTToken struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type RefreshToken struct {
	Refresh string `json:"refresh" valid:"-"`
}
