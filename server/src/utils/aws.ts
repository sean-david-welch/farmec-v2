import { getSignedUrl } from '@aws-sdk/s3-request-presigner';
import type { PutObjectCommandInput } from '@aws-sdk/client-s3';
import { S3Client, PutObjectCommand, DeleteObjectCommand } from '@aws-sdk/client-s3';

import secrets from './secrets';

const s3Client = new S3Client({
  region: 'eu-west-1',
  credentials: {
    accessKeyId: secrets.aws_access_key,
    secretAccessKey: secrets.aws_secret,
  },
});

export const generatePresignedUrl = async (folder: string, image: string): Promise<Record<string, string>> => {
  const bucketName = 'farmec-bucket';
  const imageKey = `${folder}/${image}`;
  const imageUrl = `https://${bucketName}.s3.amazonaws.com/${imageKey}`;

  const putObjectParams: PutObjectCommandInput = {
    Bucket: bucketName,
    Key: imageKey,
  };

  const command = new PutObjectCommand(putObjectParams);

  try {
    const presignedUrl = await getSignedUrl(s3Client, command, {
      expiresIn: 60,
    });

    const uploadDetails = {
      uploadUrl: presignedUrl,
      imageUrl: imageUrl,
    };

    return uploadDetails;
  } catch (error) {
    console.error('Error creating presigned URL', error);
    throw error;
  }
};

export const generateDeletePresignedUrl = async (key: string): Promise<string> => {
  const bucketName = 'farmec-bucket';

  const deleteObjectParams = {
    Bucket: bucketName,
    Key: key,
  };

  const command = new DeleteObjectCommand(deleteObjectParams);

  try {
    const presignedUrl = await getSignedUrl(s3Client, command, {
      expiresIn: 60,
    });

    return presignedUrl;
  } catch (error) {
    console.error('Error creating presigned URL for deletion', error);
    throw error;
  }
};
