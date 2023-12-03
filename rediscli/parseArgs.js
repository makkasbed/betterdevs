exports.parseArgs= function (args) {
    let json = {};

    let isSet = false;
    let key = "";
    let val = "";

    for (let i = 0; i < args.length; i++) {
        
        const element = args[i];
        if (element.charAt(0)==="-") {
            key = element.substr(1,element.length);
        }else{
           val = element;
           isSet = true;
        }

        if (isSet===true) {
            json[key] = val;
            isSet = false;
        }
        
    }
    return json;
    
}