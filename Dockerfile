FROM dannypsnl/elz-test-env

ENV GO_FILES='$(find . -iname '*.go' -type f | grep -v /vendor/)'
# All the .go files, excluding vendor/

RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/mattn/goveralls

WORKDIR /go/src/github.com/elz-lang/elz

RUN rm -r *
COPY . .

CMD go test -v ./... && go build && .scripts/coverage --coveralls
