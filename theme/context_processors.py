from googleapiclient.http import HttpRequest

from catalog.models import Supplier, SupplierQuerySet


def suppliers(request: HttpRequest) -> dict[str, SupplierQuerySet]:
    """
    Make suppliers available to all templates including in the Navbar context

    :param request: The HttpRequest object
    """
    suppliers_list: SupplierQuerySet = Supplier.objects.publish().only('id', 'name', 'slug').order_by('order')
    return {
        'suppliers': suppliers_list,
    }
