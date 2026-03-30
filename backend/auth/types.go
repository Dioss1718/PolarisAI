package auth

import "github.com/diya-suryawanshi/cloud/rbac"

type LoginRequest struct {
	EmployeeID string `json:"employeeId"`
	Password   string `json:"password"`
}

type Session struct {
	Token          string            `json:"token"`
	EmployeeID     string            `json:"employeeId"`
	Name           string            `json:"name"`
	Role           rbac.Role         `json:"role"`
	Features       map[string]string `json:"features"`
	AllowedActions []string          `json:"allowedActions"`
}
