package user

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *userRepository) Create(ctx context.Context, user *model.UserModel) error {
	query := `INSERT INTO users (id, email, username, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
