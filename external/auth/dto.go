package auth

type SendRegisterDonorAccountEmailRequestDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
