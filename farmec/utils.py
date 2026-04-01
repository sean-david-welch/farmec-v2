import resend
from django.conf import settings


class EmailClient:
    def __init__(self):
        resend.api_key = settings.RESEND_TOKEN
        self.recipient = settings.EMAIL_USER

    def send_contact_notification(self, name: str, email: str, message: str) -> None:
        subject: str = f"New Contact Form from {name}--{email}"
        params: resend.Emails.SendParams = {
            "from": "Farmec Ireland Ltd <noreply@farmec.ie>",
            "to": [self.recipient],
            "subject": subject,
            "text": message,
            "html": f"<p>{message}</p>",
        }
        resend.Emails.send(params)

    def send_warranty_notification(self, dealer: str, owner_name: str, machine_model: str, serial_number: str) -> None:
        subject: str = f"New Warranty Claim - {owner_name} / {machine_model}"
        text: str = (
            f"A new warranty claim has been submitted.\n\n"
            f"Dealer: {dealer}\n"
            f"Owner: {owner_name}\n"
            f"Machine Model: {machine_model}\n"
            f"Serial Number: {serial_number}\n"
        )
        html: str = (
            f"<p>A new warranty claim has been submitted.</p>"
            f"<ul>"
            f"<li><strong>Dealer:</strong> {dealer}</li>"
            f"<li><strong>Owner:</strong> {owner_name}</li>"
            f"<li><strong>Machine Model:</strong> {machine_model}</li>"
            f"<li><strong>Serial Number:</strong> {serial_number}</li>"
            f"</ul>"
        )
        params: resend.Emails.SendParams = {
            "from": "Farmec Ireland Ltd <noreply@farmec.ie>",
            "to": [self.recipient],
            "subject": subject,
            "text": text,
            "html": html,
        }
        resend.Emails.send(params)

    def send_registration_notification(self, dealer_name: str, owner_name: str, machine_model: str, serial_number: str) -> None:
        subject: str = f"New Machine Registration - {owner_name} / {machine_model}"
        text: str = (
            f"A new machine registration has been submitted.\n\n"
            f"Dealer: {dealer_name}\n"
            f"Owner: {owner_name}\n"
            f"Machine Model: {machine_model}\n"
            f"Serial Number: {serial_number}\n"
        )
        html: str = (
            f"<p>A new machine registration has been submitted.</p>"
            f"<ul>"
            f"<li><strong>Dealer:</strong> {dealer_name}</li>"
            f"<li><strong>Owner:</strong> {owner_name}</li>"
            f"<li><strong>Machine Model:</strong> {machine_model}</li>"
            f"<li><strong>Serial Number:</strong> {serial_number}</li>"
            f"</ul>"
        )
        params: resend.Emails.SendParams = {
            "from": "Farmec Ireland Ltd <noreply@farmec.ie>",
            "to": [self.recipient],
            "subject": subject,
            "text": text,
            "html": html,
        }
        resend.Emails.send(params)
