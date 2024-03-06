package auth

type SignUpPayload struct {
	FirstName   string `json:"firstName"   binding:"required"`
	LastName    string `json:"lastName"    binding:"required"`
	Password    string `json:"password"    binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	IsActive    bool   `json:"isActive"`
}

type LoginPayload struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}
