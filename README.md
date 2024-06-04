
# json-response-helper

`json-response-helper` is a simple Go package designed to facilitate the creation and handling of JSON responses in HTTP servers. It provides a function to standardize the structure and content of JSON responses, making it easier to ensure consistency across your HTTP endpoints.


## Features

- Constructs JSON responses with a consistent structure.
- Supports dynamic status codes.
- Automatically sets the appropriate HTTP headers.
- Encodes and writes JSON responses to the `http.ResponseWriter`.


## Installation

To install `json-response-helper`, use `go get`:

```bash
  go get github.com/bubaigcect/json-response-helper
```
    
## Usage/Examples

To use `json-response-helper` in your project, import the package and call `Response` from your HTTP handler functions.

### Example

```go
package main

import (
	"net/http"
	parsejson "github.com/bubaigcect/json-response-helper"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"example": "data"}
	parsejson.Response(w, "Request successful", data, http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

```


## Response Structure

The JSON response structure created by `Response` is defined by the `Response` struct from the `json-response-helper/common` package.

### Response Fields

- `Status` (`bool`): Indicates the success or failure of the request.
- `Code` (`int`): The HTTP status code.
- `Message` (`string`): A descriptive message about the response.
- `Data` (`interface{}`): The response data.
- `ResponseCode` (`int`, optional): A custom response code if provided.

### Example Response

```json
{
    "status": true,
    "code": 200,
    "message": "Request successful",
    "data": {
        "example": "data"
    }
}
```


## License

[MIT](https://choosealicense.com/licenses/mit/)


## 🔗 Links

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/bubaigcect/)

## 🚀 About Me

I'm a seasoned full-stack developer with 9+ years of experience in building scalable web applications. My expertise includes frontend development with React and backend development using Golang, Node.js, Python, and PHP. I am dedicated to delivering top-notch solutions that fulfill client requirements and support business goals, leveraging my diverse skill set and unwavering commitment to quality.
## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## Contact

For any questions or feedback, please reach out to us at bubaisaha.developer@gmail.com
