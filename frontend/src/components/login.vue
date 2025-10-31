<template>
  <div class="user-avatar">
    <el-dropdown>
      <el-avatar
        :size="40"
        :src="userInfo.pic"
      />
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item
            v-if="userInfo.token===''"
            @click="showLoginDialog"
          >
            <el-icon><User /></el-icon>登录
          </el-dropdown-item>

          <el-dropdown-item
            v-if="userInfo.token!==''"
            @click="showSetting"
          >
            <el-icon><Setting /></el-icon>设定
          </el-dropdown-item>

          <el-dropdown-item>
            <a
              href="https://github.com/DeeCen/yourMusic"
              target="_blank"
              style="text-decoration: none; color: inherit; display: flex; align-items: center;"
            ><el-icon><Link /></el-icon>GitHub</a>
          </el-dropdown-item>
          <el-dropdown-item @click="showAbout">
            <el-icon><InfoFilled /></el-icon>关于
          </el-dropdown-item>
          <el-dropdown-item
            v-if="userInfo.token!==''"
            @click="logout"
          >
            <el-icon><CircleClose /></el-icon>退出登录
          </el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>

  <el-dialog
    v-model="isShowAbout"
    title="免责声明"
    width="75%"
  >
    <p>0. 本软件是酷狗第三方客戶端，并非酷狗官方，需要更完善的功能请下载官方客户端体验.</p>
    <p>1. 本项目仅供学习使用，请尊重版权，请勿利用此项目从事商业行为及非法用途！</p>
    <p>2. 使用本项目的过程中可能会产生版权数据。对于这些版权数据，本项目不拥有它们的所有权。为了避免侵权，使用者务必在 24 小时内清除使用本项目的过程中所产生的版权数据。</p>
    <p>3.由于使用本项目产生的包括由于本协议或由于使用或无法使用本项目而引起的任何性质的任何直接、间接、特殊、偶然或结果性损害（包括但不限于因商誉损失、停工、计算机故障或故障引起的损害赔偿，或任何及所有其他商业损害或损失）由使用者负责。</p>
    <p>4. 禁止在违反当地法律法规的情况下使用本项目。对于使用者在明知或不知当地法律法规不允许的情况下使用本项目所造成的任何违法违规行为由使用者承担，本项目不承担由此造成的任何直接、间接、特殊、偶然或结果性责任。</p>
    <p>5. 音乐平台不易，请尊重版权，支持正版。</p>
    <p>6. 本项目仅用于对技术可行性的探索及研究，不接受任何商业（包括但不限于广告等）合作及捐赠。</p>
    <p>7. 如果官方音乐平台觉得本项目不妥，可联系本项目更改或移除。</p>
  </el-dialog>

  <el-dialog
    v-model="isLoginDialogVisible"
    title="登录"
    width="70%"
  >
    <el-tabs
      class="login-tabs"
      type="border-card"
      :stretch="true"
    >
      <el-tab-pane>
        <template #label>
          <span><el-icon><Iphone /></el-icon> 手机号登录</span>
        </template>
        <el-form>
          <el-form-item>
            <el-input
              v-model="mobile"
              :prefix-icon="Iphone"
              placeholder="请输入手机号"
              maxlength="11"
            />
          </el-form-item>
          <el-form-item>
            <el-row :gutter="10">
              <el-col :span="16">
                <el-input
                  v-model="code"
                  :prefix-icon="Message"
                  placeholder="请输入验证码"
                  maxlength="6"
                />
              </el-col>
              <el-col :span="8">
                <el-button
                  :disabled="sendCodeCD"
                  @click="sendMobileCode"
                >
                  {{ sendCodeTips }}
                </el-button>
              </el-col>
            </el-row>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              style="width: 100%;"
              @click="loginByMobile"
            >
              登录
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <!--<el-tab-pane>
                <template #label>
                    <span
                        ><el-icon><ChatDotRound /></el-icon> 微信登录</span
                    >
                </template>
                <div style="text-align: center">
                    <p>请使用微信扫描二维码登录</p>
                    <div style="width: 200px; height: 200px; background-color: #f0f0f0; margin: 0 auto;"></div>
                </div>
            </el-tab-pane>-->
    </el-tabs>
  </el-dialog>

  <el-dialog
    v-model="isShowSetting"
    title="设定"
    style="max-width: 50vw;"
  >
    <el-form
      label-position="left"
      label-width="auto"
    >
      <el-form-item
        label="音质"
        label-position="right"
      >
        <el-radio-group
          v-model="userSettingQuality"
          size="default"
          fill="#6cf"
          @change="userChangeQuality"
        >
          <el-radio-button
            label="128"
            value="128"
          />
          <el-radio-button
            label="320"
            value="320"
          />
          <el-radio-button
            label="flac"
            value="flac"
          />
          <!-- <el-radio-button label="HiRes" value="HiRes" />-->
        </el-radio-group>
      </el-form-item>

      <el-form-item
        label="歌词显示"
        label-position="right"
      >
        <el-switch
          v-model="isSettingOpenLyric"
          style="--el-switch-on-color: #13ce66;"
        />
      </el-form-item>

      <el-form-item
        v-show="isSettingOpenLyric"
        label="歌词颜色"
        label-position="right"
      >
        <el-color-picker
          v-model="selectLyricColor"
          size="small"
        />
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import { computed, ref } from 'vue';
import {
    User,
    InfoFilled,
    Link,
    Iphone,
    Message,
    CircleClose,
    Setting
} from '@element-plus/icons-vue';
import {
    userInfo,
    updateUserInfoToLocalStorage,
    readUserInfoFromLocalStorage,
    removeUserInfoToLocalStorage,
    getUserSelectQuality,
    setUserSelectQuality
} from './sharedVar.js';
import { ElMessage } from 'element-plus';
import { LoginByMobile, LoginByToken,SendMobileCode } from '@/wailsjs/go/main/App.js';
import { isSettingOpenLyric,selectLyricColor } from './lyric.js';

const isLoginDialogVisible = ref(false);
const isShowAbout = ref(false);
const mobile = ref(``);
const code = ref(``);
const sendCodeCD = ref(0);
const isShowSetting = ref(false);
const userSettingQuality = ref(getUserSelectQuality());
const sendCodeTips = computed({
    get: () => '获取验证码' + (sendCodeCD.value>0 ? `(${sendCodeCD.value})`:'' ) ,
    set: () => {}
});

const startCD = () => {
    sendCodeCD.value = 10;
    const timer = setInterval(() => {
        if (sendCodeCD.value<=0) {
            clearInterval(timer);
            return;
        }

        sendCodeCD.value--;
    },1000);
};

const showLoginDialog = () => {
    isLoginDialogVisible.value = true;
};

const showAbout = () => {
    isShowAbout.value = true;
};
const loginByMobile = async () => {
    if (mobile.value.length !== 11) {
        ElMessage.warning(`请输入11位手机号`);
        return;
    }

    if (code.value.length !== 6) {
        ElMessage.warning(`请输入6位验证码`);
        return;
    }

    const resp = await LoginByMobile(mobile.value, code.value);
    if (resp.errMsg) {
        ElMessage.warning(resp.errMsg);
        return;
    }

    updateUserInfoToLocalStorage(resp.data);
    if (userInfo.token==='') {
        ElMessage.success( '登录失败' );
        return;
    }
    ElMessage.success( '登录成功' );
    isLoginDialogVisible.value = false;
};

const getYYYYMMDD = () => {
    const date = new Date();
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}${month}${day}`;
};

const lastAutoGetVipDayKey = `lastAutoGetVipDayKey`;
const getLastAutoGetVipDay = () => {
    return String(localStorage.getItem(lastAutoGetVipDayKey));
};

const setLastAutoGetVipDay = () => {
    localStorage.setItem(lastAutoGetVipDayKey, getYYYYMMDD());
};

const isAutoGetVipDay = () => {
    const today = getYYYYMMDD();
    return getLastAutoGetVipDay()!==today;
};

const loginByToken = async (autoGetVip) => {
    console.log('start loginByToken');
    if (userInfo.token==='' || userInfo.userid===`0` || userInfo.userid===``) {
        console.log('stop loginByToken');
        return;
    }

    console.log('run loginByToken');
    const resp = await LoginByToken(userInfo.dfid,userInfo.userid,userInfo.token,autoGetVip?1:0);
    updateUserInfoToLocalStorage(resp.data);
};

const sendMobileCode = async () => {
    if (mobile.value.length !== 11) {
        ElMessage.warning(`请输入11位手机号`);
        return;
    }

    if (sendCodeCD.value>0) {
        return;
    }

    const msg = await SendMobileCode(mobile.value);
    if (msg) {
        ElMessage.warning(msg);
        return;
    }

    startCD();
    ElMessage.success(`验证码已发送`);
};

const showSetting = () => {
    isShowSetting.value = true;
};

const userChangeQuality = () => {
    setUserSelectQuality(userSettingQuality.value);
    ElMessage.success(`切换音质成功,下一首歌生效`);
};

const logout = () => {
    removeUserInfoToLocalStorage();
};

const isGetVip = isAutoGetVipDay();
if (readUserInfoFromLocalStorage() && isGetVip) {
    loginByToken(true);//每天只刷新一次token
    setLastAutoGetVipDay();
}
if (userInfo.token==='') {
    showLoginDialog();
}
</script>

<style>
.login-tabs{
  width: 100%;
  text-align: center;
}
</style>
