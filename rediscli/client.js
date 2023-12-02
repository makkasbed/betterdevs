#! /usr/bin/env node
const { program } = require('commander');
const { config } = require('dotenv');
const net = require('net');
const quit = require('./commands/quit');


config();

program.option('-h, --host <string>');

program.command('quit')
       .description('Quit the rediscli')
       .action(quit);



let host = "";
let port = "";



const client = net.createConnection(process.env.PORT, process.env.HOST,()=>{
    console.log(`Connected to Redis Server ${process.env.HOST}`);
    client.write('\r\nping hello\r\n');
});

client.on("data",(data)=>{
    console.log(`Received: ${data}`);
});

client.on("error",(err)=>{
    console.log(`Error: ${err.message}`);
});

client.on("close",()=>{
    console.log("Connection closed");
});

program.parse();

const options = program.opts();
host = options.host;
console.log('Host ', host);
