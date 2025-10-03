package skills

import (
	"api/ent"
	"context"
)

func GenSkills(c *ent.Client) ([]Skill, error) {
	ctx := context.Background()
	allSkills, err := c.Skill.Query().All(ctx)

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

func GenSkill(skill_id uint, c *ent.Client) (Skill, error) {
	ctx := context.Background()
	skill, err := c.Skill.Get(ctx, skill_id)

	if err != nil {
		return Skill{}, err
	}

	return Skill{
		Name: skill.Name,
	}, nil
}

func GenCreateSkill(skill Skill, c *ent.Client) (Skill, error) {
	tx, err := c.Tx(context.Background())
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
