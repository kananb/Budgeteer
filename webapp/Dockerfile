#syntax=docker/dockerfile:1

FROM node:alpine AS client
WORKDIR /client
COPY client ./
RUN npm install --production --no-audit
RUN npm run build

FROM golang:alpine AS server
WORKDIR /server
COPY server ./
RUN go mod download
RUN go build -o budgeteer.exe

FROM alpine:latest AS app
WORKDIR /app
COPY --from=client /client/build ./client/build
COPY --from=server /server/budgeteer.exe ./

ENTRYPOINT [ "./budgeteer.exe" ]
