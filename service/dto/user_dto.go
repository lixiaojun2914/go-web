package dto

type UserLoginDTO struct {
	//Name     string `json:"name" binding:"required,email,first_is_a" message:"用户名填写错误" required_err:"用户名不能为空"`
	Name     string `json:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码不能为空"`
}

type UserUpdateDTO struct {
	ID       uint   `json:"id" uri:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	RealName string `json:"real_name" form:"real_name"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

type UserListDTO struct {
	Paginate
}
