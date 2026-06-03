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

## 8. Unit Tests — Comprehensive Coverage for All Views and Forms

No systematic test coverage exists for Django views or forms across the project.

**What's needed:**
- Per-app `tests.py` with a `ViewsTests` class per module covering all major behaviours: status codes, template used, context data, redirects, permission checks, HTMX responses, and form submission (valid and invalid) where applicable
- Use `model_bakery.baker` for fixture generation — no hand-rolled `setUp` factories
- Use Django's `TestCase` as the base class; `pytest-django` markers where needed
- Priority order: `support/` (warranty, machine registration, parts required), then `catalog/`, `content/`, `team/`, `legal/`

---

## 10. EC2 — Add Swap Space

Server has 924MB RAM and no swap. Memory exhaustion causes site slowdowns and kills Docker/SSH responsiveness.

**What's needed:**
- Add 2GB swap file (persistent across reboots):
```bash
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

---

## 11. EC2 — Incident 2026-06-03: Site Slow / Memory Exhaustion

**What happened:**
- `unattended-upgrade` spawned 4 stuck Python processes consuming ~40% RAM
- Gunicorn workers (3) had been running since 2026-05-18 with gradual memory leak (~32% RAM)
- Combined left only 40MB available from 924MB total — no swap to fall back on
- Docker daemon starved, `docker ps` hung, SSH login slow

**How fixed:**
1. Killed stuck `unattended-upgrade` processes (`kill -9`)
2. Restarted Docker container (`docker restart $(docker ps -q)`) — freed Gunicorn worker memory
3. Available memory recovered from 40MB → 250MB, load dropped from ~5 → ~2

**Prevention:** Add swap (see #10) and consider setting `--max-requests` on Gunicorn workers to auto-recycle on memory leak.

---

## 9. Bot Protection — reCAPTCHA or Equivalent on Public Forms

Unauthenticated forms (primarily the contact form) have no bot protection and are vulnerable to spam submissions.

**What's needed:**
- Add reCAPTCHA v3 (invisible, score-based) or hCaptcha to all public-facing forms — at minimum the contact form, and also warranty claim and machine registration
- Validate the token server-side in the view before processing the form; reject submissions below the score threshold
- Store the site key in settings (from `.env`) and pass to templates via context — secret key never leaves the server
- Consider a lightweight honeypot field as a fallback if a third-party captcha service is undesirable

---
