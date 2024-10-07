<script setup>
import { useRoute} from 'vue-router';
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
const route = useRoute();

const item = ref(null);
let Userid = 0;

onMounted(async () => {

        axios.get("http://127.0.0.1:8080/profile", { params: {"user_id":route.query.user_id}}).then((res) => {
            if (res.data.msg === "success") {
                item.value = res.data.data.post_list
                console.log(item.value);
            }
            else {
                alert(res.data.msg);
            };
        }).catch((err) => { alert(err); })
    
});
const props = defineProps({
  post: Object
});

const genderStyle =  (gender) => {
  switch(gender) {
    case 'male':
      return { color: 'blue', fontSize: 'smaller' };
    case 'female':
      return { color: 'pink', fontSize: 'smaller' };
    case 'other':
      return { color: 'green', fontSize: 'smaller' };
    default:
      return {}; // 默认样式，可以根据需要设置
  }
};
</script>

<template>
     <div v-for="post in item">
  <header>
  <div class="container">
    <div class="profile">
      <div class="profile-image">
        <img :src="post.ImgURL">
      </div>
      <div class="profile-user-settings">
        <h1 class="profile-user-name">{{post.Username}}</h1>
        
        
        <button class="btn profile-edit-btn" @click="goToEditProfile">Edit Profile</button>

      </div>

      <!-- <div class="profile-stats">

        <ul>
          <li><span class="profile-stat-count">164</span> posts</li>
          <li><span class="profile-stat-count">188</span> followers</li>
          <li><span class="profile-stat-count">206</span> following</li>
        </ul>

      </div> -->
      <h1  :style="genderStyle(post.Gender)">{{ post.Gender }} - {{post.Region}}</h1>
      <div class="profile-bio">

        <p>{{post.OtherInfo}}</p>

      </div>

    </div>
    <!-- End of profile section -->

  </div>
  <!-- End of container -->

  </header>
<div v-if="post.ContactTele != '' || post.ContactQQ != '' || post.ContactWechat != '' || post.ContactOther != ''">
  <div class="contact-card">
    <div class="card-gradient">
      <div class="card-content">
        <!-- <h2>123</h2> -->
        <div v-if="post.ContactTele != ''"><p><strong>电话:</strong>{{post.ContactTele}}</p></div>
        <div v-if="post.ContactQQ != ''"><p><strong>QQ:</strong>{{post.ContactQQ}}</p></div>
        <div v-if="post.ContactWechat != ''"><p><strong>微信:</strong>{{post.ContactWechat}}</p></div>
        <div v-if="post.ContactOther != ''"><p><strong>其他联络方式:</strong>{{post.ContactOther}}</p></div>
      </div>
    </div>
  </div>
</div>


</div>


  <!-- THIS -->
   
  <RouterView />
</template>

<script>
export default {
  methods: {
    goToEditProfile() {
      this.$router.push({ path: '/editprofile' , query: this.$route.query});
    }
  },
  data() {
    return {
      item: {
        value: {
          ContactOther: "",
          ContactQQ: "",
          ContactTele: "",
          ContactWechat: "",
          CreatedAt: "2024-10-02T20:05:53.964+08:00",
          Gender: "male",
          ImgURL: "https://img0.baidu.com/it/u=2265288557,2266592009&fm=253&fmt=auto&app=138&f=JPEG?w=800&h=800",
          OtherInfo: "3",
          Region: "1",
          UserID: 111,
          Username: "Jack"
        }
      }
    };
  }
};
</script>

<style scoped>

.contact-card {
  margin-left: 40%;
  width: 300px;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.card-gradient {
  background-image: linear-gradient(to right, #6dd5ed, #2193b0);
  padding: 20px;
}

.card-content {
  color: white;
  text-align: center;
}

.card-content h2 {
  margin: 0;
  font-size: 24px;
}

.card-content p {
  margin: 10px 0;
  font-size: 16px;
}

/*

All grid code is placed in a 'supports' rule (feature query) at the bottom of the CSS (Line 310). 
        
The 'supports' rule will only run if your browser supports CSS grid.

Flexbox and floats are used as a fallback so that browsers which don't support grid will still recieve a similar layout.

*/

/* Base Styles */

:root {
  
    font-size: 10px;
}

*,
*::before,
*::after {
    box-sizing: border-box;
}

body {
    font-family: "Open Sans", Arial, sans-serif;
    min-height: 100vh;
    background-color: #fafafa;
    color: #262626;
    padding-bottom: 3rem;
}

img {
  margin-left: 5.6em;
    display: block;
}

.container {
    max-width: 93.5rem;
    margin: 0 auto;
    padding: 0 2rem;
}

.btn {
    display: inline-block;
    font: inherit;
    background: none;
    border: 0.5rem solid #000000;
    color: inherit;
    padding: 0;
    cursor: pointer;
}

.btn:focus {
    outline: 0.5rem auto #4d90fe;
}

.visually-hidden {
    position: absolute !important;
    height: 1px;
    width: 1px;
    overflow: hidden;
    clip: rect(1px, 1px, 1px, 1px);
}

/* Profile Section */

.profile {
    padding: 5rem 0;
}

.profile::after {
    content: "";
    display: block;
    clear: both;
}

.profile-image {
    float: left;
    width: calc(33.333% - 1rem);
    display: flex;
    justify-content: center;
    align-items: center;
    margin-right: 3rem;
}

.profile-image img {
    border-radius: 50%;
    width: 20vw;
    height: 20vw;
}

.profile-user-settings,
.profile-stats,
.profile-bio {
    float: left;
    width: calc(66.666% - 2rem);
}

.profile-user-settings {
    margin-top: 1.1rem;
}

.profile-user-name {
    display: inline-block;
    font-size: 3.2rem;
    font-weight: 300;
}

.profile-edit-btn {
    font-size: 1.4rem;
    line-height: 1.8;
    border: 0.1rem solid #dbdbdb;
    border-radius: 0.3rem;
    padding: 0 2.4rem;
    margin-left: 2rem;
}

.profile-settings-btn {
    font-size: 2rem;
    margin-left: 1rem;
}

.profile-stats {
    margin-top: 2.3rem;
}

.profile-stats li {
    display: inline-block;
    font-size: 1.6rem;
    line-height: 1.5;
    margin-right: 4rem;
    cursor: pointer;
}

.profile-stats li:last-of-type {
    margin-right: 0;
}

.profile-bio {
    font-size: 1.6rem;
    font-weight: 400;
    line-height: 1.5;
    margin-top: 2.3rem;
}

.profile-real-name,
.profile-stat-count,
.profile-edit-btn {
    font-weight: 600;
}

/* Gallery Section */

.gallery {
    display: flex;
    flex-wrap: wrap;
    margin: -1rem -1rem;
    padding-bottom: 3rem;
}

.gallery-item {
    position: relative;
    flex: 1 0 22rem;
    margin: 1rem;
    color: #fff;
    cursor: pointer;
}

.gallery-item:hover .gallery-item-info,
.gallery-item:focus .gallery-item-info {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.3);
}

.gallery-item-info {
    display: none;
}

.gallery-item-info li {
    display: inline-block;
    font-size: 1.7rem;
    font-weight: 600;
}

.gallery-item-likes {
    margin-right: 2.2rem;
}

.gallery-item-type {
    position: absolute;
    top: 1rem;
    right: 1rem;
    font-size: 2.5rem;
    text-shadow: 0.2rem 0.2rem 0.2rem rgba(0, 0, 0, 0.1);
}

.fa-clone,
.fa-comment {
    transform: rotateY(180deg);
}

.gallery-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

/* Loader */

.loader {
    width: 5rem;
    height: 5rem;
    border: 0.6rem solid #999;
    border-bottom-color: transparent;
    border-radius: 50%;
    margin: 0 auto;
    animation: loader 500ms linear infinite;
}

/* Media Query */

@media screen and (max-width: 40rem) {
    .profile {
        display: flex;
        flex-wrap: wrap;
        padding: 4rem 0;
    }

    .profile::after {
        display: none;
    }

    .profile-image,
    .profile-user-settings,
    .profile-bio,
    .profile-stats {
        float: none;
        width: auto;
    }

    .profile-image img {
        width: 7.7rem;
    }

    .profile-user-settings {
        flex-basis: calc(100% - 10.7rem);
        display: flex;
        flex-wrap: wrap;
        margin-top: 1rem;
    }

    .profile-user-name {
        font-size: 2.2rem;
    }

    .profile-edit-btn {
        order: 1;
        padding: 0;
        text-align: center;
        margin-top: 1rem;
    }

    .profile-edit-btn {
        margin-left: 0;
    }

    .profile-bio {
        font-size: 1.4rem;
        margin-top: 1.5rem;
    }

    .profile-edit-btn,
    .profile-bio,
    .profile-stats {
        flex-basis: 100%;
    }

    .profile-stats {
        order: 1;
        margin-top: 1.5rem;
    }

    .profile-stats ul {
        display: flex;
        text-align: center;
        padding: 1.2rem 0;
        border-top: 0.1rem solid #dadada;
        border-bottom: 0.1rem solid #dadada;
    }

    .profile-stats li {
        font-size: 1.4rem;
        flex: 1;
        margin: 0;
    }

    .profile-stat-count {
        display: block;
    }
}

/* Spinner Animation */

@keyframes loader {
    to {
        transform: rotate(360deg);
    }
}

/*

The following code will only run if your browser supports CSS grid.

Remove or comment-out the code block below to see how the browser will fall-back to flexbox & floated styling. 

*/

@supports (display: grid) {
    .profile {
        display: grid;
        grid-template-columns: 1fr 2fr;
        grid-template-rows: repeat(3, auto);
        grid-column-gap: 3rem;
        align-items: center;
    }

    .profile-image {
        grid-row: 1 / -1;
    }

    .gallery {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(22rem, 1fr));
        grid-gap: 2rem;
    }

    .profile-image,
    .profile-user-settings,
    .profile-stats,
    .profile-bio,
    .gallery-item,
    .gallery {
        width: auto;
        margin: 0;
    }

    @media (max-width: 40rem) {
        .profile {
            grid-template-columns: auto 1fr;
            grid-row-gap: 1.5rem;
        }

        .profile-image {
            grid-row: 1 / 2;
        }

        .profile-user-settings {
            display: grid;
            grid-template-columns: auto 1fr;
            grid-gap: 1rem;
        }

        .profile-edit-btn,
        .profile-stats,
        .profile-bio {
            grid-column: 1 / -1;
        }

        .profile-user-settings,
        .profile-edit-btn,
        .profile-settings-btn,
        .profile-bio,
        .profile-stats {
            margin: 0;
        }
    }
}


</style>