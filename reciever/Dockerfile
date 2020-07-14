# Start with the golang image
FROM golang:latest

# Set the work directory
WORKDIR /go/src/app

#Copy all files into the image
COPY . .

#Build the app
RUN go build -o reciever .

# Expose port 8558 to the outside world
EXPOSE 8558

# Command to run the executable
CMD ["./reciever"]