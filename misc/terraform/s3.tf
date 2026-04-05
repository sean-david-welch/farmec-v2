resource "aws_s3_bucket" "farmec_media" {
  bucket = "farmec.ie"
}

resource "aws_s3_bucket_website_configuration" "farmec_media" {
  bucket = aws_s3_bucket.farmec_media.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }
}

resource "aws_s3_bucket_cors_configuration" "farmec_media" {
  bucket = aws_s3_bucket.farmec_media.id

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "PUT", "POST", "DELETE", "HEAD"]
    allowed_origins = ["*", "https://www.farmec.ie"]
    max_age_seconds = 3000
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "farmec_media" {
  bucket = aws_s3_bucket.farmec_media.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
    bucket_key_enabled = false
  }
}

resource "aws_s3_bucket" "farmec_backups" {
  bucket = "farmec-backups"
}

resource "aws_s3_bucket_public_access_block" "farmec_backups" {
  bucket = aws_s3_bucket.farmec_backups.id

  block_public_acls       = true
  ignore_public_acls      = true
  block_public_policy     = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_server_side_encryption_configuration" "farmec_backups" {
  bucket = aws_s3_bucket.farmec_backups.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
    bucket_key_enabled = false
  }
}
