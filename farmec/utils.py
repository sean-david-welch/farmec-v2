import resend
from django.conf import settings
from django.template.loader import render_to_string
from django.utils import timezone


class EmailClient:
    def __init__(self):
        resend.api_key = settings.RESEND_TOKEN
        self.recipient = settings.EMAIL_USER

    def send_contact_notification(self, name: str, email: str, message: str) -> None:
        subject: str = f"New Contact Form from {name} - {email}"
        context = {
            "name": name,
            "email": email,
            "message": message,
            "generated_date": timezone.now().strftime("%d %b %Y, %H:%M"),
        }
        html: str = render_to_string("support/email/contact.html", context)
        params: resend.Emails.SendParams = {
            "from": "Farmec Ireland Ltd <noreply@farmec.ie>",
            "to": [self.recipient],
            "subject": subject,
            "text": message,
            "html": html,
        }
        resend.Emails.send(params)

    def send_warranty_notification(self, claim, parts) -> None:
        subject: str = f"New Warranty Claim - {claim.owner_name} / {claim.machine_model}"
        context = {
            "claim": claim,
            "parts": parts,
            "generated_date": timezone.now().strftime("%d %b %Y, %H:%M"),
        }
        html: str = render_to_string("support/email/warranty_claim.html", context)
        text: str = (
            f"A new warranty claim has been submitted.\n\n"
            f"Dealer: {claim.dealer}\n"
            f"Owner: {claim.owner_name}\n"
            f"Machine Model: {claim.machine_model}\n"
            f"Serial Number: {claim.serial_number}\n"
        )
        params: resend.Emails.SendParams = {
            "from": "Farmec Ireland Ltd <noreply@farmec.ie>",
            "to": [self.recipient],
            "subject": subject,
            "text": text,
            "html": html,
        }
        resend.Emails.send(params)

    def send_registration_notification(self, reg) -> None:
        subject: str = f"New Machine Registration - {reg.owner_name} / {reg.machine_model}"
        context = {
            "reg": reg,
            "generated_date": timezone.now().strftime("%d %b %Y, %H:%M"),
        }
        html: str = render_to_string("support/email/machine_registration.html", context)
        text: str = (
            f"A new machine registration has been submitted.\n\n"
            f"Dealer: {reg.dealer_name}\n"
            f"Owner: {reg.owner_name}\n"
            f"Machine Model: {reg.machine_model}\n"
            f"Serial Number: {reg.serial_number}\n"
        )
        params: resend.Emails.SendParams = {
            "from": "Farmec Ireland Ltd <noreply@farmec.ie>",
            "to": [self.recipient],
            "subject": subject,
            "text": text,
            "html": html,
        }
        resend.Emails.send(params)
