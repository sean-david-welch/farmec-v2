import logging
from unittest.mock import patch

from django.core.files.uploadedfile import SimpleUploadedFile
from django.test import SimpleTestCase

from farmec.logging import describe_files, log_error

logger = logging.getLogger('farmec.tests')


class LogErrorTest(SimpleTestCase):
    def test_log_error__logs_at_error_level(self):
        with self.assertLogs('farmec.tests', level='ERROR') as logs:
            log_error(logger, 'Something %s', 'broke')
        self.assertIn('Something broke', logs.output[0])

    def test_log_error__sets_sentry_context(self):
        with patch('farmec.logging.sentry_sdk.set_context') as set_context:
            with self.assertLogs('farmec.tests', level='ERROR'):
                log_error(logger, 'Something broke', context={'a': 1})
        set_context.assert_called_once_with('farmec.tests', {'a': 1})

    def test_log_error__no_context_skips_sentry_context(self):
        with patch('farmec.logging.sentry_sdk.set_context') as set_context:
            with self.assertLogs('farmec.tests', level='ERROR'):
                log_error(logger, 'Something broke')
        set_context.assert_not_called()


class DescribeFilesTest(SimpleTestCase):
    def test_describe_files(self):
        f = SimpleUploadedFile('a.jpg', b'abc', content_type='image/jpeg')
        self.assertEqual(describe_files([f]), [{'name': 'a.jpg', 'size': 3, 'content_type': 'image/jpeg'}])
