# Migration Tasks

## 1. Template & Code Cleanup
- Fix template name mismatches in `support/` views (`warrantyclaim_*`, `machineregistration_*`, `partsrequired_*`)
- Audit and remove unused templates, views, and URLs
- Remove unused JS and CSS (check for dead code in `components.js` and any unreferenced CSS)
- Remove or archive legacy Go/React code once Django is confirmed live

## 2. S3 Image Uploads
- Convert `URLField` image fields to `ImageField` across all models (`Supplier`, `Machine`, `Product`, `Spareparts`, `Blog`, `Employee`, `Carousel`, etc.)
- Confirm `django-storages` S3 backend is correctly configured in `settings.py`
- Write and run a data migration to populate new `ImageField` values from existing URL strings
- Test upload flow end-to-end in admin (Unfold file picker → S3 → public URL)

## 3. Database
- Stay on SQLite unless a specific need arises (concurrent writes, full-text search, production scale)
- If migrating: use `pgloader` or `python manage.py dumpdata` → `loaddata` into Postgres
- Postgres benefits worth reconsidering: better concurrent writes, `ArrayField`, native full-text search, easier managed backups on RDS

## 4. Unfold Admin Cleanup
- Audit sidebar navigation — ensure all models are grouped logically (catalog, content, support, team, legal)
- Add `list_display`, `search_fields`, `list_filter` to all `ModelAdmin` classes
- Add inline support where useful (e.g. `Lineitems` inline on `Machine`, `Spareparts` inline on `Supplier`)
- Review `readonly_fields` for auto-managed fields (`uid`, `created`, `modified`)
- Add `ordering` and `date_hierarchy` where appropriate

## 5. Deployment — EC2
- **Docker**: Containerise the Django app (`Dockerfile` + `docker-compose.yml` for local parity)
- **Docker Compose** (production): `web` (gunicorn), `nginx` (reverse proxy + static files), optional `redis` if Celery is added later
- **Deployment**: Use `docker compose pull && docker compose up -d` for zero-downtime-ish deploys via SSH
- **Static files**: `collectstatic` → S3 or nginx-served volume
- **Process manager**: Let Docker restart policy handle process supervision (no need for systemd/supervisor)
- **Secrets**: Pass via `.env` file on EC2, not committed to repo
- **CI/CD** (optional): GitHub Actions → SSH deploy on push to `main`

## 6. Test Suite
- Use `pytest-django` + `model_bakery` for all tests
- Cover all public views (status 200, correct template, context keys)
- Cover form submissions (valid + invalid cases) for contact, warranty, machine registration
- Cover email sending (mock Resend client, assert called with correct args)
- Cover S3 upload path (mock `boto3`/`storages`, assert file written)
- Cover admin access (superuser can reach all Unfold admin pages)
- Target: one test file per app in `tests/` directory, shared fixtures in `conftest.py`
