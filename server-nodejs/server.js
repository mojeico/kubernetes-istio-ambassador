


const tracer = require('dd-trace').init({
    profiling: true,
    env: 'snd',
    service: 'nodejs-server',
    version: 'v1'
})

const express = require('express');
const app = express();


var i = 0



// define handler for /hello URL
app.get("/servernodejs", function(req, res) {

    i++

    res.statusCode = 200;
    res.end(`response number - ${i}`);

});

const server = app.listen(3001);
