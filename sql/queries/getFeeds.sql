-- name: GetFeeds :many
SELECT feeds.name AS feed_name,
feeds.url AS url,
users.name AS user_name
FROM feeds
LEFT JOIN users
ON feeds.user_id = users.id
ORDER BY users.name;
