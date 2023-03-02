<template>
    <v-card class="content"
            variant="outlined">
        <div v-for="(field, k, i) in guessResult"
             :key="k"
             :value="field">
            <SquareContent v-if="field.name"
                           :pokemon="removeSpecialCharactersExceptDashFromString(field.name)"
                           :color="color" />
            <SquareContent v-else-if="field.habitat"
                           :guessResult="field.guessState"
                           :habitat="field.habitat"
                           :flipDelay="i-1" />
            <SquareContent v-else
                           :guessResult="field.guessState"
                           :guessText="field.text"
                           :flipDelay="i-1" />
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
    border-width: 2px;
}
</style>