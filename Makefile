pb:
	cd pb && protoc -I=. \
	-I${GOPATH}/src \
	--gogofaster_out=:. \
	receipt.proto

clean:
	rm pb/*.pb.go
	rm pb/*.json
	rm pb/*.gw.go

.PHONY: pb
