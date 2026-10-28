import logging
from collections.abc import Iterable
from typing import Any

import sentry_sdk
from django.core.files.uploadedfile import UploadedFile

LOGGING: dict[str, Any] = {
    'version': 1,
    'disable_existing_loggers': False,
    'formatters': {
        'verbose': {'format': '{asctime} {levelname} {name} {message}', 'style': '{'},
    },
    'handlers': {
        'console': {'class': 'logging.StreamHandler', 'formatter': 'verbose'},
    },
    'root': {'handlers': ['console'], 'level': 'INFO'},
    'loggers': {
        'django.request': {'handlers': ['console'], 'level': 'WARNING', 'propagate': False},
        # Bot scanners sending bad Host headers; pure noise
        'django.security.DisallowedHost': {'handlers': [], 'propagate': False},
    },
}


def log_error(
    logger: logging.Logger, message: str, *args: Any,
    context: dict[str, Any] | None = None, tags: dict[str, str] | None = None, exc_info: bool = False,
) -> None:
    """Log at ERROR and attach structured context to the Sentry event.

    :param logger: Logger to emit on, normally ``logging.getLogger(__name__)``.
    :param message: Log message, ``%``-style format string.
    :param args: Values interpolated into ``message``.
    :param context: Extra data attached to the Sentry event under ``logger.name``.
    :param tags: Searchable Sentry tags, e.g. ``{'feature': 'warranty'}``.
    :param exc_info: Include the current exception traceback.
    """
    if context:
        sentry_sdk.set_context(logger.name, context)
    for key, value in (tags or {}).items():
        sentry_sdk.set_tag(key, value)
    logger.error(message, *args, exc_info=exc_info)


def describe_files(files: Iterable[UploadedFile]) -> list[dict[str, Any]]:
    """
    Describe uploads for log context.

    :param files: Uploaded files, e.g. ``request.FILES.getlist(name)``.
    :returns: One dict per file with ``name``, ``size`` and ``content_type``.
    """
    return [{'name': f.name, 'size': f.size, 'content_type': f.content_type} for f in files]
