module github.com/wal-g/storages

go 1.13

require (
	cloud.google.com/go/storage v1.10.0
	github.com/Azure/azure-pipeline-go v0.2.2 // indirect
	github.com/Azure/azure-storage-blob-go v0.8.0
	github.com/Azure/go-autorest/autorest/adal v0.9.2 // indirect
	github.com/aws/aws-sdk-go v1.34.3
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mattn/go-ieproxy v0.0.0-20191113090002-7c0f6868bffe // indirect
	github.com/ncw/swift v1.0.49
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.11.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/wal-g/tracelog v0.0.0-20190824100002-0ab2b054ff30
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	google.golang.org/api v0.32.0
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

// TODO: remove after merge https://github.com/googleapis/google-api-go-client/pull/686
replace google.golang.org/api => github.com/agneum/google-api-go-client v0.32.1-0.20201002085333-2eedfd9dda7e
