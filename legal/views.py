from django.views.generic import ListView

from .models import Privacy, PrivacyQuerySet, Terms, TermsQuerySet


class TermsListView(ListView):
    model: type[Privacy] = Privacy
    template_name: str = 'legal/policies.html'
    context_object_name: str = 'policies'
    queryset: PrivacyQuerySet = Privacy.objects.publish().order_by('-created')
    terms_queryset: TermsQuerySet = Terms.objects.publish().order_by('-created')

    def get_context_data(self, **kwargs) -> dict[str, PrivacyQuerySet| TermsQuerySet]:
        context = super().get_context_data(**kwargs)
        context.update(terms=self.terms_queryset)
        return context
