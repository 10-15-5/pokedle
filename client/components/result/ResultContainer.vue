<template>
    <div
        class="card space-x-2 p-2 sm:space-x-[2px] sm:!p-1"
        :class="isAllFieldsCorrectButIncorrectPokemon ? '!bg-red-300 dark:!bg-red-900' : ''"
    >
        <div v-for="(field, k, i) in guessResult.fields" :key="k" :value="field">
            <ResultSquare
                :pokemon="removeSpecialCharactersExceptDashFromString(field.name)"
                :color="guessResult.color"
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
import { guessState, TotalResultCardFlipDelay } from '../../constants.js';
import { useStore } from '../../stores/store.js';
import { onMounted, ref } from 'vue';
const store = useStore();

const props = defineProps({
    pokemonName: String,
    guessResult: Object,
});

const isAllFieldsCorrectButIncorrectPokemon = ref(false);

//Set background to 'red' if all fields are correct, but pokemon is incorrect
onMounted(() => {
    if (!props.guessResult) return false;
    if (props.guessResult.isCorrectGuess) return false;

    const arr = Object.values(props.guessResult.fields);

    const isAllFieldsCorrectGuess = arr.every(
        (e, idx) => idx < 1 || e.guessState === guessState.CorrectGuess
    );

    const delay = store.isGameWon ? 0 : TotalResultCardFlipDelay;

    setTimeout(() => {
        isAllFieldsCorrectButIncorrectPokemon.value = isAllFieldsCorrectGuess;
    }, delay);
});
</script>

<style scoped></style>
