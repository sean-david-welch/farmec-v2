import { configDotenv } from 'dotenv';

configDotenv({
  path: '.env',
});

const secrets = {
  database_url: process.env.DATABASE_URL,

  project_id: process.env.FIREBASE_PROJECT_ID,
  private_key_id: process.env.FIREBASE_PRIVATE_KEY_ID,
  private_key: process.env.FIREBASE_PRIVATE_KEY,
  client_email: process.env.FIREBASE_CLIENT_EMAIL,
  client_id: process.env.FIREBASE_CLIENT_ID,
  auth_uri: process.env.FIREBASE_AUTH_URI,
  token_uri: process.env.FIREBASE_TOKEN_URI,
  auth_provider_x509_cert_url: process.env.FIREBASE_AUTH_CERT_URL,
  client_x509_cert_url: process.env.FIREBASE_CLIENT_CERT_URL,

  aws_access_key: process.env.AWS_ACCESS_KEY,
  aws_secret: process.env.AWS_SECRET,

  youtube_api_key: process.env.YOUTUBE_API_KEY,
};

export default secrets;
