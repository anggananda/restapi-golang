package tests

import (
	"context"
	"errors"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupMockKHSData() []models.Khs {
	tDummy := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	return []models.Khs{
		{
			ID: 1, NIM: "A001", NamaMHS: "Budi", Semester: "1", Tahun: "2023", Dilihat: tDummy,
			Unit: models.Unit{
				FktKode: "F01", JrsKose: "J01", PrdKode: "P01", Fakultas: "Teknik", Jurusan: "Informatika",
			},
		},
		{
			ID: 2, NIM: "A002", NamaMHS: "Ani", Semester: "2", Tahun: "2023", Dilihat: tDummy,
			Unit: models.Unit{
				FktKode: "F02", JrsKose: "J02", PrdKode: "P02", Fakultas: "Ekonomi", Jurusan: "Akuntansi",
			},
		},
		{
			ID: 3, NIM: "A003", NamaMHS: "Caca", Semester: "1", Tahun: "2022", Dilihat: tDummy,
			Unit: models.Unit{
				FktKode: "F01", JrsKose: "J01", PrdKode: "P01", Fakultas: "Teknik", Jurusan: "Informatika",
			},
		},
	}
}

func TestGetKHSFiltered_Success(t *testing.T) {
	mockData := setupMockKHSData()
	mockRepo := mocks.NewKhsMockRepository(mockData, nil)
	service := services.NewKHSService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Preparing %d mock records for successful retrieval.", len(mockData))
	t.Logf("Call: GetKHSFiltered with page=%d, limit=%d.", page, limit)

	result, total, err := service.GetKHSFiltered(ctx, "", "", "", "", "", "", page, limit)

	t.Logf("Result: Received %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Should be no error on success")
	assert.Equal(t, int64(3), total, "Total data should be 3")
	assert.Len(t, result, 3, "Result length should match mock data length")
	assert.Equal(t, "Budi", result[0].NamaMHS, "Verify the first data entry is correct")
}

func TestGetKHSFiltered_EmptyResult(t *testing.T) {
	emptyData := []models.Khs{}
	mockRepo := mocks.NewKhsMockRepository(emptyData, nil)
	service := services.NewKHSService(mockRepo)
	ctx := context.Background()

	t.Logf("Setup: Preparing 0 mock records (Empty Data scenario).")

	result, total, err := service.GetKHSFiltered(ctx, "F99", "", "", "", "", "", 1, 10)

	t.Logf("Result: Received %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Should be no error when the result is empty")
	assert.Equal(t, int64(0), total, "Total data should be 0")
	assert.Empty(t, result, "The result should be an empty slice")
}

func TestGetKHSFiltered_RepositoryError(t *testing.T) {
	expectedErr := errors.New("mongodb connection lost")

	mockRepo := mocks.NewKhsMockRepository(nil, expectedErr)
	service := services.NewKHSService(mockRepo)
	ctx := context.Background()

	t.Logf("Setup: Mocking repository to return error: '%v'.", expectedErr)

	result, total, err := service.GetKHSFiltered(ctx, "", "", "", "", "", "", 1, 10)

	t.Logf("Result: Service returned error: '%v'.", err)

	assert.Error(t, err, "Should return an error")
	assert.Nil(t, result, "Result should be nil if an error occurs")
	assert.Equal(t, int64(0), total, "Total count should be 0")
	assert.Equal(t, expectedErr, err, "Error should be the same as the mock error")
}

func TestGetKHSFiltered_WithFilterParams(t *testing.T) {
	mockData := setupMockKHSData()
	mockRepo := mocks.NewKhsMockRepository(mockData, nil)
	service := services.NewKHSService(mockRepo)
	ctx := context.Background()

	filterFakultas := "F01"
	filterTahun := "2023"
	page := 2
	limit := 1

	t.Logf("Setup: Testing filter pass-through with Fkt=%s, Tahun=%s, Page=%d, Limit=%d.", filterFakultas, filterTahun, page, limit)

	result, total, err := service.GetKHSFiltered(
		ctx,
		filterFakultas,
		"",
		"",
		filterTahun,
		"1",
		"Budi",
		page,
		limit,
	)

	t.Logf("Result: Total records returned by mock (ignoring filters): %d.", total)

	assert.NoError(t, err)

	assert.Equal(t, int64(3), total)

	assert.Len(t, result, 3)
}
