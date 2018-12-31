# SOAP Absence
SOAP Client for Solution X302S Attendance Device

## How To Run

Running API service

```
$ make run-api
```

Running Cron

```
$ make run-cron
```

Build API binary

```
$ make build-api
```

Build Cron Binary

```
$ make build-cron
```

## Config

1. Database using `mysql` and don't forget to add `parseTime=true` param on the DSN string

## Feature

### API
- Fetch All/Single User Attendance Log on specified date range

### Cron
- Synchronize User on Device to `mysql`
- Synchronize new added device
- Pull the attendance log on all registered devices


Author:
[@reyhanfahlevi](https://github.com/reyhanfahlevi)
