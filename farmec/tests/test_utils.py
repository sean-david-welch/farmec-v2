from django.test import SimpleTestCase

from farmec.utils import parse_int


class ParseIntTest(SimpleTestCase):
    def test_parse_int__valid(self):
        self.assertEqual(parse_int('3', 1), 3)

    def test_parse_int__none_returns_default(self):
        self.assertEqual(parse_int(None, 7), 7)

    def test_parse_int__empty_returns_default(self):
        self.assertEqual(parse_int('', 7), 7)

    def test_parse_int__invalid_returns_default(self):
        self.assertEqual(parse_int('abc', 7), 7)

    def test_parse_int__float_string_returns_default(self):
        self.assertEqual(parse_int('2.5', 7), 7)
