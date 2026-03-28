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


@register.inclusion_tag('components/to_top_button.html')
def to_top_button():
    """Render a scroll-to-top button."""
    return {}
