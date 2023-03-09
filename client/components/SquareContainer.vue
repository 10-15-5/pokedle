<template>
    <div class="card space-x-2 p-2 sm:space-x-[2px] sm:!p-1">
        <div v-for="(field, k, i) in guessResult" :key="k" :value="field">
            <SquareContent
                v-if="field.name"
                :pokemon="removeSpecialCharactersExceptDashFromString(field.name)"
                :color="color"
            />
            <SquareContent
                v-else-if="field.habitat"
                :guessResult="field.guessState"
                :habitat="field.habitat"
                :flipDelay="i - 1"
            />
            <SquareContent
                v-else
                :guessResult="field.guessState"
                :guessText="field.text"
                :flipDelay="i - 1"
            />
        </div>
    </div>
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

<style scoped></style>
