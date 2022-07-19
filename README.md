### crawler


### data source

https://api.metro.taipei/taipeiarenainfoboard/queuenumber.aspx

# Technical Stack

- web crawler: [https://github.com/gocolly/colly](https://github.com/gocolly/colly)
- data storage: [google sheet w/ golang api](https://developers.google.com/sheets/api/quickstart/go)

# Run

```
export GOOGLE_APPLICATION_CREDENTIAL=/Users/che-chia/Downloads/taipei-arena-iceland-63dc831f0e4c.json

 go run cmd/main.go
```
