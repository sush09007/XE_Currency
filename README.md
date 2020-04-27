# XE_Currency
The XE Currency Data API is a REST-ful (or REST-like, depending how strictly you interpret REST) web-services API.


Please ensure you have config.json file to run the project.

1. # Add file config.json
2. {
    "xe_account": {
       "xe_url": "https://xecdapi.xe.com/v1/convert_from.json/",
       "xe_account_id": "xe_account_id",
       "xe_account_key": "xe_account_key"
    },
    "postgres": {
        "host":"localhost",
        "port":"5432",
        "user":"user",  // change username
        "password":"password", //change password
        "dbname":"dbname",
        "sslmode":"disable"
    },
    "currency":[
        "AED", "CUP", "AFN"
    ]
}
3. Run using command: go run main.go



# Benchmark_InitJob-4     2758207762 ns/op
PASS
ok      XE_Currency     2.766s

# For 10 Currency Total time taken:=2.31708774s

#   For test coverage
go test -timeout 30s xe-currency/service -v -coverprofile=/tmp/vscode-goUGSjWE/go-code-cover
=== RUN   Test_Init
time="27-04-2020 18:08:43" level=info msg="Initialized config" fields.msg="initialize successfully"
--- PASS: Test_Init (0.00s)
=== RUN   TestAPIWithURL
--- PASS: TestAPIWithURL (0.00s)
=== RUN   TestAPI
--- PASS: TestAPI (0.00s)
=== RUN   TestAPI_InvalidCurrency
--- PASS: TestAPI_InvalidCurrency (0.00s)
=== RUN   TestAPI_Authorization
--- PASS: TestAPI_Authorization (0.00s)
PASS
coverage: 16.5% of statements
ok      xe-currency/service     0.009s  coverage: 16.5% of statements
