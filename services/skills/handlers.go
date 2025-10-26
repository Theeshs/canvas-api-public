package skills

import (
	"api/ent"
	"api/ent/techsctack"
	"api/services/user"
	"api/utils"
	"context"
)

type SkillHandler struct {
	client               *ent.Client
	techStackUtilHandler utils.TechStackUtils
}

func NewSkillHandler(client *ent.Client) *SkillHandler {
	return &SkillHandler{
		client:               client,
		techStackUtilHandler: utils.TechStackName(""),
	}
}

func (h *SkillHandler) GenSkills() ([]Skill, error) {
	ctx := context.Background()
	allSkills, err := h.client.Skill.Query().All(ctx)

	if err != nil {
		return []Skill{}, err
	}

	result := make([]Skill, len(allSkills))

	for i, e := range allSkills {
		result[i] = Skill{
			Name: e.Name,
		}
	}
	return result, nil
}

func (h *SkillHandler) GenSkill(skill_id uint) (Skill, error) {
	ctx := context.Background()
	skill, err := h.client.Skill.Get(ctx, skill_id)

	if err != nil {
		return Skill{}, err
	}

	return Skill{
		Name: skill.Name,
	}, nil
}

func (h *SkillHandler) GenCreateSkill(skill Skill) (Skill, error) {
	tx, err := h.client.Tx(context.Background())
	ctx := context.Background()
	if err != nil {
		return Skill{}, err
	}
	newSkill, err := tx.Skill.
		Create().
		SetName(skill.Name).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return Skill{}, err
	}
	if err := tx.Commit(); err != nil {
		return Skill{}, err
	}
	return Skill{Name: newSkill.Name}, err
}

func (h *SkillHandler) GenTechStack(user_id uint) ([]TeckStack, error) {
	ctx := context.Background()
	allTeckStacks, err := h.client.TechSctack.
		Query().
		Where(techsctack.UserID(user_id)).
		WithSkill().
		All(ctx)
	if err != nil {
		return []TeckStack{}, err
	}

	// Group by TechStack category
	grouped := make(map[string][]SkillInfo)
	for _, t := range allTeckStacks {
		grouped[t.Name] = append(grouped[t.Name], SkillInfo{
			Name: t.Edges.Skill.Name,
			Icon: t.Edges.Skill.Icon, // assuming Skill has Icon field
		})
	}

	var result []TeckStack
	for category, skills := range grouped {
		result = append(result, TeckStack{
			TeckStackName: category,
			SkillSet:      skills,
		})
	}

	return result, nil
}

func (h *SkillHandler) GenCreateTechStack(stack TechStackCreateRequest) (TechStackCreateResponse, error) {
	ctx := context.Background()
	name := h.techStackUtilHandler.GetTechStackName(stack.TeckStackName)
	user, err := user.NewUserHandler(h.client).FetchUserByID(stack.UserID)
	if err != nil {
		return TechStackCreateResponse{}, err
	}
	skill, err := h.client.Skill.Get(ctx, stack.SkillID)
	if err != nil {
		return TechStackCreateResponse{}, err
	}

	tx, err := h.client.Tx(context.Background())
	if err != nil {
		return TechStackCreateResponse{}, err
	}
	newTechStack, err := tx.TechSctack.
		Create().
		SetName(name).
		SetUser(user).
		SetSkill(skill).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return TechStackCreateResponse{}, err
	}
	if err := tx.Commit(); err != nil {
		return TechStackCreateResponse{}, err
	}
	return TechStackCreateResponse{
		TechStackName: newTechStack.Name,
		SkillID:       newTechStack.SkillID,
		UserID:        newTechStack.UserID,
		ID:            newTechStack.ID,
	}, err

}
