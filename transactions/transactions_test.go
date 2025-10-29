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
				{Id: 1, UserID: 4, Currency: 1, Amount: 500, Timestamp: 25},
				{Id: 2, UserID: 1, Currency: 2, Amount: 300, Timestamp: 12},
				{Id: 3, UserID: 2, Currency: 1, Amount: 150, Timestamp: 8},
				{Id: 4, UserID: 1, Currency: 3, Amount: -300, Timestamp: 40},
				{Id: 5, UserID: 2, Currency: 2, Amount: 200, Timestamp: 1},
				{Id: 6, UserID: 3, Currency: 2, Amount: 350, Timestamp: 6},
				{Id: 7, UserID: 2, Currency: 2, Amount: -50, Timestamp: 15},
				{Id: 8, UserID: 4, Currency: 3, Amount: 75, Timestamp: 10},
				{Id: 9, UserID: 3, Currency: 1, Amount: -75, Timestamp: 20},
				{Id: 11, UserID: 1, Currency: 1, Amount: 100, Timestamp: 5},
				{Id: 12, UserID: 3, Currency: 1, Amount: 80, Timestamp: 18},
				{Id: 14, UserID: 2, Currency: 1, Amount: -200, Timestamp: 22},
				{Id: 15, UserID: 1, Currency: 3, Amount: -25, Timestamp: 2},
				{Id: 18, UserID: 4, Currency: 1, Amount: 25, Timestamp: 35},
				{Id: 20, UserID: 1, Currency: 2, Amount: -100, Timestamp: 30},
			},
		},
		{
			name: "filter by user ID",
			filter: Filter{
				UserID: intPtr(4),
			},
			expected: []Transaction{
				{Id: 1, UserID: 4, Currency: 1, Amount: 500, Timestamp: 25},
				{Id: 8, UserID: 4, Currency: 3, Amount: 75, Timestamp: 10},
				{Id: 18, UserID: 4, Currency: 1, Amount: 25, Timestamp: 35},
			},
		},
		{
			name: "filter by time range",
			filter: Filter{
				MinTime: intPtr(20),
				MaxTime: intPtr(35),
			},
			expected: []Transaction{
				{Id: 1, UserID: 4, Currency: 1, Amount: 500, Timestamp: 25},
				{Id: 9, UserID: 3, Currency: 1, Amount: -75, Timestamp: 20},
				{Id: 14, UserID: 2, Currency: 1, Amount: -200, Timestamp: 22},
				{Id: 18, UserID: 4, Currency: 1, Amount: 25, Timestamp: 35},
				{Id: 20, UserID: 1, Currency: 2, Amount: -100, Timestamp: 30},
			},
		},
		{
			name: "filter by ID",
			filter: Filter{
				Id: intPtr(6),
			},
			expected: []Transaction{
				{Id: 6, UserID: 3, Currency: 2, Amount: 350, Timestamp: 6},
			},
		},
		{
			name: "filter by user and time range combined",
			filter: Filter{
				UserID:  intPtr(1),
				MinTime: intPtr(10),
				MaxTime: intPtr(35),
			},
			expected: []Transaction{
				{Id: 2, UserID: 1, Currency: 2, Amount: 300, Timestamp: 12},
				{Id: 20, UserID: 1, Currency: 2, Amount: -100, Timestamp: 30},
			},
		},
		{
			name: "filter by currency",
			filter: Filter{
				Currency: intPtr(2),
			},
			expected: []Transaction{
				{Id: 2, UserID: 1, Currency: 2, Amount: 300, Timestamp: 12},
				{Id: 5, UserID: 2, Currency: 2, Amount: 200, Timestamp: 1},
				{Id: 6, UserID: 3, Currency: 2, Amount: 350, Timestamp: 6},
				{Id: 7, UserID: 2, Currency: 2, Amount: -50, Timestamp: 15},
				{Id: 20, UserID: 1, Currency: 2, Amount: -100, Timestamp: 30},
			},
		},
		{
			name: "filter with no matches",
			filter: Filter{
				UserID: intPtr(5),
			},
			expected: []Transaction{},
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
