<template>
    <div class="confession-container">
      <div class="header">
        <el-button type="primary" @click="openNewConfession">创建表白</el-button>
      </div>
      <el-card v-for="confession in confessions" :key="confession.id" class="confession-card">
        <div class="confession-content">
          <p>{{ confession.content }}</p>
          <span class="timestamp">{{ confession.timestamp }}</span>
        </div>
        <div class="actions">
           <el-button type="text" @click="editConfession(confession.id)">修改</el-button>
           <el-button type="danger" @click="deleteConfession(confession.id)">删除</el-button>
        </div>
      </el-card>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        confessions: [
          { id: 1, content: "我喜欢你！", timestamp: "2024-10-01 10:00" },
          { id: 2, content: "不是，这么逆天", timestamp: "2024-10-02 12:30" },
          // 其他表白内容
        ],
      };
    },
    methods: {
      openNewConfession() {
        // 跳转到新建表白界面
        this.$router.push({ name: 'NewConfession' });
      },
      editConfession(id) {
        // 编辑表白逻辑
        console.log(`Editing confession with id: ${id}`);
      },
      deleteConfession(id) {
        // 删除表白逻辑
        this.confessions = this.confessions.filter(confession => confession.id !== id);
        this.$message.success('表白已删除！');
      },
    },
  };
  </script>
  
  <style scoped>
  .confession-container {
    padding: 20px;
  }
  .header {
    margin-bottom: 20px;
  }
  .confession-card {
    margin-bottom: 15px;
  }
  .confession-content {
    margin-bottom: 10px;
  }
  .actions {
    display: flex;
    justify-content: flex-end;
  }
  .timestamp {
    font-size: 12px;
    color: #999;
  }
  </style>
  