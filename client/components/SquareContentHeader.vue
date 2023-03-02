<template>
    <v-card class="content"
            variant="outlined">
        <p v-for="(fieldTitle, _, index) in guessFieldTitles"
           :key="fieldTitle"
           :value="fieldTitle"
           class="square-content">{{ fieldTitle }} <v-card v-if="getGuessFieldTooltipsFromIndex(index)"
                    class="tooltip"
                    variant="outlined"> {{ getGuessFieldTooltipsFromIndex(index) }} </v-card>
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
    gap: 8px;
    padding: 8px 0 8px 8px;
    border-width: 2px;
    position: static;

}

.square-content {
    width: 60px;
    height: 40px;
    display: flex;
    justify-content: center;
    border-bottom: 2px solid;
    align-items: center;
}

p {
    font-size: 15px;
    font-weight: bolder;
    font-family: pkmEmerald;
    text-align: center;
}

.tooltip {
    opacity: 0;
    text-align: center;
    padding: 6px;
    margin-top: -102px;
    border-width: 2px;
    position: absolute;
    z-index: 1;
    transition: opacity 0.3s;
    visibility: hidden;
}



/* Show the tooltip text when you mouse over the tooltip container */
.square-content:hover .tooltip {
    visibility: visible;
    opacity: 1;
    transition-delay: 0.8s;
}
</style>