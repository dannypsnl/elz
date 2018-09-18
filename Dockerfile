FROM dannypsnl/elz-test-env

RUN go get github.com/go-playground/overalls

RUN go get github.com/dannypsnl/assert

WORKDIR /go/src/github.com/elz-lang/elz

RUN rm -r *

ENV PATH=$PATH:$GOPATH/bin/

CMD go build && overalls -project=github.com/elz-lang/elz
