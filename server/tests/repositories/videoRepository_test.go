package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetVideos(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database")
	}
	defer db.Close()

	title1 := "Video Title 1"
	description1 := "Description for Video 1"
	videoID1 := "vID1"
	thumbnailURL1 := "http://example.com/thumbnail1.jpg"

	title2 := "Video Title 2"
	description2 := "Description for Video 2"
	videoID2 := "vID2"
	thumbnailURL2 := "http://example.com/thumbnail2.jpg"

	videos := []types.Video{
		{
			ID:           "1",
			SupplierID:   "12",
			WebURL:       "http://example.com/video1",
			Title:        &title1,
			Description:  &description1,
			VideoID:      &videoID1,
			ThumbnailURL: &thumbnailURL1,
			Created:      time.Now(),
		},
		{
			ID:           "2",
			SupplierID:   "12",
			WebURL:       "http://example.com/video2",
			Title:        &title2,
			Description:  &description2,
			VideoID:      &videoID2,
			ThumbnailURL: &thumbnailURL2,
			Created:      time.Now(),
		},
	}

	supplierId := videos[0].SupplierID

	rows := sqlmock.NewRows([]string{"id", "supplierId", "web_url", "title", "description", "video_id", "thumbnail_url", "created"})
	for _, video := range videos {
		rows.AddRow(video.ID, video.SupplierID, video.WebURL, video.Title, video.Description, video.VideoID, video.ThumbnailURL, video.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Video" WHERE "supplierId" = \$1 `).WillReturnRows(rows)

	repo := repository.NewVideoRepository(db)
	retrieved, err := repo.GetVideos(supplierId)
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(videos))
		assert.Equal(test, videos, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}

func TestCreateVideo(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database")
	}
	defer db.Close()

	title1 := "Video Title 1"
	description1 := "Description for Video 1"
	videoID1 := "vID1"
	thumbnailURL1 := "http://example.com/thumbnail1.jpg"

	video := &types.Video{
		ID:           "1",
		SupplierID:   "12",
		WebURL:       "http://example.com/video1",
		Title:        &title1,
		Description:  &description1,
		VideoID:      &videoID1,
		ThumbnailURL: &thumbnailURL1,
		Created:      time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "Video" \("id", "supplierId", "web_url", "title", "description", "video_id", "thumbnail_url", "created"\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), video.WebURL, video.Title, video.Description, video.VideoID, video.ThumbnailURL, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewVideoRepository(db)
	err = repo.CreateVideo(video)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}
