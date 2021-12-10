FROM golang:1.17.3-alpine3.14

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /docker-kraicklist

EXPOSE 3001

CMD [ "/docker-kraicklist" ]