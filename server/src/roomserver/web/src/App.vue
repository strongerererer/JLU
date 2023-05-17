<script >

import axios from 'axios';
import VueTypedJs from 'vue-typed-js'


export default {
        data() {
        return {
            message: '#微信读书QA',
            items: [],
            roomid: 0,
            chatid: 0,
            token: "",
            chat: "",
            autoscroll: true,
            nickname: this.randName(),
            headurl: "../public/heads/head1.png",
            joinModal: null,
            headModal: null,
            typedAni: false
        }
        },
        
        components: {
          VueTypedJs
        },

        created() {
          console.log(2);
                      
        },

        mounted() {

          this.joinModal = new bootstrap.Modal(document.getElementById('staticBackdrop'));
          this.headModal = new bootstrap.Modal(document.getElementById('modalHead'));

          console.log(1);

          this.start();
        },
        beforeUnmount() {
            clearInterval(this.timer);
        },
    
        methods: {
            async fetchData() {

                if (this.roomid == 0) return;
                
                try {

                    const response = await axios.get('/room_chatlist?roomid=' + this.roomid + "&chatid=" + this.chatid + "&token=" + this.token);
                    this.chatid = response.data.data["chatId"];

                    const addlist = response.data.data["chatList"];
                    if (addlist == undefined || addlist.length == 0) {
                      return
                    }
                    
                    for (var item of addlist) {
                      item["ani"] = this.typedAni;
                      this.items.push(item);
                    }
                    
                    this.$nextTick(() => {
                      this.typedAni = true;
                      if (this.autoscroll) {
                        this.$refs.msg_end.scrollIntoView({ block: "end", inline: "end" });
                      }
                    })

                } catch (error) {
                    console.error(error);
                }
            },

            async createRoom() {
                try {
                    const response = await axios.get('/room_create');
                    this.roomid = response.data.data["id"];
                } catch (error) {
                    console.error(error);
                }
            },
            async joinGame() {
               
                this.joinModal.hide();

                var response = await axios.get('/room_join?name=' + this.nickname + "&head=" + this.headurl);
                if (response.data["errcode"] != 0) return

                this.roomid = response.data.data["id"];;
                this.token = response.data.data["token"];

            },

            randName() {
                return "游客" + parseInt(Math.random()*(99999-10000)+10000);
            },

            async sendChat() {
                var response = await axios.get('/room_chat?roomid=' + this.roomid + "&token=" + this.token + "&chat=" + this.chat);
                this.chat = "";    
                this.fetchData();
                
            },

            scrollEvent() {

                const page_content = this.$refs.page_content;

                if (page_content.scrollTop + page_content.offsetHeight >= page_content.scrollHeight) {
                    this.autoscroll = true;
                } else {
                    this.autoscroll = false;
                }

                console.log(this.autoscroll);
            },

            setHead(e) {
              this.headurl = e.target.src;
              this.headModal.hide();
              this.joinModal.show();
            },

            start() {

              this.items = []
              this.roomid =0
              this.chatid = 0
              this.token = ""
              this.typedAni = false;
              this.fetchData();
              
              this.timer = setInterval(() => {
                this.fetchData();
              }, 1000);

              this.joinModal.show();
              this.headModal.hide();


            }, 

            restart() {
              clearInterval(this.timer);

              this.start();
            }
        }
    }

</script>

<template>

<div class="container-fluid pe-0 ps-0 page-main">

<!-- Modal -->
  <div class="modal fade" id="staticBackdrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content bg-dark">
        <div class="div-center modal-header">
          <svg version="1.0" xmlns="http://www.w3.org/2000/svg" class="me-2"
          width="24.000000pt" height="24.000000pt" viewBox="0 0 1174.000000 1280.000000"
                            preserveAspectRatio="xMidYMid meet">
                            <g transform="translate(0.000000,1280.000000) scale(0.100000,-0.100000)"
                            fill="#e5e0d8" stroke="none">
                            <path d="M9038 12713 c-106 -104 -272 -300 -463 -548 -314 -408 -527 -668
                            -834 -1020 -613 -702 -1201 -1288 -1435 -1429 -84 -51 -300 -98 -600 -131
                            -161 -18 -277 -52 -406 -122 -276 -148 -594 -405 -1230 -994 -258 -239 -370
                            -332 -504 -421 -239 -157 -409 -228 -935 -389 -551 -168 -862 -298 -1117 -465
                            -146 -96 -352 -292 -389 -371 -50 -104 40 -262 229 -406 210 -160 478 -304
                            981 -527 176 -78 332 -148 346 -156 l27 -13 -100 -293 c-149 -437 -217 -724
                            -238 -998 -11 -143 -32 -237 -89 -400 -40 -112 -61 -156 -267 -560 -151 -297
                            -304 -701 -455 -1205 -114 -381 -216 -818 -234 -1010 -14 -143 -90 -252 -384
                            -546 -343 -342 -693 -619 -876 -692 -37 -15 381 -16 5084 -16 l5124 -1 -7 122
                            c-16 297 -91 658 -151 722 -44 47 -50 92 -49 346 1 259 20 445 75 726 54 277
                            102 422 185 558 94 156 135 318 135 546 1 353 -82 645 -319 1128 -107 216
                            -144 309 -182 457 -64 250 -85 546 -77 1070 7 414 25 624 73 840 26 115 43
                            164 104 290 102 212 191 483 275 843 83 355 168 587 327 894 91 175 178 316
                            283 456 89 120 235 271 310 322 237 159 376 362 451 659 63 247 38 466 -80
                            687 -145 273 -290 407 -439 407 -66 0 -89 -9 -201 -77 -117 -71 -201 -68 -393
                            12 -212 89 -214 88 -212 -96 1 -135 -9 -183 -51 -240 -44 -59 -122 -84 -163
                            -51 -39 30 -43 97 -17 305 36 285 22 576 -37 784 -45 158 -191 379 -433 655
                            -209 238 -450 435 -532 435 -12 0 -57 -36 -110 -87z"/>
                            </g>
          </svg>

          <span class="ps-2">微信读书QA系统</span>

        </div>
        <div class="modal-body">
          <div class="row" style="justify-content: center">
            <div class="col-2" style="max-width: 60px;">
              <button type="button" class="btn btn-link m-0 p-0" data-bs-toggle="modal" data-bs-target="#modalHead">
              <img width="45" height="45" :src="this.headurl" class="rounded-pill img-head">
              </button>
            </div>
            <div class="col-10">
              <input type="text" class="rounded-lg input-nickname p-3 mb-3" placeholder="游戏昵称" aria-label="Username" aria-describedby="basic-addon1" v-model="nickname" @keydown.enter="joinGame">
            </div>
          </div>
          <div class="div-center">
            <button type="button" class="btn btn-outline-light" @click="joinGame">进入系统</button>
          </div>

        </div>
      </div>
    </div>
  </div>


  <div class="modal fade" id="modalHead" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content bg-dark">
        <div class="modal-header">请选择头像</div>
        <div class="modal-body">
          <ul class="list-group list-group-horizontal flex-wrap div-center" >
            <img v-for="n in 6" class="list-group-item rounded-pill m-2 p-0 img-head" width="75" height="75" :src='"../public/heads/head" + n + ".png"'  @click="setHead" >
          </ul>
        </div>
      </div>
    </div>
  </div>


    <div class="page-header">





      <div class="row pt-3">
        <div class="col-1 col-md-1">
          <button type="button" class="btn btn-link m-0 p-0 ps-2 text-normal" @click="restart">
          <svg stroke="currentColor" fill="currentColor" stroke-width="0" viewBox="0 0 24 24" height="24" width="24" xmlns="http://www.w3.org/2000/svg" data-darkreader-inline-fill="" 
          data-darkreader-inline-stroke="" style="--darkreader-inline-fill:currentColor; --darkreader-inline-stroke:currentColor;">
          <path fill="none" d="M0 0h24v24H0z"></path><path d="M17.77 3.77L16 2 6 12l10 10 1.77-1.77L9.54 12z">
          </path></svg>
        </button>

        </div>
        <div class="col-9 col-md-9">
          <div class="" style="padding-top: 2px;">
            <span class="">{{ message }}</span>
          </div>
        </div>
        <div class="col-2 col-md-2">
          <div class="d-flex justify-content-end">

            <button type="button" class="btn btn-link m-0 p-0 text-normal">
            <svg stroke="currentColor" fill="currentColor" stroke-width="0" viewBox="0 0 512 512" height="24" width="24" xmlns="http://www.w3.org/2000/svg" 
            data-darkreader-inline-fill="" data-darkreader-inline-stroke="" style="--darkreader-inline-fill:currentColor; --darkreader-inline-stroke:currentColor;">
            <path d="M444.7 230.4l-141.1-132c-1.7-1.6-3.3-2.5-5.6-2.4-4.4.2-10 3.3-10 8v66.2c0 2-1.6 3.8-3.6 4.1C144.1 195.8 85 300.8 64.1 409.8c-.8 4.3 5 8.3 7.7 4.9 51.2-64.5 
            113.5-106.6 212-107.4 2.2 0 4.2 2.6 4.2 4.8v65c0 7 9.3 10.1 14.5 5.3l142.1-134.3c2.6-2.4 3.4-5.2 3.5-8.4-.1-3.2-.9-6.9-3.4-9.3z">
            </path></svg>
            </button>


          <button type="button" class="btn btn-link m-0 p-0 text-normal">
            <svg stroke="currentColor" fill="currentColor" stroke-width="0" viewBox="0 0 24 24" height="25" width="25" xmlns="http://www.w3.org/2000/svg" 
            data-darkreader-inline-fill="" data-darkreader-inline-stroke="" style="--darkreader-inline-fill:currentColor; --darkreader-inline-stroke:currentColor;">
            <path d="M12 10c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0-6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0 12c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"></path>
            </svg>
          </button>
          </div>
        </div>
      </div>
    </div>

    <div class ="page-content" ref="page_content" @scroll="scrollEvent">
            <ul class="list-group list-group-flush pb-4">
                <li class="list-group-item pt-1 pb-4 ps-2 pe-2" :class="{'bg-normal':index%2==1, 'bg-normal-light': index%2==0}" v-for="(item,index) in items" :key="index">
                  <div class="m-0 p-0 row">
                  <div class="col-auto p-0">
                    <img width="45" height="45" v-if='item.Head==""' src="../public/heads/0.webp" class="rounded-pill">
                    <img width="45" height="45" v-else :src="item.Head" class="rounded-pill img-head">
                  </div>
                  <div class="col-10">
                  <b>{{ item.Name }}</b>
                  <span class="" v-if="item.MsgType!=0"></span>
                  <span class="badge text-bg-normal ms-1" v-else-if="item.PrivateId==0">ai</span>
                  <span class="badge text-bg-player ms-1" v-else>@{{ item.PrivateName }}</span>
                  <br/>

                  <VueTyped
                    :strings="[item.Content]"
                    :typeSpeed="50"
                    :backSpeed="0"
                    :loop="false"
                    v-if="item.ani"
                    >
                  </VueTyped>
                  <div v-else>
                    {{ item.Content }}
                  </div>
                  
                  </div>
                  </div>
                </li>
            </ul>
            <div ref="msg_end" style="height:0px; overflow:hidden"></div>
    </div>

    <div class="div-center page-footer">
            <div class="div-msg rounded-pill p-3">
                <input type="text" class="sendmessage" placeholder="发送消息" aria-label="Username" aria-describedby="basic-addon1" v-model="chat" @keydown.enter="sendChat" >
                <div class="input-group-append" >
                  <a @click="sendChat">
                  <svg stroke="currentColor" fill="currentColor" stroke-width="0" viewBox="0 0 24 24" height="25" width="25" xmlns="http://www.w3.org/2000/svg" style="color: rgb(88, 158, 224);">
                  <path fill="none" d="M0 0h24v24H0z"></path><path d="M2.01 21L23 12 2.01 3 2 10l15 2-15 2z">
                  </path></svg>
                  </a>
                </div>
            </div>
            
    </div>
  </div>

</template>

