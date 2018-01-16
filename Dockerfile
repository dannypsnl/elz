FROM dannypsnl/elz-test-env

RUN cd /go/src/github.com/elz-lang/elz/ && rm -r *
COPY . /go/src/github.com/elz-lang/elz

CMD cd /go/src/github.com/elz-lang/elz && go test -v ./...
