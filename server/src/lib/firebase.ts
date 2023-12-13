import type { ServiceAccount } from 'firebase-admin';
import { initializeApp, cert, getApps } from 'firebase-admin/app';

import secrets from '../utils/secrets';

const activeApps = getApps();
const serviceAccount = {
  type: 'service_account',
  project_id: secrets.project_id,
  private_key_id: secrets.private_key_id,
  private_key: secrets.private_key,
  client_email: secrets.client_email,
  client_id: secrets.client_id,
  auth_uri: secrets.auth_uri,
  token_uri: secrets.token_uri,
  auth_provider_x509_cert_url: secrets.auth_provider_x509_cert_url,
  client_x509_cert_url: secrets.client_x509_cert_url,
};

export const app =
  activeApps.length === 0
    ? initializeApp({
        credential: cert(serviceAccount as ServiceAccount),
      })
    : activeApps[0];
