<template>
  <header>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />
    <div class="wrapper">
      <HelloWorld msg="You did it!" />
    </div>
  </header>

  <main>
    <TheWelcome />
    <!-- 帖子列表 -->
    <div class="posts">
      <!-- 帖子块 -->
      <div class="post" v-for="post in posts" :key="post.id">
        <!-- 渐变边框 -->
        <div class="gradient-border">
          <div class="post-content">
            <h3>{{ post.title }}</h3>
            <p>{{ post.content }}</p>
            <p>发送人: <router-link :to="`/profile?user_id=${post.userId}`">{{ post.username }}</router-link></p>
            <p>创建时间: {{ post.createdAt }}</p>
            <button-component icon="fa-eye" @click="openModal(post)">查看完整帖子</button-component>
          </div>
        </div>
      </div>
    </div>

    <!-- 模态框 -->
    <div class="modal-overlay" v-if="selectedPost">
      <div class="modal" @click="close">
        <div class="gradient-border">
          <h2>{{ selectedPost.title }}</h2>
          <p>{{ selectedPost.content }}</p>
          <p>发送人: <router-link :to="`/profile?user_id=${selectedPost.userId}`" @click.native="close">{{ selectedPost.username }}</router-link></p>
          <p>创建时间: {{ selectedPost.createdAt }}</p>
          <button-component icon="fa-ban" color="facebook" @click="addToBlacklist">拉入黑名单</button-component>
          <button-component icon="fa-flag" color="twitter" @click="reportPost">举报</button-component>
          <button-component icon="fa-times" color="google" @click="close">关闭</button-component>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup>
import { ref } from 'vue';
import HelloWorld from './components/HelloWorld.vue';
import TheWelcome from './components/TheWelcome.vue';
import ButtonComponent from './components/Button.vue'; // 导入按钮组件
import { useRouter } from 'vue-router';

const posts = ref([
  // 假设这里有一些帖子数据
]);
const selectedPost = ref(null);
const router = useRouter();

const openModal = (post) => {
  selectedPost.value = post;
};

const close = () => {
  selectedPost.value = null;
};

const addToBlacklist = () => {
  // 实现拉入黑名单的逻辑
  console.log('拉入黑名单', selectedPost.value.userId);
};

const reportPost = () => {
  // 实现举报的逻辑
  console.log('举报帖子', selectedPost.value.id);
};
</script>

<style scoped>
/* ...其他样式... */

.posts {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.post {
  width: 70vw; /* 设置宽度为视口宽度的70% */
  max-width: 70vw; /* 设置最大宽度为视口宽度的70% */
  margin-bottom: 1rem; /* 设置帖子块之间的间距 */
}

/* 渐变边框样式 */
.gradient-border {
  --border-width: 1px;
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 1rem;
  border-radius: 10px;
  overflow: hidden;
  width: 100%; /* 帖子块宽度充满容器 */
}

.gradient-border::after {
  content: "";
  position: absolute;
  z-index: -1;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    60deg,
    hsl(224, 85%, 66%),
    hsl(269, 85%, 66%),
    hsl(314, 85%, 66%),
    hsl(359, 85%, 66%),
    hsl(44, 85%, 66%),
    hsl(89, 85%, 66%),
    hsl(134, 85%, 66%),
    hsl(179, 85%, 66%)
  );
  background-size: 300% 300%;
  background-position: 0 50%;
  border-radius: 10px;
  animation: moveGradient 4s alternate infinite;
}

@keyframes moveGradient {
  50% {
    background-position: 100% 50%;
  }
}
</style>