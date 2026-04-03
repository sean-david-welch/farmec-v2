from typing import Any, Type
from urllib.parse import urlparse, parse_qs, ParseResult
from django.conf import settings
from googleapiclient.discovery import build, Resource
from catalog.models import Video


def populate_youtube_metadata(sender: Type[Video], instance: Video, **kwargs: Any) -> None:
    """
    Populate YouTube metadata on a Video instance before it is saved.

    Queries the YouTube Data API v3 to fill ``video_id``, ``title``,
    ``description``, and ``thumbnail_url`` whenever ``web_url`` contains a
    recognisable YouTube URL.

    Supported URL formats:

    - ``https://youtu.be/<id>``
    - ``https://www.youtube.com/watch?v=<id>``
    - ``https://www.youtube.com/embed/<id>``
    - ``https://www.youtube.com/v/<id>``

    :param sender: The model class that sent the signal.
    :param instance: The Video instance about to be saved.
    :param kwargs: Additional keyword arguments passed by the signal.
    :rtype: None
    """
    if not instance.web_url:
        return
    parsed: ParseResult = urlparse(instance.web_url)
    video_id: str | None
    if parsed.hostname == 'youtu.be':
        video_id = parsed.path.lstrip('/')
    elif parsed.hostname in ('www.youtube.com', 'youtube.com') and parsed.path == '/watch':
        video_id = parse_qs(parsed.query).get('v', [None])[0]
    elif parsed.hostname in ('www.youtube.com', 'youtube.com') and parsed.path.startswith(('/embed/', '/v/')):
        video_id = parsed.path.split('/')[2] or None
    else:
        video_id = None
    if video_id:
        youtube: Resource = build('youtube', 'v3', developerKey=settings.YOUTUBE_API_KEY)
        response: dict[str, Any] = youtube.videos().list(part='id,snippet', id=video_id).execute()
        if response.get('items'):
            item: dict[str, Any] = response['items'][0]
            instance.video_id = item['id']
            instance.title = item['snippet']['title']
            instance.description = item['snippet']['description']
            instance.thumbnail_url = item['snippet']['thumbnails']['medium']['url']
