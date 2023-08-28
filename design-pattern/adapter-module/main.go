package main

import (
	"fmt"
	"adapter-module/adapter"
)

func main() {
	fmt.Println("*** Example Adapter ***")
	client := &adapter.Client{}

	fetchAdapter := &adapter.FetchAdapter{
		Instance: &adapter.Fetch{},
	}

	client.Get(fetchAdapter, "https://www.google.com")

	axiosAdapter := &adapter.AxiosAdapter{
		Instance: &adapter.Axios{},
	}
	client.Get(axiosAdapter, "https://www.bornhup.com")

	fmt.Print("*** End of Adapter ***\n\n\n")
}