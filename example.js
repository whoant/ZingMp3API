const Zing = require('./modules/ZingMp3');

Zing.getFullInfo('ZOOI7Z87')
    .then(data => console.log(data))
    .catch(err => console.log(err))

// Zing.getHome()
//     .then(data => console.log(data))
//     .catch(err => console.log(err))