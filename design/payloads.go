package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var QuoteQuery = Type("QuoteQuery", func() {
	Description("Used to describe the quotes to return")

	Attribute("symbols", ArrayOf(String), "Symbols for which to receive quotes", func() {
		Example([]string{"BTC", "BAT", "ETH"})
	})
	Attribute("base", String, "The base currency in which to quote price. Default is BTC", func() {
		Enum("USD", "BTC", "ETH")
	})
	Attribute("start", Int, "The start of the timeseries in Unix seconds")
	Attribute("end", Int, "The end of the timeseries in Unix seconds")
	Attribute("sort", String, "The sort order. Default is desc", func() {
		Enum("asc", "desc")
	})
	Attribute("limit", UInt, "Number of values to return. Default is unlimited")

	Required("symbols")
})
