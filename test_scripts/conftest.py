"""
Pytest configuration for migration tests.

Configures Django settings and database for testing.
"""

import os
import django

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'settings')

django.setup()
