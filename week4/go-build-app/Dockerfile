FROM golang:1.18.3-alpine3.16

# WORKDIR /the/workdir/path
WORKDIR /app

# COPY -> Ketika ada file yang tercompress akan utuh ditujuannya
# ADD  -> Ketika ada file yang tercompress akan diektraksi
COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]
