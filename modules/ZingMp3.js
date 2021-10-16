let request = require('request-promise');
const { FileCookieStore } = require('tough-cookie-file-store');
const fs = require('fs');

const encrypt = require('./encrypt');

const URL_API = 'https://zingmp3.vn';
const API_KEY = '88265e23d4284f25963e6eedac8fbfa3';
const SECRET_KEY = '2aa2d1c561e809b267f3638c4a307aab';
const VERSION = '1.4.2';

const cookiePath = 'ZingMp3.json';

if (!fs.existsSync(cookiePath)) fs.closeSync(fs.openSync(cookiePath, 'w'));

let cookiejar = request.jar(new FileCookieStore(cookiePath));

request = request.defaults({
    baseUrl: URL_API,
    qs: {
        apiKey: API_KEY,
    },
    gzip: true,
    json: true,
    jar: cookiejar,
});
class ZingMp3 {
    constructor() {
        this.time = null;
    }

    getFullInfo(id) {
        return new Promise(async (resolve, reject) => {
            try {
                let data = await Promise.all([
                    this.getInfoMusic(id),
                    this.getStreaming(id),
                ]);
                resolve({ ...data[0], streaming: data[1] });
            } catch (err) {
                reject(err);
            }
        });
    }

    getSectionPlaylist(id) {
        return this.requestZing({
            path: '/api/v2/playlist/getSectionBottom',
            qs: {
                id,
            },
        });
    }

    getDetailPlaylist(id) {
        return this.requestZing({
            path: '/api/v2/page/get/playlist',
            qs: {
                id,
            },
        });
    }

    getDetailArtist(alias) {
        return this.requestZing({
            path: '/api/v2/page/get/artist',
            qs: {
                alias,
            },
            haveParam: 1,
        });
    }

    getInfoMusic(id) {
        return this.requestZing({
            path: '/api/v2/song/get/info',
            qs: {
                id,
            },
        });
    }

    getStreaming(id) {
        return this.requestZing({
            path: '/api/v2/song/get/streaming',
            qs: {
                id,
            },
        });
    }

    getHome(page = 1) {
        return this.requestZing({
            path: '/api/v2/page/get/home',
            qs: {
                page,
            },
        });
    }

    getChartHome() {
        return this.requestZing({
            path: '/api/v2/page/get/chart-home',
        });
    }

    getWeekChart(id) {
        return this.requestZing({
            path: '/api/v2/page/get/week-chart',
            qs: { id },
        });
    }

    getNewReleaseChart() {
        return this.requestZing({
            path: '/api/v2/page/get/newrelease-chart',
        });
    }

    getTop100() {
        return this.requestZing({
            path: '/api/v2/page/get/top-100',
        });
    }

    search(keyword) {
        return this.requestZing({
            path: '/api/v2/search/multi',
            qs: {
                q: keyword,
            },
            haveParam: 1,
        });
    }

    async getCookie() {
        if (!cookiejar._jar.store.idx['zingmp3.vn']) await request.get('/');
    }

    // haveParam = 1 => string hash will have suffix
    requestZing({ path, qs, haveParam }) {
        return new Promise(async (resolve, reject) => {
            try {
                await this.getCookie();
                let param = new URLSearchParams(qs).toString();

                let sig = this.hashParam(path, param, haveParam);

                const data = await request({
                    uri: path,
                    qs: {
                        ...qs,
                        ctime: this.time,
                        sig,
                    },
                });

                if (data.err) reject(data);
                resolve(data.data);
            } catch (error) {
                reject(error);
            }
        });
    }

    hashParam(path, param = '', haveParam = 0) {
        this.time = Math.floor(Date.now() / 1000);
        // this.time = '1634406003';

        let strHash = `ctime=${this.time}`;
        if (haveParam === 0) strHash += param;
        const hash256 = encrypt.getHash256(strHash);
        return encrypt.getHmac512(path + hash256, SECRET_KEY);
    }
}

module.exports = new ZingMp3();
