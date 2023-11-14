import { S3Client, PutObjectCommand } from '@aws-sdk/client-s3';
import type { PutObjectCommandInput } from '@aws-sdk/client-s3';

import { getSignedUrl } from '@aws-sdk/s3-request-presigner';

const s3Client = new S3Client({
  region: 'eu-west-1',
  credentials: {
    accessKeyId: process.env.AWS_ACCESS_KEY!,
    secretAccessKey: process.env.AWS_SECRET!,
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
