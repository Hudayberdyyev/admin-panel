<template>
    <div class="new-container">
        <div class="new-operations">
            <div class="new-lang">
                <div class="operation-name">Dil:</div>
                <div class="new-lang-items">
                    <div class="lang-item-style" :class="chosedLang == 'TM'?'lang-active':''" @click="changeLang('TM')">TM</div>
                    <div class="lang-item-style" :class="chosedLang == 'RU'?'lang-active':''" @click="changeLang('RU')">RU</div>
                    <div class="lang-item-style" :class="chosedLang == 'EN'?'lang-active':''"
                    @click="changeLang('EN')">EN</div>
                </div>
            </div>
            <div class="new-instruments">
                <div class="operation-name">Gurallar:</div> 
                <div class="new-instrument-list">
                    <div class="new-instrument-items" style="font-weight:bold;" @click="textChange('b')">
                        B
                    </div>
                    <div class="new-instrument-items">
                        A <span style="font-size:10px;"  @click="textChange('big')">&#9650</span>
                    </div>
                    <div class="new-instrument-items">
                        A <span style="font-size:10px;" @click="textChange('small')">&#9660</span>
                    </div>
                    <div class="new-instrument-items" style="font-style:italic;" @click="textChange('i')">
                        K
                    </div>
                    <div class="new-instrument-items" style="text-decoration:line-through;" @click="textChange('s')">
                        A
                    </div>
                    <div class="new-instrument-items" style="text-decoration:underline;" @click="textChange('ins')">
                        A
                    </div>
                    <div class="new-instrument-items">
                        <img src="./assets/img/2087927.png" alt="" style="width:12px;" @click="textChange('sup')"> 
                    </div>
                    <div class="new-instrument-items">
                        <img src="./assets/img/1828084.png" alt="" style="width:14px;" @click="textChange('sub')">
                    </div>
                    
                </div>
            </div>
            <div class="new-status">
                <div class="operation-name">Status:</div>
                    <button class="new-status-check" :class="status_0 ? 'new-status-active':''" @click="status(0)">0</button>
                    <button class="new-status-check"  :class="status_1 ? 'new-status-active':''" @click="status(1)">1</button>
            </div>
            <div class="new-category">
                <div class="operation-name">Kategoriýalar:</div>
                <div>
                    <input type="radio" name="category" @click="categoryCheck(0)" checked><input type="text" placeholder="Täze kategoriýa" class="input-menu" :disabled="!category_0" v-model="newCategory">
                </div>
                <div>
                    <input type="radio" name="category" @click="categoryCheck(1)">
                    <select id="" class="select-menu" :disabled="!category_1">
                        <option value="" v-for="(item,idx) in category" :key=idx>{{ item.value }}</option>
                    </select>
                </div>
            </div>
            <div class="new-tags">
                <div class="operation-name">Tegler:</div>
                <div>
                    <div>
                        <input type="text" class="input-menu" placeholder="Täze teg..." v-model="newTag" @keydown.enter="addTag"> 
                        <button class="new-tag-add" @click="addTag">Goş</button>
                    </div>
                    <div class="new-tag-list">
                        <div class="new-tag-items" v-for="(item,idx) in hashtag" :key="idx" @click="delTag(idx)">
                            <div class="tag-delete"><img src="./assets/img/korzina.png" alt="" style="height:15px;"></div>
                            <span class="new-tag-item-name">{{ item.value }}</span>
                        </div>
                        
                    </div>
                </div>    
            </div>
            <div class="new-save-exit">
                <button class="save-exit-button" @click="saveNew">Save</button>
                <button class="save-exit-button" @click="$emit('exitFun')">Exit</button>
            </div>
        </div>
        <div class="new-editor">
            <div class="new-main">
                <div class="new-main-image">
                    <div class="new-main-h1">Esasy surat:</div>
                    <div style="margin:5px 8px;">
                        <input type="file" id="file1" accept="image/*" style="width:0.1px;height:0.1px;">
                        <label for="file1" class="new-main-surat">Suraty ýükle!</label>
                    </div>
                </div>
                <div class="new-main-name">
                    <div class="new-main-h1">Sözbaşy:</div>
                    <div style="margin:5px 10px;width:95%;">
                        <input type="text" class="new-main-text">
                    </div>
                </div>
            </div>
            <div class="new-content">
                <div id="new-redit" contenteditable @click="selectingFun()" @keydown="selectingFun()">
                    
                </div>
            </div>
        </div>
    </div>
</template>

<script>




export default {
    data(){
        return{
            status_0:true,
            status_1:false,
            category:[
                {value : 'Sport'},
                {value : 'Tehnologiýalar'},
                {value : 'Syýasat'},
                {value : 'Medeniýet'},
                {value : 'Ykdysadyýet'}
            ],
            category_0:true,
            category_1:false,
            newCategory:'',
            chosedLang:'RU',
            selNode:null,
            hashtag:[
                {value : "Turkmenistan"},
                {value : "Watan"},
                {value : "Ahal"},
                {value : "2021"}
            ],
            newTag:''
        }
    },
    mounted(){
        this.newStart()
    },
    methods:{
        newStart(){
            var v=''
            const mass = json.data
            mass.forEach(el => {
                v = v + '<'+el.tag+'>'+el.text+'</'+el.tag+'>'
            })
            document.getElementById('new-redit').innerHTML = v
        },
        status(k){
            this.status_0 = false; this.status_1 = false;
            if (k) { this.status_1 = true; }
            else { this.status_0 = true; }
        },
        categoryCheck(k){
            this.category_0 = false; this.category_1 = false;
            if (k) { this.category_1 = true; }
            else { this.category_0 = true; }
            
        },
        uploadImage(e){
            
            
            //k.readAsDataURL(file)
            
        },
        changeLang(k){
            this.chosedLang = k
        },
        saveNew(){
            const v = document.getElementById('new-redit').innerHTML
            
        },
        selectingFun(){
            if (!document.getSelection().toString().length) {return ;}
            this.selNode = document.getSelection()
        },
        textChange(k){
            let range = this.selNode.getRangeAt(0)
            let newNode = document.createElement(k)
            newNode.innerHTML = this.selNode.toString()
            range.deleteContents()
            range.insertNode(newNode)
            
        },
        addTag(){
            this.hashtag.push({value:this.newTag})
            this.newTag = ''
        },
        delTag(idx){
            this.hashtag.splice(idx,1)
        }
    }
}

</script>

<style src="./assets/css/newComponent.css"></style>