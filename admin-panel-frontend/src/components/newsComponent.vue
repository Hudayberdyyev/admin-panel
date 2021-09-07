<template>
    <div class="news-container">
        <div class="news-all" v-if="!addNewFromAdmin">
            <div class="news-operations">
                <div class="news-operations-row">
                    <div>
                        <img src="./assets/img/149147.png" alt="" class="news-operation-items" title="Täzelik goş" @click="addNew">
                        <img src="./assets/img/126483.png" alt="" class="news-operation-items" title="Täzeligi üýtget">
                        <img src="./assets/img/126468.png" alt="" class="news-operation-items" title="Täzeligi ýok et">
                        <img src="./assets/img/check-icon_icon-icons.com_76169.png" alt="" class="news-operation-items" title="Ähli täzelikleri belle" @click="checkAll()">
                        <img src="./assets/img/delete-icon_icon-icons.com_76172.png" alt="" class="news-operation-items" title="Ähli bellikleri aýyr" @click="discheckAll()">
                    </div>
                    <div style="display:flex;">
                        <img src="./assets/img/126492.png" alt="" class="news-operation-items" title="Öňki sahypa" @click="decPage()">
                        <div class="news-list-page">{{ page }}</div>
                        <img src="./assets/img/126490.png" alt="" class="news-operation-items" title="indiki sahypa" @click="incPage()">
                    </div>
                </div>
                <div class="news-operations-row">
                    <div>
                        Filter:
                    </div>
                    <div>
                        <select class="select-menu" v-model="categoriesID" @click="changedID()">
                            <option value="0">Ähli kategoriýalar</option>
                            <option v-for="(el, idx) in categories" :key="idx" :value="el.id">{{ el.text }}</option>
                        </select>
                    </div>
                    <div>
                        <select class="select-menu" v-model="authorsID" @click="changedID()">
                            <option value="0">Ähli awtorlar</option>
                            <option v-for="(el2, idx) in authors" :key="idx" :value="el2.id">{{ el2.name }}</option>
                        </select>
                    </div>
                </div>
            </div>
            <div class="news-list" id="news-list">
                <div class="news-list-items" v-for="(el, idx) in newsList" :key="idx" @click="showNew(el,idx)" :class="(border == el.id ? 'bordered' : '')">
                    <input type="checkbox" class="list-items-check">
                    <div class="news-mini-name">
                        {{el.title}}
                    </div>
                    <div class="news-mini-inf">
                        <div>
                            <img src="./assets/img/mbripreview_99511.png" alt="" class="news-mini-inf-icons">
                            <span class="news-mini-inf-text">{{el.view_count}}</span>
                            <img src="./assets/img/calendar_date_day_event_month_schedule_time_icon_123230.png" alt="" class="news-mini-inf-icons">
                            <span class="news-mini-inf-text">{{el.publish_date}}</span>
                            <span class="news-mini-inf-status">status: {{el.status}}</span>
                        </div>    
                    </div>
                </div>
            </div>
        </div>
        <div class="news-active" v-if="!addNewFromAdmin">
            <div class="news-active-inf">
                <div>
                    <b><big>{{selectedNewCategory}}</big></b>
                </div>
                <div>
                    <i v-for="(el,idx) in selectedNewTags" :key="idx">#{{el.name}}</i>
                </div>
            </div>
            <div class="news-active-text" id="news-active"></div>
        </div>
    <new-component v-if="addNewFromAdmin" @exitFun="restartNews"></new-component>
    </div>
</template>

<script>

import newComponent from './newComponent.vue';
import { get } from './fetchGET'

export default{
    components:{
        newComponent
    },
    props:['addNewFromAdmin'],
    data(){
        return{
            categories:'',
            authors:'',
            categoriesID:0,
            authorsID:0,
            newsList:[],
            page:1,
            selectedNewTags:[],
            selectedNewContent:[],
            selectedNewCategory:'',
            border:''
        }
    },
    mounted(){
        this.newsGet()
        this.newsStart()
    },
    updated(){
        this.newsStart()
    },
    destroyed(){
        this.exitNews()
    },
    methods:{
        async newsGet(){
            
            let categoriesResponse = await get('/api/tm/news/categories')
            categoriesResponse = await categoriesResponse.json()
            this.categories = categoriesResponse.data

            let authorsResponse = await get('/api/tm/news/authors')
            authorsResponse = await authorsResponse.json()
            this.authors = authorsResponse.data
            
            this.watchFilter()
        },
        newsStart () {
            if (this.addNewFromAdmin) { return ; }
            
        },
        addNew(){
            this.addNewFromAdmin = true
        },
        restartNews(){
            this.addNewFromAdmin = false
        },
        exitNews(){
            this.$emit('exitNews')
        },
        checkAll(){
            let elements = document.querySelectorAll('input[type$="checkbox"]')
            for (let elem of elements) {
                elem.checked = true
            }
        },
        discheckAll(){
            let elements = document.querySelectorAll('input[type$="checkbox"]')
            for (let elem of elements) {
                elem.checked = false
            }
        },
        incPage(){
            this.page = this.page + 1 
            this.watchFilter()
        },
        decPage(){
            this.page = this.page > 1 ? this.page - 1 : this.page
            this.watchFilter()
        },
        changedID(){
            this.watchFilter()
            this.page = 1
        },
        async showNew(el, idx){
            this.border = el.id
            document.getElementById('news-active').innerHTML = '<h2>Ýüklenilýär...</h2>'
            this.selectedNewCategory = el.Categories.text

            let tagResponse = await get(`/api/${el.hl}/news/${el.news_text_id}/tags`)
            tagResponse = await tagResponse.json()
            this.selectedNewTags = tagResponse.data

            let contentResponse = await get(`/api/tm/news/${el.news_text_id}/content`)
            contentResponse = await contentResponse.json()
            this.selectedNewContent = contentResponse.data
            
            var v = '<b><big>' + el.title + '</big></b>'
            this.selectedNewContent.forEach(elem => {
                if (elem.tag == "img") { v = v + '<'+elem.tag+' src="' + elem.attributes[0].value + '" style="width:100%">'}
                else {v = v + '<'+elem.tag+'>'+elem.value+'</'+elem.tag+'>'}
            })
            document.getElementById('news-active').innerHTML = v
        },
        async watchFilter(){
            document.getElementById('news-active').innerHTML = ''
            document.getElementById('news-list').scrollTop = 0
           let response = await get(`/api/tm/news/all/author/${this.authorsID}/category/${this.categoriesID}?limit=10&page=${this.page}`)
            response = await response.json()
            this.newsList = response.data
       }
    },
    computed:{
        
    },
    watch:{
       
    }
}


</script>

<style src="./assets/css/newsComponent.css"></style>