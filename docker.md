-- DOCKER -- 

Dockerfile is where you set and create your environment variables (the conditions needed to run the gym app)

net/http is usually the standard to call apis

need a server to be running the same version as go 

have an exposed api endpoint that sends back a response, the app is exposed by using the port environment variables


docker build -t gym (-t means tagging the image) <- this builds the docker image
docker run -t 8080:8080 -tid gym

Create a docker-compose.yml file with any service you'd like
Run docker-compose up -d
Stop the newly created service
Remove the service from the docker-compose.yml file
Run docker-compose down --remove-orphans

path packages/postgres

docker compose up 