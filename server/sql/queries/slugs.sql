-- name: GetSupplierBySlug :one
SELECT id, name, logo_image, marketing_image, description,
       social_facebook, social_twitter, social_instagram,
       social_youtube, social_linkedin, social_website,
       created, slug
FROM Supplier
WHERE slug = ? LIMIT 1;

-- name: GetMachineBySlug :one
SELECT id, supplier_id, name, machine_image, description,
       machine_link, created, slug
FROM Machine
WHERE slug = ? LIMIT 1;

-- name: GetProductBySlug :one
SELECT id, machine_id, name, product_image, description,
       product_link, slug
FROM Product
WHERE slug = ? LIMIT 1;

-- name: GetBlogBySlug :one
SELECT id, title, date, main_image, subheading, body,
       created, slug
FROM Blog
WHERE slug = ? LIMIT 1;

-- name: GetSparePartsBySlug :one
SELECT id, supplier_id, name, parts_image,
       spare_parts_link, slug
FROM SpareParts
WHERE slug = ? LIMIT 1;