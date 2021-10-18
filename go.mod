module github.com/liubaninc/sdk-go

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/gin-gonic/gin v1.7.4
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/tendermint/spm v0.1.6
	github.com/tendermint/tendermint v0.34.14
	google.golang.org/genproto v0.0.0-20211013025323-ce878158c4d4
	google.golang.org/grpc v1.41.0
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/postgres v1.1.2
	gorm.io/driver/sqlite v1.1.6
	gorm.io/gorm v1.21.16
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk v0.44.0 => ./cosmos-sdk

replace github.com/tendermint/spm v0.1.6 => github.com/tendermint/spm v0.1.4
