<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute, useRouter, RouterLink, RouterView } from 'vue-router';

const route = useRoute();
const posts = ref([]);
const router = useRouter();


const reportReason_1 = ref('');
const characterLimit = 250; // 定义字符限制

function limitCharacters() {
  if (reportReason_1.value.length > characterLimit) {
    reportReason_1.value = reportReason_1.value.substring(0, characterLimit);
  }
}



onMounted(async () => {
        try {
          const response = await axios.get('http://127.0.0.1:8080/main', {
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




function InBlack(id) {
  // console.log(id)
  if (id == route.query.user_id){
    alert("请不要拉黑自己")
    return 
  }
      axios.post("http://127.0.0.1:8080/blacklist/new", {
          user_id: String(route.query.user_id), 
          blocked_id : String(id)
        })
        .then(response => {
          // console.log(Number(id))
          // console.log(5555)
          location.reload();
          console.log(response.data)
        })
        .catch(error => {
          console.error('Failed to create post:', error);
          // 处理发帖失败的逻辑，例如显示错误消息
          alert("无法拉黑，请稍后再试。")
        });
}


const reportReason = ref('');
const reportSenderId = ref(null);
const reportModalVisible = ref(false);

function openReportModal(senderId) {
  reportSenderId.value = senderId;
  reportModalVisible.value = true;
  // 显示模态框的逻辑，例如通过修改CSS来显示
  document.getElementById('reportModal').style.display = 'block';
}

function closeReportModal() {
  reportModalVisible.value = false;
  // 隐藏模态框的逻辑
  document.getElementById('reportModal').style.display = 'none';
}

async function submitReport() {
  const postData = {
      reporter_id: parseInt(route.query.user_id),
      post_id: parseInt(reportSenderId.value),
      reason: reportReason.value
  };
  try {
    const response = await axios.post("http://127.0.0.1:8080/report", postData);
    console.log(response.data);
    alert('举报成功！');
    console.log(postData.reporter_id)
    closeReportModal();
  } catch (error) {
    // console.log(postData)
    console.error('Failed to submit report:', error);
    alert('举报失败，请稍后再试。');
    // closeReportModal();
  }
}

function navigateToProfile(userID) {
      // 假设您有一个名为 'UserProfile' 的路由，并且需要传递 'user_id' 作为查询参数
      router.push({
        name: 'profilec', // 路由名称，确保在路由配置中定义了该名称
        query: { user_id: userID } // 查询参数
      });
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
          <div v-for="post in posts" :key="post.PostID">
          <tr>
            <td style="margin-left: 5em;">
              <div class="badge badge-success badge-success-alt" style="margin-right: 0.05em;" @click="toggleExpand(post)">{{ post.isExpanded ? '收起' : '展开' }}</div>
              <div v-if="post.SenderID != route.query.user_id" class="badge badge-success badge-success-alt_1" style="margin-right: 0.05em; margin-top: 0.5em;" @click="InBlack(post.SenderID)">拉黑</div>
              <div v-if="post.SenderID != route.query.user_id" class="badge badge-success badge-success-alt_1" style="margin-right: 0.05em; margin-top: 0.5em;" @click="openReportModal(post.PostID)">举报</div>

            </td>
            <td>
              <a href="#">
                <div class="d-flex align-items-center">
                  <!-- <div class="avatar avatar-blue mr-3">EB</div> -->

                  <div class="">
                    <!-- <p class="font-weight-bold mb-0">{{ post.SenderID }}</p> -->
                    <p class="text-muted mb-0" @click="navigateToProfile(post.SenderID)">{{ post.Username }}</p>
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
              <td>{{ post.CreatedAt.slice(0, 10)+"  "+post.CreatedAt.slice(11, 16) }}</td>
            </td>
          </tr>
        </div>

        </tbody>
      </table>
    </div>
  </div>
</div>

<div id="reportModal" class="modal" style="display: none;">
  <div class="modal-content">
    <span class="close" @click="closeReportModal">&times;</span>
    <div class="form-container">
      <p>请输入举报理由：</p>
      <textarea
        v-model="reportReason_1"
        id="reportReason_1"
        placeholder="请输入举报原因，不得超过250字"
        @input="limitCharacters"
      ></textarea>

      <button class="submit-button" @click="submitReport">提交举报</button>
    </div>
  </div>
</div>

</template>

<script>
export default {
  setup() {
    // function navigateToProfile(userID) {
    //   // 假设您有一个名为 'UserProfile' 的路由，并且需要传递 'user_id' 作为查询参数
    //   router.push({
    //     name: 'profilec', // 路由名称，确保在路由配置中定义了该名称
    //     query: { user_id: userID } // 查询参数
    //   });
    // }
  },
    mounted() {
        let script = document.createElement('script');
        script.type = 'text/javascript';
        script.src = 'src/js/main.js';
        console.log("当前载入:"+script.src);
        document.body.appendChild(script);
    },
}
</script>

<style scoped>
body {
  margin-left: 20em;
  background: #f7f7f7;
}

/* 模态框样式 */
.modal {
  display: none; /* Hidden by default */
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
  max-width: 500px; /* 限制最大宽度 */
  height: auto; /* 根据内容自动调整高度 */
  position: relative; /* 为绝对定位的按钮提供参照 */
}

.modal-content .form-container {
  display: flex;
  flex-direction: column;
  align-items: stretch; /* 使输入框和按钮占满容器宽度 */
}

.modal-content textarea {
  width: 100%; /* 使输入框占满容器宽度 */
  height: 150px; /* 固定输入框的高度 */
  margin-bottom: 10px; /* 在输入框和按钮之间添加一些间隔 */
  resize: none; /* 禁止调整输入框大小 */
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