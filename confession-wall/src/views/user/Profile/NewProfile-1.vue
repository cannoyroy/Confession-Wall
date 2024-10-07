<script setup>
import { useRoute} from 'vue-router';
import 'element-plus/dist/index.css';
const route = useRoute();

</script>

<template>
    <div>
      <el-radio-group v-model="gender">
        <el-radio :value="male">男</el-radio>
        <el-radio :value="female">女</el-radio>
        <el-radio :value="other">其他</el-radio>
      </el-radio-group>
  
      <el-input v-model="contact.phone" placeholder="联系电话"></el-input>
      <el-input v-model="contact.qq" placeholder="QQ"></el-input>
      <el-input v-model="contact.wechat" placeholder="微信"></el-input>
      <el-input v-model="contact.other" placeholder="其他联系方式"></el-input>
  
      <el-cascader
        v-model="selectedRegion"
        :options="regionData"
        placeholder="选择地区"
        @change="handleRegionChange"
      ></el-cascader>
  
      <el-input v-model="avatarUrl" placeholder="头像链接"></el-input>
  
      <el-input
        type="textarea"
        v-model="introduction"
        placeholder="自我介绍"
      ></el-input>
  
      <el-button type="primary" @click="handleSubmit">设置完成</el-button>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        gender: '',
        male: false,
        female: false,
        other: false,
        contact: {
          phone: '',
          qq: '',
          wechat: '',
          other: ''
        },
        selectedRegion: [],
        regionData: [], // 地区数据
        avatarUrl: '',
        introduction: ''
      };
    },
    methods: {
      handleSubmit() {
        // 提交数据的逻辑
        console.log('提交的数据:', this.$data);
      },
      handleRegionChange(value) {
        // 处理地区选择变化的逻辑
      }
    }
  };
</script>