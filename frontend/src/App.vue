<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">
        <div class="headerDiv">
          <Search />
          <Login />
        </div>
      </el-header>

      <Player />
    </el-container>

    <Agreement />
  </div>
</template>

<script setup>
import Search from './components/component-search.vue';
import Login from './components/component-login.vue';
import Agreement from './components/component-agreement.vue';
import Player from './components/component-player.vue';
import { initSharedVarOnce } from './components/sharedVar.js';
import { addBgColor } from '@/src/components/bgColor.js';

initSharedVarOnce().then(() => {
    console.log(`initSharedVar success`);
}).catch((e) => {
    console.log(`initSharedVar error`, e);
});
addBgColor('app');
</script>

<style>
.common-layout {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}

.header{
    background-color: rgba(255,255,255, .95);
    box-shadow: 0 2px 10px #0000008a;
}

.headerDiv{
    display: flex;
    margin: 15px 5vw 0;
}
</style>
