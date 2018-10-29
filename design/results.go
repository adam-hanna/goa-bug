package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var Quote = ResultType("application/vnd.prices.quote", func() {
	Description("A Quote provides pricing information for a crypto currency")
	TypeName("Quote")

	Attributes(func() {
		Attribute("id", String)
		Attribute("from", String)
		Attribute("to", String)
		Attribute("rate", Float64)
		Attribute("exchange", String)
		Attribute("timestamp", Int64)
		Attribute("vol", UInt64)
	})

	View("default", func() {
		Attribute("id")
		Attribute("from")
		Attribute("to")
		Attribute("rate")
		Attribute("exchange")
		Attribute("timestamp")
		Attribute("vol")
	})

	View("tiny", func() {
		Attribute("from")
		Attribute("to")
		Attribute("rate")
		Attribute("timestamp")
	})

	Required("from", "to", "rate", "timestamp")
})
