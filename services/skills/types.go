package skills

type Skill struct {
	Name string `json:"name"`
}

type SkillInfo struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type TeckStack struct {
	TeckStackName string      `json:"techstack_name"`
	SkillSet      []SkillInfo `json:"skills_set"`
}

type TechStackCreateRequest struct {
	TeckStackName string `json:"techstack_name"`
	SkillID       uint   `json:"skill_id"`
	UserID        uint   `json:"user_id"`
}

// Responses

type TechStackCreateResponse struct {
	ID            uint   `json:"id"`             // the newly created tech stack ID
	TechStackName string `json:"techstack_name"` // name
	SkillID       uint   `json:"skill_id"`       // skill reference
	UserID        uint   `json:"user_id"`        // user reference
}
