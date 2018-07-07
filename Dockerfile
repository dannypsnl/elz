FROM dannypsnl/elz-test-env

RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/go-playground/overalls

RUN go get github.com/dannypsnl/assert

WORKDIR /go/src/github.com/elz-lang/elz

RUN rm -r *

ENV PATH=$PATH:$GOPATH/bin/

CMD cd cmd/elz && go build && cd ../.. && \
    overalls -project=github.com/elz-lang/elz
