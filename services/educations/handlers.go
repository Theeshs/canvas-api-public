package educations

import (
	"api/ent"
	"api/ent/education"
	"api/services/user"
	"api/utils"
	"context"
	"time"
)

type EducationHandler struct {
	client *ent.Client
}

func NewEducationHandler(client *ent.Client) *EducationHandler {
	return &EducationHandler{
		client: client,
	}
}

func (eh *EducationHandler) GenUserEducations(user_id uint) ([]Education, error) {
	edu, err := eh.client.Education.Query().Where(education.UserID(user_id)).All(context.Background())

	if err != nil {
		return nil, err
	}

	result := make([]Education, len(edu))

	for i, e := range edu {
		result[i] = Education{
			InstitueName:      e.InstituteName,
			StartDate:         e.StartDate.Format("2006-01-02"),
			EndDate:           e.EndDate.Format("2006-01-02"),
			ModeOfStudy:       e.ModeOfStudy,
			DegreeType:        e.DegreeType,
			AreaOfStudy:       e.AreaOfStudy,
			CurrentlyStudying: e.CurrentlyStudying,
			Description:       e.Description,
			UserID:            e.UserID,
		}
	}

	return result, nil
}

func (eh *EducationHandler) GenUserEducation(education_id uint, client *ent.Client) (Education, error) {
	edu, err := eh.client.Education.Query().Where(education.ID(education_id)).Only(context.Background())

	if err != nil {
		return Education{}, err
	}

	return Education{
		InstitueName:      edu.InstituteName,
		StartDate:         edu.StartDate.Format("2006-01-02"),
		EndDate:           edu.EndDate.Format("2006-01-02"),
		ModeOfStudy:       edu.ModeOfStudy,
		DegreeType:        edu.DegreeType,
		AreaOfStudy:       edu.AreaOfStudy,
		CurrentlyStudying: edu.CurrentlyStudying,
		Description:       edu.Description,
		UserID:            edu.UserID,
	}, nil
}

func (eh *EducationHandler) GenCreateUserEducation(edu Education, client *ent.Client) (Education, error) {
	tx, err := eh.client.Tx(context.Background())
	if err != nil {
		return Education{}, err
	}
	userAvailable, err := user.NewUserHandler(eh.client).FetchUserByID(uint(edu.UserID))
	startDate, _ := utils.ConvertJsonDate(edu.StartDate)
	endDate, _ := utils.ConvertJsonDate(edu.EndDate)

	newEducation, err := tx.Education.Create().
		SetInstituteName(edu.InstitueName).
		SetStartDate(startDate).SetEndDate(endDate).
		SetCurrentlyStudying(edu.CurrentlyStudying).
		SetUser(userAvailable).
		SetDegreeType(edu.DegreeType).
		SetDescription(edu.Description).
		SetCreatedAt(time.Now()).
		SetAreaOfStudy(edu.AreaOfStudy).
		SetModeOfStudy(edu.ModeOfStudy).Save(context.Background())

	if err != nil {
		tx.Rollback()
		return Education{}, err
	}

	if err := tx.Commit(); err != nil {
		return Education{}, err
	}

	return Education{
		InstitueName:      newEducation.InstituteName,
		StartDate:         newEducation.StartDate.Format("2006-01-02"),
		EndDate:           newEducation.EndDate.Format("2006-01-02"),
		ModeOfStudy:       newEducation.ModeOfStudy,
		DegreeType:        newEducation.DegreeType,
		AreaOfStudy:       newEducation.AreaOfStudy,
		CurrentlyStudying: newEducation.CurrentlyStudying,
		Description:       newEducation.Description,
		UserID:            newEducation.UserID,
	}, nil
}
