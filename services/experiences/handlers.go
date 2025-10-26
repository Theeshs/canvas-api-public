package experiences

import (
	"api/ent"
	"api/ent/experience"
	"api/services/user"
	"api/utils"
	"context"
	"fmt"
	"time"
)

type ExperienceHandler struct {
	client *ent.Client
}

func NewExperienceHandler(client *ent.Client) *ExperienceHandler {
	return &ExperienceHandler{
		client: client,
	}
}

func (eh *ExperienceHandler) GenUserExperiences(user_id uint) ([]Experience, error) {
	exp, err := eh.client.Experience.Query().
		Where(experience.UserID(user_id)).
		Order(ent.Desc(experience.FieldStartDate)).
		All(context.Background())

	if err != nil {
		return nil, err
	}
	result := make([]Experience, len(exp))

	for i, e := range exp {
		result[i] = Experience{
			CompanyName:  e.CompanyName,
			StartDate:    e.StartDate.Format("2006-01-02"),
			EndDate:      e.EndDate.Format("2006-01-02"),
			CurrentPlace: e.CurrentPlace,
			Position:     e.Position,
			UserID:       e.UserID,
			Description:  e.Description,
		}
	}

	return result, nil
}

func (eh *ExperienceHandler) GenUserExperience(exp_id uint) (Experience, error) {
	exp, err := eh.client.Experience.Query().Where(experience.ID(exp_id)).Only(context.Background())

	if err != nil {
		return Experience{}, err
	}
	return Experience{
		CompanyName:  exp.CompanyName,
		StartDate:    exp.StartDate.Format("2006-01-02"),
		EndDate:      exp.EndDate.Format("2006-01-02"),
		CurrentPlace: exp.CurrentPlace,
		Position:     exp.Position,
		UserID:       exp.UserID,
		Description:  exp.Description,
	}, nil
}

func (eh *ExperienceHandler) GenCreateExperience(experience Experience) (Experience, error) {
	ctx := context.Background()
	tx, err := eh.client.Tx(ctx)
	if err != nil {
		return Experience{}, err
	}

	userAvailable, err := user.NewUserHandler(eh.client).FetchUserByID(experience.UserID)

	if err != nil {
		return Experience{}, err
	}

	startDate, err1 := utils.ConvertJsonDate(experience.StartDate)
	endDate, err2 := utils.ConvertJsonDate(experience.EndDate)
	fmt.Print(err1)

	if err1 != nil || err2 != nil {
		combinedErr := fmt.Errorf("experience error: %v, education error: %v", err1, err2)
		return Experience{}, combinedErr
	}
	newExperience, err := tx.Experience.Create().
		SetCompanyName(experience.CompanyName).
		SetStartDate(startDate).
		SetEndDate(endDate).
		SetPosition(experience.Position).
		SetCurrentPlace(experience.CurrentPlace).
		SetUser(userAvailable).
		SetDescription(experience.Description).
		SetCreatedAt(time.Now()).Save(ctx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return Experience{}, fmt.Errorf("create err: %v, rollback err: %v", err, rbErr)
		}
		return Experience{}, err
	}

	if err := tx.Commit(); err != nil {
		return Experience{}, err
	}

	start := newExperience.StartDate.Format("2006-01-02")
	end := newExperience.EndDate.Format("2006-01-02")

	return Experience{
		CompanyName:  newExperience.CompanyName,
		StartDate:    start,
		EndDate:      end,
		CurrentPlace: newExperience.CurrentPlace,
		Position:     newExperience.Position,
		UserID:       newExperience.UserID,
		Description:  newExperience.Description,
	}, nil
}

func (eh *ExperienceHandler) GenUserExperiencesWithSkills(user_id uint) ([]ExperienceWithSkills, error) {
	ctx := context.Background()
	exps, err := eh.client.Experience.
		Query().
		WithUserSkillAssociation(func(q *ent.UserSkillAssociationQuery) {
			q.WithSkill()
		}).
		Order(ent.Desc(experience.FieldStartDate)).
		All(ctx)

	if err != nil {
		return nil, err
	}
	result := make([]ExperienceWithSkills, len(exps))

	for i, e := range exps {
		skillNames := make([]string, 0, len(e.Edges.UserSkillAssociation))

		for _, usa := range e.Edges.UserSkillAssociation {
			if usa.Edges.Skill != nil {
				skillNames = append(skillNames, usa.Edges.Skill.Name)
			}
		}
		result[i] = ExperienceWithSkills{
			CompanyName:  e.CompanyName,
			StartDate:    e.StartDate.Format("2006-01-02"),
			EndDate:      e.EndDate.Format("2006-01-02"),
			CurrentPlace: e.CurrentPlace,
			Position:     e.Position,
			UserID:       e.UserID,
			Description:  e.Description,
			Skills:       skillNames,
		}
	}

	return result, nil
}

func (eh *ExperienceHandler) GenAddSkillsToExperience(association SkillAssociation) error {
	ctx := context.Background()

	tx, err := eh.client.Tx(ctx)
	if err != nil {
		return err
	}

	if _, err := eh.client.UserSkillAssociation.Create().
		SetExperienceID(association.ExperienceID).
		SetSkillID(association.SkillID).
		Save(ctx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}
