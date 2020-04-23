protoc-gen-gocosmos:
	@echo "Installing protoc-gen-gocosmos..."
	@go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos

proto-gen:
	./protoc-gen.sh

