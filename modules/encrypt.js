const crypto = require('crypto');


const getHash256 = (a) => {
    return crypto.createHash('sha256').update(a).digest('hex');
}
const getHmac512 = (str, key) => {
    let hmac = crypto.createHmac("sha512", key);
    return hmac.update(Buffer.from(str, 'utf8')).digest("hex");
}

module.exports = {getHash256, getHmac512};