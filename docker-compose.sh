set -e
cd "$(dirname "$0")"
cd ..
git clone https://github.com/tradingAI/proto.git
cd proto && make proto

cd ..
cd go
go test -v ./...
