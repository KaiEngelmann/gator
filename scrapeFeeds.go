package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kaiengelmann/gator/internal/database"
)

func scrapeFeeds(s *AppState) error {
	ctx := s.ctx
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("error getting next feed: %w", err)
	}
	now := time.Now()
	timeParams := database.MarkFeedFetchedTimeParams{
		LastFetchedAt: sql.NullTime{Time: now, Valid: true},
		UpdatedAt:     now,
		ID:            feed.ID,
	}
	err = s.db.MarkFeedFetchedTime(ctx, timeParams)
	if err != nil {
		return fmt.Errorf("MarkFeedFetchedTime(%s): %w", feed.Url, err)
	}
	url := feed.Url
	resp, err := fetchFeed(s.ctx, url)
	if err != nil {
		return fmt.Errorf("fetchFeed(%s): %w", feed.Url, err)
	}
	for _, item := range resp.Channel.Item {
		pub := parsePubDate(item.PubDate)
		itemDescription := toNullString(item.Description)
		postParms := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         url,
			Description: itemDescription,
			PublishedAt: pub,
			FeedID:      feed.ID,
		}
		_, err = s.db.CreatePost(ctx, postParms)
		if err != nil {
			return fmt.Errorf("error creating post: %w", err)
		}

	}

	return nil
}
func parsePubDate(s string) sql.NullTime {
	layouts := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC822Z,
		"Mon, 02 Jan 2006 15:04:05 MST",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return sql.NullTime{Time: t, Valid: true}
		}
	}
	return sql.NullTime{Time: time.Time{}, Valid: false}
}
func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}
