// Package transactions provides data structures and filtering funtionality for finacial transactions.
// It includes transaction filtering with multipel criteria and sorting capabilites.
package transactions

import (
	"sort"
)

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
	Id        *int `json:"id"`
	MinTime   *int `json:"min_time"`
	MaxTime   *int `json:"max_time"`
	UserID    *int `json:"user_id"`
	Currency  *int `json:"currency"`
	MinAmount *int `json:"min_amount"`
	MaxAmount *int `json:"max_amount"`
}

// FilterTransactions applys the filter criterea to a slice of transactions and returns filterd results.
// The function filters by time range, user ID, currency limits, and transacion ID.
// Results are automaticaly sorted by transaction ID in asending order.
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
//   - Currency: Exact currency match filter (nil = no filter)
//   - Id: Exact transaction ID match (nil = no filter)
//   - MinAmount/MaxAmount: Filters by amount range (inclusive)
func FilterTransactions(transactions []Transaction, filter Filter) []Transaction {
	filtered := make([]Transaction, 0, len(transactions))

	// Aply filters
	for _, txn := range transactions {
		// Filter by time rang
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
		if filter.Currency != nil && txn.Currency != *filter.Currency {
			continue
		}

		// Filter by transaction ID
		if filter.Id != nil && txn.Id != *filter.Id {
			continue
		}

		// Filter by amount range
		if filter.MinAmount != nil && txn.Amount < *filter.MinAmount {
			continue
		}
		if filter.MaxAmount != nil && txn.Amount > *filter.MaxAmount {
			continue
		}

		filtered = append(filtered, txn)

	}

	// Sort filterd by Id using inlne comparision function
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Id < filtered[j].Id
	})

	return filtered
}

// - A slice of 15 Transaction structs with test data
func GetSampleTransactions() []Transaction {
	return []Transaction{
		{Id: 11, UserID: 1, Currency: 1, Amount: 100, Timestamp: 5},
		{Id: 7, UserID: 2, Currency: 2, Amount: -50, Timestamp: 15},
		{Id: 3, UserID: 2, Currency: 1, Amount: 150, Timestamp: 8},
		{Id: 15, UserID: 1, Currency: 3, Amount: -25, Timestamp: 2},
		{Id: 2, UserID: 1, Currency: 2, Amount: 300, Timestamp: 12},
		{Id: 9, UserID: 3, Currency: 1, Amount: -75, Timestamp: 20},
		{Id: 5, UserID: 2, Currency: 2, Amount: 200, Timestamp: 1},
		{Id: 12, UserID: 3, Currency: 1, Amount: 80, Timestamp: 18},
		{Id: 1, UserID: 4, Currency: 1, Amount: 500, Timestamp: 25},
		{Id: 20, UserID: 1, Currency: 2, Amount: -100, Timestamp: 30},
		{Id: 8, UserID: 4, Currency: 3, Amount: 75, Timestamp: 10},
		{Id: 14, UserID: 2, Currency: 1, Amount: -200, Timestamp: 22},
		{Id: 6, UserID: 3, Currency: 2, Amount: 350, Timestamp: 6},
		{Id: 18, UserID: 4, Currency: 1, Amount: 25, Timestamp: 35},
		{Id: 4, UserID: 1, Currency: 3, Amount: -300, Timestamp: 40},
	}
}
