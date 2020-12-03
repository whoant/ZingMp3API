const Zing = require('./modules/ZingMp3');

const id = 'ZO00E6FO';

Zing.getFullInfo(id)
    .then(data => console.log(data))
    .catch(err => console.log(err));