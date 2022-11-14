# building frontend static files
FROM node:18.12.1 AS frontendBuilder

WORKDIR /Lift-Tracker

COPY web/reactjs/ .

WORKDIR /Lift-Tracker/web/reactjs
RUN npm install 
RUN npm run build

# building backend executable 
FROM golang:1.19.3 AS backendBuilder

WORKDIR /Lift-Tracker

COPY . .

RUN go build ./cmd/Lift-Tracker

CMD [ "./Lift-Tracker.exe" ]