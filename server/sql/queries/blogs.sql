-- name: GetBlogs :many
SELECT id, title, date, main_image, subheading, body, created
FROM Blog
ORDER BY created DESC;
