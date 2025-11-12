package projects

import (
	"api/ent"
	"api/ent/project"
	"api/services/user"
	"context"
	"fmt"
)

type ProjectsHandler struct {
	client *ent.Client
}

func NewProjectHandler(client *ent.Client) *ProjectsHandler {
	return &ProjectsHandler{
		client: client,
	}
}

func (ph *ProjectsHandler) GenCreateProject(project ProjectCreateRequest) (ProjectResponse, error) {
	tx, err := ph.client.Tx(context.Background())
	ctx := context.Background()
	if err != nil {
		return ProjectResponse{}, err
	}
	user, err := user.NewUserHandler(ph.client).FetchUserByID(project.UserID)
	if err != nil {
		return ProjectResponse{}, err
	}

	newProject, err := tx.Project.Create().
		SetUser(user).
		SetProjectName(project.ProjectName).
		SetDescription(project.Description).
		SetURL(project.URL).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return ProjectResponse{}, err
	}

	if err := tx.Commit(); err != nil {
		return ProjectResponse{}, err
	}

	return ProjectResponse{
		ProjectName: newProject.ProjectName,
		Description: newProject.Description,
		UserID:      newProject.UserID,
		ID:          newProject.ID,
	}, err
}

func (ph *ProjectsHandler) GenProjectList(user_id uint) ([]ProjectResponse, error) {
	ctx := context.Background()
	allProjects, err := ph.client.Project.Query().Where(project.UserID(user_id)).All(ctx)

	if err != nil {
		return []ProjectResponse{}, err
	}

	result := make([]ProjectResponse, len(allProjects))

	for i, e := range allProjects {
		result[i] = ProjectResponse{
			ProjectName: e.ProjectName,
			Description: e.Description,
			URL:         e.URL,
			ID:          e.ID,
			UserID:      e.UserID,
			SkillSet:    []string{},
		}
	}
	return result, nil
}

func (ph *ProjectsHandler) GenProjectByID(userID, projectID uint) (ProjectResponse, error) {
	project, err := ph.client.Project.Query().
		Where(project.UserID(userID), project.ID(projectID)).
		Only(context.Background())
	if err != nil {
		return ProjectResponse{}, fmt.Errorf("failed to fetch project: %w", err)
	}

	return ProjectResponse{
		ProjectName: project.ProjectName,
		Description: project.Description,
		URL:         project.URL,
		ID:          project.ID,
		UserID:      project.UserID,
		SkillSet:    []string{},
	}, nil
}

func (ph *ProjectsHandler) GenDeleteProjectByID(userID, projectID uint) error {
	count, err := ph.client.Project.Delete().
		Where(project.ID(projectID), project.UserID(userID)).
		Exec(context.Background())
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no project found to delete")
	}
	return nil
}
