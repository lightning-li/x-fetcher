FROM golang:1.23.1-alpine as builder

WORKDIR /app
COPY . .

RUN go mod tidy && \
    go build -o fetch_x .

FROM chromedp/headless-shell:latest
WORKDIR /app
COPY --from=builder /app/fetch_x .
Run chmod +x /app/fetch_x
COPY entrypoint.sh .
ENV PATH /app:$PATH
ENTRYPOINT ["entrypoint.sh"]
