package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("netflixtechblog.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15"),
		colly.CacheDir("./cache"),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "pr=1; sz=1905; _ga=GA1.2.1286073688.1705508717; _gid=GA1.2.1387226610.1705508717; lightstep_guid/medium-web=93e43bb965b9b592; lightstep_session_id=ae8048b846610c6c; tz=-480; _dd_s=rum=0&expire=1705509687088; sid=1:ZhbXnnuUuCTMvE8pQ3D7pJrJyvNKFuD8q/h1WfFRFP2JBdhc+aoWOTsf6oiMEHAP; uid=lo_f0b1be58c288")
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://netflixtechblog.com")
}
