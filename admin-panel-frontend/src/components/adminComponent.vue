<template>
    <div class="admin-container">
        <div class="admin-menu-fix"></div>
        <div class="admin-menu" id="menu">
            <div>
                <div class="admin-name">
                    <div><b>Belet</b> News</div>
                </div>
                <div class="admin-menu-items admin-active-item" id="main" @click="show('main')">
                    <div style="display:flex;">
                        <img src="./assets/img/126496.png" alt="" class="admin-menu-items-icon">
                        <div class="admin-show">Esasy</div>    
                    </div>    
                </div>
                <div class="admin-menu-items" id="news" @click="show('news')">
                    <div style="display:flex;">
                        <img src="./assets/img/149346.png" alt="" class="admin-menu-items-icon">
                        <span class="admin-show">Täzelikler</span>
                    </div>    
                </div>
                <div class="admin-menu-items" id="stat" @click="show('stat')">
                    <div style="display:flex;">
                        <img src="./assets/img/126484.png" alt="" class="admin-menu-items-icon">
                        <span class="admin-show">Statistika</span>
                    </div>    
                </div>
                <div class="admin-menu-items" id="mails" @click="show('mails')">
                    <div style="display:flex;">
                        <img src="./assets/img/126516.png" alt="" class="admin-menu-items-icon">
                        <span class="admin-show">Уведомления</span>
                    </div>    
                </div>
                <div class="admin-menu-items" id="users" @click="show('users')">
                    <div style="display:flex;">
                        <img src="./assets/img/126486ad.png" alt="" class="admin-menu-items-icon">
                        <span class="admin-show">Dolandyryjylar</span>
                    </div>    
                </div>
            </div>
        </div>
        <div class="admin-container-2">
            <div class="admin-nav">
                <div class="admin-nav-1-block">
                    <div style="display:flex; justify-content:space-between;">
                        <img src="./assets/img/126474.png" alt="" class="admin-search-icon" id="search-icon">
                        <div class="admin-search-panel">
                            <input type="text" class="admin-search-input" placeholder="gözleg...">
                            <button class="admin-search-button">OK</button>
                        </div>
                    </div>    
                </div> 
                <div class="admin-nav-2-block">
                    <div style="display:flex; justify-content: space-around;">
                        <button class="admin-button" @click="addNew=true;show('news');">+ Täzelik goş</button>
                        <img src="./assets/img/6575677.jpg" alt="" class="admin-user-icon">
                        <img src="./assets/img/126467.png" alt="" class="admin-exit-icon" @click="exitFun">
                    </div>
                </div>
            </div>
            <div class="admin-content">
                <main-component v-if="comps.main"></main-component>
                <news-component v-if="comps.news" :addNewFromAdmin="addNew" @exitNews="falseAddNew"></news-component>
                <stat-component v-if="comps.stat"></stat-component>
                <sos-component v-if="comps.mails"></sos-component>
                <users-component v-if="comps.users"></users-component>
            </div>
        </div>
    </div>
</template>



<script>

import mainComponent from './mainComponent.vue';
import newsComponent from './newsComponent.vue';
import statComponent from './statComponent.vue';
import sosComponent from './sosComponent.vue';
import usersComponent from './usersComponent.vue';

export default {
    components:{
        mainComponent,
        newsComponent,
        statComponent,
        sosComponent,
        usersComponent
    },
    data(){
        return{
            comps:{
                main:true,
                news:false,
                stat:false,
                mails:false,
                users:false
            },
            addNew:false    
        }
    },
    mounted(){
        
    },
    methods:{
        
        exitFun(){
            this.$emit('exitFun',true)
        },
        show(k){

            for (let key in this.comps) {
                this.comps[key] = false
                document.getElementById(key).setAttribute("class","admin-menu-items")
            }
            this.comps[k] = true

            const v = document.getElementById(k)
            v.setAttribute("class","admin-menu-items admin-active-item")
            
            
        },
        falseAddNew(){
            this.addNew = false
        }
    }
}


</script>

<style src="./assets/css/adminComponent.css"></style>
