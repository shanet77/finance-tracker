package core

import (
	"net/http"
	"net/url"
)

// HTTPClient defines methods of an HTTP client.
type HTTPClient interface {
	// Do executes the HTTP request to the API server.
	Do(*http.Request) (*http.Response, error)
}

var client HTTPClient

// UseClient specifies which API server to use.
func UseClient(s HTTPClient) {
	client = s
}

type httpClient struct{}

func (*httpClient) Do(r *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(r)
}

// NewHTTPClient uses the default http.Client.
func NewHttpClient(httpClient *http.Client) (*HTTPClient, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	b, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, BaseURL: b}
	c.PriceHistory = &PriceHistoryService{client: c}
	c.Account = &AccountsService{client: c}
	c.MarketHours = &MarketHoursService{client: c}
	c.Quotes = &QuotesService{client: c}
	c.Instrument = &InstrumentService{client: c}
	c.Chains = &ChainsService{client: c}
	c.Mover = &MoverService{client: c}
	c.TransactionHistory = &TransactionHistoryService{client: c}
	c.User = &UserService{client: c}
	c.Watchlist = &WatchlistService{client: c}

	return c, nil
}
