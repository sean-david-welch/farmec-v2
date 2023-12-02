import utils from '../../styles/Utils.module.css';

const ContactForm = () => {
  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget as HTMLFormElement);

    const body = {
      name: formData.get('name'),
      email: formData.get('email'),
      message: formData.get('message'),
    };

    const response = await fetch('http://localhost:4321/api/contact', {
      method: 'POST',
      body: JSON.stringify(body),
    });
  };

  return (
    <form onSubmit={handleSubmit} class={utils.contactForm}>
      <label for="name">Name:</label>
      <input type="text" name="name" required={true} placeholder="name" />
      <label for="email">Email:</label>
      <input type="email" name="email" required={true} placeholder="email" />
      <label for="message">Message:</label>
      <textarea name="message" placeholder="Enter your message here..." cols={30} rows={11} required={true} />

      <div class={utils.recaptcha}></div>

      <button class={utils.btnForm} type="submit">
        Submit
      </button>
    </form>
  );
};

export default ContactForm;
