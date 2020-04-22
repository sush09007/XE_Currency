# XE_Currency
The XE Currency Data API is a REST-ful (or REST-like, depending how strictly you interpret REST) web-services API.


As of now only Insert in Db works.Work on update id going on.

Please ensure you have config.json file to run the project.

1. Add file config.json at root
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
    }
}
3. Run using command: go run XE_Currency.go
