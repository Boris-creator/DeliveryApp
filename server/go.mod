module playground.com/server

go 1.24.0

require (
	github.com/fasthttp/router v1.5.4
	github.com/go-playground/validator/v10 v10.25.0
	github.com/swaggo/fasthttp-swagger v1.0.2
	github.com/valyala/fasthttp v1.59.0
	google.golang.org/grpc v1.71.0
	playground.com/geosuggest v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/time v0.10.0 // indirect
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/caarlos0/env/v11 v11.3.1
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/jackc/pgx/v5 v5.7.2
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/savsgio/gotils v0.0.0-20240704082632-aef3928b8a38 // indirect
	github.com/swaggo/files/v2 v2.0.1 // indirect
	github.com/swaggo/swag v1.16.3
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	playground.com/proto v0.0.0-00010101000000-000000000000
)

replace playground.com/proto => ../proto

replace playground.com/geosuggest => ../geosuggest
