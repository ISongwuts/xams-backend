package models

import "github.com/golang-jwt/jwt/v4"

type (
	UserResponse struct {
		UserID    string `json:"user_id"`
		Email     string `json:"email"`
		Prename   string `json:"prename"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		BranchID  uint   `json:"branch_id"`
		Role      Role   `json:"role"`
		Token     string `json:"token"`
	}

	UserClaims struct {
		UserID    string 
		Email     string
		jwt.RegisteredClaims
	}

)
