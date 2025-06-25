package response_test

import (
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/response"
	"github.com/stretchr/testify/require"
)

func TestPaginated(t *testing.T) {
	testCases := []struct {
		name               string
		items              interface{}
		totalItems         int
		page               int
		limit              int
		expectedSuccess    bool
		expectedTotalPages int
	}{
		{
			name:               "BasicPagination",
			items:              []string{"item1", "item2", "item3"},
			totalItems:         25,
			page:               1,
			limit:              10,
			expectedSuccess:    true,
			expectedTotalPages: 3,
		},
		{
			name:               "LastPage",
			items:              []string{"item1", "item2"},
			totalItems:         22,
			page:               3,
			limit:              10,
			expectedSuccess:    true,
			expectedTotalPages: 3,
		},
		{
			name:               "SinglePage",
			items:              []string{"item1"},
			totalItems:         1,
			page:               1,
			limit:              10,
			expectedSuccess:    true,
			expectedTotalPages: 1,
		},
		{
			name:               "EmptyResult",
			items:              []string{},
			totalItems:         0,
			page:               1,
			limit:              10,
			expectedSuccess:    true,
			expectedTotalPages: 0,
		},
		{
			name:               "ExactDivision",
			items:              []string{"item1", "item2", "item3", "item4", "item5"},
			totalItems:         50,
			page:               2,
			limit:              5,
			expectedSuccess:    true,
			expectedTotalPages: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := response.NewPaginated(tc.items, tc.totalItems, tc.page, tc.limit)

			require.Equal(t, tc.expectedSuccess, result.Success)
			require.Equal(t, tc.items, result.Data)
			require.NotNil(t, result.Pagination)
			require.Equal(t, tc.page, result.Pagination.Page)
			require.Equal(t, tc.limit, result.Pagination.Limit)
			require.Equal(t, tc.totalItems, result.Pagination.TotalItems)
			require.Equal(t, tc.expectedTotalPages, result.Pagination.TotalPages)
		})
	}
}

func TestSuccess(t *testing.T) {
	testCases := []struct {
		name string
		data interface{}
	}{
		{
			name: "WithStringData",
			data: "success message",
		},
		{
			name: "WithMapData",
			data: map[string]interface{}{"key": "value", "count": 42},
		},
		{
			name: "WithSliceData",
			data: []int{1, 2, 3, 4, 5},
		},
		{
			name: "WithNilData",
			data: nil,
		},
		{
			name: "WithStructData",
			data: struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			}{ID: 1, Name: "test"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := response.NewSuccess(tc.data)

			require.True(t, result.Success)
			require.Equal(t, tc.data, result.Data)
		})
	}
}

func TestError(t *testing.T) {
	testCases := []struct {
		name    string
		message string
	}{
		{
			name:    "SimpleErrorMessage",
			message: "Something went wrong",
		},
		{
			name:    "DetailedErrorMessage",
			message: "User not found with ID: 123",
		},
		{
			name:    "EmptyErrorMessage",
			message: "",
		},
		{
			name:    "ValidationErrorMessage",
			message: "Invalid input: email is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := response.NewError(tc.message)

			require.False(t, result.Success)
			require.Equal(t, tc.message, result.Message)
		})
	}
}

func TestPaginationStruct(t *testing.T) {
	pagination := response.Pagination{
		Page:       2,
		Limit:      15,
		TotalItems: 100,
		TotalPages: 7,
	}

	require.Equal(t, 2, pagination.Page)
	require.Equal(t, 15, pagination.Limit)
	require.Equal(t, 100, pagination.TotalItems)
	require.Equal(t, 7, pagination.TotalPages)
}
