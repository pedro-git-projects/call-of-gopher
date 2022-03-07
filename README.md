# Call Of Gopher

# What is it

Call of Gopher is a fullstack web application.

Its backend is a RESTful API that recieves an input with the name, age, birthplace, residence and occupation of an investigator and responds with a complete character sheet in JSON form.  
 
The frontend is a simple Svelte application that makes a post request to the backend, fetches the response and displays it in a asynchronous manner.

All calculations follow the rules for the 7nth edition of Chaosium's Call of Cthullu RPG.

---

# Supported endpoint actions

| Method | URL Pattern         | Action                     |
|--------|---------------------|----------------------------|
| GET    | /v1/healthcheck     | Displays availability and system info      |
| POST   | /v1/investigator    | Creates a new investigator |

---

# Makefile
 
Use *make help* to list all available commands.

Use *make run* to start the backend server.

Use *make dev* to start the frontend development server.

---

# Dependencies

## julienschmidt/httprouter

When building API endpoints without thrid party libraries we face the limitation that the http.ServeMux does not allow routing to different handlers based on the request method. It doesn't provide support for clean URLs with interpolated paramters either.

Thus the choice was made to use julienschmidt's httprouter for it solves both this problems and is extremely performant because it uses a radix sort algorithm for URL matching

## rs/cors

Is a net/http handler implementing Cross Origin Resource Sharing W3 specification in Golang.

---

# Miscellaneous

## URL prefixing

As one can see in the endpoints, all of them are prefixed with v1. As in real business users often need to change endpoints functionality overtime, sometimes breaking backwards compatibility, versioning needs to be implemented. There are two common approaches for doing so:

- Prefixing URLs, like done here

- Using custom Accept and Content-Type headers on requests and responses, ie: Accept: application/vnd.go-of-cthullu-v1.

Even though custom headers are arguably "purer", I think URL prefixes come out on top regarding ease of use, thus the API was written this way.


---

# Why Go?

Go was chosen because of its great standard library, which includes almost everything one needs to build scalable web applications. Using less imports not only improves the development and maintenance process, but highly improves on the security of the application.

It is also a compiled, statically typed and concurrency first language, providing thus great error handling and performance. 

---

# Why Svelte

By abandoning the usual virtual DOM diffing techniques and adopting a compiler that generates vanilla JavaScript at build time, Svelte greatly reduces runtime overhead and improves performance. It is a match made in heaven for Go. 

---

# Todo 

The next step is to implement a frontend that is able to send the POST request to the backend as well as recuperate the JSON response and display it in a meaningful manner for the end user.
