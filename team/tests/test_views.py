from django.test import TestCase
from django.urls import reverse_lazy
from django.contrib.auth.models import User
from model_bakery import baker

from team.models import Employee
from content.models import Timeline


class EmployeeListViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.employee = baker.make(Employee, name='John Doe', publish=True)
        cls.timeline = baker.make(Timeline, publish=True)
        cls.url = reverse_lazy('team:employee_list')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_employee_list__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_employee_list__context(self):
        response = self.client.get(self.url)
        with self.subTest('employee in context'):
            self.assertIn(self.employee, response.context['employees'])
        with self.subTest('timeline in context'):
            self.assertIn(self.timeline, response.context['timelines'])

    def test_employee_list__published_only(self):
        hidden_employee = baker.make(Employee, publish=False)
        hidden_timeline = baker.make(Timeline, publish=False)
        response = self.client.get(self.url)
        with self.subTest('unpublished employee not in context'):
            self.assertNotIn(hidden_employee, response.context['employees'])
        with self.subTest('unpublished timeline not in context'):
            self.assertNotIn(hidden_timeline, response.context['timelines'])

    def test_employee_list__anonymous(self):
        self.client.logout()
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)
