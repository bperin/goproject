package transactions

import (
	"reflect"
	"testing"
)

func TestFilterTransactions(t *testing.T) {
	transactions := GetSampleTransactions()

	testCases := []struct {
		name     string
		filter   Filter
		expected []Transaction
	}{
		{
			name:   "no filter",
			filter: Filter{},
			expected: []Transaction{
				{Id: 1, UserID: 1, Currency: 1, Amount: 100, Timestamp: 1},
				{Id: 2, UserID: 1, Currency: 2, Amount: 200, Timestamp: 2},
				{Id: 3, UserID: 2, Currency: 1, Amount: 150, Timestamp: 3},
				{Id: 4, UserID: 1, Currency: 1, Amount: 300, Timestamp: 4},
				{Id: 5, UserID: 2, Currency: 2, Amount: 250, Timestamp: 5},
			},
		},
		{
			name: "filter by user ID",
			filter: Filter{
				UserID: intPtr(1),
			},
			expected: []Transaction{
				{Id: 1, UserID: 1, Currency: 1, Amount: 100, Timestamp: 1},
				{Id: 2, UserID: 1, Currency: 2, Amount: 200, Timestamp: 2},
				{Id: 4, UserID: 1, Currency: 1, Amount: 300, Timestamp: 4},
			},
		},
		{
			name: "filter by time range",
			filter: Filter{
				MinTime: intPtr(2),
				MaxTime: intPtr(4),
			},
			expected: []Transaction{
				{Id: 2, UserID: 1, Currency: 2, Amount: 200, Timestamp: 2},
				{Id: 3, UserID: 2, Currency: 1, Amount: 150, Timestamp: 3},
				{Id: 4, UserID: 1, Currency: 1, Amount: 300, Timestamp: 4},
			},
		},
		{
			name: "filter by ID",
			filter: Filter{
				Id: intPtr(3),
			},
			expected: []Transaction{
				{Id: 3, UserID: 2, Currency: 1, Amount: 150, Timestamp: 3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FilterTransactions(transactions, tc.filter)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// Helper function to create int pointers
func intPtr(i int) *int {
	return &i
}
