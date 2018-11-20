package binds

// Login login info bind
type Login struct {
	UserName    string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password 	string `form:"password" json:"password" xml:"password"  binding:"required"`
	Type 		string `form:"type" json:"type" xml:"type"`
}
