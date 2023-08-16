FROM golang:latest
WORKDIR /app
COPY . ./
RUN go build -o geocoder
EXPOSE 3333
CMD ["./geocoder"]