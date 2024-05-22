// write a code to listen server on port 5000
const http = require('http')
const server = http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'text/plain' })
  res.end("Hello World!")
})

server.listen(5000, () => {
  console.log("server is listening on port 5000...")
})