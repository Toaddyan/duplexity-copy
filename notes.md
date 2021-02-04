# Definitions
## Data plane = User Plane 
- The content that we are proxying
## Control Plane 
- Management of our app

# Control Plane 
## Flow of Mom
1. Mommy will want to talk to Proxy Service to find out where usernode is 
2. Proxy Service will respond by saying where the usernode is. 
## Flow of User
1. Usernode will  say I want to connect to backend service
2. Backend service will give it a websocket service to connect to 