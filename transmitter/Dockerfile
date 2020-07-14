# Start from the python3 image
FROM python:3

# Set the work directory
WORKDIR /usr/src/app

#Instal the requests module
RUN pip install requests

#Copy all files into the image
COPY . .

# Run the transmitter
CMD ["python3", "transmitter.py"]



