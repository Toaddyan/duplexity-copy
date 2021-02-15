# duplexity
 
 Websocket has 2 end points
 Dialer end point is user end point. controller is control plane end point

 Brand new Laptop first time connecting to duplexity
 Client should first establish identity. 
 ### Auth 
 - Who am I? client should talk to auth end point first.  (/auth)
 - call back will call store code into the redis DB. 
 - Auth endpoint will flush everything back to the client. Client will get its ID token

### Websocket 
 - Talk to the control end point FIRST. have dialer wait. 
 - Control will talk to the backend first @/register of backend with id and websocket information.

## Backend
- /register
- /locate
- /remove

## Redis Pub Sub Set-up 
- topics: 
    - websockets
        - disconnect 

### Register
 - Takes in a big object:
     - array 
        - Hostname 
        - resource destination. 
        - EX:  Key = cat1.duplexity.io, localhost:9090
        - EX2: Todd1.dupelxity.io, localhost:9090
    - backend will do the ok for /dialer and let dialer go

## Decisions 
We are deciding to hit the auth service directly 
- pros: easy 
- cons: scalable 

We are deciding for the proxy to hit the backend service instead of the REDIS DB 
- Pros: cleaner architecture 
- Cons: slower load time for mommy. 

websocket Hit control first 


ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDI4lur1lVGKRJf6nisjAPjsjTn9V/JGdHBdkDV/NPa/yOfnG7tcPe+BUpAbeLdsWKwJUfJBlVC8e7BctrrDiaofYQUzIbf82Rs471Bb69EwbZSmIdeYGfkoUxiDTJnVVwsgrP/PeUxDyNr3lut/Tpn2qfi4ekEONiedLW6pFvz+pikkHYBKZTvQ69A1ZcwA+AKEf0rFQ5njyfQh3x7jIlR5866EzWlp6ZYl7mv6RxQEpu1eHxsMk02RG/pnff2FQMlmLhQzT+hWhNzy6baKdj3K08JH3IEMcx+4UrtFB0Dd+gzxF5IilbpAXpwRmRUiOZrdigo80M32DnqqcLG9S6n thyan@ad3.ucdavis.edu@pc42


The websockets themselves have many clients to them. 
If i'm trying to target a specific client, then we need to have all of these websockets subscribed to a channel. 
