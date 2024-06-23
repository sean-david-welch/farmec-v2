-- name: GetBlogs :many
SELECT id, title, date, main_image, subheading, body, created
FROM Blog
ORDER BY created DESC;

-- name: GetBlogByID :one
SELECT *
FROM Blog
WHERE id = ?;

-- name: CreateBlog :exec
INSERT into Blog
(id, title, date, main_image, subheading, body, created)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateBlogNoImage :exec
UPDATE Blog
SET title = ?, date = ?, subheading = ?, body = ?
WHERE id = ?;

-- name: UpdateBlog :exec
UPDATE Blog
SET title = ?, date = ?, main_image = ?, subheading = ?, body = ?
WHERE id = ?;

-- name: DeleteBlog :exec
DELETE from Blog
WHERE id = ?