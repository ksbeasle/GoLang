# Base Image
# Go App - Version go1.13.8 darwin/amd64
#FROM golang:1.12.0-alpine3.9
FROM golang:1.13.8
# Make an app directory to hold src files 
RUN mkdir /app

# Add everything in root directory into /app
ADD . /app

# Sepcifiy we want to execute commands inside our app now
WORKDIR /app

# Pull any dependencies
RUN go mod download

# Now we can run our go app
RUN go build -o main .

# Start command which executes the binary executable
CMD ["/app/main"]

# docker build -t go-docker-example .
# docker images
# docker run -p 7070:7071 -it go-docker-example