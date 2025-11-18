package main

import (
	"fmt"
	"time"
)

func agg(s *AppState, cmd UserCommand) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: agg <duration>")
	}

	delay, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration %q: %w", cmd.args[0], err)
	}
	minDelay := 60 * time.Second
	if delay < minDelay {
		return fmt.Errorf("minimum allowed delay is %s", minDelay)
	}
	fmt.Printf("Collecting feeds every %s\n", delay)

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	if err := scrapeFeeds(s); err != nil {
		fmt.Println("Error scraping feeds:", err)
	}

	for {
		select {
		case <-s.ctx.Done():
			fmt.Println("Aggregator shutting down")
			return nil
		case <-ticker.C:
			if err := scrapeFeeds(s); err != nil {
				fmt.Println("Error scraping feeds:", err)
			}
		}
	}
}
