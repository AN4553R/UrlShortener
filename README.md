# URL Shortening Service (URLSS)
> A super fast URL shortener service written in Go.

## Motivation
This app takes a long URL and shrinks it down to fewer characters.

### Insight
Input: `https://www.example.co.uk:443/blog/article/search?docid=720&hl=en#dayone`

Output: `http://localhost:8080/ZNkw23qpiss`


## Benefits
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

## Downsides
- Security issues
  - You have no idea where a link is sending you when it’s in a shortened form. This can be helped by using a custom domain, but even then, it could seem suspicious.
- Unreliable service: 
  - links become defunct after timeout

## Requirements
### Functional
- The service should generate a uniqe alias for the provided address
- The service should redirect user to the original URL when he calls the short link
- The short link should have a lifetime that the user specifies on creation
- The short link should track clicks count
### Non-functional

