# Migration Tasks

## 1. Security Hardening
- Run `manage.py check --deploy` and resolve all warnings
- `DEBUG=False`, `SECRET_KEY` from env, `ALLOWED_HOSTS` locked down
- Enable `SECURE_SSL_REDIRECT`, `SECURE_HSTS_SECONDS`, `SESSION_COOKIE_SECURE`, `CSRF_COOKIE_SECURE`
- Rate limiting on unauthenticated POST endpoints (contact form, warranty claims, machine registration)
- Honeypot field on public forms
- Error logging to file or stdout for production

## 2. Test Suite
- `pytest-django` + `model_bakery`
- Cover public views, form submissions, email sending, S3 uploads, admin access
- One test file per app in `tests/`, shared fixtures in `conftest.py`

## 3. Deployment — EC2
- `Dockerfile` + `docker-compose.yml` (`web` gunicorn, `nginx` reverse proxy)
- SSL via certbot + Let's Encrypt
- `collectstatic` → S3 or nginx volume
- `.env` on EC2 only
- Switch nginx from Go binary to gunicorn

## 4. Go/React Cleanup (post-deployment)
- Remove `/server` and `/client` source trees
- Stop and remove Go binary, Air config, Node/npm tooling
- Remove Go runtime if unused
