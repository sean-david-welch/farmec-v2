from catalog.models import Supplier


def suppliers(request):
    """Make suppliers available to all templates."""
    suppliers_list = Supplier.objects.publish().only('id', 'name').order_by('order')
    return {
        'suppliers': suppliers_list,
    }
