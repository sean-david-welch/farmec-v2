from django.views.generic import ListView

from content.models import Timeline, TimelineQuerySet
from team.models import Employee, EmployeeQuerySet


class EmployeeListView(ListView):
    model: type[Employee] = Employee
    template_name: str = 'pages/about.html'
    context_object_name: str = 'employees'
    queryset: EmployeeQuerySet = Employee.objects.publish().order_by('order')

    def get_context_data(self, **kwargs) -> dict[str, EmployeeQuerySet | TimelineQuerySet]:
        context = super().get_context_data(**kwargs)
        context.update(timelines=Timeline.objects.publish().order_by('order'))
        return context
