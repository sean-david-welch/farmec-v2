from catalog.models import Supplier, SupplierQuerySet


def suppliers(request):
    """Make suppliers available to all templates."""
    suppliers_list: SupplierQuerySet = Supplier.objects.publish().only('id', 'name', 'slug').order_by('-created')
    return {
        'suppliers': suppliers_list,
    }
