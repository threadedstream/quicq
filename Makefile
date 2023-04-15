generate:
	protoc --go_out=. --go_opt=Mproto/quicq.proto=pkg/proto/quicq/v1 proto/quicq.proto