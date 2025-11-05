<template>
  <el-dialog
    v-model="isShowAgreement"
    title="使用者条款"
    width="70%"
    :modal="true"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    :show-close="false"
  >
    <p>1. 本程式是酷狗第三方客户端，并非酷狗官方，需要更完善的功能请下载官方客户端体验.</p>
    <p>2. 本项目仅供学习交流使用，您在使用过程中应尊重版权，不得用于商业或非法用途。</p>
    <p>3. 在使用本项目的过程中，可能会产生版权内容。 本项目不拥有这些版权内容的所有权。 为了避免侵权行为，您需在 24 小时内清除本项目产生的版权内容。</p>
    <p>4. 本专案的开发者不对因使用或无法使用本项目所导致的任何损害承担责任，包括但不限于资料遗失、停工、电脑故障或其他经济损失。</p>
    <p>5. 您不得在违反当地法律法规的情况下使用本项目。 因违反法律法规所导致的任何法律后果由使用者承担。</p>
    <p>6. 本计划仅用于技术探索和研究，不接受任何商业合作、广告或捐赠。 若官方音乐平台对此专案存有疑虑，可随时联络开发者移除相关内容。</p>
    <p style="color:red; font-weight: bold;">
      同意继续使用本项目，您即接受以上条款声明内容。
    </p>
    <div style="text-align: right;">
      <el-button
        type="success"
        round
        @click="userStatusWithAgreement(true)"
      >
        我已知悉并同意以上条款
      </el-button>
      <el-button
        type="danger"
        round
        @click="userStatusWithAgreement(false)"
      >
        不同意
      </el-button>
    </div>
  </el-dialog>
</template>

<script setup>
import { ref } from 'vue';
import { ElMessage } from 'element-plus';

const isShowAgreementKey = `isShowAgreement`;
const isShowAgreement = ref(localStorage.getItem(isShowAgreementKey)!==`1`);

const userStatusWithAgreement = (isAgree) => {
    localStorage.setItem(isShowAgreementKey,  isAgree ? `1`:`0`);
    isShowAgreement.value = !isAgree;
    if (!isAgree) {
        ElMessage.error(`同意使用者條款方可使用本软件`);
    }
};
</script>

<style></style>
