<template>
    <v-card class="content" variant="outlined">
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
    background-color: white;
    gap: 8px;
    padding: 8px 0 8px 8px;
}

.square-content {
    width: 60px;
    height: 40px;
    display: flex;
    justify-content: center;
    border-bottom: 2px solid rgb(0, 0, 0);
    align-items: center;
    color: rgb(0, 0, 0);
}

p {
    font-size: 14px;
    font-weight: bolder;
    color: rgb(0, 0, 0);
    font-family: pkmEmerald;
}
</style>