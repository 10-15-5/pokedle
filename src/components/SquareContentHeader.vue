<template>
    <v-card class="content">
        <p v-for="(fieldTitle) in guessFieldTitles"
           :key="fieldTitle"
           :value="fieldTitle"
           class="square-content">{{ fieldTitle }}</p>
    </v-card>
</template>

<script setup>
import { guessFieldTitles } from '../constants.js';
</script>

<style scoped>
.content {
    display: flex;
    flex-direction: row;
    justify-content: center;
    gap: 6px;
    padding: 8px;
    background-color: rgb(255, 255, 255, .9);
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