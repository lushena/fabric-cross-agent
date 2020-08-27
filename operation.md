# operate fabric

cd /opt/gopath/src/github.com/hyperledger/fabric-cross-agent/test/integration/e2e
go test -v  github.com/hyperledger/fabric-sdk-go/test/integration/e2e -test.run TestE2E testLocal=true
