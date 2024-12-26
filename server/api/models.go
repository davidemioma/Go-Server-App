package main

import (
	"go-server-tutorial/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func databaseUserToUser(user database.User) User{
	var role string

	if user.Role.Valid {
		role = user.Role.String
	}

	return User{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		Role: role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}