import { ref, watch } from 'vue';

const regex = /\[(\d{2}):(\d{2})\.(\d{2})\](.+)/;
let currLyricArr = [];
let currLyricPlayer = null;
let lyricElm = null;
let autoShowLyricTimer = 0;
let lastShowLyricText = '';


const saveSelectLyricColorHistory = () => {
    localStorage.setItem(`selectLyricColor`, selectLyricColor.value);
};

const findShowLyricText = (currentMs) => {
    currentMs += currLyricDelayMs.value;
    let left = 0;
    let right = currLyricArr.length - 1;

    while (left <= right) {
        const mid = left + Math.trunc((right - left) / 2);
        const ms = currLyricArr[mid]['ms'];
        const nextMs = typeof currLyricArr[mid + 1] !== 'undefined' ? currLyricArr[mid + 1]['ms'] : 4294967296;

        switch (true) {
            case ms <= currentMs && currentMs < nextMs:
                return currLyricArr[mid].lyric;

            case currentMs >= nextMs:left = mid + 1; break;
            default:right = mid - 1; break;
        }

    }

    return '';
};

const loadSelectLyricColorHistory = () => {
    const old = localStorage.getItem(`selectLyricColor`);

    return typeof old == 'string' && old.length === 7 ? old : '#ff0097';
};

export const currLyricDelayMs = ref(0);

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

    if (lyricElm && lastShowLyricText !== text) {
        lyricElm.innerText = text;
        lastShowLyricText = text;
    }
};

const startTimerOnce = () => {
    if (autoShowLyricTimer > 0) {
        return;
    }

    autoShowLyricTimer = setInterval(() => {
        if (!isSettingOpenLyric.value) {
            showTextToElement('用户关闭歌词显示');

            return;
        }

        const currentMs = currLyricPlayer?.audio?.currentTime * 1000;
        const showText = findShowLyricText(currentMs);

        showTextToElement(showText);
    }, 500);
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
            const minutes = parseInt(match[1]) || 0;
            const seconds = parseInt(match[2]) || 0;
            const milliseconds = parseInt(match[3]) || 0;
            const ms = (minutes * 60 + seconds) * 1000 + milliseconds * 10;

            ret.push({ ms, lyric });
        }
    }

    return ret;
};

export const autoShowLyric = (arr, aPlayer) => {
    currLyricArr = [];
    currLyricDelayMs.value = 0;
    if (!arr || arr.length === 0 || !aPlayer) {
        showTextToElement('参数异常不展示歌词');

        return;
    }

    currLyricArr = arr;
    currLyricPlayer = aPlayer;
    startTimerOnce();
};

// 用户选择颜色重新染色
watch(selectLyricColor, () => {
    lyricElm = null;
    saveSelectLyricColorHistory();
});
