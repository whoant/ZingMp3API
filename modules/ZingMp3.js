let request = require('request-promise');
const encrypt = require('./encrypt');

const API_KEY = '38e8643fb0dc04e8d65b99994d3dafff';
const SERCRET_KEY = '10a01dcf33762d3a204cb96429918ff6';

request = request.defaults({
    baseUrl: 'https://zingmp3.vn/api',
    qs: {
        api_key: API_KEY
    },
    gzip: true,
    json: true
});

class ZingMp3 {


    static getInfoDetail(id) {
        return new Promise(async (resolve, reject) => {
            const option = {
                nameAPI: '/song/get-song-detail',
                qs: {
                    id
                },
                param: 'id=' + id
            };

            try {
                const data = await ZingMp3.requestZing(option);
                if (data.err) reject(data);
                resolve(data);
            } catch (error) {
                reject(error);
            }
        })
    }

    static getSongInfo(id) {
        return new Promise(async (resolve, reject) => {

            const option = {
                nameAPI: '/song/get-song-info',
                qs: {
                    id
                },
                param: 'id=' + id
            };

            try {
                const data = await ZingMp3.requestZing(option);
                if (data.err) reject(data);
                resolve(data);
            } catch (error) {
                reject(error);
            }
        })
    }

    static requestZing({nameAPI, param, qs})
    {
        let sig = this.hashParam(nameAPI, param);
        return request({
            uri: nameAPI,
            qs: {
                ctime: this.time,
                sig,
                ...qs
            }
        });
    }

    static hashParam(nameAPI, param = '')
    {
        this.time = Math.floor(Date.now() / 1000);
        const hash256 = encrypt.getHash256(`ctime=${this.time}${param}`);
        return encrypt.getHmac512(nameAPI + hash256, SERCRET_KEY);
    }
}

module.exports = ZingMp3;