/// <reference types="astro/client" />

interface ImportMetaEnv {
  readonly FIREBASE_PROJECT_ID: string;
  readonly FIREBASE_PRIVATE_KEY_ID: string;
  readonly FIREBASE_PRIVATE_KEY: string;
  readonly FIREBASE_CLIENT_EMAIL: string;
  readonly FIREBASE_CLIENT_ID: string;
  readonly FIREBASE_AUTH_URI: string;
  readonly FIREBASE_TOKEN_URI: string;

  readonly FIREBASE_AUTH_CERT_URL: string;
  readonly FIREBASE_CLIENT_CERT_URL: string;

  readonly DATABASE_URL: string;
  readonly YOUTUBE_API_KEY: string;

  readonly AWS_ACCESS_KEY: string;
  readonly AWS_SECRET: string;

  readonly FB_WEB_API_KEY: string;
  readonly FB_PROJECT_ID: string;
  readonly FB_AUTH_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
