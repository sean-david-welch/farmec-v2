FROM python:3.13-slim

# WeasyPrint system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    libpango-1.0-0 \
    libpangoft2-1.0-0 \
    libpangocairo-1.0-0 \
    libgdk-pixbuf-2.0-0 \
    libcairo2 \
    libffi8 \
    fonts-liberation \
    && rm -rf /var/lib/apt/lists/*

# Install uv
COPY --from=ghcr.io/astral-sh/uv:latest /uv /usr/local/bin/uv

WORKDIR /app

# Install dependencies (cached layer)
COPY pyproject.toml uv.lock ./
RUN uv sync --frozen --no-dev

# Copy application code
COPY . .

EXPOSE 8000

CMD ["sh", "-c", "uv run python manage.py collectstatic --noinput && uv run gunicorn farmec.wsgi:application --bind 0.0.0.0:8000 --workers 3 --timeout 120 --access-logfile - --error-logfile -"]
