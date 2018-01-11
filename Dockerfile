FROM dannypsnl/elz-test-env

COPY . /go/src/github.com/elz-lang/elz

CMD cd /go/src/github.com/elz-lang/elz && go test -v ./...
