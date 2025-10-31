<br />
<p align="center">
    <img src="https://github.com/DeeCen/yourMusic/raw/main/docs/logo.png" alt="Logo" width="156" height="156">
  <h2 align="center" style="font-weight: 600">yourMusic</h2>
  <p align="center">
    一款开源简洁的酷狗概念版第三方客户端, written by go & wails.
    <br />
    <a href="https://github.com/DeeCen/yourMusic" target="blank"><strong>🌎 GitHub仓库</strong></a>&nbsp;&nbsp;|&nbsp;&nbsp;
    <a href="https://github.com/DeeCen/yourMusic/releases" target="blank"><strong>📦️ 下载</strong></a>
    <br />
    <br />
  </p>


</p>

![images](https://github.com/DeeCen/yourMusic/raw/main/docs/1.png)
![images](https://github.com/DeeCen/yourMusic/raw/main/docs/2.png)

## ❤️ 前言

早在`虾米`音乐还在的时候, 我只认一个音乐APP

哎, `虾米`没了后我就不再用手机听歌了

.....

.....

.....

很多年之后, 朋友推荐我: `有一款音乐APP做的不错, 叫[酷狗概念版]`

噢, 我又开始用手机听歌了

.....

.....

.....

直到2025年, 因为换了手机然后重装APP后发现: `上面很多广告, 手机也开始变烫`

好吧, 我又不再用手机听歌了

.....

后来无意间发现了一款叫<a href="https://github.com/MoeKoeMusic/MoeKoeMusic" target="blank"><strong>MoeKoeMusic</strong></a>的软件, 该软件使用了electron开发并对接酷狗的API, 真的很赞👍

但是, 它包体很大(200+M)&我的网速很慢, 而我对音乐的要求又很简单: `听个响`



这时, 我想起了一句来源于<a href="https://github.com/tinygo-org/tinygo" target="blank"><strong>tinygo</strong></a>的话:

if [Python](https://micropython.org/) can run on microcontrollers, then certainly [Go](https://golang.org/) should be able to run on even lower level micros.



于是, 我"抄袭"了<a href="https://github.com/MoeKoeMusic/MoeKoeMusic" target="blank"><strong>MoeKoeMusic</strong></a>`(包括这个文档)`




## ✨ 特性

- ✅ 使用 Vue.js 开发

- 🔴 酷狗概念版账号登录（手机验证码方式）

- 🔒 数据安全, 所有数据只保存到 localStorage

- 📃 支持歌曲搜索, 音质切换, 歌词显示

- 🚫 无任何广告

- 🔗 官方服务器直连, 无任何第三方 API

- ✔️ 每日自动领取VIP, 登录就是VIP

- 🎨 随机多彩背景

- ⚙️ 多平台支持 (非windows平台需自行构建)

- <img src="https://go.dev/images/favicon-gopher.png"  width="20"/> 得益于go的强大, 打包后本软件只有1个文件且不足20M, 感谢go!

  

## 📦️ 安装

### 方式1. windows平台可下载后直接使用

访问本项目的 [Releases](https://github.com/DeeCen/yourMusic/releases) 页面下载exe文件直接使用



### 方式2. Mac/Linux平台需使用者自行构建
0. 环境准备 
   [nodejs](https://nodejs.org)
   [go](https://go.dev/dl/)
   [wails](https://wails.io/)

1. 克隆本仓库
```sh
git clone https://github.com/DeeCen/yourMusic.git
```

2. 安装前端依赖
```sh
cd yourMusic/frontend
npm i
```

3. 使用 yourMusic/docs/APlayer.min.js 替换 frontend/node_modules/aplayer/dist/APlayer.min.js
  [why?](https://github.com/DIYgod/APlayer/issues/801)


4. 编译项目
```sh
cd yourMusic
go mod tidy
wails build
# 成功后 yourMusic/build/bin/ 会生成可执行文件
```



## ⭐ 支持项目

如果您觉得这个项目对您有帮助, 欢迎给我一个 Star 让我知道真有人在用.



## ✅ 反馈

如有任何问题或建议，欢迎提交 issue 或 pull request。



## ⚠️ 免责声明
0. 本程序是酷狗第三方客户端，并非酷狗官方，需要更完善的功能请下载官方客户端体验.
1. 本项目仅供学习使用，请尊重版权，请勿利用此项目从事商业行为及非法用途！
2. 使用本项目的过程中可能会产生版权数据。对于这些版权数据，本项目不拥有它们的所有权。为了避免侵权，使用者务必在 24 小时内清除使用本项目的过程中所产生的版权数据。
3. 由于使用本项目产生的包括由于本协议或由于使用或无法使用本项目而引起的任何性质的任何直接、间接、特殊、偶然或结果性损害（包括但不限于因商誉损失、停工、计算机故障或故障引起的损害赔偿，或任何及所有其他商业损害或损失）由使用者负责。        
1. 禁止在违反当地法律法规的情况下使用本项目。对于使用者在明知或不知当地法律法规不允许的情况下使用本项目所造成的任何违法违规行为由使用者承担，本项目不承担由此造成的任何直接、间接、特殊、偶然或结果性责任。    
2. 音乐平台不易，请尊重版权，支持正版。
3. 本项目仅用于对技术可行性的探索及研究，不接受任何商业（包括但不限于广告等）合作及捐赠。
4. 如果官方音乐平台觉得本项目不妥，可联系本项目更改或移除。
            

## 📜 开源许可

本项目仅供个人学习研究使用，禁止用于商业及非法用途。

基于 [GNU General Public License v2.0 (GPL-2.0)](https://github.com/DeeCen/yourMusic/blob/main/LICENSE) 许可进行开源。



## 👍 灵感来源

1 [MoeKoeMusic/MoeKoeMusic](https://github.com/MoeKoeMusic/MoeKoeMusic)
2 [MakcRe/KuGouMusicApi](https://github.com/MakcRe/KuGouMusicApi) 