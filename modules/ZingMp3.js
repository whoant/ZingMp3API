let request = require("request-promise");
const encrypt = require('./encrypt');

const API_KEY = 'kI44ARvPwaqL7v0KuDSM0rGORtdY1nnw';
const SERCRET_KEY = '882QcNXV4tUZbvAsjmFOHqNC1LpcBRKW';
const URL_API = 'https://zingmp3.vn';

let cookiejar = request.jar();

request = request.defaults({
    qs: {
        apiKey: API_KEY
    },
    gzip: true,
    json: true,
    jar: cookiejar
});
class ZingMp3 {


    static getFullInfo(id){
        return new Promise(async(resolve, reject) => {
            try {
                let data = await Promise.all([this.getInfoMusic(id), this.getStreaming(id)]);
                resolve({...data[0], streaming: data[1]});
            } catch (err) {
                reject(err);
            }
        });
    }

    static getInfoMusic(id) {
        return new Promise(async (resolve, reject) => {

            const option = {
                path: '/api/v2/song/getInfo',
                qs: {
                    id
                },
                param: 'id=' + id
            };
            
            try {
                const data = await this.requestZing(option);
                if (data.err) reject(data);
                resolve(data.data);
            } catch (error) {
                reject(error);
            }
        })
    }
    static getStreaming(id) {
        return new Promise(async (resolve, reject) => {

            const option = {
                path: '/api/v2/song/getStreaming',
                qs: {
                    id
                },
                param: 'id=' + id
            };
            
            try {
                const data = await this.requestZing(option);
                if (data.err) reject(data);
                resolve(data.data);
            } catch (error) {
                reject(error);
            }
        })
    }

    static getHome(page = 1) {
        return new Promise(async (resolve, reject) => {

            const option = {
                path: '/api/v2/home',
                qs: {
                    page
                },
                param: 'page=' + page
            };
            
            try {
                const data = await this.requestZing(option);
                if (data.err) reject(data);
                resolve(data.data);
            } catch (error) {
                reject(error);
            }
        })
    }

    static async getCookie(){
        if (!cookiejar._jar.store.idx['zingmp3.vn']) await request.get(URL_API);
        
    }

    static async requestZing({path, param, qs})
    {
        await this.getCookie();
        
        let sig = this.hashParam(path, param);
        return request({
            uri: URL_API + path,
            qs: {
                ctime: this.time,
                sig,
                ...qs
            },
        });
    }

    static hashParam(path, param = '')
    {
        this.time = Math.floor(Date.now() / 1000);
        const hash256 = encrypt.getHash256(`ctime=${this.time}${param}`);
        return encrypt.getHmac512(path + hash256, SERCRET_KEY);
    }
}

module.exports = ZingMp3;