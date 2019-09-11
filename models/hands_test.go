package models

var Cards_FourOfAKind = Cards{
	&Card{
		Value: One,
		Suit:  Spade,
	},
	&Card{
		Value: One,
		Suit:  Heart,
	},
	&Card{
		Value: One,
		Suit:  Diamond,
	},
	&Card{
		Value: One,
		Suit:  Club,
	},
	&Card{
		Value: Three,
		Suit:  Club,
	},
}

var Cards_TwoPair = Cards{
	&Card{
		Value: One,
		Suit:  Spade,
	},
	&Card{
		Value: One,
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
		Value: One,
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
