package users

type ProfilePayload struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"  binding:"required"`
	Age       int    `json:"age"       binding:"required"`
	Email     string `json:"email"     binding:"required"`
	Adress    string `json:"adress"    binding:"required"`
}
