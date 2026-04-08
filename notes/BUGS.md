# Bugs & Pending Work

## 1. Form Validation & Error Handling — Warranty Claim & Machine Registration

Public-facing support forms currently have no client-side validation or server-side error feedback.

**What's needed:**
- Server-side: add `clean()` / `clean_<field>()` methods to `WarrantyclaimForm` and `MachineregistrationForm`, return field-level errors
- Template: render `{{ form.field.errors }}` (or equivalent) next to each input so the user knows what to fix
- Required field indicators and user-friendly error messages
- Consider basic JS validation (HTML5 `required`, `pattern`) as a first line of defence

---

## 2. SEO — Identify Target Search Queries

Need to define the keyword set the site should rank for before any on-page or meta work.

**What's needed:**
- Compile a list of target queries (e.g. "farm machinery Ireland", supplier/machine-specific terms, local dealer queries)
- Map queries to specific pages (supplier detail, machine detail, blog posts, etc.)
- See `SEO.md` for any existing notes — expand it once the query list is agreed

---

## 3. EC2 — AWS CLI Not Configured for S3 Backup

Nightly backup cron (`0 2 * * *`) copies the SQLite DB to `s3://farmec-backups/` but has been silently failing since at least 2026-04-05 — AWS credentials are not configured on the server.

**What's needed:**
- Preferred: attach an IAM instance profile to the EC2 instance scoped to `s3:PutObject` on `s3://farmec-backups/*` — no credentials in `.env` required
- Quick fix alternative: add `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` to `.env` on the server
- Either way: add `AWS_DEFAULT_REGION=eu-west-1` to `.env` on the server (currently missing)

---

## 4. Inline JS — Move to Dedicated JS Files

Inline `<script>` blocks scattered across templates should be extracted into dedicated static JS files.

**What's needed:**
- Audit templates for inline `<script>` blocks
- Move logic into appropriately named files under `theme/static/js/`
- Reference them via `{% static %}` tags
- Ensures CSP compatibility, cacheability, and easier maintenance

---

## 5. Google Maps — Move to Backend Rendering

Google Maps is currently rendered client-side (API key exposed in templates/JS).

**What's needed:**
- Move map generation to the backend — use the Maps Static API or embed URL server-side
- Remove the JS Maps SDK from the frontend entirely
- Store the API key only in `.env` / server environment, never in templates or static files

---

## 6. Carousel — Move to HTMX Rendering

The homepage carousel is currently driven by client-side JS.

**What's needed:**
- Replace JS carousel logic with an HTMX-driven approach (e.g. `hx-get` polling or swap on interaction)
- Render slide markup server-side from `Carousel` model data
- Remove or minimise the JS dependency for slide transitions

---

## 7. Warranty Parts Required — Move to HTMX Rendering

The parts-required section of the warranty claim flow is rendered/updated client-side.

**What's needed:**
- Replace JS-driven parts list with HTMX requests against a Django view
- Use `HTMXViewMixin` / `django-template-partials` for partial responses
- Ensure add/remove part interactions update the DOM via HTMX swaps, not manual JS DOM manipulation

---
