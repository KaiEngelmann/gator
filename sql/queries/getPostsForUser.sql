-- name: GetPostsForUser :many
SELECT posts.*,
users.name AS user_name,
feeds.name AS feed_name
FROM posts
JOIN feeds ON posts.feed_id = feeds.id
JOIN users ON feeds.user_id = users.id
WHERE users.name = $1
ORDER BY posts.updated_at DESC
LIMIT $2;
