
cmd=one
# pull push chat mchat
.PHONY:pb
pb:
	cd proto && protoc -I . --go_out=plugins=grpc:. *.proto
	cd proto && rm -rf out.bin

.PHONY:cli
cli:
	cd cli && go run main.go -cmd=${cmd}