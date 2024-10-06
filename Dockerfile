FROM --platform=amd64 golang:1.23.2-bullseye

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o /mimir

CMD [ "sh", "-c", "/mimir ${APP}" ]
