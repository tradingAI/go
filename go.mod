module github.com/tradingAI/go

go 1.13

require (
	github.com/garyburd/redigo v1.6.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.3.0
	github.com/minio/minio-go/v6 v6.0.50
	github.com/stretchr/testify v1.5.1
	github.com/tradingAI/proto/gen/go/common v0.0.0-00010101000000-000000000000
	github.com/tradingAI/tweb v0.1.24
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/tradingAI/proto/gen/go/tweb => ../proto/gen/go/tweb

replace github.com/tradingAI/proto/gen/go/common => ../proto/gen/go/common

replace github.com/tradingAI/proto/gen/go/model => ../proto/gen/go/model

replace github.com/tradingAI/proto/gen/go/scheduler => ../proto/gen/go/scheduler
