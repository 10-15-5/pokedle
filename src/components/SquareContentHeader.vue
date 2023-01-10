<template>
  <div class="content">
    <p v-for="(fieldTitle) in guessFieldTitles"
       :key="fieldTitle"
       :value="fieldTitle"
       class="square-content">{{ fieldTitle }}</p>
  </div>
</template>

<script setup>
import { guessFieldTitles } from '../constants.js';
</script>

<style scoped>
.content {
  display: flex;
  flex-direction: row;
  justify-content: center;
  gap: 10px;
  width: 100%;
}

.square-content {
  width: 60px;
  display: flex;
  justify-content: center;
  border-bottom: 2px solid rgb(0, 0, 0);
  align-items: center;
  color: rgb(0, 0, 0);
}

p {
  font-size: smaller;
  color: rgb(0, 0, 0);
  font-weight: bolder;
}
</style>