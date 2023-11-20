package main

import (
	"github.com/tidwall/gjson"
)

type dateVal struct {
	date  string
	price string
}

type qMap map[string][]dateVal

func parse(data string) qMap {
	all := gjson.Get(data, "@this").Map()
	res := qMap{}

	for tick, _ := range all {
		dates := gjson.Get(string(data), tick+".values.#.datetime").Array()
		closes := gjson.Get(string(data), tick+".values.#.close").Array()
		series := []dateVal{}

		for i, date := range dates {
			series = append(series, dateVal{date: date.String(), price: closes[i].String()})
		}

		res[tick] = series
	}

	return res
}
