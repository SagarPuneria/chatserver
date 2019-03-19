# Client Server
A client server program developed using golang.

## Project details: Updates and revision.
There are 10000 clients (C) and there is 1 server (S).

- Server (S)  is a TCP server listening on a port 8080.
- All 10000 Clients (C) connect to server (S).
- Each client maintains 2 version numbers.
- Current Version & Previous Version.
- Server maintains Current Version and Previous Version
- After connecting to the Server, each client sends Current Version to the Server.
- If the Client’s Current Version does not match Server’s Current Version, then Server sends a current Version update.
- Client takes Server supplied current version. Moves its current version to previous version. And applies the supplied current version to its current version.
- Then sends the current CV & PV back to server.
- Goal: All 10,000 clients should start with random CV (Current Version) and PV (Previous Version) and should get updated to Server’s CV as fast as possible (say < 1 min).

## Preliminary
This instruction will run on the local machine for above problem statement and will help you to get the results of the project.

### Application Requirements
This project runs on Linux, Windows, Darwin machines and uses go1.12 in its creation.

### The Project Structure
1. chatserver/server/server.go - The main point of server program.
2. chatserver/client/client.go - The main point of client program.
3. chatserver/ipvalidation/ipcheck.go - Contains logic for validating ipaddress and port.
4. chatserver/util/util.go - Contains logic for random number, address, reading console input and exceptional handling.

### Instructions for build and run the client-server application
Step 1 - Open the Terminal <br />
Step 2 - Go to the stored project source directory <br />
Step 3 - Set GOPATH evironment variable which determines the location of this project workspace <br />
Step 4 - Run ```go build server.go``` <br />
Step 5 - Run ```./server``` <br />
Step 6 - Run ```go build client.go``` <br />
Step 7 - Run ```./client``` <br />
Step 8 - Enter valid server ip address <br />
Step 8 - Enter valid server port address else the default port will be 8080 <br />