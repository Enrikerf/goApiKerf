

To compile on debug mode
reference: [attach debbuger](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html)

> go build -gcflags="-N -l" -o debugApp.sh

if the name is app the compiled will be named /app/goApiKerf (the name of de project)







------------------------
docker build . -t go --> will use Dockerfile

docker run -it --name go go
docker run interactivo darnombre go a_la_image_go

------

with docker-compose

docker-compose up 
docker ps to see container name 
docker exec -it trygo_go_<NUMBER> /bin/bash

// go run hello-world.go --> ejecuta el programa sin necesidad de compilar?
// go build hello-world.go --> compila el programa a la plataforma desde la que se llama
// env GOOS=darwin GOARCH=amd64 go build --> para compilar a mac ya que docker lo está compilando desde un linux y el ejecutable que crea lo pone ahí por defecto