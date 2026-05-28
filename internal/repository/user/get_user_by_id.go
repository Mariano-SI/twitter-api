package user

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *userRepository) GetUserById(ctx context.Context, userId string) (*model.UserModel, error) {
	query := "SELECT id, email, username, password, created_at, updated_at FROM users WHERE id = ?"

	row := r.db.QueryRowContext(ctx, query, userId)

	var result model.UserModel

	err := row.Scan(&result.ID, &result.Username, &result.Email, &result.Password, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
