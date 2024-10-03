<script setup>
import { useRoute} from 'vue-router';
const route = useRoute();
const postlist = ref(null);
let Userid = 0;
function GetProfile() {
        axios.get("http://127.0.0.1:8080/profile", { params: {"user_id":Userid}}).then((res) => {
            if (res.data.msg === "success") {
                postlist.value = res.data.data.post_list
            }
            else {
                alert(res.data.msg);
            };
        }).catch((err) => { alert(err); })
    };
</script>

<template>
  <h1 style="margin-left: 2em;">Newprofile页面</h1>
<button @click="GetProfile">GetProfile</button>
<div class="NEW-PROFILE" >
            <div class="feedback-container">
            <div v-for="item in postlist" class="feedback-item">
                <div class="oneBlock">
                    <p>姓名:{{item.Username}}</p>
                    <p>性别:{{item.Gender}}</p>
                    <p>联系电话:{{item.ContactTele}}</p>
                    <p>qq:{{item.ContactQQ}}</p>
                    <p>微信:{{item.ContactWechat}}</p>
                    <p>其他联系方式：{{item.ContactOther}}</p>
                    <p>发布时间:{{item.CreateAt}}</p>
                    <img src="{{item.ImgUrl}}" alt="头像">
                </div>
                </div>
            </div>
        </div>
</template>

<script>
</script>

<style>
</style>