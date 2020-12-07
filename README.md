

docker build . -t go --> will use Dockerfile

docker run -it --name go go
docker run interactivo darnombre go a_la_image_go

------

with docker-compose

docker-compose up 
docker ps to see container name 
docker exec -it trygo_go_<NUMBER> /bin/bash