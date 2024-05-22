/*
const http = require('http')

const server = http.createServer((req, res) => {
// Set CORS headers
res.setHeader('Access-Control-Allow-Origin', '*'); // Replace * with your client's origin
res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE');
res.setHeader('Access-Control-Allow-Headers', 'Content-Type');

res.writeHead(200, { 'Content-Type': 'text/plain' });
res.write(JSON.stringify({name:"Vaibhav"}))
res.end()

})

server. listen(3000,()=>{
    console.log("listening on port 3000 ...")
})
*/


const express = require('express');
const app = express()

app.get('/',(req,res)=>{
    console.log('user hit the resource')
    const obj = {
        name:"Vaibhav",
        age: 21,
        fav_color: "olive-green"
    }
    res.json(obj)
    res.status(200).send('Home Page')
})

app.get('/about',(req,res)=>{
    res.status(200).send('About Page')
})

app.all('*',(req,res)=>{
    res.status(404).send('<h1>Page Not Found</h1>')
})

app.listen(3000,()=>{
    console.log('listening on port 3000...')
})