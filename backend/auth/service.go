package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"

	"github.com/diya-suryawanshi/cloud/rbac"
	"golang.org/x/crypto/bcrypt"
)

type userRecord struct {
	EmployeeID   string
	Name         string
	Role         rbac.Role
	PasswordHash []byte
}

type Service struct {
	mu       sync.RWMutex
	users    map[string]userRecord
	sessions map[string]Session
}

func NewService() *Service {
	return &Service{
		users: map[string]userRecord{
			"AT001": {
				EmployeeID:   "AT001",
				Name:         "Admin Operator",
				Role:         rbac.Admin,
				PasswordHash: mustHash("admin123"),
			},
			"AT002": {
				EmployeeID:   "AT002",
				Name:         "DevOps Operator",
				Role:         rbac.DevOps,
				PasswordHash: mustHash("devops123"),
			},
			"AT003": {
				EmployeeID:   "AT003",
				Name:         "Security Operator",
				Role:         rbac.Security,
				PasswordHash: mustHash("security123"),
			},
		},
		sessions: make(map[string]Session),
	}
}

func mustHash(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hash
}

func (s *Service) Login(employeeID, password string) (Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[employeeID]
	if !ok {
		return Session{}, errors.New("invalid employee id or password")
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		return Session{}, errors.New("invalid employee id or password")
	}

	token, err := generateToken()
	if err != nil {
		return Session{}, err
	}

	session := Session{
		Token:          token,
		EmployeeID:     user.EmployeeID,
		Name:           user.Name,
		Role:           user.Role,
		Features:       featureMap(user.Role),
		AllowedActions: rbac.AllowedActionCategories(user.Role),
	}

	s.sessions[token] = session
	return session, nil
}

func (s *Service) Resolve(token string) (Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, ok := s.sessions[token]
	return session, ok
}

func featureMap(role rbac.Role) map[string]string {
	features := []rbac.Feature{
		rbac.FeatureRunGovernance,
		rbac.FeatureCloudGraph,
		rbac.FeatureSimulationStudio,
		rbac.FeatureGovernanceAction,
		rbac.FeatureGitOpsView,
		rbac.FeatureGitOpsMerge,
		rbac.FeatureExplainability,
		rbac.FeatureBillShock,
		rbac.FeatureFeedbackLoop,
		rbac.FeatureNotifications,
	}

	out := make(map[string]string, len(features))
	for _, f := range features {
		out[string(f)] = string(rbac.FeatureAccess(role, f))
	}
	return out
}

func generateToken() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
