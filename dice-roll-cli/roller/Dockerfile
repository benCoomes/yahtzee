FROM golang:1.18

WORKDIR /usr/src/app

# pre-copy go.mod to these layers only need to be rebuilt if dependencies change.
COPY go.mod ./
RUN go mod download && go mod verify

# copy the rest of the app
COPY . .
RUN go build -v -o /usr/local/bin/roller ./...

ENTRYPOINT ["roller"]

# roll 6 dice by default
CMD ["x", "x", "x", "x", "x", "x"]