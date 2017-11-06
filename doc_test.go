package simular_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thingful/simular"
)

func ExampleRegisterStubRequests() {
	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			"GET",
			"http://example.com/",
			simular.NewStringResponder(200, "ok"),
		),
		simular.NewStubRequest(
			"GET",
			"http://another.com/",
			simular.NewStringResponder(200, "also ok"),
		),
	)

	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	fmt.Println(string(body))

	resp, err = http.Get("http://another.com/")
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	fmt.Println(string(body))

	if err = simular.AllStubsCalled(); err != nil {
		// handle error properly in real code
		panic(err)
	}

	// Output:
	// ok
	// also ok
}

func ExampleRegisterStubRequest_WithHeader() {
	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			"GET",
			"http://example.com/",
			simular.NewStringResponder(200, "ok"),
			simular.WithHeader(
				&http.Header{
					"Authorization": []string{"Bearer api-key"},
				},
			),
		),
	)

	_, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("Error without header")
	}

	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		// handle error properly
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer api-key")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	fmt.Println(string(body))

	if err = simular.AllStubsCalled(); err != nil {
		// handle error properly in real code
		panic(err)
	}

	// Output:
	// Error without header
	// ok
}
