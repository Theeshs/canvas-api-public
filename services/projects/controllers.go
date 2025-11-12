package projects

import (
	"api/ent"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	handler *ProjectsHandler
}

func NewProjectConteroller(client *ent.Client) *ProjectController {
	return &ProjectController{
		handler: NewProjectHandler(client),
	}
}

func (pc *ProjectController) GetProjects(c *fiber.Ctx) error {
	projects, err := pc.handler.GenProjectList(1)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to list projects",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Projects retrived successfully",
		"skill":   projects,
	})
}

func (pc *ProjectController) CreateProject(c *fiber.Ctx) error {
	project := new(ProjectCreateRequest)
	if err := c.BodyParser(project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	createdProject, err := pc.handler.GenCreateProject(*project)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to create project",
			"stack": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project created successfully",
		"project": createdProject,
	})
}

func (pc *ProjectController) GetProjectByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Convert string to uint
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID",
		})
	}
	projectID := uint(idUint64)
	project, err := pc.handler.GenProjectByID(1, projectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to get project",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project retrived successfully",
		"skill":   project,
	})
}

func (pc *ProjectController) DeleteProjectByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Convert string to uint
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID",
		})
	}
	projectID := uint(idUint64)

	err = pc.handler.GenDeleteProjectByID(1, projectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to delete project",
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
