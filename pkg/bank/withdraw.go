package bank

const withdrawLimit = 20_000_00

func Withdraw(card types.Card, amount types.Amount) types.Card {
	
	if amount > withdrawLimit {
		return card
	}
	
	if amount < 0 {
		return card
	}
	
	if amount > card.Balance {
		return card
	}

	card.Balance -= amount;
	return card
}