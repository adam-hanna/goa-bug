package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

// API describes the global properties of the API server.
var _ = API("api", func() {
	Title("Web v1 API")
	Description("The back-end API for the web v1 tool for investigating arbitrage")
	Server("localhost:8080", func() {
		Description("Development host")
	})
})

// Service describes a service
var _ = Service("prices", func() {
	Description("The prices service returns pricing information for various crypto")

	// Method describes a service method (endpoint)
	Method("quote", func() {
		// Payload describes the method payload
		Payload(QuoteQuery)

		// Result describes the method result
		Result(CollectionOf(Quote), func() {
			View("tiny")
		})

		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			GET("/prices")

			Param("symbols")
			Param("base")
			Param("start")
			Param("end")

			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})
