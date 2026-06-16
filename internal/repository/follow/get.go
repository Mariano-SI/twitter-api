package follow

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *followRepository) Get(ctx context.Context, followerId, followedId string) (*model.FollowModel, error) {
	query := `SELECT id, follower_id, followed_id, created_at FROM follows WHERE follower_id = ? AND followed_id = ?`

	row := r.db.QueryRowContext(ctx, query, followerId, followedId)

	var result model.FollowModel
	err := row.Scan(&result.ID, &result.FollowerID, &result.FollowedID, &result.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
