<script setup>
import { useRoute} from 'vue-router';
const route = useRoute();
function EditProfile() {
        let data = {
            "user_id": Userid,
            "gender": "Gender",
            "contact_tele": "Contact_tele",
            "contact_qq":"Contact_qq",
            "contact_wechat":"Contact_wechat",
            "contact_other":"Contact_other",
            "region":"Region",
            "img_url":"Img_url",
            "other_info":"Other_info",
        }
        if (data.content.trim() !== "") {

            axios.put("http://127.0.0.1:8080/profile/edit", data).then((res) => {
                if (res.data.msg === "success") {
                    alert("修改成功");
                    
                    
                }
                else {
                    alert(res.data.msg);
                };
            }).catch((err) => { alert(err); })
        }
        else {
            alert("内容不能为空")
        }
    };
</script>

<template>
  <h1 style="margin-left: 2em;">EditProfile页面</h1>
  <RouterView />
  性别：<input type="text" v-model="Gender" placeholder="请输入性别">
  联系方式：
  <br>
  手机号：<input type="text" v-model="Contact_tele" placeholder="请输入手机号">
  QQ号：<input type="text" v-model="Contact_qq" placeholder="请输入QQ号">
  微信号：<input type="text" v-model="Contact_wechat" placeholder="请输入微信号">
  其他：<input type="text" v-model="Contact_other" placeholder="请输入其他联系方式">
  <br>
  地区：<input type="text" v-model="Region" placeholder="请输入地区">
  头像：<input type="text" v-model="Img_url" placeholder="请输入头像地址">
  其他信息：<input type="text" v-model="Other_info" placeholder="请输入其他信息">
  <br>
  <br>
<button  @click="EditProfile">修改</button>
<script>
</script>

<style>
</style>