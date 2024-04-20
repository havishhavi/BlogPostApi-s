package dto

type Register struct {
	// validation and checking as well
	Name            string `binding:"required,min=5,max=20"`
	Email           string `binding:"required,min=5,max=20,email"`
	Password        string `binding:"required,min=5,max=10"`
	PasswordConfirm string `binding:"required,min=5,max=10,eqfield=Password"`
	Mobile          string `binding:"required,numeric,len=10"`
}

type Login struct {
	//Login details email, password
	Email    string `binding:"required,min=5,max=20,email"`
	Password string `binding:"required,min=5,max=10"`
}
