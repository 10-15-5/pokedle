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
  margin: 0px 10px 0px 10px;
}

.square-content {
  width: 60px;
  display: flex;
  justify-content: center;
  border-bottom: 2px solid white;
  align-items: center;
}

p {
  font-size: small;
  color: white;
  font-weight: 500;
}
</style>