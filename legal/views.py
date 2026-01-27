from django.views.generic import ListView

from .models import Privacy, PrivacyQuerySet, Terms, TermsQuerySet


class PrivacyListView(ListView):
    model: type[Privacy] = Privacy
    template_name: str = 'legal/privacy_list.html'
    context_object_name: str = 'privacy'
    queryset: PrivacyQuerySet = Privacy.objects.publish().order_by('-created')


class TermsListView(ListView):
    model: type[Terms] = Terms
    template_name: str = 'legal/terms_list.html'
    context_object_name: str = 'terms'
    queryset: TermsQuerySet = Terms.objects.publish().order_by('-created')
