language: go

go:
- 1.x
- "1.10"
- 1.11.x
- master

script: 
  - go test -v ./...
