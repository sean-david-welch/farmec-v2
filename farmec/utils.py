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
