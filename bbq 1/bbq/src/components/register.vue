<script setup>
import { ref } from 'vue'
import axios from "axios";
import router from '@/router';
import {username} from "@/App.vue";
import{password}from "@/App.vue";
import{name}from "@/App.vue";
import{user_type}from "@/App.vue";
function registration() {
        let data = {
            "name": String(name.value),
            "username": String(username.value),
            "password": String(password.value),
            "user_type": Number(user_type.value)
        }
        axios.post("/api/reg", data).then((res) => {
            if (res.data.msg ==='success') {
                alert('注册成功');
                
            }
            else {
                alert(res.data.msg);
            };
        }).catch((err) => { alert(err);})
    };
    function ToLogin() {router.push('/login')}
</script>
<template>
    <!--注册页面-->
    <div class="registration-container">
    <form class="registration-form">
        <div class="form-group">
            <label for="name">姓名：</label>
            <input id="name" v-model="name" autocomplete="off" />
        </div>
        <div class="form-group">
            <label for="username">用户：</label>
            <input id="username" v-model="username" autocomplete="off" />
        </div>
        <div class="form-group">
            <label for="password">密码：</label>
            <input id="password" type="password" v-model="password" autocomplete="off" />
        </div>
        <div class="form-group">
            <input name="usertype" type="radio" value=1 v-model="usertype" id="student" />
            <label for="student">学生</label>
        </div>
        <div class="form-group">
            <input name="usertype" type="radio" value=2 v-model="usertype" id="admin" />
            <label for="admin">管理员</label>
        </div>
        <button type="button" class="btn-submit" @click="registration">注册</button>
        <button type="button" class="btn-login" @click="ToLogin">登录</button>
    </form>
</div>
</template>
<style>
.registration-container {
    max-width: 400px;
    margin: 50px auto;
    padding: 20px;
    background: #f7f7f7;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.registration-form {
    display: flex;
    flex-direction: column;
}

.form-group {
    margin-bottom: 15px;
}

.form-group label {
    margin-bottom: 5px;
    font-weight: bold;
}

.registration-form input[type="text"],
.registration-form input[type="password"] {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
    box-sizing: border-box;
}

.registration-form input[type="radio"] {
    margin-right: 10px;
}
</style>
