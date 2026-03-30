# Deployment Notes

## Architecture

```
Browser → nginx (SSL termination) → gunicorn (Docker, port 8000) → Django
```

- nginx handles HTTPS, static files, and proxies dynamic requests to the container
- gunicorn runs inside Docker, bound to `127.0.0.1:8000` on the host
- SQLite database and staticfiles are bind-mounted from the host into the container
- GitHub Actions triggers deploys on push to `main`

---

## Issues Encountered & Fixes

### 1. GitHub Actions secrets not passed to deploy script
**Symptom:** `.env` on EC2 was all blank, causing Django 400 errors.
**Cause:** `appleboy/ssh-action` requires an explicit `env:` block to make secrets available as env vars in the script. The `envs:` parameter alone is not enough.
**Fix:** Add `env:` mapping block to the deploy step in `.github/workflows/deploy.yml`.

### 2. `docker compose restart` does not reload env_file
**Symptom:** Updated `.env` on server, restarted container, but env vars inside were still blank.
**Cause:** `docker compose restart` restarts the container process but does not recreate it — env_file is not re-read.
**Fix:** Always use `docker compose up -d` to apply env changes (recreates the container).

### 3. Secret key with `$` signs breaks docker compose
**Symptom:** Warnings like `The "xyz" variable is not set. Defaulting to a blank string.`
**Cause:** Docker compose reads `.env` for variable substitution in the compose file. Values containing `$something` are treated as variable references.
**Fix:** Generate secret keys using only alphanumeric characters and safe symbols (`-_=+`). No `$` signs.

```bash
python -c "import secrets, string; chars = string.ascii_letters + string.digits + '-_=+'; print(''.join(secrets.choice(chars) for _ in range(50)))"
```

### 4. Redirect loop (`ERR_TOO_MANY_REDIRECTS`)
**Symptom:** Browser showed infinite redirect loop after enabling `SECURE_SSL_REDIRECT`.
**Cause:** nginx terminates SSL and proxies to gunicorn as HTTP. Django saw HTTP and redirected to HTTPS — loop.
**Fix:** Add to `settings.py`:
```python
SECURE_PROXY_SSL_HEADER = ('HTTP_X_FORWARDED_PROTO', 'https')
```
nginx already sends `X-Forwarded-Proto: $scheme`, so Django correctly identifies the connection as HTTPS.

### 5. Docker build cache serving stale code
**Symptom:** `docker compose build` completed instantly (all CACHED), but new code wasn't in the container.
**Fix:** Force a full rebuild when cache is suspected stale:
```bash
docker compose build --no-cache && docker compose up -d
```

### 6. Disk full on EC2 (7.6GB root volume)
**Cause:** systemd journal logs accumulated to ~900MB with no rotation limit. Old Go server files (~160MB) left on disk. Docker build cache (~750MB) built up over time.
**Fix:**
```bash
sudo journalctl --vacuum-size=50M
docker system prune -af
```
**Prevention:** Set a journal size limit (already applied):
```
/etc/systemd/journald.conf → SystemMaxUse=100M
```

### 7. `collectstatic` during Docker build wiped by bind mount
**Symptom:** Static files missing at runtime — bind mount overwrote the staticfiles directory built into the image.
**Fix:** Run `collectstatic` in the container CMD at startup, not during `docker build`. Static files are written into the bind-mounted directory at runtime.

---

## Useful Diagnostic Commands

```bash
# Check container status
docker compose ps

# Tail container logs
docker compose logs web --tail=50 -f

# Check env vars inside container
docker compose exec web env | grep -E 'ALLOWED|DEBUG|SECRET'

# Test gunicorn directly (bypassing nginx)
curl -s -o /dev/null -w "%{http_code}" -H "Host: farmec.ie" -H "X-Forwarded-Proto: https" http://localhost:8000

# Check disk usage
df -h
sudo du -sh /home/seanwelch/* /var/lib/docker /var/log

# Free up disk space
sudo journalctl --vacuum-size=50M
docker system prune -af

# Check nginx config
sudo nginx -t
sudo nginx -s reload

# Django shell inside container
docker compose exec web uv run python manage.py shell

# Run migrations manually
docker compose exec web uv run python manage.py migrate
```

---

## Robustness Improvements (TODO)

### Infrastructure
- **Resize EBS volume to 20GB** — 7.6GB is too small for Docker. Do this in AWS Console → EC2 → Volumes → Modify, then `sudo growpart /dev/nvme0n1 1 && sudo resize2fs /dev/root`.
- **Set up log rotation** — journald limit already applied. Also add logrotate for nginx logs if they grow large.
- **Add a swap file** — helps prevent OOM kills on the t-series instance during Docker builds.

### Deployment
- **Add a Docker health check** to `docker-compose.yml`:
```yaml
healthcheck:
  test: ["CMD", "curl", "-f", "http://localhost:8000/health/"]
  interval: 30s
  timeout: 10s
  retries: 3
```
- **Add a `/health/` endpoint** in Django that returns 200 — lets nginx and Docker verify the app is up before routing traffic.
- **Pin the `appleboy/ssh-action` version** — using `@master` means the action can change unexpectedly. Pin to a specific tag like `@v1.0.3`.

### Database
- **Automate database backups** — add a cron job or just command to copy the SQLite file to S3 daily:
```bash
aws s3 cp /home/seanwelch/farmec-v2/database/database.db s3://farmec.ie/backups/database-$(date +%Y%m%d).db
```
- **Consider migrating to PostgreSQL** if the site grows — SQLite has write-locking limitations under concurrent load.

### Secrets
- **Never use `get_random_secret_key()`** for production without checking for `$` characters — use the safe generation command above instead.
- **Rotate credentials** any time they appear in plaintext (chat, logs, etc.).
