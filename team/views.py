from django.views.generic import ListView

from .models import Employee, EmployeeQuerySet


class EmployeeListView(ListView):
    model: type[Employee] = Employee
    template_name: str = 'team/employee_list.html'
    context_object_name: str = 'employees'
    queryset: EmployeeQuerySet = Employee.objects.publish().order_by('-created')
