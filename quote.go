package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func fetchQuery(symbols string) (map[string][]dateVal, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.twelvedata.com",
		Path:   "time_series",
	}

	query := u.Query()

	query.Set("symbol", symbols)
	query.Set("interval", "1day")
	query.Set("apikey", "[[INSERT API KEY HERE from twelvedata.com]]")

	u.RawQuery = query.Encode()

	res := map[string][]dateVal{}
	resp, err := http.Get(u.String())

	if err != nil {
		return res, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return res, err
	}

	return parse(string(body)), nil
}
