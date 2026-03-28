from django import template

register = template.Library()


@register.inclusion_tag('components/account_button.html', takes_context=True)
def account_button(context):
    """Render the account/login button with dropdown menu."""
    user = context.get('user')
    is_authenticated = user.is_authenticated if user else False
    is_admin = user.is_staff if is_authenticated else False

    return {
        'is_authenticated': is_authenticated,
        'is_admin': is_admin,
        'user': user,
    }


@register.inclusion_tag('components/mobile_login.html', takes_context=True)
def mobile_login(context, on_click=""):
    """Render the mobile login/account button."""
    user = context.get('user')
    is_authenticated = user.is_authenticated if user else False
    is_admin = user.is_staff if is_authenticated else False

    return {
        'is_authenticated': is_authenticated,
        'is_admin': is_admin,
        'on_click': on_click,
    }


@register.inclusion_tag('components/download_pdf.html')
def download_pdf(pdf_type, warranty_claim=None, registration=None):
    """Render a PDF download button."""
    return {
        'pdf_type': pdf_type,
        'warranty_claim': warranty_claim,
        'registration': registration,
    }


@register.inclusion_tag('components/map.html')
def google_map(lat=53.49, lng=-6.54, zoom=10):
    """Render a Google Map embed."""
    return {
        'lat': lat,
        'lng': lng,
        'zoom': zoom,
    }


@register.inclusion_tag('components/social_links.html')
def social_links(facebook=None, twitter=None, instagram=None, linkedin=None, website=None, youtube=None):
    """Render social media links."""
    return {
        'facebook': facebook,
        'twitter': twitter,
        'instagram': instagram,
        'linkedin': linkedin,
        'website': website,
        'youtube': youtube,
    }


@register.inclusion_tag('components/timeline_card.html', takes_context=True)
def timeline_card(context, timeline):
    """Render a timeline card."""
    user = context.get('user')
    is_admin = user.is_staff if user and user.is_authenticated else False

    return {
        'timeline': timeline,
        'is_admin': is_admin,
    }


@register.inclusion_tag('components/to_top_button.html')
def to_top_button():
    """Render a scroll-to-top button."""
    return {}
