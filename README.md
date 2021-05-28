# cryptocompare
Crypto pairs pricing REST API | Golang Challenge

## Endpoint

	/api/prices?fsyms=BTC&tsyms=GBP
  
Used to retrieve pricing of provided pairs, receiving any number of comma separated assets.

E.g.: fsyms=BTC,ETH,ADA

E.g.: tsyms=USD,GBP,CAD,DASH

## Data

The price is always retrieved from cryptocompare.com when it's available. Otherwise the data is retrieved from a MySQL DB.

A cron scheduler runs every minute refreshing the pairs in MySQL DB.

The refreshed pairs are parameters provided in the server environment.
