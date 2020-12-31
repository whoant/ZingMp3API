const Zing = require('./modules/ZingMp3');


Zing.getSectionPlaylist('6Z87F988')
    .then(data => console.log(data))
    .catch(err => console.log(err))