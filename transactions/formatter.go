package transactions

import "time"

type CampaignTransactionFormat struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormattTransaction(transaction Transactions) CampaignTransactionFormat {
	formatt := CampaignTransactionFormat{}
	formatt.ID = transaction.ID
	formatt.Name = transaction.User.Name
	formatt.Amount = transaction.Amount
	formatt.CreatedAt = transaction.CreatedAt

	return formatt
}

func FormatTransactions(transactions []Transactions) []CampaignTransactionFormat {
	if len(transactions) == 0 {
		return []CampaignTransactionFormat{}
	}

	transactionsFormatter := []CampaignTransactionFormat{}

	for _, transaction := range transactions {
		formatter := FormattTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
