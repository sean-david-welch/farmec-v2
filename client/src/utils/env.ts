const config = {
    baseUrl: import.meta.env.BASE_URL,

    firebaseApiKey: import.meta.env.FB_WEB_API_KEY,
    firebaseAuthDomain: import.meta.env.FB_AUTH_URL,
    firebaseProjectId: import.meta.env.FB_PROJECT_ID,

    mapsKey: import.meta.env.PUBLIC_MAPS_KEY,
    recaptchaKey: import.meta.env.RECAPTCHA_PUBLIC_KEY,

    stripePublicKey: import.meta.env.STRIPE_PUBLIC_KEY,
};

export default config;
