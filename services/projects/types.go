package projects

type Skill struct {
	Name string `json:"name"`
}

type ProjectResponse struct {
	ID          uint     `json:"id"`
	ProjectName string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	UserID      uint     `json:"user_id"`
	SkillSet    []string `json:"skill_set"`
}

type ProjectCreateRequest struct {
	ProjectName string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	UserID      uint   `json:"user_id"`
}
