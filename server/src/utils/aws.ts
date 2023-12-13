import { getSignedUrl } from '@aws-sdk/s3-request-presigner';
import type { PutObjectCommandInput } from '@aws-sdk/client-s3';
import { S3Client, PutObjectCommand } from '@aws-sdk/client-s3';

import secrets from './secrets';

const s3Client = new S3Client({
  region: 'eu-west-1',
  credentials: {
    accessKeyId: secrets.aws_access_key,
    secretAccessKey: secrets.aws_secret,
  },
});

export const generatePresignedUrl = async (
  bucketName: string,
  key: string,
  expiresInSeconds: number = 60
): Promise<string> => {
  const putObjectParams: PutObjectCommandInput = {
    Bucket: bucketName,
    Key: key,
  };

  const command = new PutObjectCommand(putObjectParams);

  try {
    const presignedUrl = await getSignedUrl(s3Client, command, {
      expiresIn: expiresInSeconds,
    });

    return presignedUrl;
  } catch (error) {
    console.error('Error creating presigned URL', error);
    throw error;
  }
};
