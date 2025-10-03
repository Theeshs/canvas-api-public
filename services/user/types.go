package user

type User struct {
	Email              string  `json:"email"`
	Password           string  `json:"password"`
	UserName           string  `json:"username"`
	FirstName          *string `json:"first_name"`
	LastName           *string `json:"last_name"`
	DOB                *string `json:"dob"`
	GithubUsername     *string `json:"github_username"`
	Description        *string `json:"description"`
	MobileNumber       *int    `json:"mobile_number"`
	AddresBlock        *string `json:"address_block"`
	AddressStreet      *string `json:"address_street"`
	ResidentialCountry *string `json:"recidential_country"`
	Nationality        *string `json:"nationality"`
}

type EmailMessage struct {
	UserEmail string `json:"email"`
	Name      string `json:"name"`
	Message   string `json:"message"`
}
