package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/kaiengelmann/gator/internal/database"
)

func browse(s *AppState, cmd UserCommand, user database.User) error {
	var limit int32 = 2
	if len(cmd.args) == 1 {
		n, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid number: %w", err)
		}
		limit = int32(n)
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: browse <number>")
	}
	getPostsParms := database.GetPostsForUserParams{
		Name:  user.Name,
		Limit: limit,
	}
	return printPosts(s, getPostsParms)
}

func printPosts(s *AppState, p database.GetPostsForUserParams) error {
	posts, err := s.db.GetPostsForUser(s.ctx, p)
	if err != nil {
		return fmt.Errorf("get posts: %w", err)
	}
	if len(posts) == 0 {
		fmt.Println("No posts found. Try running: agg")
	}
	for _, post := range posts {
		fmt.Printf("- %s\n %s\n %s\n published: %s\n\n",
			post.Title,
			nullOr(post.Description),
			post.Url,
			timeOr(post.PublishedAt),
		)
	}
	return nil
}

func timeOr(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format(time.RFC3339)
	}
	return "(unknown)"
}
func nullOr(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return "(none)"
}
