module github.com/liubaninc/sdk-go

go 1.16

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/swaggo/gin-swagger v1.3.2
	github.com/swaggo/swag v1.7.3
	github.com/tendermint/spm v0.1.6
	github.com/tendermint/tendermint v0.34.14
	github.com/urfave/cli/v2 v2.3.0 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211015200801-69063c4bb744 // indirect
	golang.org/x/tools v0.1.7 // indirect
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
