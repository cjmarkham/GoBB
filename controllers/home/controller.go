package home

type Controller struct {

}

func ProvideHomeController () (*Controller, error) {
	return &Controller{}, nil
}
