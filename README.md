# ZingMp3 API

ZingMp3API is a NodeJS library.

## Reverse API

More infomation, you can read my [post](https://vovanhoangtuan4-2.medium.com/tôi-đã-lấy-api-zingmp3-như-thế-nào-55f5fa555eda)

## Installation

```bash
npm i zingmp3-api
```

## Usage

### 1.Get home

```javascript
ZingMp3.getHome(1);
```

### 2. Get info music

Example: https://zingmp3.vn/bai-hat/City-Girls-Chris-Brown-Young-Thug/ZWBOW9CO.html

```javascript
ZingMp3.getFullInfo('ZWBOW9CO');
```

### 3. Get details playlist or album

Example: https://zingmp3.vn/album/Top-100-Nhac-R-B-Au-My-Hay-Nhat-Various-Artists/ZWZB96C8.html

```javascript
ZingMp3.getDetailPlaylist('ZWZB96C8');
```

### 4: Search

Example: Your keyword: Sơn Tùng MTP

```javascript
ZingMp3.search('Sơn Tùng MTP');
```

### 5: Get Chart

```javascript
ZingMp3.getChartHome();
```

### 6: Get Chart Week

Example: https://zingmp3.vn/zing-chart-tuan/Bai-hat-Viet-Nam/IWZ9Z08I.html

```javascript
ZingMp3.getWeekChart('IWZ9Z08I');
```

## Contact

[Facebook](https://www.facebook.com/vovanhoangtuan/)
