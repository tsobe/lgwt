package arrays

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	return Reduce(transactions, 0.0, func(balance float64, transaction Transaction) float64 {
		if transaction.To == name {
			balance += transaction.Sum
		}
		if transaction.From == name {
			balance -= transaction.Sum
		}
		return balance
	})
}
