const net = require('node:net');


exports.connect = async function(host, port){
    try {
        const client = await net.createConnection(host,port);
        
        return client;
    } catch (error) {
        console.log(`Unable to connect at this time ${host}:${port}`);
    }  
}

exports.command = async function (host,port, cmd) {
    try {
        const client = await net.createConnection(host,port);
        const response = await client.write(cmd);
        console.log(response);
    } catch (error) {
        console.log(`Unable to execute command ${cmd}`);
    }
}

