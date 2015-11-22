# /bin/bash
# It require go version >= 1.5
go build -buildmode=c-archive sp.go
gcc test.c sp.a -lpthread -o test
