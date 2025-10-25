package public

import (
	"api/services/educations"
	"api/services/experiences"
	"api/services/projects"
)

type UserDetails struct {
	Email              string                             `json:"email"`
	UserName           string                             `json:"username"`
	FirstName          *string                            `json:"first_name"`
	LastName           *string                            `json:"last_name"`
	DOB                *string                            `json:"dob"`
	GithubUsername     *string                            `json:"github_username"`
	Description        *string                            `json:"description"`
	MobileNumber       *int                               `json:"mobile_number,omitempty"`
	AddresBlock        *string                            `json:"address_block"`
	AddressStreet      *string                            `json:"address_street"`
	ResidentialCountry *string                            `json:"recidential_country"`
	Nationality        *string                            `json:"nationality"`
	Experiences        []experiences.ExperienceWithSkills `json:"experiences,omitempty"`
	Education          []educations.Education             `json:"educations,omitempty"`
	Projects           []projects.ProjectResponse         `json:"projects,omitempty"`
}

// // Error implements error.
// func (u UserDetails) Error() string {
// 	panic("unimplemented")
// }
