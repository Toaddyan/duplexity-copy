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
- cons: not scalable 

We are deciding for the proxy to hit the backend service instead of the REDIS DB 
- Pros: cleaner architecture 
- Cons: slower load time for mommy. 

websocket Hit control first 

The websockets themselves have many clients to them. 
If i'm trying to target a specific client, then we need to have all of these websockets subscribed to a channel. 


## Broker 

Who is the broker? 
The websocket should be a broker to the clients 

===
# Duplexity Speedrun

## Client Connecting to Duplexity

1. Client has clientID and pipe(s)
2. Client connects control websocket with discoveryRequest
    - clientID
3. Control websocket responds back with discoveryResponse
    - dataPlaneURI
4. Client connects to data websocket
    - clientID
    - pipe(s)
    Websocket does a registerPipe on backend
5. a) Data websocket's authorizer sends pipesAreRegisteredResponse via hub's control plane
   b) Client waits until it receives pipesAreRegisteredResponse
6. Client is all good

## Client Disconnects from Duplexity

1. Client has clientID
2. User does a Control + C
3. Client sends a clientDisconnectRequest
   This is only a courtsey
4. Control does a hub.unregister <- client
5. hub.unregsiter does a removePipeRequest on backend


## Force Disconnect a client from Duplexity

1. Force clientID to disconnect off of Duplexity
2. Send to hub a forceDisconnectRequest via readChan
