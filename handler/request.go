package handler

import (
	"fmt"
	"strings"
)

func errorParamter(messages []string) error {
	if len(messages) > 0 {
		return fmt.Errorf(strings.Join(messages, ", "))
	}
	return nil
}

// CREATE Opening

type CreateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float32 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	errorMessages := []string{}
	if r.Role == "" {
		errorMessages = append(errorMessages, "Role is required")
	}
	if r.Company == "" {
		errorMessages = append(errorMessages, "Company is required")
	}
	if r.Location == "" {
		errorMessages = append(errorMessages, "Location is required")
	}
	if r.Link == "" {
		errorMessages = append(errorMessages, "Link is required")
	}
	if r.Remote == nil {
		errorMessages = append(errorMessages, "Remote is required")
	}
	if r.Salary <= 0 {
		errorMessages = append(errorMessages, "Salary needs to be greater than 0")
	}

	return errorParamter(errorMessages)
}

// UPDATE Opening

type UpdateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float32 `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one field is required")
}
