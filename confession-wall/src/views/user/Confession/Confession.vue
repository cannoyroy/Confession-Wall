<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute, useRouter, RouterLink, RouterView } from 'vue-router';

const route = useRoute();
const posts = ref([]);
const router = useRouter();

const postContent = ref('');
const characterLimit = 250; // 定义字符限制

function limitCharacters() {
  if (postContent.value.length > characterLimit) {
    postContent.value = postContent.value.substring(0, characterLimit);
  }
}

onMounted(async () => {
        try {
          const response = await axios.get('http://127.0.0.1:8080/confessions/manage', {
            params: { user_id: route.query.user_id } // 假设 user_id 来自查询参数
          });

          posts.value = response.data.data.post_list.map(post => ({
              ...post,
              isExpanded: false
            }));
          console.log(response.data.data.post_list)
          // console.log(posts.value);
        } catch (error) {
          console.error('Failed to fetch posts:', error);
        }
});

function toggleExpand(post) {
  post.isExpanded = !post.isExpanded;
}


</script>

<template>
  
  <div class="container">
  <div class="row py-5">
    <div class="col-12">
      <table id="example" class="table table-hover responsive nowrap" style="width:100%">
        <!-- <thead>
          <tr>
            <th>555</th>
            <th>Phone Number</th>
            <th>Profession</th>
            <th>Date of Birth</th>
            <th>App Access</th>
            <th>Actions</th>
          </tr>
        </thead> -->
        <tbody>
          <div class="badge badge-success badge-success-alt" style="margin-right: 0.05em; margin-top: 0.5em;" @click="showNewModal = true">创建</div>
          <div v-for="post in posts" :key="post.PostID">
          <tr>
            <td style="margin-left: 5em;">
              <div class="badge badge-success badge-success-alt" style="margin-right: 0.05em;" @click="toggleExpand(post)">{{ post.isExpanded ? '收起' : '展开' }}</div>
              <div class="badge badge-success badge-success-alt_1" style="margin-right: 0.05em; margin-top: 0.5em;" @click="editFormData.id = post.PostID; showEditModal = true">修改</div>
              <div class="badge badge-success badge-success-alt_1" style="margin-right: 0.05em; margin-top: 0.5em;" @click="Delete(post.PostID)">删除</div>
            </td>
            <td>
              <a href="#">
                <div class="d-flex align-items-center">
                  <!-- <div class="avatar avatar-blue mr-3">EB</div> -->

                  <div class="">
                    <!-- <p class="font-weight-bold mb-0">{{ post.SenderID }}</p> -->
                    <!-- <p class="text-muted mb-0" @click="navigateToProfile(post.SenderID)">{{ post.Username }}</p> -->
                  </div>
                </div>
              </a>
            </td>

            <!-- <td>(7565667 8768</td> -->
            <td>
              <div v-if="post.isExpanded">{{ post.Content }}</div>
              <div v-else>{{ post.Content.slice(0, 60) }}...</div>
            </td>
            <!-- <td>09/0411/1996</td> -->
            
            <td>
              <div class="dropdown">
                <!-- <button class="btn btn-sm btn-icon" type="button" id="dropdownMenuButton2" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <i class="bx bx-dots-horizontal-rounded" data-toggle="tooltip" data-placement="top"
                        title="Actions"></i>
                    </button> -->
                <div class="dropdown-menu" aria-labelledby="dropdownMenuButton2">
                  <!-- <a class="dropdown-item" href="#"><i class="bx bxs-pencil mr-2"></i>拉黑该用户</a> -->
                  <!-- <a class="dropdown-item text-danger" href="#"><i class="bx bxs-trash mr-2"></i> 拉黑该用户 </a> -->
                </div>
              </div>
              <td>{{ post.ScheduledAt.slice(0, 10)+"  "+post.ScheduledAt.slice(11, 16) }}</td>
              <td><span v-if="post.Anonymous === 1">匿名中</span>
              <span v-else>非匿名</span> <!-- 或者你想显示的其他内容 --></td>
              <td><<span v-if="isTimeToPublish([post.ScheduledAt])">
                  正常发布
                </span><span v-else>
              未到发布时间
            </span></td>

              <span v-if="post.ReportStatus === 1">帖子举报审核中</span>
              <span v-if="post.ReportStatus === 2">帖子被举报成功</span>
              
            </td>
          </tr>
        </div>

        </tbody>
      </table>
    </div>
  </div>
</div>


    <!-- 弹窗模态框 -->
    <div v-if="showEditModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="showEditModal = false">&times;</span>
        <h2>修改</h2>
        <form @submit.prevent="submitEdit">

          <label for="postContent">帖子内容:</label>
          <textarea id="postContent" v-model="editFormData.postContent" @input="truncatePostContent" required></textarea>

          <label for="scheduledTime">定时发送时间 (格式: YYYY-MM-DD HH:mm):</label>
          <input type="text" id="scheduledTime" v-model="editFormData.scheduledTime" required>
          
          <label for="isAnonymous">是否匿名:</label>
          <select id="isAnonymous" v-model="editFormData.isAnonymous">
            <option value="true">是</option>
            <option value="false">否</option>
          </select>
          
          <button type="submit">提交</button>
        </form>
      </div>
    </div>


    <!-- 弹窗模态框 -->
    <div v-if="showNewModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="showNewModal = false">&times;</span>
        <h2>创建</h2>
        <form @submit.prevent="New">

          <label for="postContent">帖子内容:</label>
          <textarea id="postContent" v-model="newFormData.postContent" placeholder="不得超过250字" @input="truncatePostContent" required></textarea>

          <label for="scheduledTime">定时发送时间 (格式: YYYY-MM-DD HH:mm):</label>
          <input type="text" id="scheduledTime" v-model="newFormData.scheduledTime" required>
          
          <label for="isAnonymous">是否匿名:</label>
          <select id="isAnonymous" v-model="newFormData.isAnonymous">
            <option value="true">是</option>
            <option value="false">否</option>
          </select>
          
          <button type="submit">提交</button>
        </form>
      </div>
    </div>


</template>

<script>
export default {
  data() {
    return {
      showEditModal: false,
      showNewModal: false,
      editFormData: {
        id: 1, //帖子ID
        postContent: 'none', // 假设帖子内容在post.content属性中
        scheduledTime: '2024-09-26 20:54', // 默认值
        isAnonymous: true // 默认值
      },
      newFormData: {
        postContent: 'none', // 假设帖子内容在post.content属性中
        scheduledTime: '2024-09-26 20:54', // 默认值
        isAnonymous: false // 默认值
      }
    };
  },
    mounted() {
        let script = document.createElement('script');
        script.type = 'text/javascript';
        script.src = 'src/js/main.js';
        console.log("当前载入:"+script.src);
        document.body.appendChild(script);
    },
    methods: {
      truncatePostContent() {
        if (this.newFormData.postContent.length > 250) {
          this.newFormData.postContent = this.newFormData.postContent.substring(0, 250);
        }
        if (this.editFormData.postContent.length > 250) {
          this.editFormData.postContent = this.editFormData.postContent.substring(0, 250);
        }
      },
      isTimeToPublish(scheduledAt) {
        const scheduledTime = new Date(scheduledAt);
        const currentTime = new Date();
        // 比较当前时间是否大于等于发布时间
        return currentTime >= scheduledTime;
      },
      submitEdit() {
        console.log(this.editFormData.isAnonymous == true ? 1 : 0)
        
        const postData = {
          post_id: this.editFormData.id,
          content: this.editFormData.postContent,
          scheduled_time: this.editFormData.scheduledTime,
          anonymous: this.editFormData.isAnonymous == true ? 1 : 0
        };

        axios.put("http://127.0.0.1:8080/confessions/manage/edit", postData)
        .then(response => {
          console.log('Post created successfully:', response.data);
          alert("帖子修改成功")
          // this.$message.alert("修改成功")
          // 处理发帖成功的逻辑，例如显示成功消息或跳转到帖子列表
          this.content = ''
          this.post_id = ''
          this.showEditModal = ''

        });
        // .catch(error => {
        //   console.error('Fail', error);
        //   alert("无法修改，请确认您想修改的帖子是属于您的或者帖子是否存在。")
        // });
      },
      Delete(postData) {
        console.log(postData)
        axios.delete("http://127.0.0.1:8080/confessions/manage/delete", {params:{post_id: postData}})
        .then(response => {
          console.log('Post created successfully:', response.data);
          alert("帖子删除成功")
          // 处理发帖成功的逻辑，例如显示成功消息或跳转到帖子列表
          this.post_id = ''
        })
      },
      New(){
        const query = this.$route.query;
      const postData = {
        user_id: parseInt(query.user_id),
        content: this.newFormData.postContent,
        ScheduledAt: this.newFormData.scheduledTime,
        Anonymous: this.newFormData.isAnonymous == true ? 1 : 0
      };

      axios.post("http://127.0.0.1:8080/confessions/new", postData)
        .then(response => {
          console.log('Post created successfully:', response.data);
          alert("PostID:"+" | 帖子发送成功")
          // 处理发帖成功的逻辑，例如显示成功消息或跳转到帖子列表
          this.content = '';
          this.characterCount = 0;
          this.showNewModal = 0;
        })
      }
  }
}
</script>

<style scoped>
body {
  margin-left: 20em;
  background: #f7f7f7;
}


/* 模态框样式 */
.modal {
  display: block; /* Hidden by default */
  position: fixed; /* Stay in place */
  z-index: 1; /* Sit on top */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0,0,0); /* Fallback color */
  background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto; /* 15% from the top and centered */
  padding: 20px;
  border: 1px solid #888;
  width: 80%; /* Could be more or less, depending on screen size */
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}


.modal-content .submit-button {
  width: 100%; /* 使按钮占满容器宽度 */
  padding: 10px;
  background-color: #b2d7ff;
  color: white;
  border: none;
  cursor: pointer;
}

.modal-content .submit-button:hover {
  background-color: #0056b3;
}

.table {
  margin-left: 6em;
  border-spacing: 0 0.3rem !important;
  width: 50vw;
  height: 50vh;
}

.table .dropdown {
  display: inline-block;
}

.table td,
.table th {
  vertical-align: middle;
  margin-bottom: 2px;
  border: none;
}

.table thead tr,
.table thead th {
  border: none;
  font-size: 12px;
  letter-spacing: 1px;
  text-transform: uppercase;
  background: transparent;
}

.table td {
  word-wrap: break-word; /* 允许长单词或非标准字符在到达边界时换行 */
  overflow-wrap: break-word; /* CSS3 的属性，与 word-wrap 类似，推荐使用 */
  width: 40em;
  /* max-height: 10em; */
  background: rgba(255, 255, 255, 0.5);
  /* margin-right: 1em; */
  /* margin-left: 5em; */
}

.table td:first-child {
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
  width: 2%;
}

.table td:nth-child(2) {
  /* border-top-left-radius: 10px;
  border-bottom-left-radius: 10px; */
  width: 6%;
}

.table td:nth-child(3) {
  /* border-top-left-radius: 10px; */
  /* border-bottom-left-radius: 10px; */
  width: 40%;
}

.table td:last-child {
  border-top-right-radius: 10px;
  border-bottom-right-radius: 10px;
  width: 5%;
}

.avatar {
  width: 2.75rem;
  height: 2.75rem;
  line-height: 3rem;
  border-radius: 50%;
  display: inline-block;
  background: transparent;
  position: relative;
  text-align: center;
  color: #868e96;
  font-weight: 700;
  vertical-align: bottom;
  font-size: 1rem;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.avatar-sm {
  width: 2.5rem;
  height: 2.5rem;
  font-size: 0.83333rem;
  line-height: 1.5;
}

.avatar-img {
  width: 100%;
  height: 100%;
  -o-object-fit: cover;
  object-fit: cover;
}

.avatar-blue {
  background-color: #c8d9f1;
  color: #467fcf;
}

table.dataTable.dtr-inline.collapsed
  > tbody
  > tr[role="row"]
  > td:first-child:before,
table.dataTable.dtr-inline.collapsed
  > tbody
  > tr[role="row"]
  > th:first-child:before {
  top: 28px;
  left: 14px;
  border: none;
  box-shadow: none;
}

table.dataTable.dtr-inline.collapsed > tbody > tr[role="row"] > td:first-child,
table.dataTable.dtr-inline.collapsed > tbody > tr[role="row"] > th:first-child {
  padding-left: 48px;
}

table.dataTable > tbody > tr.child ul.dtr-details {
  width: 100%;
}

table.dataTable > tbody > tr.child span.dtr-title {
  min-width: 50%;
}

table.dataTable.dtr-inline.collapsed > tbody > tr > td.child,
table.dataTable.dtr-inline.collapsed > tbody > tr > th.child,
table.dataTable.dtr-inline.collapsed > tbody > tr > td.dataTables_empty {
  padding: 0.75rem 1rem 0.125rem;
}

div.dataTables_wrapper div.dataTables_length label,
div.dataTables_wrapper div.dataTables_filter label {
  margin-bottom: 0;
}

@media (max-width: 767px) {
  div.dataTables_wrapper div.dataTables_paginate ul.pagination {
    -ms-flex-pack: center !important;
    justify-content: center !important;
    margin-top: 1rem;
  }
}

.btn-icon {
  background: #fff;
}
.btn-icon .bx {
  font-size: 20px;
}

.btn .bx {
  vertical-align: middle;
  font-size: 20px;
}

.dropdown-menu {
  padding: 0.25rem 0;
}

.dropdown-item {
  padding: 0.5rem 1rem;
}

.badge {
  padding: 0.5em 0.75em;
}

.badge-success-alt {
  background-color: #d7f2c2;
  color: #7bd235;
}

.badge-success-alt_1 {
  background-color: #f2c2c2;
  color: #d2354a;
}

.table a {
  color: #212529;
}

.table a:hover,
.table a:focus {
  text-decoration: none;
}

table.dataTable {
  margin-top: 12px !important;
}

.icon > .bx {
  display: block;
  min-width: 1.5em;
  min-height: 1.5em;
  text-align: center;
  font-size: 1.0625rem;
}

.btn {
  font-size: 0.9375rem;
  font-weight: 500;
  padding: 0.5rem 0.75rem;
}

.avatar-blue {
      background-color: #c8d9f1;
      color: #467fcf;
    }

    .avatar-pink {
      background-color: #fcd3e1;
      color: #f66d9b;
    }
</style>