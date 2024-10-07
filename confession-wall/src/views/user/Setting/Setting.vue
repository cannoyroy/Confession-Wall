<script setup>
import { useRoute} from 'vue-router';
const route = useRoute();
</script>

<template>
  <div>
    <button @click="toggleFont">Toggle Font</button>
  </div>
</template>

<script>
export default {
  name: 'FontToggleComponent',
  methods: {
    toggleFont() {
      // 检查当前根元素的字体是否是我们想要切换到的字体
      const currentFont = getComputedStyle(document.documentElement).getPropertyValue('--font-main');
      
      // 定义两种字体，根据当前字体切换到另一种
      const fonts = ["'Poppins'", "'Arial', sans-serif"];
      const newFontIndex = fonts.indexOf(currentFont) === 0 ? 1 : 0;
      const newFont = fonts[newFontIndex];

      // 更新CSS变量以改变全局字体
      document.documentElement.style.setProperty('--font-main', newFont);
    }
  }
}
</script>
