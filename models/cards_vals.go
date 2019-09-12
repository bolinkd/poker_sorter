package models

var Cards_FourOfAKind = Cards{
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

var Cards_TwoPair = Cards{
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

var Cards_NoPairs = Cards{
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

var Cards_ValidMatchingSuits = Cards{
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

var Cards_MatchingSuitsMissingOne = Cards{
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

var Cards_ValidHighCard = Cards{
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

var Cards_InvalidHighCard = Cards{}

var Cards_SequenceValid = Cards{
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

var Cards_SequenceInvalid = Cards{
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

var Cards_FullHouse = Cards{
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

var Cards_RoyalFlush = Cards{
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

var Cards_LowStraight = Cards{
	&Card{
		Value: Ace,
		Suit:  Diamond,
	},
	&Card{
		Value: Two,
		Suit:  Spade,
	},
	&Card{
		Value: Three,
		Suit:  Heart,
	},
	&Card{
		Value: Four,
		Suit:  Heart,
	},
	&Card{
		Value: Five,
		Suit:  Heart,
	},
}

//Hand 1: 6♣, 4♥, K♠, 5♣, 10♠
var Cards_MatchingPair1 = Cards{
	&Card{
		Value: Ten,
		Suit:  Club,
	},
	&Card{
		Value: Ten,
		Suit:  Heart,
	},
	&Card{
		Value: Eight,
		Suit:  Spade,
	},
	&Card{
		Value: Eight,
		Suit:  Club,
	},
	&Card{
		Value: Three,
		Suit:  Spade,
	},
}

//Hand 2: 10♦, 3♥, K♠, 4♠, 5♦
var Cards_MatchingPair2 = Cards{
	&Card{
		Value: Ten,
		Suit:  Diamond,
	},
	&Card{
		Value: Ten,
		Suit:  Heart,
	},
	&Card{
		Value: Eight,
		Suit:  Spade,
	},
	&Card{
		Value: Six,
		Suit:  Spade,
	},
	&Card{
		Value: Eight,
		Suit:  Diamond,
	},
}
