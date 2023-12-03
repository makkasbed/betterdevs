exports.parseResp=function (res) {
    let str = res;
    if (str.charAt(0)==="$") {
        str = str.split('\r\n')[1];
    }
    return str.replaceAll('$','').replaceAll('+','').replaceAll('-','').replace(/\d\r\n/g,'').replaceAll('\r\n','');
}