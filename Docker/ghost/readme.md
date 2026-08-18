# Ghost Blog


## Contianer Setup
For dev, run with:
`docker run -d -e NODE_ENV=development -e url=http://localhost:3001 -p 3001:2368 -v ghost-vol:/var/lib/ghost ghost`