from django.views.generic import ListView

from .models import Privacy, PrivacyQuerySet, Terms, TermsQuerySet


class PrivacyListView(ListView):
    model: type[Privacy] = Privacy
    template_name: str = 'legal/policies.html'
    context_object_name: str = 'policies'
    queryset: PrivacyQuerySet = Privacy.objects.publish().order_by('-created')

    def get_context_data(self, **kwargs) -> dict[str, PrivacyQuerySet| TermsQuerySet]:
        context = super().get_context_data(**kwargs)
        context['terms'] = Terms.objects.publish().order_by('-created')
        return context
