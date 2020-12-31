let request = require("request-promise");
const {FileCookieStore} = require('tough-cookie-file-store');
const fs = require('fs');

const encrypt = require('./encrypt');

const URL_API = 'https://zingmp3.vn';
const API_KEY = 'kI44ARvPwaqL7v0KuDSM0rGORtdY1nnw';
const SERCRET_KEY = '882QcNXV4tUZbvAsjmFOHqNC1LpcBRKW';
const cookiePath = 'ZingMp3.json';

if (!fs.existsSync(cookiePath)) fs.closeSync(fs.openSync(cookiePath, 'w'));

let cookiejar = request.jar(new FileCookieStore(cookiePath));

request = request.defaults({
    qs: {
        apiKey: API_KEY
    },
    gzip: true,
    json: true,
    jar: cookiejar
});
class ZingMp3 {

    static getFullInfo(id) {
        return new Promise(async (resolve, reject) => {
            try {
                let data = await Promise.all([this.getInfoMusic(id), this.getStreaming(id)]);
                resolve({ ...data[0], streaming: data[1] });
            } catch (err) {
                reject(err);
            }
        });
    }

    static getSectionPlaylist(id) {
        return this.requestZing({
            path: '/api/v2/playlist/getSectionBottom',
            qs: {
                id
            }
        });

    }

    static getDetailPlaylist(id) {
        return this.requestZing({
            path: '/api/v2/playlist/getDetail',
            qs: {
                id
            }
        });

    }

    static getInfoMusic(id) {
        return this.requestZing({
            path: '/api/v2/song/getInfo',
            qs: {
                id
            }
        });
    }

    static getStreaming(id) {
        return this.requestZing({
            path: '/api/v2/song/getStreaming',
            qs: {
                id
            }
        });
    }

    static getHome(page = 1) {
        return this.requestZing({
            path: '/api/v2/home',
            qs: {
                page
            }
        });
    }

    static async getCookie() {
        if (!cookiejar._jar.store.idx['zingmp3.vn']) await request.get(URL_API);
    }

    static requestZing({ path, qs }) {
        return new Promise(async (resolve, reject) => {
            try {
                await this.getCookie();
                let param = new URLSearchParams(qs).toString();

                let sig = this.hashParam(path, param);

                const data = await request({
                    uri: URL_API + path,
                    qs: {
                        ctime: this.time,
                        sig,
                        ...qs
                    },
                });

                if (data.err) reject(data);
                resolve(data.data);
            } catch (error) {
                reject(error);
            }
        });
    }

    static hashParam(path, param = '') {
        this.time = Math.floor(Date.now() / 1000);
        const hash256 = encrypt.getHash256(`ctime=${this.time}${param}`);
        return encrypt.getHmac512(path + hash256, SERCRET_KEY);
    }
}

module.exports = ZingMp3;