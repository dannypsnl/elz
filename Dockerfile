FROM dannypsnl/elz-test-env

RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/mattn/goveralls
RUN go get github.com/modocache/gover

WORKDIR /go/src/github.com/elz-lang/elz

RUN rm -r *
COPY . .

ENV PATH=$PATH:$GOPATH/bin/
ENV COVER_TOKEN=q1SjUrwywcLURT4eok5Ki37gXZeGZiwg4

CMD go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' ./... | xargs sh -c && \
    go build && \
    gover && \
    goveralls -coverprofile=gover.coverprofile -service=travis-ci -repotoken $COVER_TOKEN
