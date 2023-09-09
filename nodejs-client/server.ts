

/// ###############################
/// ######## ADD THIS PART ########
/// ###############################

// https://docs.datadoghq.com/tracing/other_telemetry/connect_logs_and_traces/nodejs/
// https://docs.datadoghq.com/tracing/trace_collection/dd_libraries/nodejs/?tab=containers


// tracer.ts


const http = require('http');
const express = require('express');
const app = express();

var message = "ERROR !!! not get response !!!"


app.get("/hello1", function(req, res) {


    http.get('http://nodejs-server-service:3000/servernodejs', res => {
        let data = [];
        const headerDate = res.headers && res.headers.date ? res.headers.date : 'no response date';
        console.log('Status Code:', res.statusCode);
        console.log('Date in Response header:', headerDate);
        res.on('data', chunk => {
            data.push(chunk);
        });
        res.on('end', () => {
            console.log("INFO", ` ${Buffer.concat(data).toString()}`);
            message = ` ${Buffer.concat(data).toString()}`
        });
    }).on('error', err => {
        console.log('Error: ', err.message);
    });
    res.statusCode = 200;
    res.setHeader('Content-Type', 'text/plain');
    res.end(message);
});


const server = app.listen(3000);