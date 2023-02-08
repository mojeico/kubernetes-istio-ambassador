
// This line must come before importing any instrumented module.



// todo read about it - https://docs.datadoghq.com/tracing/other_telemetry/connect_logs_and_traces/nodejs

const tracer = require('dd-trace').init({
    logInjection: true
});

/*const formats = require('dd-trace/ext/formats');

class Logger {
    log(level, message) {
        const span = tracer.scope().active();
        const time = new Date().toISOString();
        const record = { time, level, message };

        //console.log(` SPAN is - ${span}`)
        if (span) {
            tracer.inject(span.context(), formats.LOG, record);
        }

        console.log(JSON.stringify(record));
    }
}

module.exports = Logger;*/


const http = require('http');

const express = require('express');
const app = express();



var message = "ERROR !!! not get response !!!"


// define handler for /goodbye URL
app.get('/goodbye', function(req, res) {

    //tracer.init()

    //const my_logger = new Logger();
    //my_logger.log("INFO", 'Client get request ');

    console.log("INFO", 'Client get request ');

    res.send("goodbye");

});

// define handler for /hello URL
app.get("/hello", function(req, res) {

    //const my_logger = new Logger();

    //my_logger.log("INFO", 'Client get request ');

    console.log("INFO", 'Client get request ');

    http.get('http://nodejs-server-service:80/servernodejs', res => {
        let data = [];
        const headerDate = res.headers && res.headers.date ? res.headers.date : 'no response date';
        console.log('Status Code:', res.statusCode);
        console.log('Date in Response header:', headerDate);

        res.on('data', chunk => {
            data.push(chunk);
        });

        res.on('end', () => {
            console.log('Response ended: ');
            //const users = JSON.parse(Buffer.concat(data).toString());

            //my_logger.log("INFO", ` ${Buffer.concat(data).toString()}`);
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