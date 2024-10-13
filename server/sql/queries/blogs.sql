-- name: GetBlogs :many
select id, title, date, main_image, subheading, body, created
from Blog
order by created desc;

-- name: GetBlogByID :one
select id, title, date, main_image, subheading, body, created
from Blog
where id = ?;

-- name: CreateBlog :exec
insert into Blog (id, title, date, main_image, subheading, body, created)
values (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateBlogNoImage :exec
update Blog
set title      = ?,
    date       = ?,
    subheading = ?,
    body       = ?
where id = ?;

-- name: UpdateBlog :exec
update Blog
set title      = ?,
    date       = ?,
    main_image = ?,
    subheading = ?,
    body       = ?
where id = ?;

-- name: DeleteBlog :exec
delete
from Blog
where id = ?;
