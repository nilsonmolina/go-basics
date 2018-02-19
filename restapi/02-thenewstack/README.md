# Website
[thenewstack.io](https://thenewstack.io/make-a-restful-json-api-go/)

***

In this post, we will not only cover how to use Go to create a RESTful JSON API, but we will also talk about good RESTful design. If you have ever consumed an API in the past that doesn’t follow good design, then you end up writing bad code to consume a bad API. Hopefully, after this article you will have a better idea of what a well behaved API should look like.

# What is a JSON API?
Before JSON, there was XML. Having used XML and JSON both, there is no question that JSON is the clear winner. I’m not going to cover in depth the concept of a JSON API, as it is detailed quite well on jsonapi.org.

# A Basic Web Server
A RESTful service starts with fundamentally being a web service first. Here is a really basic web server that responds to any requests by simply outputting the request url:
```
package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
```
Running this example will spin up a server on port 8080, and can be accessed at http://localhost:8080


# Things We Didn’t Do
While we are off to a great start, there is a lot left to do. Things we haven’t addressed are:

Version Control – What if we need to modify the API and that results in a breaking change? Might we add /v1/prefix to all our routes to start with?
Authentication – Unless this is a free/public API, we probably need some authentication. I suggest learning about JSON web tokens
eTags – If you are building something that needs to scale, you will likely need to implement eTags

# What Else is Left?
As with all projects, things start off small but can quickly spiral out of control. If I was going to take this to the next level and make it production ready, these are just some of the additional things to do:

Lots of refactoring!
Create packages for several of these files, such as some JSON helpers, decorators, handlers, and more.
Testing… oh yes, you can’t forget that. We didn’t do ANY testing here. For a production system, this is a must.
