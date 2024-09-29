<script setup>
import { ref } from 'vue'
import axios from "axios";
import router from '@/router';
import {username} from "@/App.vue";
import{password}from "@/App.vue";
import{name}from "@/App.vue";
import{user_type}from "@/App.vue";
let Userid = 0
let Usertype = 0
function Login() {
        let data = {
            "username": String(username.value),
            "password": String(password.value)
        }
        
        axios.post("/api/login", data).then((res) => {
            console.log(res);
            if (res.data.msg === 'success') {
                alert('登录成功');
                Userid = res.data.data.user_id;
                Usertype = res.data.data.user_type;
                if (Usertype === 2) {
                    ToAdmin();
                }
                else {
                    ToStuPost();
                };
            }
            else {
                alert(res.data.msg);
            };
        }).catch((err) => { alert(err);})
    };
</script>
<template>
<body>
    <div class="login-container">
        <h1 class="logo">bbq</h1>
        <form>
            <input type="text" placeholder="账号" required v-model="username">
            <input type="password" placeholder="密码" required v-model="password">
            <button type="button"@click="Login" >登录</button>
        </form>
        <a href="#" class="forgot-password">忘记密码？</a>
        <div class="signup">
            没有账号？<a href="#">注册</a>
        </div>
    </div>
</body>
</template>

<style>
        body {
            font-family: Arial, sans-serif;
            background-color: #fafafa;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .login-container {
            background-color: #fff;
            border: 1px solid #dbdbdb;
            border-radius: 1px;
            padding: 20px 40px;
            text-align: center;
            width: 350px;
        }
        .logo {
            font-family: 'Brush Script MT', cursive;
            font-size: 40px;
            margin-bottom: 20px;
        }
        input {
            width: 100%;
            padding: 10px;
            margin-bottom: 10px;
            border: 1px solid #dbdbdb;
            border-radius: 3px;
            box-sizing: border-box;
        }
        button {
            width: 100%;
            padding: 10px;
            background-color: #0095f6;
            color: white;
            border: none;
            border-radius: 3px;
            font-weight: bold;
            cursor: pointer;
        }
        .forgot-password {
            color: #00376b;
            font-size: 12px;
            text-decoration: none;
            margin-top: 12px;
            display: inline-block;
        }
        .signup {
            margin-top: 20px;
            font-size: 14px;
        }
        .signup a {
            color: #0095f6;
            text-decoration: none;
            font-weight: bold;
        }
    </style>


