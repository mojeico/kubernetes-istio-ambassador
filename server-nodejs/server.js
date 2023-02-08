




const express = require('express');
const app = express();


var i = 0



// define handler for /hello URL
app.get("/servernodejs", function(req, res) {


    res.statusCode = 200;
    res.end(`from v3 `);

});

const server = app.listen(3001);
