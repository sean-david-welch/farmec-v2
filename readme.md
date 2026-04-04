# Farmec Web Application v2

Web platform for Farmec Ireland, an agricultural equipment supplier. Built with Django.

## Technology Stack

- **Backend:** Django 6 (Python 3.13)
- **Server:** Gunicorn behind Nginx
- **Database:** SQLite
- **File Storage:** AWS S3 (`static.farmec.ie`, `eu-west-1`)
- **Email:** Resend
- **Deployment:** Docker + GitHub Actions on AWS EC2
- **Admin:** Unfold (custom Django admin)

## Django Apps

| App | Purpose |
|-----|---------|
| `catalog` | Suppliers, Machines, Products, Spare Parts, Videos |
| `content` | Blogs, Carousels, Exhibitions, Timelines |
| `support` | Warranty Claims, Parts Required, Machine Registrations |
| `team` | Employee records |
| `legal` | Privacy and Terms documents |
| `theme` | Shared templates, static files, context processors |

## Local Development

```bash
# Install dependencies
uv sync

# Run migrations
uv run python manage.py migrate

# Start dev server
uv run python manage.py runserver

# Run tests
pytest
```

## Deployment

Pushes to `main` trigger an automatic deploy via GitHub Actions:

1. SSH to EC2
2. Write `.env` from GitHub secrets
3. `git pull origin main`
4. `docker compose build`
5. `docker compose run --rm web python manage.py migrate`
6. `docker compose up -d`
7. `nginx -s reload`

See `DEPLOYMENT.md` for infrastructure notes, common issues, and diagnostic commands.

## Common Commands

```bash
just ssh        # SSH into EC2
just db-push    # Copy local database to EC2
just db-pull    # Copy EC2 database to local
just provision  # Run Ansible provisioning playbook
```

## Database Backups

A cron job runs daily at 2am on EC2, backing up the SQLite database to S3. To manage or test it:

```bash
# View installed cron jobs
crontab -l

# Edit cron jobs manually
crontab -e

# Test the S3 backup manually
set -a && source ~/farmec-v2/.env && set +a && \
aws s3 cp ~/farmec-v2/database/database.db \
  s3://farmec.ie/backups/database-$(date +%Y%m%d)-test.db

# List backups in S3
aws s3 ls s3://farmec.ie/backups/
```
