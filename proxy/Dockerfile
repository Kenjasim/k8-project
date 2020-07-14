# Start with the golang image
FROM golang:latest

# Set the work directory
WORKDIR /go/src/app

#Copy all files into the image
COPY . .

#Build the app
RUN go build -o proxy .

# Expose port 8558 to the outside world
EXPOSE 6285

# Command to run the executable
CMD ["./proxy"]