module github.com/tradingAI/go

go 1.13

require (
	github.com/garyburd/redigo v1.6.0
	github.com/go-ini/ini v1.55.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.3.0
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/minio/minio-go/v6 v6.0.50
	github.com/stretchr/testify v1.5.1
	github.com/tradingAI/proto/gen/go/common v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/tradingAI/proto/gen/go/tweb => ../proto/gen/go/tweb

replace github.com/tradingAI/proto/gen/go/common => ../proto/gen/go/common

replace github.com/tradingAI/proto/gen/go/model => ../proto/gen/go/model

replace github.com/tradingAI/proto/gen/go/scheduler => ../proto/gen/go/scheduler
