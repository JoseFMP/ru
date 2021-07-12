package quotes

type QuoteMode uint8

const QuoteDoNotIgnore QuoteMode = 0
const QuoteIgnorePMS QuoteMode = 1
const QuoteIgnorePMSAndRU QuoteMode = 2

func (quote QuoteMode) AsString() string {
	return quotesAsString[quote]
}

var quotesAsString = map[QuoteMode]string{
	QuoteDoNotIgnore:    "DoNotIgnore",
	QuoteIgnorePMS:      "QuoteIgnorePMS",
	QuoteIgnorePMSAndRU: "QuoteIgnorePMSAndRU",
}
