# Solving Linear Equations of two variables in GoLang

A RESTful service to solve linear equations with two unknowns, written in Go.

### Installation

``` cmd
go get github.com/abhilashr1/linearequations 
go get github.com/gorilla/mux
cd %gopath%/src/github.com/abhilashr1/linearequations
%gopath%/bin/linearequations
[For Windows] %gopath%/bin/linearequations.exe
```
    
### Usage

To use the application locally, run the linearequations executable, which gets the HTTP Application Server running. Next, send a POST call as below: 

``` bash
curl --request POST \
--url http://localhost/api/linear \
--header 'content-type: application/json' \
--data '{		
    "First" : "x-y=5",
    "Second": "X+Y=3"
}'  
```

For instructions, send a GET call to the help endpoint

``` bash
curl --request GET \
--url https://golinearequations.azurewebsites.net/api/help

```

To query in other languages

#### C#

``` cs 
    var client = new RestClient("http://localhost/api/linear");
    var request = new RestRequest(Method.POST);
    request.AddHeader("content-type", "application/json");
    request.AddParameter("application/json", "{\"First\" : \"x-y=5\", \"Second\": \"X+Y=3\" }", ParameterType.RequestBody);
    IRestResponse response = client.Execute(request);
```

#### Go

``` go
package main

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
)

func main() {
    url := "http://localhost/api/linear"
    payload := strings.NewReader("{\"First\" : \"x-y=5\",\"Second\": \"X+Y=3\"}")
    req, _ := http.NewRequest("POST", url, payload)
    req.Header.Add("content-type", "application/json")
    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(res)
    fmt.Println(string(body))
}
```

## Output

``` bash
    $ curl --request POST --url http://localhost/api/linear --header 'content-type: application/json' --data '{ "First" : "
    x-y=5", "Second" : "x+y=3" }'
    {"x":"4.000000","y":"-1.000000"}
```
