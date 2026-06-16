package user

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *userRepository) Update(ctx context.Context, user *model.UserModel) error {
	query := `
		UPDATE users
		SET description = ?, profile_picture_url = ?, profile_picture_key = ?, updated_at = ?
		WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		user.Description, user.ProfilePictureUrl, user.ProfilePictureKey, user.UpdatedAt, user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
