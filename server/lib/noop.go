package lib

type NoOpS3Client struct{}

func NewNoOpS3Client() *NoOpS3Client {
	return &NoOpS3Client{}
}

func (client *NoOpS3Client) GeneratePresignedUrl(folder, image string) (string, string, error) {
	return "https://example.com/fake-presigned-url", "https://example.com/fake-image-url", nil
}

func (client *NoOpS3Client) DeleteImageFromS3(imageUrl string) error {
	return nil
}
