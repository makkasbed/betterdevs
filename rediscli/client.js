const net = require('net');
const { parseArgs } = require('./parseArgs');
const { parseResp } = require('./parseResp');
const prompt = require('prompt-sync')({sigint: true});


const client = new net.Socket();
let port = 6379;
let host = '127.0.0.1';

let command = '';


const args = parseArgs(process.argv.slice(2));


if (args.hasOwnProperty("h")) {
    host = args["h"];
    if (args.hasOwnProperty("p")) {
        port = args["p"];
    }
}

client.connect(port, host, function() {
   
    client.write('\r\nPING\r\n');
    
});

client.on('data', function(data) {
    console.log(parseResp(data.toString()));
    command = prompt(`${host}:${port}>`);
    if (command == 'quit') {
        process.exit();
    }else{
        client.write(`\r\n${command}\r\n`);
    }
    
});

client.on('close', function() {
    console.log('Connection closed');
});


