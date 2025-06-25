package pagination_test

import (
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/pagination"
	"github.com/stretchr/testify/require"
)

func TestParamsDefaults(t *testing.T) {
	testCases := []struct {
		name          string
		params        pagination.Params
		expectedPage  int
		expectedLimit int
	}{
		{
			name:          "ZeroValues",
			params:        pagination.Params{Page: 0, Limit: 0},
			expectedPage:  pagination.DefaultPage,
			expectedLimit: pagination.DefaultLimit,
		},
		{
			name:          "OnlyPageSet",
			params:        pagination.Params{Page: 5, Limit: 0},
			expectedPage:  5,
			expectedLimit: pagination.DefaultLimit,
		},
		{
			name:          "OnlyLimitSet",
			params:        pagination.Params{Page: 0, Limit: 25},
			expectedPage:  pagination.DefaultPage,
			expectedLimit: 25,
		},
		{
			name:          "BothSet",
			params:        pagination.Params{Page: 3, Limit: 50},
			expectedPage:  3,
			expectedLimit: 50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.params.Defaults()

			require.Equal(t, tc.expectedPage, tc.params.Page)
			require.Equal(t, tc.expectedLimit, tc.params.Limit)
		})
	}
}

func TestParamsOffset(t *testing.T) {
	testCases := []struct {
		name           string
		params         pagination.Params
		expectedOffset int
	}{
		{
			name:           "FirstPage",
			params:         pagination.Params{Page: 1, Limit: 10},
			expectedOffset: 0,
		},
		{
			name:           "SecondPage",
			params:         pagination.Params{Page: 2, Limit: 10},
			expectedOffset: 10,
		},
		{
			name:           "ThirdPageLargerLimit",
			params:         pagination.Params{Page: 3, Limit: 25},
			expectedOffset: 50,
		},
		{
			name:           "HighPageSmallLimit",
			params:         pagination.Params{Page: 10, Limit: 5},
			expectedOffset: 45,
		},
		{
			name:           "MaxLimitSecondPage",
			params:         pagination.Params{Page: 2, Limit: pagination.MaxLimit},
			expectedOffset: pagination.MaxLimit,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			offset := tc.params.Offset()
			require.Equal(t, tc.expectedOffset, offset)
		})
	}
}

func TestConstants(t *testing.T) {
	require.Equal(t, 1, pagination.DefaultPage)
	require.Equal(t, 10, pagination.DefaultLimit)
	require.Equal(t, 100, pagination.MaxLimit)
}

func TestParamsStructFields(t *testing.T) {
	params := pagination.Params{
		Page:  5,
		Limit: 20,
	}

	require.Equal(t, 5, params.Page)
	require.Equal(t, 20, params.Limit)
}

func TestParamsWithDefaults(t *testing.T) {
	params := pagination.Params{}
	params.Defaults()

	offset := params.Offset()

	require.Equal(t, pagination.DefaultPage, params.Page)
	require.Equal(t, pagination.DefaultLimit, params.Limit)
	require.Equal(t, 0, offset)
}
