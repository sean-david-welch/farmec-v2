# Migration Tasks

## 1. Template & Code Cleanup
- Remove Go/React source code (`/server`, `/client`) once Django is confirmed live on EC2

## 2. S3 Image Uploads
- Test upload flow end-to-end in admin (Unfold file picker → S3 → public URL)

## 3. Database
- Staying on SQLite

## 4. PDF Downloads for Warranty Claims & Machine Registrations
- Generate a printable PDF for each `Warrantyclaim` and `Machineregistration` record from the admin panel
- Use `weasyprint` or `xhtml2pdf` to render a Django template to PDF (WeasyPrint produces cleaner output)
- Add a custom Unfold admin action and/or a detail-view button that triggers the PDF download
- Create dedicated print templates (`warranty_pdf.html`, `registration_pdf.html`) — clean layout with company branding, all relevant fields, and any attached images
- Ensure the PDF is served as an attachment (`Content-Disposition: attachment`) so it downloads directly

## 5. Unfold Admin Cleanup
- Audit sidebar navigation — ensure all models are grouped logically (catalog, content, support, team, legal)
- Add `list_display`, `search_fields`, `list_filter` to all `ModelAdmin` classes
- Add inline support where useful (e.g. `Lineitems` inline on `Machine`, `Spareparts` inline on `Supplier`)
- Review `readonly_fields` for auto-managed fields (`uid`, `created`, `modified`)
- Add `ordering` and `date_hierarchy` where appropriate

## 5. Production Security Hardening
- Run `manage.py check --deploy` and resolve all warnings before go-live
- Ensure `DEBUG=False`, `SECRET_KEY` sourced from env, `ALLOWED_HOSTS` locked down
- Enable `SECURE_SSL_REDIRECT`, `SECURE_HSTS_SECONDS`, `SESSION_COOKIE_SECURE`, `CSRF_COOKIE_SECURE`
- Add rate limiting to unauthenticated POST endpoints (contact form, warranty claims, machine registration) — use `django-ratelimit` or nginx `limit_req`
- Add a honeypot field to public forms as basic spam protection
- Configure Django file-based error logging for production (at minimum log `ERROR` level to a file or stdout for Docker)

## 6. Deployment — EC2
- **Docker**: Containerise the Django app (`Dockerfile` + `docker-compose.yml` for local parity)
- **Docker Compose** (production): `web` (gunicorn), `nginx` (reverse proxy + static files), optional `redis` if Celery is added later
- **SSL**: Set up certbot + Let's Encrypt via nginx for HTTPS
- **Deployment**: `docker compose pull && docker compose up -d` over SSH; consider a GitHub Actions workflow for one-command deploys on push to `main`
- **Static files**: `collectstatic` → S3 or nginx-served volume
- **Secrets**: `.env` file on EC2 only, not committed to repo
- **Cutover**: Switch nginx to route traffic to gunicorn (Django) instead of the Go binary; verify all routes before decommissioning Go

## 7. EC2 Cleanup (post-deployment)
- Stop and disable the Go server process (systemd unit or whatever manages it)
- Remove Go binary, build artifacts, and Air hot-reload config
- Remove Node/npm and React build tooling if no longer needed
- Remove Go runtime if not used by anything else on the instance
- Clean up any Go-specific environment variables, cron jobs, or log files
- Optionally snapshot the EC2 instance before cleanup as a rollback point

## 8. Test Suite
- Use `pytest-django` + `model_bakery` for all tests
- Cover all public views (status 200, correct template, context keys)
- Cover form submissions (valid + invalid cases) for contact, warranty, machine registration
- Cover email sending (mock Resend client, assert called with correct args)
- Cover S3 upload path (mock `boto3`/`storages`, assert file written)
- Cover admin access (superuser can reach all Unfold admin pages)
- Target: one test file per app in `tests/` directory, shared fixtures in `conftest.py`
