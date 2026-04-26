from django.test import TestCase, Client
from django.urls import reverse_lazy
from django.contrib.auth.models import User
from model_bakery import baker

from catalog.models import Supplier, Machine, Product, Spareparts


class SupplierListViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.supplier = baker.make(Supplier, name='Test Supplier', slug='test-supplier', publish=True)
        cls.url = reverse_lazy('catalog:supplier_list')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_supplier_list__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_supplier_list__published_only(self):
        hidden = baker.make(Supplier, publish=False)
        response = self.client.get(self.url)
        with self.subTest('published supplier in context'):
            self.assertIn(self.supplier, response.context['suppliers'])
        with self.subTest('unpublished supplier not in context'):
            self.assertNotIn(hidden, response.context['suppliers'])

    def test_supplier_list__anonymous(self):
        self.client.logout()
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)


class SupplierDetailViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.supplier = baker.make(Supplier, name='Test Supplier', slug='test-supplier', publish=True)
        cls.machine = baker.make(Machine, supplier=cls.supplier, publish=True)
        cls.url = reverse_lazy('catalog:supplier_detail', kwargs={'slug': cls.supplier.slug})

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_supplier_detail__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_supplier_detail__context(self):
        response = self.client.get(self.url)
        with self.subTest('supplier in context'):
            self.assertEqual(response.context['supplier'], self.supplier)
        with self.subTest('related machines in context'):
            self.assertIn(self.machine, response.context['machines'])

    def test_supplier_detail__unpublished_returns_404(self):
        hidden = baker.make(Supplier, slug='hidden', publish=False)
        response = self.client.get(reverse_lazy('catalog:supplier_detail', kwargs={'slug': hidden.slug}))
        self.assertEqual(response.status_code, 404)

    def test_supplier_detail__anonymous(self):
        self.client.logout()
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)


class MachineDetailViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.supplier = baker.make(Supplier, publish=True)
        cls.machine = baker.make(Machine, supplier=cls.supplier, slug='test-machine', publish=True)
        cls.product = baker.make(Product, machine=cls.machine, publish=True)
        cls.url = reverse_lazy('catalog:machine_detail', kwargs={'slug': cls.machine.slug})

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_machine_detail__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_machine_detail__context(self):
        response = self.client.get(self.url)
        with self.subTest('machine in context'):
            self.assertEqual(response.context['machine'], self.machine)
        with self.subTest('related products in context'):
            self.assertIn(self.product, response.context['products'])

    def test_machine_detail__unpublished_returns_404(self):
        hidden = baker.make(Machine, slug='hidden-machine', publish=False)
        response = self.client.get(reverse_lazy('catalog:machine_detail', kwargs={'slug': hidden.slug}))
        self.assertEqual(response.status_code, 404)


class SparePartsIndexViewTest(TestCase):
    @classmethod
    def setUpTestData(cls):
        super().setUpTestData()
        cls.user = User.objects.create_user(username='testuser', password='testpass')
        cls.supplier = baker.make(Supplier, publish=True)
        cls.sparepart = baker.make(Spareparts, supplier=cls.supplier, publish=True)
        cls.url = reverse_lazy('catalog:spareparts')

    def setUp(self):
        super().setUp()
        self.client.force_login(self.user)

    def test_spareparts_index__returns_200(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 200)

    def test_spareparts_index__only_suppliers_with_parts(self):
        supplier_no_parts = baker.make(Supplier, publish=True)
        response = self.client.get(self.url)
        with self.subTest('supplier with parts in context'):
            self.assertIn(self.supplier, response.context['spareparts'])
        with self.subTest('supplier without parts not in context'):
            self.assertNotIn(supplier_no_parts, response.context['spareparts'])
