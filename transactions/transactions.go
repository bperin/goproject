// Package transactions provides data structures and filtering functionality for financial transactions.
// It includes transaction filtering with multiple criteria and sorting capabilities.
package transactions

import "sort"

// Transaction represents a single financial transaction with metadata.
type Transaction struct {
	Id        int `json:"id"`
	UserID    int `json:"user_id"`
	Currency  int `json:"currency"`
	Amount    int `json:"amount"`
	Timestamp int `json:"timestamp"`
}

// Filter represents the filtering criteria for transactions.
// Uses pointer fields (*int) for optional filters - nil means no filter applied.
// !!!IMPORTANT I HAD THE NULL POINTER ON THE OUTSIDE IT WOULD NOT HAVE RAN!
type Filter struct {
	Id       *int `json:"id"`
	MinTime  *int `json:"min_time"`
	MaxTime  *int `json:"max_time"`
	UserID   *int `json:"user_id"`
	Currency int  `json:"currency"`
}

// FilterTransactions applies the filter criteria to a slice of transactions and returns filtered results.
// The function filters by time range, user ID, currency limits, and transaction ID.
// Results are automatically sorted by transaction ID in ascending order.
//
// Parameters:
//   - transactions: Slice of Transaction structs to filter
//   - filter: Filter struct containing optional filtering criteria
//
// Returns:
//   - A new slice containing only transactions that match all specified filter criteria
//
// Filtering rules:
//   - MinTime/MaxTime: Filters by timestamp range (inclusive)
//   - UserID: Exact match filter (nil = no filter)
//   - Currency: Excludes transactions with currency >= 1000
//   - Id: Exact transaction ID match (nil = no filter)
func FilterTransactions(transactions []Transaction, filter Filter) []Transaction {
	filtered := make([]Transaction, 0, len(transactions))

	// Apply filters
	for _, txn := range transactions {
		// Filter by time range
		if filter.MinTime != nil && txn.Timestamp < *filter.MinTime {
			continue
		}
		if filter.MaxTime != nil && txn.Timestamp > *filter.MaxTime {
			continue
		}

		// Filter by user ID
		if filter.UserID != nil && txn.UserID != *filter.UserID {
			continue
		}

		// Filter by currency
		if txn.Currency >= 1000 {
			continue
		}

		// Filter by transaction ID
		if filter.Id != nil && txn.Id != *filter.Id {
			continue
		}

		filtered = append(filtered, txn)

	}

	// Sort filtered by Id using inline comparison function
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Id < filtered[j].Id
	})

	return filtered
}

// GetSampleTransactions returns a predefined set of sample transaction data for testing and demonstration.
// The sample data includes 5 transactions with varying user IDs, currencies, amounts, and timestamps.
//
// Returns:
//   - A slice of 5 Transaction structs with test data
func GetSampleTransactions() []Transaction {
	return []Transaction{
		{Id: 1, UserID: 1, Currency: 1, Amount: 100, Timestamp: 1},
		{Id: 2, UserID: 1, Currency: 2, Amount: 200, Timestamp: 2},
		{Id: 3, UserID: 2, Currency: 1, Amount: 150, Timestamp: 3},
		{Id: 4, UserID: 1, Currency: 1, Amount: 300, Timestamp: 4},
		{Id: 5, UserID: 2, Currency: 2, Amount: 250, Timestamp: 5},
	}
}
