const http = require('http');

function object_creator(name,age) {
const obj = {
	name,
	age,
	introduction : function () {
		return `Hello my name is ${this.name}`
		}
	}
	return JSON.stringify(obj);
}

http.createServer(function (req, res) {
  res.writeHead(200, {'Content-Type': 'text/plain'});
  const obj = object_creator("abc",21);
  res.write(obj);
  res.end();
}).listen(8080);
console.log("listening on port 8080....");
