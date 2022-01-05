package entities

type UserCreateParams struct {
	Name     string `yaml:"name" json:"name" form:"name" query:"name"`
	Email    string `yaml:"email" json:"email" form:"email" query:"email"`
	Gender   string `yaml:"gender" json:"gender" form:"gender" query:"gender"`
	Age      uint   `yaml:"age" json:"age" form:"age" query:"age"`
	Birthday string `yaml:"birthday" json:"birthday" form:"birthday" query:"birthday"`
	Password string `yaml:"password" json:"password" form:"password" query:"password"`
}
