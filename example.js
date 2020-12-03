const ZingMp3 = require('./modules/ZingMp3');

ZingMp3.getSongInfo('ZOI6BFA9')
    .then(({data}) => {

        console.log(data);
    })
    .catch(err => console.log(err));