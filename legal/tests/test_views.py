from django.test import TestCase
from django.urls import reverse_lazy
from django.contrib.auth.models import User
from model_bakery import baker

from legal.models import Privacy, Terms


class TermsListViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.privacy = baker.make(Privacy, title='Privacy Policy', publish=True)
        cls.terms = baker.make(Terms, title='Terms & Conditions', publish=True)
        cls.url = reverse_lazy('legal:privacy_list')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_privacy_list__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_privacy_list__context(self):
        response = self.client.get(self.url)
        with self.subTest('privacy in context'):
            self.assertIn(self.privacy, response.context['policies'])
        with self.subTest('terms in context'):
            self.assertIn(self.terms, response.context['terms'])

    def test_privacy_list__published_only(self):
        hidden_privacy = baker.make(Privacy, publish=False)
        hidden_terms = baker.make(Terms, publish=False)
        response = self.client.get(self.url)
        with self.subTest('unpublished privacy not in context'):
            self.assertNotIn(hidden_privacy, response.context['policies'])
        with self.subTest('unpublished terms not in context'):
            self.assertNotIn(hidden_terms, response.context['terms'])

    def test_privacy_list__anonymous(self):
        self.client.logout()
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)
