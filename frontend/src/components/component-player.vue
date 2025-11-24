<template>
  <el-main class="main">
    <div
      id="aPlayer"
      style="width:60vw;margin:0 auto;"
    />
  </el-main>
  <el-footer class="footer">
    <div
      class="playlist-icon"
      @click="showPlaylist"
    >
      <el-icon class="playListAnimate">
        <Document />
      </el-icon>
    </div>
  </el-footer>

  <el-dialog
    v-model="isShowPlaylistPage"
    title=""
    width="70%"
    center
  >
    <el-table
      :data="showPlaylistData"
      style="max-width: 80vw;"
    >
      <el-table-column
        prop="idx"
        label="序号"
        min-width="1"
      />
      <el-table-column
        prop="FileName"
        label="歌曲"
        min-width="5"
        show-overflow-tooltip
      />
      <el-table-column
        prop="SingerName"
        label="歌手"
        min-width="3"
        show-overflow-tooltip
      />
      <el-table-column
        label="操作"
        min-width="2"
      >
        <template #default="scope">
          <el-row :gutter="1">
            <el-col :span="6">
              <el-button
                circle
                size="small"
                type="primary"
                @click="playSong(scope.row)"
              >
                <el-icon><VideoPlay /></el-icon>
              </el-button>
            </el-col>

            <el-col :span="6">
              <el-button
                circle
                size="small"
                type="danger"
                @click="removeSong(scope.row)"
              >
                <el-icon><Delete /></el-icon>
              </el-button>
            </el-col>

            <el-col :span="6">
              <el-button
                circle
                size="small"
                type="success"
                @click="copySongURL(scope.row)"
              >
                <el-icon><CopyDocument /></el-icon>
              </el-button>
            </el-col>

            <el-col
              :span="6"
            >
              <el-button
                circle
                size="small"
                type="success"
                color="#0c8918"
                @click="downloadSong(scope.row)"
              >
                <el-icon><Download /></el-icon>
              </el-button>
            </el-col>
          </el-row>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container">
      <el-pagination
        :current-page="showPlaylistPage"
        :page-size="pageSize"
        size="small"
        layout="pager, total"
        :total="showPlaylistTotal"
        @current-change="changePlaylistPage"
      />
    </div>
  </el-dialog>
</template>

<script setup>
import { computed, onMounted, ref, toRaw, watch } from 'vue';
import APlayer from 'aplayer';
import 'aplayer/dist/APlayer.min.css';
import {
    Document,
    CopyDocument,
    Delete,
    Download,
    VideoPlay
} from '@element-plus/icons-vue';
import {
    playSongList,
    copySongURL,
    userInfo,
    getUserSelectQuality,
    savePlayListDelay,
    getSongHash,
    downloadSong
} from './sharedVar.js';
import { GetSongURL } from '@/wailsjs/go/main/App';
import { autoShowLyric, parseLyric } from '@/src/components/lyric.js';
import { addBgColor } from '@/src/components/bgColor.js';
import { ElMessage } from 'element-plus';
import { WindowSetTitle } from '@/wailsjs/runtime/runtime.js';

const pageSize = 10;

/** @type {APlayer} */
let aPlayer = null;
// 播放器获取歌曲地址连续出错次数,多次失败后暂停播放
let aPlayerGetAudioErrCount = 0;
// 是否展示自定义播放列表页
const isShowPlaylistPage = ref(false);
// 当前自定义播放列表查看的是第几页
const showPlaylistPage = ref(1);
// 自定义播放列表当前页数据, 一页一页的查看
const showPlaylistData = ref([]);
// 自定义播放列表一共有多少数据,用于分页计算
const showPlaylistTotal = computed({
    get:() => playSongList.value.length,
    set:() => {}
});

// 获取APlayer的播放器列表
const getAPlayerList = () => {
    return aPlayer.list;
};

// 展示自定义播放列表函数
const showPlaylist = () => {
    changePlaylistPage(1);
    isShowPlaylistPage.value = true;
};

// 自定义播放列表翻页函数
const changePlaylistPage = (page) => {
    const s = Math.max(0, page - 1) * pageSize;

    if (s >= playSongList.value.length) {
        showPlaylistData.value = [];

        return;
    }

    showPlaylistPage.value = page;
    const showPageList = playSongList.value.slice(s, s + pageSize);

    for (let i = 0; i < showPageList.length; i++) {
        showPageList[i]['idx'] = s + i + 1;
    }
    showPlaylistData.value = showPageList;
};

// 查找播放器列表的歌曲下标,用于播放器操作删除和播放
const findIdxFromPlayerList = (song) => {
    /** @type {api.Song[]} */
    const audios = getAPlayerList().audios;

    for (const songKey in audios) {
        if (audios[songKey].FileHash === song.FileHash) {
            return songKey;
        }
    }

    return -1;
};

// 切换播放此歌曲
const playSong = (song) => {
    const idx = findIdxFromPlayerList(song);

    console.log(`playSong`, song, idx);
    if (idx >= 0) {
        getAPlayerList().switch(idx);
    }
};

// 删除播放列表
const removeSong = (song) => {
    removePlaySongList(song);
    removeShowPlaylistPage(song);
};

// 删除自定义列表的数据
const removePlaySongList = (song) => {
    for (let i = 0; i < playSongList.value.length; i++) {
        const curSong = playSongList.value[i];

        if (curSong.FileHash === song.FileHash) {
            playSongList.value.splice(i, 1);
            break;
        }
    }
};

// 删除自定义列表的当前分页数据
const removeShowPlaylistPage = (song) => {
    for (let i = 0; i < showPlaylistData.value.length; i++) {
        const curSong = showPlaylistData.value[i];

        if (curSong.FileHash === song.FileHash) {
            showPlaylistData.value.splice(i, 1);
            break;
        }
    }
};


// 把纪录播放列表同步到播放器
const syncPlaySongListToPlayer = () => {
    console.log('start syncPlaySongListToPlayer', playSongList.value.length);
    const hasMap = new(Map);
    const playerList = getAPlayerList();
    const playerAudios = playerList.audios || [];

    for (let i = 0; i < playSongList.value.length; i++) {
        const song = playSongList.value[i];

        hasMap.set(song.FileHash, toRaw(song));
    }

    /**
     * 播放器删除使用下标,如果正向删除下标会对不上
     * 这里使用反向遍历删除
     */
    for (let i = playerAudios.length - 1; i >= 0; i--) {
        /** @type {api.Song} */
        const hash = playerAudios[i].FileHash;

        if (!hasMap.has(hash)) {
            playerList.remove(i); // 纪录播放器要删除的
        } else {
            hasMap.delete(hash);
        }
    }

    // hasMap 剩下是需要新增的
    hasMap.forEach((song) => {
        playerList.add(song);
    });

    // 自动备份
    savePlayListDelay();
};

/*
* APlayer bug: Illegal customType: customGetSongURL
* 是因为播放器不支持 async function, 请修改源代码并清理vite缓存:
* "[object Function]" === Object.prototype.toString.call (xxx)
* 为
* "function" === typeof (xxx)
*/

// 动态计算歌曲播放地址和歌词
const customGetSongURL = async function (audioElement, audio, player) {
    console.log('1 执行更新获取歌曲地址', audio.name);
    const quality = getUserSelectQuality();
    const hash = getSongHash(audio, quality);

    if (audio.url === '' || hash !== audio.urlHash) {
        console.log('2 执行更新获取歌曲地址', audio.name, quality, hash);
        const resp = await GetSongURL(userInfo.dfid, userInfo.userid, userInfo.token, hash);

        console.log('3 执行更新获取歌曲地址', audio.name, resp);
        let errMsg = ``;

        if (resp.errMsg) {
            aPlayerGetAudioErrCount++;
            errMsg = '获取播放地址错误:' + resp.errMsg;
        } else if (!resp.data || resp.data.length === 0) {
            aPlayerGetAudioErrCount++;
            errMsg = '获取播放地址失败,切换到下一首歌曲';
        }

        if (aPlayerGetAudioErrCount > 5) {
            ElMessage.error(`获取地址多次出错,已暂停播放`);
            try {
                player.pause();
            } catch (e) {
                console.log('call player.pause err', e);
            }

            return;
        }

        if (errMsg !== '') {
            ElMessage.warning(errMsg);
            if (player && typeof player.skipForward === 'function') {
                try {
                    player.skipForward();
                } catch (e) {
                    console.log('call player.skipForward err', e);
                }
            }

            return;
        }

        // 成功
        aPlayerGetAudioErrCount = 0;
        audio.url = resp.data[0];
        audio.urlHash = hash;
        audio.lrc = parseLyric(resp.lyric);

        const size = resp.size > 0 ? (resp.size / 1024 / 1024).toFixed(2) : '0';

        console.log('4 执行更新获取歌曲地址', audio.name, audio.url, size + ' MB');
    }

    // 动态更新地址和歌词
    audioElement.src = audio.url;
    autoShowLyric(audio.lrc, player);
    addBgColor('app');
    WindowSetTitle(audio.name);
    if (player && typeof player.play === 'function') {
        try {
            player.play();
        } catch (e) {
            console.log('call player.play err', e);
        }
    }
};

// 自动同步播放列表到播放器
watch(playSongList, syncPlaySongListToPlayer, { deep: true });

// 启动
onMounted(() => {
    aPlayer = new APlayer({
        container: document.getElementById('aPlayer'),
        mini: false,
        autoplay: false,
        theme: '#ff0097',
        loop: 'all',
        order: 'list',
        preload: 'auto',
        volume: 1,
        mutex: true,
        listFolded: true,
        lrcType: 1, // 开启这个功能只为了有个展示歌词的<p>元素,歌词展示不支持动态加载需要额外实现
        storageName:`APlayerConfig`,
        listMaxHeight:'75vh',

        /**
         * @typedef {Object} playerAudio
         * @property {string} artist
         * @property {string} name
         * @property {string} url
         * @property {string} cover
         */

        /** @type playerAudio[] */
        audio: [],
        customAudioType: {
            customGetSongURL: customGetSongURL
        }
    });
});
</script>

<style>
.playlist-icon {
    position: fixed;
    bottom: 20px;
    right: 20px;
    font-size: 24px;
    cursor: pointer;
}

.playListAnimate{
    animation: pulse 1s ease-in-out infinite;
}
/* 定义脉冲动画 */
@keyframes pulse {
    0% {
        transform: scale(1);
    }
    50% {
        transform: scale(1.2);
    }
    100% {
        transform: scale(1);
    }
}
</style>
