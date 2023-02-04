<template>
    <v-card class="content"
            variant="outlined">
        <div v-for="(field, i) in guessResult"
             :key="i"
             :value="field">
            <SquareContent v-if="field.name"
                           :pokemon="removeSpecialCharactersExceptDashFromString(field.name)"
                           :color="color" />
            <SquareContent v-else-if="field.habitat"
                           :guessResult="field.guessState"
                           :habitat="field.habitat" />
            <SquareContent v-else
                           :guessResult="field.guessState"
                           :guessText="field.text" />
        </div>
    </v-card>
</template>

<script setup>
import SquareContent from './SquareContent.vue';
import { removeSpecialCharactersExceptDashFromString } from '../helpers.js';
const props = defineProps({
    pokemonName: String,
    guessResult: Object,
    color: String,
});
</script>

<style scoped>
.content {
    display: flex;
    flex-direction: row;
    justify-content: center;
    gap: 8px;
    padding: 8px 0 8px 8px;
    background-color: rgb(255, 255, 255, .9);
}
</style>