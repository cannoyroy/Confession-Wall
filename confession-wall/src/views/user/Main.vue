<template>
    <div class="forum-container">
      <div v-for="post in posts" :key="post.PostID" class="post">
        <div class="post-header">
          <h2>{{ post.PostID }}</h2>
          <p>By {{ post.SenderID }} on {{ post.CreatedAt }}</p>
        </div>
        <div class="post-body">
          <p>{{ post.Content }}</p>
        </div>
        <div class="post-footer">
          <!-- 可以添加更多帖子信息，如评论数、点赞数等 -->
        </div>
      </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute, RouterLink, RouterView } from 'vue-router';

const route = useRoute();
const posts = ref([]);


onMounted(async () => {
  try {
    const response = await axios.get('http://127.0.0.1:8080/main', {
      params: { user_id: route.query.user_id } // 假设 user_id 来自查询参数
    });

    posts.value = response.data.data.post_list;
    console.log(response.data)
    // console.log(route.query.user_type);
  } catch (error) {
    console.error('Failed to fetch posts:', error);
  }
});
</script>


<script>
export default {
    mounted() {
        let script = document.createElement('script');
        script.type = 'text/javascript';
        script.src = 'src/js/main.js';
        console.log("当前载入:"+script.src);
        document.body.appendChild(script);
    },
}
</script>

<style>
.forum-container {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  padding-left: 100px;
}

.post {
  border-bottom: 1px solid #ccc;
  padding: 15px 0;
}

.post-header {
  margin-bottom: 10px;
}

.post-header h2 {
  margin: 0;
  font-size: 1.5em;
}

.post-header p {
  margin: 5px 0 0 0;
  color: #666;
  font-size: 0.9em;
}

.post-body p {
  margin: 0;
}

.post-footer {
  margin-top: 10px;
}
</style>