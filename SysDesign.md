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
