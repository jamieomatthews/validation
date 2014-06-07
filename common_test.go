package validation

type User struct {
	Name       string `form:"aName custom:"cname"`
	Age        int    `form:"age" custom:"cage"`
	Email      string `form:"email" custom:"cemail"`
	Profile    string `form:"profile" custom:"cprofile"`
	CreditCard string `form:"credit"`
	PageURL    string `form:"url"`
	NickName   string `form:"nick_name"`
	Weight     int    `form:"weight"`
}
