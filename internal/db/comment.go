package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MrBooi/go-rest-api/internal/comment"
	uuid "github.com/satori/go.uuid"
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

	row := d.Client.QueryRowContext(ctx,
		`SELECT id, slug, body, author
		 FROM comments WHERE id = $1`,
		uuid,
	)
	if err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author); err != nil {
		return comment.Comment{}, fmt.Errorf("could not get comment: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments 
		(id, slug, author, body) VALUES
		(:id, :slug, :author, :body)`,
		postRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil
}

func (d *Database) UpdateCommentById(ctx context.Context, uuid string, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     uuid,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug,
		author = :author,
		body = :body
		WHERE id = :id`,
		cmtRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) DeleteCommentById(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete comment by id from the database: %w", err)
	}
	return nil
}
