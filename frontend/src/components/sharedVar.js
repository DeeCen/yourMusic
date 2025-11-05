import { ref, reactive } from 'vue';
import { DownloadSong, GetSongURL } from '@/wailsjs/go/main/App.js';
import { ElMessage } from 'element-plus';
import { gzToStr, strToGz } from './gz.js';

const userInfoSaveKey = 'userInfo';
const userSelectQualityKey = `userSelectQuality`;
const savePlaySongListKey = 'userPlaySongList';
const songQuality128 = '128'; // 对应 FileHash
const songQuality320 = '320'; // 对应 HQ
const songQualityFlac = 'flac'; // 对应 SQ
const songQualityHiRes = 'HiRes'; // 不知道对应哪个 FileHash,为了兼容保留

let userSelectQualityVal = localStorage.getItem(userSelectQualityKey) || songQualityFlac;
let lastSavePlaySongListLen = 0;

/** @type {import('vue').Ref<api.Song[]>} */
export const playSongList = ref([]);
export const userInfo = reactive({ dfid: '', pic: '', token: '', userid: '0' });

const copyText = async (text) => {
    try {
        await navigator.clipboard.writeText(text);
        return true;
    } catch (err) {
        console.warn('回退到传统复制', err);
        return copyFallback(text);
    }
};

// 传统复制
const copyFallback = (text) => {
    try {
        const textarea = document.createElement('textarea');
        textarea.value = text;
        textarea.style.cssText = `
        position: fixed;
        left: -9999px;
        opacity: 0;
        pointer-events: none;
      `;

        document.body.appendChild(textarea);
        textarea.select();
        textarea.setSelectionRange(0, 99999);
        const successful = document.execCommand('copy');
        document.body.removeChild(textarea);
        return successful;
    } catch (err) {
        console.warn('传统复制失败:', err);
    }
    return false;
};

const resetUserInfo = (info) => {
    if (!info || info.token === '' || info.userid === '' || info.userid === '0' || info.userid === 0) {
        return false;
    }

    userInfo.dfid = String(info.dfid || '');
    userInfo.pic = String(info.pic || '');
    userInfo.token = String(info.token || '');
    userInfo.userid = String(info.userid || '0');
    localStorage.setItem(userInfoSaveKey, JSON.stringify(userInfo));
    return true;
};

export const updateUserInfoToLocalStorage = (info) => {
    console.log('updateUserInfoToLocalStorage', info);
    resetUserInfo(info);
};

export const readUserInfoFromLocalStorage = () => {
    const s = localStorage.getItem(userInfoSaveKey);
    if (!s) {
        return false;
    }

    try {
        const info = JSON.parse(s);
        return resetUserInfo(info);
    } catch (e) {
        console.log('readUserInfoFromLocalStorage error', e);
    }
    return false;
};

export const removeUserInfoToLocalStorage = () => {
    localStorage.removeItem(userInfoSaveKey);

    userInfo.dfid = ``;
    userInfo.pic = ``;
    userInfo.token = ``;
    userInfo.userid = `0`;
};

export const getSongHash = (song, quality) => {
    const HQ = song?.HQ?.Hash || '';
    const SQ = song?.SQ?.Hash || '';
    const hash = song.FileHash || '';

    // 后台音质有回退机制,给不同顺序即可
    let ret = [hash, HQ, SQ];
    switch (quality) {
        case songQualityHiRes:
        case songQualityFlac:
            ret = [SQ, HQ, hash];
            break;

        case songQuality320:
            ret = [HQ, hash, SQ];
            break;
    }

    ret = ret.filter((v) => (v ? v : null));
    return ret.join(',');
};

export const copySongURL = async (song) => {
    /**@type api.Song */
    if (!song) {
        return;
    }

    const resp = await GetSongURL(userInfo.dfid, userInfo.userid, userInfo.token, getSongHash(song, songQualityHiRes));
    if (resp.errMsg) {
        ElMessage.warning(`地址获取失败-` + resp.errMsg);
        return;
    }

    if (!resp.data || resp.data.length === 0) {
        ElMessage.warning(`地址获取失败`);
        return;
    }

    const copyOk = await copyText(resp.data[0]);
    if (copyOk) {
        const mb = resp.size > 0 ? (resp.size / 1024 / 1024).toFixed(2) : '0';
        ElMessage.success(`已复制 (${mb} MB)`);
        return;
    }

    ElMessage.warning(`复制失败,请自行复制:` + resp.data[0]);
};

export const downloadSong = async (song) => {
    /**@type api.Song */
    if (!song) {
        return;
    }
    const resp = await DownloadSong(userInfo.dfid, userInfo.userid, userInfo.token, getSongHash(song, songQualityHiRes),song.FileName);
    if (resp.is_success===false) {
        ElMessage.warning(resp.msg);
        return;
    }

    ElMessage.success(resp.msg);
};

export const setUserSelectQuality = (flag) => {
    switch (flag) {
        case songQuality128:
        case songQuality320:
        case songQualityFlac:
        case songQualityHiRes:
            userSelectQualityVal = flag;
            localStorage.setItem(userSelectQualityKey, userSelectQualityVal);
            break;
    }
};

export const getUserSelectQuality = () => {
    switch (userSelectQualityVal) {
        case songQuality320:
        case songQualityFlac:
        case songQualityHiRes:
            return userSelectQualityVal;

        default:
            return songQuality128;
    }
};

const getPlaySongList = async () => {
    const b64 = localStorage.getItem(savePlaySongListKey);
    if (!b64) {
        return [];
    }

    try {
        const json = await gzToStr(b64);
        const songList = json ? JSON.parse(json) : [];
        songList.map((song) => {
            fixSongForPlayer(song);
        });
        return songList;
    } catch (e) {
        console.log(`getPlaySongList error`, e);
    }
    return [];
};

const savePlaySongList = async () => {
    const l = playSongList.value.length;
    if (l === lastSavePlaySongListLen) {
        console.log('savePlaySongList stop 2', l, lastSavePlaySongListLen);
        return;
    }

    const saveList = playSongList.value.slice(0, 1000);
    const deepCopy = [];
    for (const song of saveList) {
        deepCopy.push({
            SingerName: song.SingerName,
            Image: song.Image,
            FileHash: song.FileHash,
            AlbumID: song.AlbumID,
            FileName: song.FileName,
            SQ: { Hash: song.SQ.Hash },
            HQ: { Hash: song.HQ.Hash }
        });
    }

    const gzStr = await strToGz(JSON.stringify(deepCopy));
    if (!gzStr) {
        console.log('savePlaySongList fail 2');
        return;
    }
    localStorage.setItem(savePlaySongListKey, gzStr);
    lastSavePlaySongListLen = l;
    console.log('savePlaySongList success 2', gzStr.length / 1024, 'KB');
};

let savePlayListDelayTimerId = 0;
export const savePlayListDelay = () => {
    clearTimeout(savePlayListDelayTimerId);
    savePlayListDelayTimerId = setTimeout(() => {
        console.log('savePlaySongList run 1');
        savePlaySongList()
            .then(() => {
                console.log('savePlaySongList end 3');
            })
            .catch((e) => {
                console.log('savePlaySongList err 3', e);
            });
    }, 5000);
};

export const initSharedVarOnce = async () => {
    if (window['_initSharedVar']) {
        return;
    }
    window['_initSharedVar'] = true;

    playSongList.value = await getPlaySongList();
    lastSavePlaySongListLen = playSongList.value.length;
    console.log('get playSongList history', lastSavePlaySongListLen);
};

// 格式化song添加额外的字段兼容播放器字段
export const fixSongForPlayer = (song) => {
    song.type = `customGetSongURL`; // 自定义获取播放地址函数
    song.artist = song.SingerName; // 歌手
    song.name = song.FileName; // 歌名
    song.url = ``; // 播放地址空, 在customGetAudioURL函数从新转为正确的地址
    song.urlHash = ``; // 存 url 对应的 hash
    song.cover = song.Image ? song.Image.replaceAll(`{size}`, 480) : ''; // 图片
    return song;
};
