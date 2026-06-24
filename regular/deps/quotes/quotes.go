package quotes

var baseURL = "https://zenquotes.io/api"

type QuotesClient struct{}

func NewQuotesClient() *QuotesClient {
	return &QuotesClient{}
}

func (c *QuotesClient) GetQuote() {
}
