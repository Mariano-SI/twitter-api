package follow

import (
	"context"
	"fmt"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *followRepository) Create(ctx context.Context, follow model.FollowModel) error {
	query := `INSERT INTO follows (id, follower_id, followed_id, created_at) VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, follow.ID, follow.FollowerID, follow.FollowedID, follow.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create follow: %w", err)
	}

	return nil
}
