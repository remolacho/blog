package post

import (
	"blog/internal/engineDB"
	"blog/internal/server/logger"
	"context"
	"strconv"
	"time"
)

type PostRepository struct {
	Data *engineDB.Data
}

func (pr *PostRepository) All(ctx context.Context) ([]Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts;
    `

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer logger.Sql(q, []string{""})

	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) Find(ctx context.Context, id uint) (Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts WHERE id = $1;
    `

	defer logger.Sql(q, []string{convertToString(id)})

	row := pr.Data.DB.QueryRowContext(ctx, q, id)

	var p Post
	err := row.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return Post{}, err
	}

	return p, nil
}

func (pr *PostRepository) FindByUser(ctx context.Context, userID uint) ([]Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts
        WHERE user_id = $1;
    `

	defer logger.Sql(q, []string{convertToString(userID)})

	rows, err := pr.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) Create(ctx context.Context, p *Post) error {
	q := `
    INSERT INTO posts (body, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `

	defer logger.Sql(q, []string{p.Body, convertToString(p.UserID)})

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Body, p.UserID, time.Now(), time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) Update(ctx context.Context, id uint, p Post) error {
	q := `
    UPDATE posts set body=$1, updated_at=$2
        WHERE id=$3;
    `

	defer logger.Sql(q, []string{p.Body, convertToString(p.UserID), convertToString(id)})

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Body, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM posts WHERE id=$1;`

	defer logger.Sql(q, []string{convertToString(id)})

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func convertToString(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}
