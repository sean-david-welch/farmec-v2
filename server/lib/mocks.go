package lib

type MockS3Client struct{}

func NewMockS3Client() *MockS3Client {
	return &MockS3Client{}
}

func (client *MockS3Client) GeneratePresignedUrl(folder, image string) (string, string, error) {
	return "https://example.com/fake-presigned-url", "https://example.com/fake-image-url", nil
}

func (client *MockS3Client) DeleteImageFromS3(imageUrl string) error {
	return nil
}
