package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mariano-SI/twitter-api/internal/model"
)

func (r *userRepository) GetProfileById(ctx context.Context, userId string) (*model.UserProfileModel, error) {
	query := `
		SELECT u.id, u.username, u.description, u.profile_picture_url,
		       COUNT(DISTINCT f_in.id)  AS followers_count,
		       COUNT(DISTINCT f_out.id) AS following_count
		FROM users u
		LEFT JOIN follows f_in  ON f_in.followed_id  = u.id
		LEFT JOIN follows f_out ON f_out.follower_id = u.id
		WHERE u.id = ?
		GROUP BY u.id, u.username, u.description, u.profile_picture_url`

	row := r.db.QueryRowContext(ctx, query, userId)

	var result model.UserProfileModel
	err := row.Scan(
		&result.ID, &result.Username, &result.Description, &result.ProfilePictureUrl,
		&result.FollowersCount, &result.FollowingCount,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
