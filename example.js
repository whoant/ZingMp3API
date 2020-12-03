const Zing = require('./modules/ZingMp3');

Zing.getInfoDetail('ZO00E6FO')
    .then(data => console.log(data))
    .catch(err => console.log(err));