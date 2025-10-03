package experiences

type Experience struct {
	CompanyName  string `json:"company_name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	CurrentPlace bool   `json:"current_place"`
	Position     string `json:"position"`
	UserID       uint   `json:"user_id"`
	Description  string `json:"description"`
}

type ExperienceWithSkills struct {
	CompanyName  string   `json:"company_name"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	CurrentPlace bool     `json:"current_place"`
	Position     string   `json:"position"`
	UserID       uint     `json:"user_id"`
	Description  string   `json:"description"`
	Skills       []string `json:"skills"`
}

type SkillAssociation struct {
	ExperienceID uint `json:"experience_id"`
	SkillID      uint `json:"skill_id"`
}
