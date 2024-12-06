test-cover:
  mkdir -p coverage
  go test ./... -coverprofile=coverage/cover.out
  go tool cover -html=coverage/cover.out
