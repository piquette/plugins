# Plugins
A repo for [marketstore][mkts] plugins

## Installation
The `plugin` option for the `buildmode` flag when executing `go install` is acting up on macOS + go version 1.10. To install all these plugins, clone this repo and run `make`.

## Bars
Run `make bars` to install. This plugin backfills historical OHLCV bars for specified symbols and then periodically pushes new bars. This is accomplished through the [finance-go][finance] `history` and `equity` packages.

### Configuration
Options for configuring the `bars` plugin are as follows-

Name | Type | Default | Description
--- | --- | --- | ---
symbols | slice of strings | ["AAPL"] | The symbols to retrieve chart bars for
start | string | 01-01-2018 | The point at which to start aggregating bars from
end | string | current time | The point at which to stop aggregating bars
interval | string | 1d | The aggregation interval for each bar

Note that start/end must be formatted as mm-dd-yyyy. The options for interval can be found in the history package [here][docs]. An example `mkts.yml` file can be found [here][bars]



[finance]: https://github.com/piquette/finance-go
[bars]: https://github.com/piquette/plugins/tree/master/bars
[docs]: https://github.com/piquette/finance-go/blob/master/history/time.go#L8
[mkts]: https://github.com/alpacahq/marketstore
