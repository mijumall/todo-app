FROM ubuntu:22.04
SHELL ["/bin/bash", "-c"]

ENV workspace /space

RUN apt update && apt full-upgrade -y && apt install -y less vim curl unzip git
WORKDIR $workspace

RUN curl -OL https://go.dev/dl/go1.19.linux-arm64.tar.gz
RUN tar -xz -f go1.* && rm go1.*

ENV GOPATH $workspace/gopath
ENV PATH $PATH:$workspace/go/bin:$GOPATH/bin

RUN go install golang.org/x/tools/cmd/godoc@latest && \
	go install golang.org/x/tools/gopls@latest && \
	go install github.com/cosmtrek/air@latest

WORKDIR $workspace/work

CMD ["godoc", "-http", "0.0.0.0:6060"]
