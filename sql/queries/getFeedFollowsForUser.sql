-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*,
users.name AS user_name,
feeds.name AS feed_name
FROM feed_follows
JOIN users ON feed_follows.user_id = users.id
JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE users.name = $1
ORDER BY feeds.updated_at DESC;
