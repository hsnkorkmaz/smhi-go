# smhi-go

## Documentation

Full API documentation can be found at [SMHI](https://opendata.smhi.se/apidocs/metobs/index.html).

## Getting Started

```bash
go get github.com/hsnkorkmaz/smhi-go
```

```go
import smhi "github.com/hsnkorkmaz/smhi-go/services"
```

# Endpoints

### Category
https://opendata-download-metobs.smhi.se/api.json
```go
response, err := smhi.GetCategoryResult()
```

### Version
https://opendata-download-metobs.smhi.se/api/version/1.0.json
```go
response, err := smhi.GetVersionResult("1.0")
```

### Parameter
https://opendata-download-metobs.smhi.se/api/version/1.0/parameter/1.json
```go
response, err := smhi.GetParameterResult("1.0", "1")
```
### Station
https://opendata-download-metobs.smhi.se/api/version/1.0/parameter/1/station/159880.json
```go
response, err := smhi.GetStationResult("1.0", "1", "159880")
```
### Station Set
https://opendata-download-metobs.smhi.se/api/version/1.0/parameter/1/station-set/all.json
```go
response, err := smhi.GetStationSetResult("1.0", "1", "all")
```

### Period - Station
https://opendata-download-metobs.smhi.se/api/version/1.0/parameter/1/station/159880/period/latest-months.json
```go
response, err := smhi.GetPeriodStationResult("1.0", "1", "159880", "latest-months")
```

### Period - Station Set
https://opendata-download-metobs.smhi.se/api/version/1.0/parameter/1/station-set/all/period/latest-hour.json
```go
response, err := smhi.GetPeriodStationSetResult("1.0", "1", "all", "latest-hour")
```


### Data - Station
https://opendata-download-metobs.smhi.se/api/version/1.0/parameter/1/station/159880/period/latest-months/data.json
```go
response, err := smhi.GetDataStationResult("1.0", "1", "159880", "latest-months")
```

### Data - Station Set
https://opendata-download-metobs.smhi.se/api/version/1.0/parameter/1/station-set/all/period/latest-hour/data.json
```go
response, err := smhi.GetDataStationSetResult("1.0", "1", "all", "latest-hour")
```

