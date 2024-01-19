package mocks

import "github.com/stretchr/testify/mock"

type MockS3Client struct {
    mock.Mock
}

func (m *MockS3Client) GeneratePresignedUrl(folder, image string) (string, string, error) {
    args := m.Called(folder, image)
    return args.String(0), args.String(1), args.Error(2)
}

func (m *MockS3Client) DeleteImageFromS3(imageUrl string) error {
    args := m.Called(imageUrl)
    return args.Error(0)
}

