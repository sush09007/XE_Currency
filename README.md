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
    },
    "currency":[
        "AED", "CUP", "AFN","AED", "CUP", "AFN", "ETB", "ALL", "AMD", "AOA", "ARS", "AZN", "BAM", "BBD", "BDT", "BGN", "IQD", "BMD", "IRR", "BIF", "BRL", "BSD", "BTN", "BYN", "CAD", "BZD", "KPW", "JOD", "COP", "CRC", "CVE", "CZK", "DOP", "DZD", "EGP", "GBP", "GEL", "AWG", "GHS", "GIP", "GTQ", "GYD", "HKD", "HNL", "HRK", "HUF", "CUC", "ILS", "IMP", "INR", "BOB", "JEP", "JMD", "KES", "KGS", "FKP", "CHF", "ERN", "GGP", "BND", "CDF", "IDR", "CLP", "GNF", "JPY", "KMF", "SPL", "PYG", "TZS", "MRU", "KYD", "KZT", "MDL", "LKR", "LRD", "LSL", "RUB", "MGA", "SHP", "MMK", "MNT", "MOP", "MUR", "MVR", "MWK", "MXN", "MYR", "NAD", "NGN", "NIO", "NPR", "NZD", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "KWD", "RON", "RSD", "SYP", "LYD", "SAR", "SBD", "SDG", "SEK", "SGD", "TWD", "SOS", "SRD", "SZL", "TJS", "TMT", "STN", "TOP", "TRY", "TVD", "TND", "MKD", "UAH", "UGX", "UYU", "UZS", "USD", "LAK", "RWF", "KRW", "BHD", "OMR", "BWP", "XCD", "CNY", "YER", "ZAR", "ZMW", "ANG", "FJD", "GMD", "HTG", "KHR", "LBP", "MAD", "MZN", "QAR", "SCR", "SLL", "THB", "TTD", "AUD", "DKK", "NOK", "SVC", "VEF", "WST", "ZWD", "EUR", "VES", "XOF", "XPF", "DJF", "ISK", "VUV", "XAF", "VND"
    ]
}
3. Run using command: go run XE_Currency.go


To-do : writing test cases
