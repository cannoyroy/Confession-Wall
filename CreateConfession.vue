<template>
    <div class="new-confession-container">
      <el-card>
        <h2>新建表白</h2>
        <el-input
          type="textarea"
          v-model="content"
          placeholder="请输入你的表白内容"
          rows="4"
        ></el-input>
        <el-button type="primary" @click="submitConfession">提交表白</el-button>
      </el-card>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        content: '',
      };
    },
    methods: {
      submitConfession() {
        // 提交表白逻辑
        if (this.content.trim()) {
          // 假设这里是发送到后端的逻辑
          console.log(`提交的表白内容: ${this.content}`);
          this.$message.success('表白提交成功！');
          this.$router.push({ name: 'Confession' }); // 返回到表白列表
        } else {
          this.$message.warning('表白内容不能为空！');
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .new-confession-container {
    padding: 20px;
  }
  </style>
  