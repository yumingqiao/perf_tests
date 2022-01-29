FROM golang:1.16 as builder

WORKDIR /workspace
COPY configs configs
COPY write_small_files.go write_small_files.go
COPY read_small_files.go read_small_files.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=off go build -o write_small_files write_small_files.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=off go build -o read_small_files read_small_files.go


FROM ubuntu:18.04
WORKDIR /
COPY --from=builder /workspace/configs configs
COPY --from=builder /workspace/write_small_files write_small_files
COPY --from=builder /workspace/read_small_files read_small_files

CMD ["/bin/sh"]