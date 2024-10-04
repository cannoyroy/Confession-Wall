<script setup>
import { useRoute} from 'vue-router';
import { ref } from "vue";
import axios from "axios";
const route = useRoute();
const username = ref("");
const password = ref("");
const name = ref("");
const usertype = ref(1);
const listPost = ref(null);
const listFeedback = ref(null);
let Userid = 0
let Usertype = 0
onMounted(() => {
axiosGetPostlist();
})
function formatTime(date) {
const year = date.getFullYear();
const month = (date.getMonth() + 1).toString().padStart(2, '0');
const day = date.getDate().toString().padStart(2, '0');
const hours = date.getHours().toString().padStart(2, '0');
const minutes = date.getMinutes().toString().padStart(2, '0');
return `${year}-${month}-${day} ${hours}:${minutes}`;
}


//管理
function axiosGetPostlist() {
        axios.get("http://127.0.0.1:8080/confessions/manage",{params:{"user_id":Userid}}).then((res) => {
            if (res.data.code === 200) {
                listPost.value = res.data.data.post_list
            }
            else {
                alert(res.data.msg);
            };
        }).catch((err) => { alert(err); })
    };
//发布新表白
function axiosPostPutupPost(event) {
        const now = new Date();
        const formattedTime = this.formatTime(now);
        let data = {
            "content": event.target.parentNode.children[0].value,
            "user_id": Userid,
            "ScheduledAt": formattedTime,
            "Anonymous": 0
        }
        if (data.content.trim() !== "") {
            axios.post("http://127.0.0.1:8080/confessions/new", data).then((res) => {
                if (res.data.code === 200) {
                    alert("发布成功");
                    event.target.parentNode.children[0].value = "";
                    
                }
                else {
                    alert(res.data.msg);
                };
            }).catch((err) => { alert(err); })
        }
        else
        {
            alert("不能发送空帖子")
        }
    };
//删除表白
function axiosDeleteDeletePost(postId) {
        let data = {
            params: {
                "user_id": Userid,
                "post_id": postId
            }
        }
        axios.delete("http://127.0.0.1:8080/confessions/manage/delete", data).then((res) => {
            if (res.data.code === 200) {
                alert("删除成功");
                
            }
            else {
                alert(res.data.msg);
            };
        }).catch((err) => { alert(err); })
    };
//修改表白
function axiosPutEditPost(postId,event) {
        const now = new Date();
        const formattedTime = this.formatTime(now);
        let data = {
            "user_id": storedUserid,
            "post_id": postId,
            "content": event.target.parentNode.children[4].value,
            "Anonymous": 0,
            scheduled_time:formattedTime,
        }
        if (data.content.trim() !== "") {

            axios.put("http://127.0.0.1:8080/confessions/manage/edit", data).then((res) => {
                if (res.data.code === 200) {
                    alert("修改成功");
                    event.target.parentNode.children[4].value = "";
                    
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
  <h1 style="margin-left: 2em;">Confessoin页面</h1>
  <RouterView />
  <div class="page" >
            <div class="oneBlock">
                <textarea></textarea>
                <button @click="axiosPostPutupPost">发布新表白</button>
            </div>
            <div v-for="item in listPost">
                <div class="oneBlock">
                    <p>帖子序号:{{item.post_id}}</p>
                    <p>楼主:{{item.sender_id}}</p>
                    <p>创建时间:{{item.create_time}}</p>
                    <p>内容:{{item.content}}</p>

                    <button v-if="item.user_id===Userid" @click="axiosDeleteDeletePost(item.id)">删除</button>
                    <textarea></textarea>
                    <button v-if="item.user_id===Userid" @click="axiosPutEditPost(item.id, $event)">修改</button>
                    
                </div>
            </div>
        </div>
</template>

<script>
</script>

<style>
.page {
  background-color: #fff0f7; 
  color: #5c1e0f; 
  font-family: 'Arial', sans-serif; 
  padding: 20px;
  border-radius: 10px; 
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); 
}

button {
  background-color: #ff69b4; 
  color: white;
  border: none;
  padding: 10px 20px;
  margin-top: 10px;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #ff1493; 
}

.oneBlock {
  background-color: #ffe4e1; 
  border: 1px solid #ff69b4; 
  padding: 15px;
  margin-bottom: 20px;
  border-radius: 5px;
}

textarea {
  width: 100%;
  height: 100px;
  margin-bottom: 10px;
  padding: 10px;
  border: 1px solid #ff69b4;
  border-radius: 5px;
}

p {
  margin: 5px 0;
}


p:first-of-type {
  font-weight: bold;
  color: #9b1e1e; 
}
</style>