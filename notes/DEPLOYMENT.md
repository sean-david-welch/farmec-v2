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
- **Consider migrating to PostgreSQL** if the site grows — SQLite has write-locking limitations under concurrent load.

### Secrets
- **Never use `get_random_secret_key()`** for production without checking for `$` characters — use the safe generation command above instead.
- **Rotate credentials** any time they appear in plaintext (chat, logs, etc.).
