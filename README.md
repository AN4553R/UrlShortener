# URL Shortening Service (URLSS)
> A super fast URL shortener service written in Go.

## Motivation
This app takes a long URL and shrinks it down to fewer characters.

### Insight
Input: `https://www.example.co.uk:443/blog/article/search?docid=720&hl=en#dayone`

Output: `http://localhost:8080/ZNkw23qpiss`


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

## System Design
### Requirements
#### Functional
- The service should generate a uniqe alias for the provided address
- The service should redirect user to the original URL when he calls the short link
- The short link should have a lifetime that the user specifies on creation
- The short link should track clicks count
#### Non-functional
- The service should be able to handle numerous requests
- Forwarding should be real-time with minimum delay
- The short link should be random in order not to be predictable
#### Estimations
The service is going to serve heavy reads since there will be a huge number of redirects compared to creating new ones. Let’s assume that the ratio between reading and writing is 50:1.
##### Traffic estimates
If we have 500k new short links every month, then we will expect 25 million (50 * 500k = 25 million) redirects for the same period. So we have 1 new link every 5 seconds: 

`500k / (30 days * 24 hours * 3600 seconds) = ~ 1 link in 5 seconds.`

And 10 redirects every second: 

`25 million / (30 days * 24 hours * 3600 seconds) = ~ 10 redirects per second.`
##### Memory estimates
Let’s say we store each address for a maximum - 1 year. Since we expect 500k new link every month, then we will have near 6 million records in the database: 
`500k record/month * 12 months = 6 million`

Let’s assume that each record in the database - approximately 1000 bytes. [The recommended maximum size for a link is 2000 characters](https://stackoverflow.com/questions/417142/what-is-the-maximum-length-of-a-url-in-different-browsers/417184#417184) and according to the standard, the URL encodes with ASCII characters, which occupy 1 byte, i.e. the link can hold  2000 bytes by recommended maximum size). So we will use half of this value as average. Then we need 6 TB of memory to store records for 1 year: 

`6 million record * 1000 bytes per record = 6 GB`

>A little summary of the nature of the model that we will be working with:
>- We need to store several million records
>- Each record is small 
>- The service is very read-heavy

## Technology stack 
This service should store relatively small amounts of data, no need to store relations between different data models, and redirections should be processed as quickly as possible, we will use Redis with Goland.
## Overview
Generating unique URL using randomization and Base62 encoding
### Data model information
- Id integer 
- Original string 
- Expiration date 
- Visits integer
### We will have 3 endpoints: 
- Create a new short link 
- Redirect to the origin URL on passing a short link 
- Get general information about the short link
