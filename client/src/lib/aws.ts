interface ImageData {
    imageFile: File;
    presignedUrl: string;
}

export const uploadFileToS3 = async (imageData: ImageData) => {
    const { imageFile, presignedUrl } = imageData;

    try {
        const response = await fetch(presignedUrl, {
            method: 'PUT',
            headers: {
                'Content-Type': imageFile.type,
            },
            body: imageFile,
        });

        if (!response.ok) {
            throw new Error(`Failed to upload file: ${response.statusText}`);
        }

        return { success: true, status: response.status };
    } catch (error) {
        console.error('Error in uploadFileToS3:', error);
        console.error('Failed imageData:', imageData);
        throw error;
    }
};
