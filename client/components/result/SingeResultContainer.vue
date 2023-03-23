<template>
    <div
        class="single-result-container"
        :class="
            isCorrect
                ? ' bg-dark-emerald  dark:bg-light-emerald'
                :'!bg-red-300 dark:!bg-red-900'
        "
    >
        <ResultSquare
            :pokemon="removeSpecialCharactersExceptDashFromString(pokemonName)"
            :color="color"
            :type="GuessType.Pokemon"
        />
    </div>
</template>

<script setup>
import ResultSquare from './ResultSquare.vue';
import { removeSpecialCharactersExceptDashFromString } from '../../helpers.js';
import { GuessType } from '../../constants.js';
import { useStore } from '../../stores/store.js';
import { onMounted, ref } from 'vue';
const store = useStore();

const props = defineProps({
    pokemonName: String,
    color: String,
    isCorrect: Boolean
});
</script>

<style scoped></style>
