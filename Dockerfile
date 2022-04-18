# Specify the base image of this go app 
FROM golang:latest

# Specify that we now need to execute any commands in this directory 
RUN mkdir /build 
WORKDIR /build

RUN export GO111MODULE=on

RUN go install github.com/cindy1408/gym/src@latest

RUN cd /build && git clone https://github.com/cindy1408/gym.git

RUN cd /build/gym/src && go build 

ENV PORT=8080

EXPOSE 8080


# Copy everything from this project to a filesystem of the container 

# Obtain the package needed to run the code (alternatively use GO modules)

# Compile the binary exe for our app

# Start the application 