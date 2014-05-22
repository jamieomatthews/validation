package validation

type User struct {
	Name    string `form:"aName custom:"cname"`
	Age     int    `form:"age" custom:"cage"`
	Email   string `form:"email" custom:"cemail"`
	Profile string `form:"profile" custom:"cprofile"`
}
