# Changelog

## 2026-05-12 — SEO & Product Pages

### Models
- Added `meta_title` (70 char) and `meta_description` (160 char) to `Supplier`, `Machine` and `Product` models
- Migrations: `0005` (Supplier + Machine), `0006` (Product)

### SEO Content
- AI-generated SEO copy applied to all 13 suppliers, 17 machines and 42 products via Django shell
- All copy targets `[Brand] + [product type] + Ireland | Farmec` pattern

### Templates
- `supplier_meta.html` — OG tags + Schema.org `Organization` + `BreadcrumbList` JSON-LD
- `machine_meta.html` — OG tags + Schema.org `Product` + `BreadcrumbList` JSON-LD
- `product_meta.html` — OG tags + Schema.org `Product` + 5-level `BreadcrumbList` JSON-LD
- All three included via `{% include %}` in their respective detail templates
- Admin SEO fieldsets added to `SupplierAdmin`, `MachineAdmin`, `ProductAdmin`

### Product Detail Pages
- New `ProductDetailView` — `/products/<slug>/`
- `product_detail.html` — vertical card layout, image links to supplier site, Supplier Website button
- Product cards on machine pages updated — image links to supplier site, "View Details" button links to product detail page
- Product image no longer scales on hover (card border still highlights red)

### Breadcrumbs
- `components/breadcrumbs.html` partial — centred, chevron separators, current page in red
- Added to `machine_detail.html`: Suppliers › [Supplier] › [Machine]
- Added to `product_detail.html`: Suppliers › [Supplier] › [Machine] › [Product]
- Breadcrumb context passed from `MachineDetailView` and `ProductDetailView`

### Slug Fixes
- SIP slug changed `sip-slovenia` → `sip`, 301 redirect added for old URL
- 10 product slugs containing `/` fixed (SIP Spider tedders and Star rakes)

### Favicon
- New square favicon (red `#dc2626` background, white F) replacing wide logo PNG
- `favicon.ico`, `favicon-32x32.png`, `favicon.png` (192×192) added to static files
- `favicon.ico` now served at root `/favicon.ico` via URL redirect
- `base.html` updated with full favicon link set including `apple-touch-icon`

### Infrastructure
- `just logs` — SSH into EC2 and tail Docker container errors
- `just logs-live` — follow live Docker container log stream
