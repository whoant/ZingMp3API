const ZingMp3 = require('./modules/ZingMp3');

const Zing = new ZingMp3();

Zing.getStreaming('ZEFE70B9').then(data => console.log(data)).catch(err => console.log(err));