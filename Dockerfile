# Dockerfile
FROM golang:1.7.5
RUN go get github.com/jinzhu/gorm && go get github.com/gorilla/mux && go get github.com/mattn/go-sqlite3
WORKDIR /server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch  
EXPOSE 3000
WORKDIR /root/
COPY --from=0 /server/app .
COPY --from=0 /server/data.sqlite3 .
CMD ["./app"] 