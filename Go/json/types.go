package json

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      string  `json:"id"`
	Profile Profile `json:"profile,omitempty"`
	Karma   int     `json:"karma,omitempty"`
}

type Profile struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type TrafficShapingPolicy struct {
	// If traffic shaping policy enabled for this backend, indicating whether it's a default backend or not
	Enabled bool `json:"enabled"`
	// Policy of traffic shaping
	Policy string `json:"policy"`
	// Order is the evaluation order when matching traffic policy
	Order int `json:"order"`
	// Weight (0-100) of traffic to redirect to the backend.
	// e.g. Weight 20 means 20% of traffic will be redirected to the backend and 80% will remain
	// with the other backend. 0 weight will not send any traffic to this backend
	Weight int `json:"weight"`
	// Match conditions
	Match [][]Match `json:"match"`
}

type Match struct {
	Header Header `json:"header"`
	Cookie Cookie `json:"cookie"`
}

type Header struct {
	// Header on which to redirect requests to this backend
	Header string `json:"header"`
	// HeaderValue on which to redirect requests to this backend
	HeaderValue string `json:"headerValue"`
}

type Cookie struct {
	// Cookie on which to redirect requests to this backend
	Cookie string `json:"cookie"`
	// CookieValue on which to redirect requests to this backend
	CookieValue string `json:"cookieValue"`
}

type sslConfiguration struct {
	Certificates map[string]string `json:"certificates"`
}

// configureCertificates JSON encodes certificates and POSTs it to an internal HTTP endpoint
// that is handled by Lua
func configureCertificates() string {
	configuration := &sslConfiguration{
		Certificates: map[string]string{},
	}
	bytes, err := json.Marshal(configuration)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

// test json omitempty
// func main() {
// user := User{
// 	ID: "test",
// }
// bytes, err := json.Marshal(user)
// if err != nil {
// 	fmt.Println(err)
// }

// t := TrafficShapingPolicy{
// 	Match: [][]Match{
// 		{
// 			{
// 				Header: Header{
// 					Header:      "mellon",
// 					HeaderValue: "coolie",
// 				},
// 			},
// 			{
// 				Cookie: Cookie{
// 					Cookie:      "hello",
// 					CookieValue: "world",
// 				},
// 			},
// 			{
// 				Header: Header{
// 					Header:      "music",
// 					HeaderValue: "post-punk",
// 				},
// 			},
// 		},
// 		{
// 			{
// 				Header: Header{
// 					Header:      "x-version",
// 					HeaderValue: "feature",
// 				},
// 			},
// 			{
// 				Cookie: Cookie{
// 					Cookie:      "y-cookie",
// 					CookieValue: "canary",
// 				},
// 			},
// 		},
// 	},
// }

// 	t1 := [][]Match{
// 		{
// 			{
// 				Header: Header{
// 					Header:      "mellon",
// 					HeaderValue: "coolie",
// 				},
// 			},
// 			{
// 				Cookie: Cookie{
// 					Cookie:      "hello",
// 					CookieValue: "world",
// 				},
// 			},
// 			{
// 				Header: Header{
// 					Header:      "music",
// 					HeaderValue: "post-punk",
// 				},
// 			},
// 		},
// 		{
// 			{
// 				Header: Header{
// 					Header:      "x-version",
// 					HeaderValue: "feature",
// 				},
// 			},
// 			{
// 				Cookie: Cookie{
// 					Cookie:      "y-cookie",
// 					CookieValue: "canary",
// 				},
// 			},
// 		},
// 	}

// 	bytes, err := json.Marshal(t1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(bytes))
// 	t2 := [][]Match{}
// 	// t2 := TrafficShapingPolicy{}
// 	if err := json.Unmarshal(bytes, &t2); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%#v\n", t2)
// }
