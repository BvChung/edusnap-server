# EduSnap REST API

Built with [Go](https://go.dev/) and utilizes Google Gemini Generative AI to provide educational information based on user images/screenshots.

# Tech
- [Go](https://go.dev/)
- [Google Gemini API](https://cloud.google.com/vertex-ai/generative-ai/docs/model-reference/gemini?_ga=2.40498155.-893235068.1676336544#multimodal)
- [Supabase Postgre Database](https://supabase.com/)

## API Endpoints

```
- Login + Register
/api/auth 

- Account information
/api/accounts 

- Make Gemini AI model request
/api/message 
```

## Running Unit Tests
```
cd to directory and run

go test . -v
```
