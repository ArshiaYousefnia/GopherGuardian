# GopherGuardian

Gopher Guardian is a simple monitoring tool written in go

HTTP and TCP protocols are supported. Each endpoint is checked in a cycle and downtimes and delays are detected.

On detection of a down endpoint, alert is sent. Normal delay metrcis are also logged to provide easier monitoring.

Currently Email and Telegram Bot is supported for Alerting.

## Install and Build
```shell
git clone https://github.com/ArshiaYousefnia/GopherGuardian.git
cd GopherGuardian
go mod tidy
go build -o out ./cmd/gopherguardian
```

## Usage
You should prepare a json config. An example config is provided in repo. A simple config is like:
```json
{
  "targets": [
    {
      "name": "http test",
      "type": "http",
      "address": "https://www.google.com",
      "interval": 1,
      "alert": {
        "email": "tester@gmail.com",
        "telegram": "3213232"
      }
    },
    {
      "name": "tcp test",
      "type": "tcp",
      "address": "example.com:80",
      "interval": 1,
      "alert": {
        "email": "tester@gmail.com",
        "telegram": "3213232"
      }
    }
  ]
}
```
Then run the monitoring server via:
```shell
./out --config <config-path> --port <port> --verbose <verbose>
```

Config is defaulted to `config.json`, port to `8080`, and verbose to `false`.
