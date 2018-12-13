# Passport validation service
Validates passport's series and number using FMS database.

## API

### Request

`GET /?series=SERIES&number=NUMBER`

Where `SERIES` is 4  digits, `NUMBER` is 6 digits of passport.

### Responses:

#### Code 200

* `{"result":"valid"}`
* `{"result":"invalid"}`

#### Code 400, 500

`{"error":"Details"}`


## Run

### Requirements:

#### Runtime:

* curl
* bzip2
* 7GB RAM minimum

#### Development:

* curl
* bzip2
* go
* 7GB RAM minimum

It consumes double amount of RAM to be able work while updating.

### How to start:

```
go build .
./invalid-passports-go [--addr=":8002"] [--source-file="/tmp/list_of_expired_passports.csv"]
```

## Update data storage

`./update-db.sh`

Should be added to cron.
It does following:
* Download and extract archive with csv to /tmp/list_of_expired_passports.csv
* Send SIGUSR1 to process.

## License

The MIT License
