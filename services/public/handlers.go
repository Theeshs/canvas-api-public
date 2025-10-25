package public

import (
	"api/ent"
	"api/ent/user"
	"api/services/educations"
	"api/services/experiences"
	"api/services/projects"
	"context"
	"fmt"
	"sync"
)

func GetPortfolioData(client *ent.Client) (UserDetails, error) {
	email := "theekshana.sandaru@gmail.com"

	user := client.User.Query().Where(user.Email(email)).OnlyX(context.Background())
	mn := int(user.MobileNumber)
	dobStr := user.Dob.Format("2006-01-02")

	var (
		exp              []experiences.ExperienceWithSkills
		edu              []educations.Education
		pro              []projects.ProjectResponse
		err1, err2, err3 error
		wg               sync.WaitGroup
	)

	wg.Add(3)

	go func() {
		defer wg.Done()
		exp, err1 = experiences.GenUserExperiencesWithSkills(user.ID, client)
	}()

	go func() {
		defer wg.Done()
		edu, err2 = educations.GenUserEducations(user.ID, client)
	}()

	go func() {
		defer wg.Done()
		pro, err3 = projects.NewProjectHandler(client).GenProjectList(user.ID)
	}()

	wg.Wait()

	if err1 != nil || err2 != nil {
		combinedErr := fmt.Errorf("experience error: %v, education error: %v, project error: %v", err1, err2, err3)
		return UserDetails{}, combinedErr
	}

	response := UserDetails{
		Email:              email,
		UserName:           user.Username,
		FirstName:          &user.FirstName,
		LastName:           &user.LastName,
		DOB:                &dobStr,
		GithubUsername:     &user.GithubUsername,
		Description:        &user.Description,
		MobileNumber:       &mn,
		AddresBlock:        &user.AddressBlock,
		AddressStreet:      &user.AddressStreet,
		ResidentialCountry: &user.RecidentialCountry,
		Nationality:        &user.Nationality,
		Experiences:        exp,
		Education:          edu,
		Projects:           pro,
	}

	return response, nil
}
