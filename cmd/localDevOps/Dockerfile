FROM golang:1.18 as builder

RUN mkdir /app
RUN chmod 700 /app

#RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN git clone -b main https://github.com/helmutkemper/flight.git /app

WORKDIR /app

ARG CGO_ENABLED=0
#RUN go mod tidy
RUN go build -o /app/main /app/cmd/server/main.go

EXPOSE 8080

CMD ["/app/cmd/server/main"]

# git clone -b main https://github.com/helmutkemper/flight.git /app
