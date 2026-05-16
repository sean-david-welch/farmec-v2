import css_inline
from django.conf import settings
from django.core.mail import EmailMultiAlternatives
from django.db.models import QuerySet
from django.template.loader import render_to_string
from django.utils import timezone

from support.models import Warrantyclaim, Partsrequired, Machineregistration


class EmailClient:
    email_css: str = (settings.BASE_DIR / 'theme' / 'static' / 'css' / 'emails.css').read_text()
    inliner: css_inline.CSSInliner = css_inline.CSSInliner(extra_css=email_css)
    from_email: str = 'Farmec Ireland Ltd <info@farmec.ie>'

    def __init__(self) -> None:
        self.recipient: str = settings.EMAIL_USER

    def send(self, subject: str, text: str, html: str, cc: list[str] | None = None) -> None:
        """
        Send an HTML email with a plain-text fallback via Django's email backend.

        :param subject: Email subject line.
        :param text: Plain-text fallback body.
        :param html: HTML body (CSS already inlined).
        :param cc: Optional list of CC recipients.
        """
        email: EmailMultiAlternatives = EmailMultiAlternatives(
            subject=subject,
            body=text,
            from_email=self.from_email,
            to=[self.recipient],
            cc=cc or [],
        )
        email.attach_alternative(html, 'text/html')
        email.send()

    def send_contact_notification(self, name: str, email: str, message: str) -> None:
        """
        Send a contact form notification.

        :param name: Sender's name.
        :param email: Sender's email address.
        :param message: Message body.
        """
        subject: str = f"New Contact Form from {name} - {email}"
        context: dict[str, str] = {
            "name": name,
            "email": email,
            "message": message,
            "generated_date": timezone.now().strftime("%d %b %Y, %H:%M"),
        }
        html: str = self.inliner.inline(render_to_string("support/email/contact.html", context))
        self.send(subject=subject, text=message, html=html)

    def send_warranty_notification(self, claim: Warrantyclaim, parts: QuerySet[Partsrequired], images: QuerySet) -> None:
        """
        Send a warranty claim submission notification.

        :param claim: Submitted ``Warrantyclaim`` instance.
        :param parts: Related ``Partsrequired`` queryset for the claim.
        :param images: Related ``WarrantyImage`` queryset for the claim.
        """
        subject: str = f"New Warranty Claim - {claim.owner_name} / {claim.machine_model}"
        context: dict[str, object] = {
            "claim": claim,
            "parts": parts,
            "images": images,
            "generated_date": timezone.now().strftime("%d %b %Y, %H:%M"),
        }
        html: str = self.inliner.inline(render_to_string("support/email/warranty_claim.html", context))
        text: str = (
            f"A new warranty claim has been submitted.\n\n"
            f"Dealer: {claim.dealer}\n"
            f"Owner: {claim.owner_name}\n"
            f"Machine Model: {claim.machine_model}\n"
            f"Serial Number: {claim.serial_number}\n"
        )
        cc: list[str] = [claim.dealer_contact] if claim.dealer_contact else []
        self.send(subject=subject, text=text, html=html, cc=cc)

    def send_registration_notification(self, reg: Machineregistration) -> None:
        """
        Send a machine registration submission notification.

        :param reg: Submitted ``Machineregistration`` instance.
        """
        subject: str = f"New Machine Registration - {reg.owner_name} / {reg.machine_model}"
        context: dict[str, str] = {
            "reg": reg,
            "generated_date": timezone.now().strftime("%d %b %Y, %H:%M"),
        }
        html: str = self.inliner.inline(render_to_string("support/email/machine_registration.html", context))
        text: str = (
            f"A new machine registration has been submitted.\n\n"
            f"Dealer: {reg.dealer_name}\n"
            f"Owner: {reg.owner_name}\n"
            f"Machine Model: {reg.machine_model}\n"
            f"Serial Number: {reg.serial_number}\n"
        )
        self.send(subject=subject, text=text, html=html)
