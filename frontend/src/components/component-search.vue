<template>
  <div class="search-bar">
    <el-input
      v-model="searchKeyword"
      placeholder="搜索音乐"
      maxlength="40"
      @keyup.enter="querySearch"
    >
      <template #append>
        <el-button @click="querySearch">
          <el-icon><SearchIcon /></el-icon>
        </el-button>
      </template>
    </el-input>
  </div>

  <el-dialog
    v-model="isShowSearchResult"
    title=""
    width="75%"
  >
    <el-table
      v-loading="isLoadingSearchResult"
      :data="searchResults"
      stripe
      style="width: 100%"
    >
      <el-table-column
        prop="FileName"
        label="歌曲"
        min-width="3"
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
        min-width="1"
      >
        <template #default="scope">
          <el-row :gutter="1">
            <el-col :span="8">
              <el-button
                circle
                size="small"
                type="primary"
                @click="addPlayList(scope.row)"
              >
                +
              </el-button>
            </el-col>

            <el-col :span="8">
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
              :span="8"
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
        :current-page="searchPage"
        :page-size="pageSize"
        size="small"
        layout="pager, total"
        :total="searchResultsTotal"
        @current-change="changePage"
      />
    </div>
  </el-dialog>
</template>

<script setup>
import { SearchMusic } from '@/wailsjs/go/main/App.js';
import { computed, ref } from 'vue';
import { CopyDocument, Search as SearchIcon,Download } from '@element-plus/icons-vue';
import { playSongList,
    userInfo,
    copySongURL,
    fixSongForPlayer,
    downloadSong
} from './sharedVar.js';
import { ElMessage } from 'element-plus';

const pageSize = 10;
const searchKeyword = ref('张学友');
const searchPage = ref(1);
const searchResultsTotal = ref(1);
const isLoadingSearchResult = ref(false);

/** @type {import('vue').Ref<api.Song[]>} */
const searchResults = ref([]);

const isShowSearchResult = computed({
    get: () => searchResults.value.length > 0,
    set: (newVisible) => {
        if (!newVisible) {
            searchResults.value = [];
            searchPage.value = 1;
        }
    }
});

const changePage = (v) => {
    searchPage.value = v;
    querySearch();
};

const querySearch = async () => {
    if (searchKeyword.value && searchPage.value>0) {
        console.log(`querySearch`,searchKeyword.value,searchPage.value);
        isLoadingSearchResult.value = true;
        const resp = await SearchMusic(userInfo.dfid, userInfo.userid, userInfo.token, searchKeyword.value, searchPage.value);
        isLoadingSearchResult.value = false;
        if (resp.errMsg) {
            ElMessage.warning(resp.errMsg);
            return;
        }

        searchResults.value = resp.data.lists;
        searchResultsTotal.value = resp.data.total;
        console.log('搜索结果数量 page=%d num=%d',searchPage.value,searchResults.value.length);
    }
};

const isRepeatPlatList = (song) => {
    for (let i = 0; i < playSongList.value.length; i++) {
        const s = playSongList.value[i];
        if (s.FileHash===song.FileHash) {
            return true;
        }
    }
    return false;
};

const addPlayList = (song) => {
    // 不重复就添加播放列表
    if (isRepeatPlatList(song)===false) {
        fixSongForPlayer(song);
        playSongList.value.push(song);
    }

    // 移除搜索结果
    for (let i = 0; i < searchResults.value.length; i++) {
        const s = searchResults.value[i];
        if (s.FileHash===song.FileHash) {
            searchResults.value.splice(i,1);
            break;
        }
    }

    // 没有数据就刷新下一页
    if (searchResults.value.length===0) {
        changePage(searchPage.value+1);
    }
};
</script>

<style>
.search-bar {
    flex-grow: 1;
    max-width: 25vw;
    margin: 0 auto;
}
</style>
