<template>
    <v-card class="content"
            variant="outlined">
        <p v-for="(fieldTitle, _, index) in guessFieldTitles"
           :key="fieldTitle"
           :value="fieldTitle"
           class="square-content">{{ fieldTitle }} 
           <!-- <v-card v-if="getGuessFieldTooltipsFromIndex(index)"
                    class="tooltip-card"
                    variant="outlined"> {{ getGuessFieldTooltipsFromIndex(index) }} </v-card> -->
        </p>
    </v-card>
</template>

<script setup>
import { guessFieldTitles, getGuessFieldTooltipsFromIndex } from '../constants.js';


</script>

<style scoped>
.content {
    display: flex;
    flex-direction: row;
    justify-content: center;
    background-color: white;
    gap: 8px;
    padding: 8px 0 8px 8px;
    border-width: 2px;
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
    font-family: pkmEmerald;
}

.tooltip {
    font-family: pkmEmerald;
    margin: 0px;
    padding: 0px;
}

.tooltip-card {
    background-color: white;
    border-width: 2px;
    color: black;
    padding: 10px;
}
</style>