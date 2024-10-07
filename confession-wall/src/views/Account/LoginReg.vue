<template>
    <div class="container">
    <div class="welcome">
        <div class="pinkbox">
        <div class="signup nodisplay">
            <h1>register</h1>
            <form autocomplete="off" @submit.prevent="register">
            <input type="text" placeholder="username" id="usernameR" v-model="usernameR" @input="limitCharacters_1" required>
            <!-- <input type="email" placeholder="email"> -->
            <input type="password" placeholder="password" v-model="passwordR" @input="limitCharacters_1" required>
            <input type="password" placeholder="confirm password" v-model="passwordRC" @input="limitCharacters_1" required>
            <button class="button submit">create account </button>
            </form>
        </div>
        <div class="signin">
            <h1>sign in</h1>
            <form class="more-padding" autocomplete="off" @submit.prevent="login">
            <input type="text" placeholder="username" v-model="username" @input="limitCharacters_2" required>
            <input type="password" placeholder="password" v-model="password" @input="limitCharacters_2" required>
            <div class="checkbox">
                <input type="checkbox" id="remember" v-model="isRemember"/><label for="remember">remember me</label>
            </div>

            <button class="button submit">login</button>
            </form>
        </div>
        </div>
        <div class="leftbox">
        <h2 class="title"><span>MOM</span>ENTO</h2>
        <p class="desc"><span>ROMANCE</span>&<br>REMEMBRANCE</p>
        <p class="account"></p>
        <img class="flower smaller" src="https://image.ibb.co/d5X6pn/1357d638624297b.jpg" alt="1357d638624297b" border="0">
        <p class="account">用户需在4-19位，密码需在9-15位</p>
        <button class="button" id="signin">login</button>
        </div>
        <div class="rightbox">
          <h2 class="title"><span>MOM</span>ENTO</h2>
          <p class="desc"><span>ROMANCE</span>&<br>REMEMBRANCE</p>
          
        <img class="flower" src="https://preview.ibb.co/jvu2Un/0057c1c1bab51a0.jpg"/>
        <p class="account">don't have an account?</p>
        <button class="button" id="signup">sign up</button>
        </div>
    </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      password: '',
      usernameR: '',
      passwordR: '',
      passwordRC: '',
      isRemember: 0
    };
  },
  mounted() {
    localStorage.setItem('token', 0); // 每次加载页面时初始化 token 为 0
    console.log('in login page')
    console.log('Token:', localStorage.getItem('token'), 'Expires at:', localStorage.getItem('expiresAt'));

    let script = document.createElement('script');
        script.type = 'text/javascript';
        script.src = 'src/js/accounts.js';
        console.log("当前载入:"+script.src);
        document.body.appendChild(script);
  },
  methods: {
    limitCharacters_1(event) {
      if (this.usernameR.length > 19) {
        // 如果输入长度超过50，则截断输入值
        event.target.value = this.usernameR.substring(0, 19);
        // 同步v-model绑定的数据
        this.usernameR = event.target.value;
      }
      if (this.passwordR.length > 15) {
        event.target.value = this.passwordR.substring(0, 15);
        // 同步v-model绑定的数据
        this.passwordR = event.target.value;
      }
      if (this.passwordRC.length > 15) {
        event.target.value = this.passwordRC.substring(0, 15);
        // 同步v-model绑定的数据
        this.passwordRC = event.target.value;
      }
    },
    limitCharacters_2(event) {
      if (this.username.length > 19) {
        // 如果输入长度超过50，则截断输入值
        event.target.value = this.username.substring(0, 19);
        // 同步v-model绑定的数据
        this.username = event.target.value;
      }
      if (this.password.length > 15) {
        event.target.value = this.password.substring(0, 15);
        // 同步v-model绑定的数据
        this.password = event.target.value;
      }
    },
    login() {
      const userData = {
        username: this.username,
        password: this.password
      };

      axios.post("http://127.0.0.1:8080/login", userData)
        .then(response => {
          // 处理登录成功的响应
          console.log(response.data);
          // 这里可以根据实际情况进行页面跳转或状态更新
          if (response.data.msg === "success") {

            const userInfo = {
              user_id: response.data.data.user_id,
              username: response.data.data.username,
              type: 1 // 1表示用户 2表示Admin
            };
            console.log(userInfo)

            if (response.data.data.user_type == 1) {
              alert("登陆成功\n点击图标随时返回主界面");
              const token = true; // 假设返回的 token 存在 response.data.token
              var cnt = 60;
              if (this.isRemember == true){
                cnt = 60*24*7;
                // alert("222")
              }else{
                // alert("333")
                cnt = 60*2;
              }
              const expiresIn = cnt; // 设置有效期为60秒（1小时）
              // alert(expiresIn)
              const expiresAt = Date.now() + expiresIn * 1000; // 计算过期时间

              localStorage.setItem('token', token);
              localStorage.setItem('expiresAt', expiresAt); // 存储过期时间
              
              userInfo.type = 1;

              axios.get("http://127.0.0.1:8080/profile", {params: { user_id: response.data.data.user_id}})
                .then(response => {
                  if (!Array.isArray(response.data.data.post_list) || response.data.data.post_list.length === 0) {
                    alert("新用户请填写基本信息");
                    this.$router.push({ name: 'newprofile', query: userInfo});
                  } else {
                    this.$router.push({ name: 'main', query: userInfo});
                  }
                })


              
            }else{

              userInfo.type = 2;
              alert("管理端登陆成功\n点击图标随时返回主界面")

              const token = true; // 假设返回的 token 存在 response.data.token
              var cnt = 60;
              if (this.isRemember == true){
                cnt = 60*24*7;
                // alert("222")
              }else{
                // alert("333")
                cnt = 60;
              }
              const expiresIn = cnt; // 设置有效期为60秒（1小时）
              // alert(expiresIn)
              const expiresAt = Date.now() + expiresIn * 1000; // 计算过期时间

              localStorage.setItem('token', token);
              localStorage.setItem('expiresAt', expiresAt); // 存储过期时间
              
              this.$router.push({ name: 'adget', query: userInfo});
            }

          } else {
            alert("登录失败\n"+response.data.msg);
          }
        })
        .catch(error => {
          // 处理登录失败的响应
          console.error(error);
        });
    },
    goToSpecifiedLink() {
      window.location.href = "http://localhost:5173/login"; // 这里替换成你要跳转的链接
    },
    // limitCharacters_1(event) {
    //   if (this.username.length > 50) {
    //     // 如果输入长度超过50，则截断输入值
    //     event.target.value = this.username.substring(0, 50);
    //     // 同步v-model绑定的数据
    //     this.username = event.target.value;
    //   }
    // },
    // limitCharacters_2(event) {
    //   if (this.name.length > 100) {
    //     event.target.value = this.name.substring(0, 100);
    //     // 同步v-model绑定的数据
    //     this.name = event.target.value;
    //   }
    // },
    // limitCharacters_3(event) {
    //   if (this.password.length > 16) {
    //     event.target.value = this.password.substring(0, 16);
    //     // 同步v-model绑定的数据
    //     this.password = event.target.value;
    //   }
    // },
    register() {
      const userData = {
        username: this.usernameR,
        password: this.passwordR,
        password_confirm: this.passwordRC
      };
      if(this.passwordR != this.passwordRC){
        alert("两次输入密码不一致")
        return
      }
      if(this.username1 == ""){
        alert("请输入用户名")
        return
      }

      axios.post("http://127.0.0.1:8080/reg", userData)
        .then(response => {
          // 处理注册成功的响应
          console.log(response.data);
          // 这里可以根据实际情况进行页面跳转或状态更新
          if (response.data.msg === "success") {
            this.usernameR = '';
            this.passwordR = '';
            this.passwordRC = '';
              // 弹出注册成功的消息
            alert("注册成功！");

              // 显示用户的相关信息
            // const userInfo = `用户名: ${response.data.username}\n姓名:${response.data.name}\n用户类型: ${response.data.user_type}`;
            // alert(userInfo);

              // 这里可以根据实际情况进行页面跳转或状态更新
          } else {
              // 如果msg不等于"success"，处理错误情况
            alert("注册失败\n"+response.data.msg);
          }
        })
        .catch(error => {
          // 处理注册失败的响应
          console.error(error);
        });
    }
  }
};
</script>
<!-- 
<script>
export default {
    mounted() {
        let script = document.createElement('script');
        script.type = 'text/javascript';
        script.src = 'src/js/accounts.js';
        console.log("当前载入:"+script.src);
        document.body.appendChild(script);
    },
}
</script> -->

<style lang="scss" scoped>
$gray: #8E9AAF;
$lavender: #CBC0D3;
$pale: #EAC7CC;
$white: #f6f6f6;

$pink: adjust-color($pale, $lightness: -20%);

@import url('https://fonts.googleapis.com/css?family=Open+Sans:300,400|Lora');

$sans-serif: 'Open Sans', sans-serif;
$serif: 'Lora', serif;

body {
  background: $lavender;
}

/* div box styling */
.container {
  margin: auto;
  width: 650px;
  height: 550px;
  position: relative;
}

.welcome {
  background: $white;
  width: 650px;
  height: 415px;
  position: absolute;
  top: 25%;
  border-radius: 5px;
  box-shadow: 5px 5px 5px rgba(0,0,0,.1);
}

.pinkbox {
  position: absolute;
  top: -10%;
  left: 5%;
  background: $pale;
  width: 320px;
  height: 500px;
  border-radius: 5px;
  box-shadow: 2px 0 10px rgba(0,0,0,.1);
  transition: all .5s ease-in-out;
  z-index: 2;
}

.nodisplay {
  display:none;
  transition: all .5s ease;
}

.leftbox, .rightbox {
  position: absolute;
  width: 50%;
  transition: 1s all ease;
}

.leftbox {
  left: -2%;
}
.rightbox {
  right: -2%;
}

/* font & button styling */
h1 {
  font-family: $sans-serif;
  text-align: center;
  margin-top: 95px;
  text-transform: uppercase;
  color: $white;
  font-size: 2em;
  letter-spacing: 8px;
}

.title {
  font-family: $serif;
  color: $gray;
  font-size: 1.8em;
  line-height: 1.1em;
  letter-spacing: 3px;
  text-align: center;
  font-weight: 300;
  margin-top: 20%;
}
.desc {
  margin-top: -8px;
}
.account {
  margin-top: 45%;
  font-size: 10px;
}
p {
  font-family: $sans-serif;
  font-size: .7em;
  letter-spacing: 2px;
  color: $gray;
  text-align: center;
}

span {
  color: $pale;
}

.flower {
  position: absolute;
  width: 120px;
  height: 120px;
  top: 46%;
  left: 29%;
  opacity: .7;
}

.smaller {
  width: 90px;
  height: 100px;
  top: 48%;
  left: 38%;
  opacity: .9;
}

button {
  padding: 12px;
  font-family: $sans-serif;
  text-transform: uppercase;
  letter-spacing: 3px;
  font-size: 11px;
  border-radius: 10px;
  margin: auto;
  outline: none;
  display: block;
  &:hover {
    background: $pale;
    color: $white;
    transition: background-color 1s ease-out;
  }
}

.button {
  margin-top: 3%;
  background: $white;
  color: $pink;
  border: solid 1px $pale;
}

/* form styling */

form {
  display: flex;
  align-items: center;
  flex-direction: column;
  padding-top: 7px;
}
.more-padding{
  padding-top: 35px;
  input {
    padding: 12px;
  }
  .submit {
    margin-top: 45px;
  }
}
.submit {
  margin-top: 25px;
  padding: 12px;
  border-color: $pink;
  &:hover {
    background: $lavender;
    border-color: adjust-color($lavender, $lightness: -5%);
  }
}

input {
  background: $pale;
  width: 65%;
  color: $pink;
  border: none;
  border-bottom: 1px solid rgba($white, 0.5);
  padding: 9px;
  margin: 7px;
  &::placeholder {
    color: rgba($white, 1);
    letter-spacing: 2px;
    font-size: 1.3em;
    font-weight: 100;
  }
  &:focus {
    color: $pink;
    outline: none;
    border-bottom: 1.2px solid rgba($pink, 0.7);
    font-size: 1em;
    transition: .8s all ease;
    &::placeholder {
      opacity: 0;
    }
  }
}

label {
  font-family: $sans-serif;
  color: $pink;
  font-size: 0.8em;
  letter-spacing: 1px;
}

.checkbox {
  display: inline;
  white-space: nowrap;
  position: relative;
  left: -62px;
  top: 5px;
}

input[type=checkbox] {
  width: 7px;
  background: $pink;
}

.checkbox input[type="checkbox"]:checked + label {
  color: $pink;
  transition: .5s all ease;
}

</style>