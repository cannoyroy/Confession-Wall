<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute, useRouter, RouterLink, RouterView } from 'vue-router';

const route = useRoute();
const posts = ref([]);
const router = useRouter();


onMounted(async () => {
        try {
          const response = await axios.get('http://127.0.0.1:8080/blacklist/manage', {
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

function OutBlack(BlockedID, UserID) {
  axios.delete("http://127.0.0.1:8080/blacklist/delete", {params:{blocked_id: BlockedID, user_id: UserID}})
        .then(response => {
          console.log('Post created successfully:', response.data);
          alert("删除成功")
          // 处理发帖成功的逻辑，例如显示成功消息或跳转到帖子列表
          location.reload();
        })
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
              <div class="badge badge-success badge-success-alt_1" style="margin-right: 0.05em; margin-top: 0.5em;" @click="OutBlack(post.BlockedID, post.UserID)">删除</div>
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
              <div>{{ post.BlockedUsername }}</div>
            </td>
            <!-- <td>09/0411/1996</td> -->
            
          </tr>
        </div>

        </tbody>
      </table>
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