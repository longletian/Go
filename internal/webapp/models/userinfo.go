package models

type Userinfo struct {
	Username string `form:"username" json:"username",binding:"required"`
	Password string `form:"password",json:"password",binding:"required"`
	Sex      int    `form:"sex",json:"sex"`
	Name string  `form:"name",json:"name"`
	Email string  `form:"email",json:"email"`
	Phone string  `form:"phone",json:"phone"`
}

func AddUser(userinfo Userinfo)(bool,error) {


}

