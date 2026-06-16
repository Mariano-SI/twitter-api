package follow

import (
	"context"
	"fmt"
)

func (r *followRepository) Delete(ctx context.Context, followId string) error {
	query := `DELETE FROM follows WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, followId)
	if err != nil {
		return fmt.Errorf("failed to delete follow: %w", err)
	}

	return nil
}
