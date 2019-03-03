# raygun4go-fork details

added functionality to customise the POST handler. This helps with go apps hosted on app engine, where posts are not allowed with default http.Client and http.Transport, but using urlfetch.Client

therefore a custom "poster" might be used in this case

for example:

```
	poster := RaygunPoster{Context: context}

	raygun, err := raygun4go.NewWithCustomPoster(RaygunAppName, RaygunToken, poster)
	if err != nil {
		log.Println("Unable to create Raygun client:", err.Error())
	}
	defer raygun.HandleError()

	go raygun.SendError(err)
  
```

# raygun4go
raygun4go adds Raygun-based error handling to your golang code. It catches all
occuring errors, extracts as much information as possible and sends the error
to Raygun via their REST-API.

## Getting Started

### Installation
```
  $ go get github.com/breekmd/raygun4goWithAppengine
```
### Options

The client returned by ``New`` has several chainable option-setting methods

Method                    | Description
--------------------------|------------------------------------------------------------
`Silent(bool)`            | If set to true, this prevents the handler from sending the error to Raygun, printing it instead.
`Request(*http.Request)`  | Adds the responsible http.Request to the error.
`Version(string)`         | If your program has a version, you can add it here.
`Tags([]string)`          | Adds the given tags to the error. These can be used for filtering later.
`CustomData(interface{})` | Add arbitrary custom data to you error. Will only reach Raygun if it works with `json.Marshal()`.
`User(string)`            | Add the name of the affected user to the error.

## Bugs and feature requests

Have a bug or a feature request? Please first check the list of
[issues](https://github.com/breekmd/raygun4goWithAppengine/issues).
