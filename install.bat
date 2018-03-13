go clean
set GOPATH=%cd%
cd src/app
go install -v
cd ../../
app