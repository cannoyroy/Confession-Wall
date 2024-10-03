<script setup>
import { useRoute} from 'vue-router';
import 'element-plus/dist/index.css';
import axios from 'axios';
const route = useRoute();
</script>

<template>
     <section class="accordion">
       <div class="item">
            <img src="https://f9ir.github.io/acc/acc/img/Location-Pin.png" alt="">
            <h3>左转男爵，右转女巫，你走中路？</h3>
       </div>
            <p>
                <div style="display: flex; justify-content: center; align-items: center;">
                    <el-radio-group v-model="gender">
                    <el-radio :value="maleValue" style="color: white;margin-right: 20px">男</el-radio>

                    <el-radio :value="otherValue" style="color: white;margin-right: 20px">其他</el-radio>

                    <el-radio :value="femaleValue" style="color: white;margin-right: 20px">女</el-radio>
                    </el-radio-group>
                </div>
            </p>
        <div class="item">
            <img src="https://f9ir.github.io/acc/acc/img/Headphones.png" alt="">
            <h3>速速报上，联络暗号</h3>
       </div>
        <p>
            <el-input v-model="contact.phone" placeholder="电话（留空即不变）"></el-input>
            <el-input v-model="contact.qq" placeholder="QQ（留空即不变）"></el-input>
            <el-input v-model="contact.wechat" placeholder="微信（留空即不变）"></el-input>
            <el-input v-model="contact.other" placeholder="其他联系方式（留空即不变）"></el-input>
        </p>
        <div class="item">
            <img src="https://f9ir.github.io/acc/acc/img/Lightbulb.png" alt="">
            <h3>来自M87星云？</h3>
       </div>
            <p>
                <!-- <el-cascader
                    v-model="selectedRegion"
                    :options="regionData"
                    placeholder="选择地区"
                    @change="handleRegionChange"
                ></el-cascader> -->
                <el-input v-model="selectedRegion" placeholder="输入地区（留空即不变）"></el-input>
            </p>
        <div class="item">
            <img src="https://f9ir.github.io/acc/acc/img/Bookmarks.png" alt="">
            <h3>“揭开神秘的面纱”</h3>
       </div>
            <p>
                <el-input v-model="avatarUrl" placeholder="头像链接（留空即不变）"></el-input>
            </p>
        <div class="item">
            <img src="https://f9ir.github.io/acc/acc/img/Lightning-Bolt.png" alt="">
            <h3>所以，你是……？</h3>
       </div>
           <p>
            <el-input
                type="textarea"
                v-model="introduction"
                placeholder="自我介绍（留空即不变）"
            ></el-input>
           </p>
           
   </section>
   <div style="display: flex; justify-content: center; align-items: center;">
        <el-button type="primary" @click="handleSubmit">设置完成</el-button>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    data() {
      return {
        user_id: this.$route.params.user_id,
        gender: '',
        maleValue: 'male', // 定义male的值
        otherValue: 'other', // 定义other的值
        femaleValue: 'female', // 定义female的值
        contact: {
          phone: '',
          qq: '',
          wechat: '',
          other: ''
        },
        selectedRegion: '',
        // selectedRegion: [],
        // regionData: ["浙江", "广东"], // 地区数据
        avatarUrl: '',
        introduction: ''
      };
    },
    methods: {
        handleSubmit(){
            const newProfile={
                user_id: Number(this.$route.query.user_id),
                gender: this.gender || 'none',
                contact_tele: this.contact.phone || 'none',
                contact_qq: this.contact.qq || 'none',
                contact_wechat: this.contact.wechat || 'none',
                contact_other: this.contact.other || 'none',
                region: this.selectedRegion || 'none',
                img_url: this.avatarUrl || 'none',
                other_info: this.introduction || 'none'
            }
            // console.log(newProfile)
            axios.put('http://127.0.0.1:8080/profile/edit', newProfile)
                .then(res => {
                    
                    console.log(res);
                    this.$message.success('修改成功');
                    this.$router.push({ name: 'profile', query: this.$route.query});
                })
                .catch(err => {
                    console.log(err);
                    console.log(newProfile);
                    this.$message.error('设置失败');
                }); 
            
        }
    }
}

</script>

<style lang="scss" scoped>
$Mahogany : #6B5567;
$AppleBlossom : #9273A6;
$GoldenSand : #eadeff;
$EggWhite : #F2A9C8;
$OrangeRoughy : #FFC7CF;
$mainfont : "JF Flat Regular";

@font-face {
    font-family: "JF Flat Regular";
    src: url('https://f9ir.github.io/acc/acc/font/JF-Flat-regular.eot');
    src: url('https://f9ir.github.io/acc/acc/font/JF-Flat-regular.eot?#iefix') format('embedded-opentype'),
    url('https://f9ir.github.io/acc/acc/font/JF-Flat-regular.svg#JF Flat Regular') format('svg'),
    url('https://f9ir.github.io/acc/acc/font/JF-Flat-regular.woff') format('woff'),
    url('https://f9ir.github.io/acc/acc/font/JF-Flat-regular.ttf') format('truetype');
    font-weight: normal;
    font-style: normal;
    background: #f6704b;
}

body{
  background: #f6704b;
}

*{
  margin: 0;
  padding: 0;
}
.accordion{
    margin:50px auto;
    width: 380px;
    background: #ccc;
    cursor: pointer;
    .item{
        height: 100px;
        h3{
            display: inline-block;
            vertical-align: middle;
            height: 100%;
            padding-left: 50px;
            font-family: $mainfont;
            font-size: 20px;
            font-weight: 400;
        }
        img{
            padding-left: 15px;
            width: 30px;
            vertical-align: middle;
        }
        h3:before{
            content: "";
            display: inline-block;
            vertical-align: middle;
            height: 100%;
        } 
    }
    .item:first-of-type{
        background: $Mahogany;
        color: $GoldenSand;
    }
    .item:nth-of-type(2){
        background: $AppleBlossom;
        color: $EggWhite;
    }
    .item:nth-of-type(3){
        background: $GoldenSand;
        color: $Mahogany;
    }
    .item:nth-of-type(4){
        background: $EggWhite;
        color: $Mahogany;
    }
    .item:last-of-type{
        background: $OrangeRoughy;
        color: $Mahogany;
    }
    p{
        font-family: $mainfont;
        font-size: 18px;
        font-weight: 400;
        padding: 15px;
        display: block;
        box-shadow: inset 0 3px 7px rgba(#000, 0.2);
    }
    p:first-of-type{
        background: $Mahogany;
        color: $GoldenSand;
    }
    p:nth-of-type(2){
        background: $AppleBlossom;
        color: $EggWhite;
    }
    p:nth-of-type(3){
        background: $GoldenSand;
        color: $Mahogany;
        
    }
    p:nth-of-type(4){
        background: $EggWhite;
        color: $OrangeRoughy;
    }
    p:last-of-type{
        background: $OrangeRoughy;
        color: $EggWhite;
    }
}
</style>