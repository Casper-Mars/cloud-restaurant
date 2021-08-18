module github.com/Casper-Mars/cloud-restaurant/interface

go 1.16

require (
	github.com/Casper-Mars/cloud-restaurant v0.0.0-20210818073626-563fc6e4cc25
	github.com/go-kratos/kratos/v2 v2.0.5
	github.com/google/wire v0.5.0
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/Casper-Mars/cloud-restaurant  => ../../
)