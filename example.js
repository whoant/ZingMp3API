const Zing = require('./modules/ZingMp3');
const fs = require('fs');

Zing.getDetailPlaylist('ZWZB969E')
    .then(data => console.log(1))
    .catch(err => console.log(err))

// Zing.getHome(1)
//     .then(data => console.log(data))
//     .catch(err => console.log(err))