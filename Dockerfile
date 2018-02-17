FROM dannypsnl/elz-test-env

WORKDIR /go/src/github.com/elz-lang/elz

RUN rm -r *
COPY . .

CMD go test -v ./... && go build
