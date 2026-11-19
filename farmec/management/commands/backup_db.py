from django.core.management.base import BaseCommand, CommandError
from django.tasks import TaskResultStatus

from farmec.tasks import backup_database


class Command(BaseCommand):
    help = 'Back up the SQLite database to S3 (run nightly from cron).'

    def handle(self, *args, **options) -> None:
        result = backup_database.enqueue()
        if result.status != TaskResultStatus.SUCCESSFUL:
            raise CommandError(
                'Backup {status}: {errors}'.format(status=result.status, errors=result.errors)
            )
        message: str = 'Backed up to {key}'.format(key=result.return_value)
        self.stdout.write(self.style.SUCCESS(message))
