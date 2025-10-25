package utils

import (
	"fmt"
	"strings"
)

type TechStackUtils interface {
	GetTechStackName(name string) string
}

type TechStackName string

const (
	BACKEND  TechStackName = "Backend Development"
	FRONTEND TechStackName = "Forntend Development"
)

func (t TechStackName) GetTechStackName(name string) string {
	fmt.Print(strings.ToLower(name))
	switch strings.ToLower(name) {
	case "backend":
		return string(BACKEND)
	case "frontend":
		return string(FRONTEND)
	default:
		return "Unknown Tech Stack"
	}
}
