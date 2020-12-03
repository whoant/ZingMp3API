let request = require('request-promise');
const encrypt = require('./encrypt');

const API_KEY = '38e8643fb0dc04e8d65b99994d3dafff';
const SERCRET_KEY = '10a01dcf33762d3a204cb96429918ff6';
const URL_API = [
    'https://zingmp3.vn/api',
    'https://beta.zingmp3.vn'
];

request = request.defaults({
    qs: {
        api_key: API_KEY,
        apiKey: API_KEY
    },
    gzip: true,
    json: true
});

class ZingMp3 {

    static getFullInfo(id){
        return new Promise(async(resolve, reject) => {
            try {
                const data = await Promise.all([this.getSongInfo(id), this.getStreaming(id)]);
                const infoSong = data[0].data;
                let res = {
                    id,
                    title: infoSong.title,
                    artists_names: infoSong.artists_names,
                    thumbnail: infoSong.thumbnail_medium,
                    lyric: infoSong.lyric,
                    streaming: data[1].data
                };
                resolve(res);
            } catch (error) {
                reject(error);
            }
        });
    }

    static getInfoDetail(id) {
        return new Promise(async (resolve, reject) => {
            const option = {
                nameAPI: '/song/get-song-detail',
                typeApi: 0,
                qs: {
                    id
                },
                param: 'id=' + id
            };

            try {
                const data = await this.requestZing(option);
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
                typeApi: 0,
                qs: {
                    id
                },
                param: 'id=' + id
            };

            try {
                const data = await this.requestZing(option);
                if (data.err) reject(data);
                resolve(data);
            } catch (error) {
                reject(error);
            }
        })
    }

    static getStreaming(id) {
        return new Promise(async (resolve, reject) => {

            const option = {
                nameAPI: '/api/v2/song/getStreaming',
                typeApi: 1,
                qs: {
                    id
                },
                param: 'id=' + id
            };

            try {
                const data = await this.requestZing(option);
                if (data.err) reject(data);
                resolve(data);
            } catch (error) {
                reject(error);
            }
        })
    }

    static requestZing({nameAPI, typeApi, param, qs})
    {
        
        let sig = this.hashParam(nameAPI, param);
        return request({
            uri: URL_API[typeApi] + nameAPI,
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