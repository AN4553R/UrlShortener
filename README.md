# URL Shortening Service (URLSS)
> A super fast URL shortener service.

## Motivation
This app takes a long URL and shrinks it down to fewer characters.
### Benefits
Aside from just being more pleasing to the eye, shortened URLs can have the following added benefits:

- More streamlined in appearance
  - If you’re trying to build a brand and are doing so through a social media campaign,
  having all your links be the same length and appear very similar is crucial.
  - Shortened links display better in emails, print material, and other places.
- Tracking capabilities
  - Clicks on the link can be tracked
- More validity in social media engagement:
  - A shortened link matching other links on your digital marketing materials and social media content lends validity to you as a business.
  - Long URLs with random characters and numbers can look shady, especially to someone who’s a new customer

### Downsides
- Security issues
  - You have no idea where a link is sending you when it’s in a shortened form. This can be helped by using a custom domain, but even then, it could seem suspicious.
- Unreliable service: 
  - links become defunct after timeout
>[System Design](./SysDesign.md)

## Technology stack 
- This service should store relatively small amounts of data, no need to store relations between different data models, and redirections should be processed as quickly as possible
- I use Go with Redis and Base 62 encoding.
## Overview
Generating unique URL using randomization and Base62 encoding
### Data model information
- Id integer 
- Original string 
- Expiration date 
- Visits integer
### Endpoints: 
- Create a new short link 
- Redirect to the origin URL on passing a short link 
- Get general information about the short link

### Insight
source: `https://www.example.co.uk:443/blog/article/search?docid=720&hl=en#dayone`

Destination: `http://localhost:8080/ZNkw23qpiss`

Request:
```bash
curl -L -X POST 'localhost:8080/encode' \
-H 'Content-Type: application/json' \
--data-raw '{
    "url": "https://www.example.co.uk:443/blog/article/search?docid=720&hl=en#dayone",
    "expires": "2022-10-04 17:18:00"
}'
```
Response:
```shell
{"success":true,"shortUrl":"http://localhost:8080/ZNkw23qpiss"}
```
## Load testing
This is just service testing results on a local machine
