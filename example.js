const Zing = require('./modules/ZingMp3');

Zing.getStreaming('ZOI867CW')
    .then(data => console.log(data))
    .catch(err => console.log(err))