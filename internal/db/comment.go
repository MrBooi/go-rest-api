package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MrBooi/go-rest-api/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(row CommentRow) comment.Comment {
	return comment.Comment{
		ID:   row.ID,
		Slug: row.Slug.String,
		Body: row.Body.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow

	row := d.Client.QueryRowContext(ctx, "SELECT id, slug, body, author FROM comments WHERE id = $1", uuid)

	if err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author); err != nil {
		return comment.Comment{}, fmt.Errorf("could not get comment: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}
