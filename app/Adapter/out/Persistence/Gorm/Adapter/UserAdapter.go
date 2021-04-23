package Adapter


type GetUsersAdapter struct {

}

func (service GetUsersAdapter) Get() string {
	return "user from adapter"
}