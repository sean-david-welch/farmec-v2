from django.apps import AppConfig
from django.db.models.signals import pre_save


class CatalogConfig(AppConfig):
    name = 'catalog'

    def ready(self) -> None:
        from catalog.models import Video
        from catalog.signals import populate_youtube_metadata
        super().ready()
        pre_save.connect(receiver=populate_youtube_metadata, sender=Video, dispatch_uid='populate_youtube_metadata')
