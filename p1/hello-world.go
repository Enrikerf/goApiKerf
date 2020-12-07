package main


func main(){
	println("hello world")
}

// go run hello-world.go --> ejecuta el programa sin necesidad de compilar?
// go build hello-world.go --> compila el programa a la plataforma desde la que se llama
// env GOOS=darwin GOARCH=amd64 go build --> para compilar a mac ya que docker lo está compilando desde un linux y el ejecutable que crea lo pone ahí por defecto