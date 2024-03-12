# EduSnap REST API

Built with [Go](https://go.dev/) and utilizes Google Gemini Generative AI to provide educational information based on user images/screenshots.

## API Endpoints

```
- Login + Register
/api/auth 

- Account information
/api/accounts 

- Make Gemini AI model request
/api/message 
```

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/BvChung/edusnap-server.git
   ```

2. Install the required dependencies:

   ```shell
   go get -d ./...
   ```

## Running Unit Tests
```
cd to directory and run

go test . -v
```
