<script setup>
import { ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export const username = ref("");
export const password = ref("");
export const name = ref("");
export const user_type = ref("1");
</script>

<template>
  
  <RouterLink to="/login">login</RouterLink>
  <RouterLink to="/reg">register</RouterLink>
  <RouterView />
</template>

<style scoped>

</style>
