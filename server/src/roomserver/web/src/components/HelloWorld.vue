<script setup>
import { ref } from 'vue'

defineProps({
  msg: String,
})

const count = ref(0)
</script>

<template>
  
  <div class=page-main>
    <div class="modal fade" id="staticBackdrop" data-backdrop="static" data-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content bg-normal">
                    <div class="modal-header div-center">
                        
                        <svg version="1.0" xmlns="http://www.w3.org/2000/svg" class="mr-2"
                            width="24.000000pt" height="24.000000pt" viewBox="0 0 1174.000000 1280.000000"
                            preserveAspectRatio="xMidYMid meet">
                            <g transform="translate(0.000000,1280.000000) scale(0.100000,-0.100000)"
                            fill="#ffffff" stroke="none">
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

                        <span class="mt-1">全球领先的人工智能杀平台</span>

                    </div>
                    <div class="modal-body">
    
                        <input type="text" class="rounded-lg input-nickname p-3 mb-3" placeholder="游戏昵称" aria-label="Username" aria-describedby="basic-addon1" v-model="nickname" @keydown.enter="joinGame">
                        <br/>
                        <div class="div-center">
                            <button type="button" class="btn btn-outline-light" @click = "joinGame">加入游戏</button>
                        </div>
                    </div>
                </div>
            </div>
    </div>



    <div class="page-header">
            <nav class="navbar navbar-dark mb-0 pb-0">
            
                <a class="navbar-brand" href="#">

                    <svg version="1.0" xmlns="http://www.w3.org/2000/svg" class="mr-2"
                            width="24.000000pt" height="24.000000pt" viewBox="0 0 1174.000000 1280.000000"
                            preserveAspectRatio="xMidYMid meet">
                            <g transform="translate(0.000000,1280.000000) scale(0.100000,-0.100000)"
                            fill="#ffffff" stroke="none">
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
                    <span class="mt-3 ml-2">{{ message }}</span>
                </a>
            </nav>
    </div>

    <div class ="page-content" ref="page_content" @scroll="scrollEvent">
            <ul class="list-group list-group-flush">
                <li class="list-group-item pt-4 pb-4 pl-5 pr-5" :class="{'bg-normal':index%2==1, 'bg-normal-light': index%2==0}" v-for="(item,index) in items">
                    {{ item.Content  }}
                </li>
            </ul>
            <div ref="msg_end" style="height:0px; overflow:hidden"></div>
    </div>

    <div class="page-footer div-center">
            <div class="div-msg rounded-lg p-3">
                <input type="text" class="sendmessage" placeholder="发送消息" aria-label="Username" aria-describedby="basic-addon1" v-model="chat" @keydown.enter="sendChat" >
                <div class="input-group-append">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="bi bi-send mt-1" viewBox="0 0 16 16">
                        <path d="M15.854.146a.5.5 0 0 1 .11.54l-5.819 14.547a.75.75 0 0 1-1.329.124l-3.178-4.995L.643 7.184a.75.75 0 0 1 .124-1.33L15.314.037a.5.5 0 0 1 .54.11ZM6.636 10.07l2.761 4.338L14.13 2.576 6.636 10.07Zm6.787-8.201L1.591 6.602l4.339 2.76 7.494-7.493Z"/>
                    </svg>
                </div>
            </div>
            
    </div>
  </div>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
