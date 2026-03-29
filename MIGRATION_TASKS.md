# Migration Tasks

## 1. Template & Code Cleanup
- Remove Go/React source code (`/server`, `/client`) once Django is confirmed live on EC2

## 2. S3 Image Uploads ✅
- ImageField with S3 upload via django-storages
- Data migrations to strip existing URL prefixes
- Templates updated to use `.url`

## 3. Database
- Staying on SQLite

## 4. PDF Downloads ✅
- Admin action "Download as PDF" on Warranty Claims and Machine Registrations
- Single record → single PDF, multiple selected → zip file
- WeasyPrint renders branded HTML templates

## 5. Unfold Admin ✅
- Sidebar navigation grouped (catalog, content, support, team, legal, users)
- `list_display`, `search_fields`, `list_filter`, `ordering` on all ModelAdmin classes
- `PartsrequiredInline` on Warranty Claims
- `readonly_fields` for auto-managed fields
- Auth admin re-registered with Unfold forms (no Groups)
- Farmec logo in top left linking to site
- Recent actions removed from dashboard

## 6. Production Security Hardening
- Run `manage.py check --deploy` and resolve all warnings before go-live
- Ensure `DEBUG=False`, `SECRET_KEY` sourced from env, `ALLOWED_HOSTS` locked down
- Enable `SECURE_SSL_REDIRECT`, `SECURE_HSTS_SECONDS`, `SESSION_COOKIE_SECURE`, `CSRF_COOKIE_SECURE`
- Add rate limiting to unauthenticated POST endpoints (contact form, warranty claims, machine registration)
- Add a honeypot field to public forms as basic spam protection
- Configure Django file-based error logging for production

## 7. Deployment — EC2
- **Docker**: Containerise the Django app (`Dockerfile` + `docker-compose.yml`)
- **Docker Compose** (production): `web` (gunicorn), `nginx` (reverse proxy + static files)
- **SSL**: certbot + Let's Encrypt via nginx
- **Static files**: `collectstatic` → S3 or nginx-served volume
- **Secrets**: `.env` file on EC2 only, not committed
- **Cutover**: Switch nginx to route traffic to gunicorn (Django)

## 8. EC2 Cleanup (post-deployment)
- Stop and disable the Go server process
- Remove Go binary, build artifacts, Air config
- Remove Node/npm and React build tooling
- Remove Go runtime if unused

## 9. Test Suite
- Use `pytest-django` + `model_bakery`
- Cover all public views, form submissions, email sending, S3 upload, admin access
- One test file per app in `tests/`, shared fixtures in `conftest.py`
