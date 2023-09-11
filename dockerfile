FROM golang:1.21-rc-alpine as builder
WORKDIR /opt/whoknowkmh-portfolio/
# add gcc for CGO
RUN apk add build-base
COPY . .
# build the main whoknowkmh-portfolio binary
RUN go build 

FROM alpine:latest
WORKDIR /opt
# copy over the binaries we built in the last step
COPY --from=builder /opt/whoknowkmh-portfolio/whoknowkmh-portfolio ./
# don't forget to copy over the SQL migration files.
COPY --from=builder /opt/whoknowkmh-portfolio/data ./data
CMD ["./whoknowkmh-portfolio"]
