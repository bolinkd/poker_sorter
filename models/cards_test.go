package models

var cards_FourOfAKind = Cards{
	&Card{
		Value: Ace,
		Suit:  Spade,
	},
	&Card{
		Value: Ace,
		Suit:  Heart,
	},
	&Card{
		Value: Ace,
		Suit:  Diamond,
	},
	&Card{
		Value: Ace,
		Suit:  Club,
	},
	&Card{
		Value: Three,
		Suit:  Club,
	},
}

var cards_TwoPair = Cards{
	&Card{
		Value: Ace,
		Suit:  Spade,
	},
	&Card{
		Value: Ace,
		Suit:  Heart,
	},
	&Card{
		Value: Two,
		Suit:  Diamond,
	},
	&Card{
		Value: Two,
		Suit:  Club,
	},
	&Card{
		Value: Three,
		Suit:  Club,
	},
}

var cards_NoPairs = Cards{
	&Card{
		Value: Ace,
		Suit:  Spade,
	},
	&Card{
		Value: Two,
		Suit:  Heart,
	},
	&Card{
		Value: Three,
		Suit:  Diamond,
	},
	&Card{
		Value: Four,
		Suit:  Club,
	},
	&Card{
		Value: Five,
		Suit:  Club,
	},
}

var cards_ValidMatchingSuits = Cards{
	&Card{
		Value: Six,
		Suit:  Heart,
	},
	&Card{
		Value: Eight,
		Suit:  Heart,
	},
	&Card{
		Value: King,
		Suit:  Heart,
	},
	&Card{
		Value: Queen,
		Suit:  Heart,
	},
	&Card{
		Value: Ace,
		Suit:  Heart,
	}}

var cards_MatchingSuitsMissingOne = Cards{
	&Card{
		Value: Six,
		Suit:  Heart,
	},
	&Card{
		Value: Eight,
		Suit:  Spade,
	},
	&Card{
		Value: King,
		Suit:  Heart,
	},
	&Card{
		Value: Queen,
		Suit:  Heart,
	},
	&Card{
		Value: Ace,
		Suit:  Heart,
	},
}

var cards_ValidHighCard = Cards{
	&Card{
		Value: Five,
		Suit:  Heart,
	},
	&Card{
		Value: King,
		Suit:  Spade,
	},
	&Card{
		Value: Six,
		Suit:  Diamond,
	},
}

var cards_InvalidHighCard = Cards{}

var cards_SequenceValid = Cards{
	&Card{
		Value: Six,
		Suit:  Heart,
	},
	&Card{
		Value: Two,
		Suit:  Spade,
	},
	&Card{
		Value: Five,
		Suit:  Heart,
	},
	&Card{
		Value: Four,
		Suit:  Diamond,
	},
	&Card{
		Value: Three,
		Suit:  Club,
	},
}

var cards_SequenceInvalid = Cards{
	&Card{
		Value: Six,
		Suit:  Heart,
	},
	&Card{
		Value: Eight,
		Suit:  Spade,
	},
	&Card{
		Value: King,
		Suit:  Heart,
	},
	&Card{
		Value: Queen,
		Suit:  Heart,
	},
	&Card{
		Value: Ace,
		Suit:  Heart,
	},
}

var cards_FullHouse = Cards{
	&Card{
		Value: Six,
		Suit:  Heart,
	},
	&Card{
		Value: Six,
		Suit:  Spade,
	},
	&Card{
		Value: Six,
		Suit:  Diamond,
	},
	&Card{
		Value: Queen,
		Suit:  Heart,
	},
	&Card{
		Value: Queen,
		Suit:  Diamond,
	},
}

var cards_RoyalFlush = Cards{
	&Card{
		Value: Ten,
		Suit:  Heart,
	},
	&Card{
		Value: Jack,
		Suit:  Heart,
	},
	&Card{
		Value: Queen,
		Suit:  Heart,
	},
	&Card{
		Value: King,
		Suit:  Heart,
	},
	&Card{
		Value: Ace,
		Suit:  Heart,
	},
}