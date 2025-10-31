import { ref, watch } from 'vue';

const regex = /\[(\d{2}):(\d{2})\.(\d{2})\](.+)/;

let currLyricIdx = 0;
let currLyricArr = [];
let currLyricPlayer = null;
let lyricElm = null;
let autoShowLyricTimer = 0;

const saveSelectLyricColorHistory = () => {
    localStorage.setItem(`selectLyricColor`, selectLyricColor.value);
};

const loadSelectLyricColorHistory = () => {
    const old = localStorage.getItem(`selectLyricColor`);
    return typeof old == 'string' && old.length === 7 ? old : '#ff0097';
};

export const isSettingOpenLyric = ref(true);
export const selectLyricColor = ref(loadSelectLyricColorHistory());

const showTextToElement = (text) => {
    if (!lyricElm || !lyricElm.parentNode) {
        lyricElm = document.getElementsByClassName('aplayer-lrc-current')[0];
        if (lyricElm) {
            lyricElm.style.fontSize = '16px';
            lyricElm.style.color = selectLyricColor.value;
            lyricElm.innerText = '歌词准备就绪';
        }
    }

    if (lyricElm) {
        lyricElm.innerText = text;
    }
};

const startTimerOnce = () => {
    if (autoShowLyricTimer > 0) {
        return;
    }

    autoShowLyricTimer = setInterval(() => {
        if (currLyricIdx >= currLyricArr.length) {
            // showTextToElement('歌词展示结束');
            console.log('歌词展示结束');
            return;
        }

        if (!isSettingOpenLyric.value) {
            showTextToElement('用户关闭歌词显示');
            return;
        }

        const show = currLyricArr[currLyricIdx];
        const currentMs = currLyricPlayer?.audio?.currentTime * 1000;
        if (currentMs >= show.ms) {
            currLyricIdx++;
            showTextToElement(show.lyric);
        }
    }, 200);
    console.log('启动歌词展示定时器=%d', autoShowLyricTimer);
};

export const parseLyric = (str) => {
    if (!str) {
        return [];
    }

    str = str.replaceAll('\r\n', '\n');
    const ret = [];
    const arr = str.split('\n');
    for (let i = 0; i < arr.length; i++) {
        const s = arr[i];
        const match = s.match(regex);
        if (match && match.length >= 5) {
            const lyric = match[4];
            const minutes = parseInt(match[1]);
            const seconds = parseInt(match[2]);
            const milliseconds = parseInt(match[3]);
            const ms = (minutes * 60 + seconds) * 1000 + milliseconds * 10;
            ret.push({ ms, lyric });
        }
    }
    return ret;
};

export const autoShowLyric = (arr, aPlayer) => {
    if (!arr || arr.length === 0 || !aPlayer) {
        currLyricIdx = 0;
        currLyricArr = [];
        showTextToElement('参数异常不展示歌词');
        return;
    }

    currLyricIdx = 0;
    currLyricArr = arr;
    currLyricPlayer = aPlayer;

    startTimerOnce();
};

// 用户选择颜色重新染色
watch(selectLyricColor, () => {
    lyricElm = null;
    saveSelectLyricColorHistory();
});
