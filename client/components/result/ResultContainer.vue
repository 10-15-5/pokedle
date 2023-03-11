<template>
    <div class="card space-x-2 p-2 sm:space-x-[2px] sm:!p-1">
        <div v-for="(field, k, i) in guessResult" :key="k" :value="field">
            <ResultSquare
                :pokemon="removeSpecialCharactersExceptDashFromString(field.name)"
                :color="color"
                :guessResult="field.guessState"
                :habitat="field.habitat"
                :flipDelay="i - 1"
                :guessText="field.text"
                :type="field.type"
            />
        </div>
    </div>
</template>

<script setup>
import ResultSquare from '../result/ResultSquare.vue';
import { removeSpecialCharactersExceptDashFromString } from '../../helpers.js';
const props = defineProps({
    pokemonName: String,
    guessResult: Object,
    color: String,
});
</script>

<style scoped></style>
