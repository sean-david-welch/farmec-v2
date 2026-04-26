from unittest.mock import patch

from django.test import TestCase, Client
from django.urls import reverse_lazy
from django.contrib.auth.models import User
from model_bakery import baker

from content.models import Blog, Carousel, Exhibition


class HomeViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.carousel = baker.make(Carousel, publish=True)
        cls.url = reverse_lazy('content:home')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)
        self.htmx_client = Client(HTTP_HX_Request='true')
        self.htmx_client.force_login(self.user)

    def test_home__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_home__context(self):
        response = self.client.get(self.url)
        with self.subTest('carousels in context'):
            self.assertIn(self.carousel, response.context['carousels'])
        with self.subTest('contact form in context'):
            self.assertIn('form', response.context)

    def test_home__htmx_contact_valid(self):
        with patch('content.views.EmailClient') as mock_email:
            response = self.htmx_client.post(self.url, data={
                'name': 'Test User',
                'email': 'test@example.com',
                'message': 'Hello there',
            })
        self.assertEqual(response.status_code, 200)
        mock_email().send_contact_notification.assert_called_once_with(
            name='Test User',
            email='test@example.com',
            message='Hello there',
        )

    def test_home__htmx_contact_invalid(self):
        response = self.htmx_client.post(self.url, data={'name': '', 'email': '', 'message': ''})
        self.assertEqual(response.status_code, 200)


class BlogListViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.blog = baker.make(Blog, title='Published Post', publish=True)
        cls.url = reverse_lazy('content:blog_list')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_blog_list__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_blog_list__published_only(self):
        hidden = baker.make(Blog, publish=False)
        response = self.client.get(self.url)
        with self.subTest('published blog in context'):
            self.assertIn(self.blog, response.context['blogs'])
        with self.subTest('unpublished blog not in context'):
            self.assertNotIn(hidden, response.context['blogs'])

    def test_blog_list__anonymous(self):
        self.client.logout()
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)


class BlogDetailViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.blog = baker.make(Blog, title='Test Post', publish=True)
        cls.url = reverse_lazy('content:blog_detail', kwargs={'id': cls.blog.pk})

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_blog_detail__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_blog_detail__context(self):
        response = self.client.get(self.url)
        self.assertEqual(response.context['blog'], self.blog)

    def test_blog_detail__unpublished_returns_404(self):
        hidden = baker.make(Blog, publish=False)
        response = self.client.get(reverse_lazy('content:blog_detail', kwargs={'id': hidden.pk}))
        self.assertEqual(response.status_code, 404)


class ExhibitionListViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.exhibition = baker.make(Exhibition, title='Test Exhibition', publish=True)
        cls.url = reverse_lazy('content:exhibition_list')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_exhibition_list__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_exhibition_list__published_only(self):
        hidden = baker.make(Exhibition, publish=False)
        response = self.client.get(self.url)
        with self.subTest('published exhibition in context'):
            self.assertIn(self.exhibition, response.context['exhibitions'])
        with self.subTest('unpublished exhibition not in context'):
            self.assertNotIn(hidden, response.context['exhibitions'])

    def test_exhibition_list__anonymous(self):
        self.client.logout()
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)
