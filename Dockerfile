FROM dannypsnl/elz-test-env

RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/mattn/goveralls
RUN go get github.com/go-playground/overalls

WORKDIR /go/src/github.com/elz-lang/elz

RUN rm -r *

ENV PATH=$PATH:$GOPATH/bin/

CMD go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' ./... | xargs sh -c && \
    go build && \
    overalls -project=github.com/elz-lang/elz
