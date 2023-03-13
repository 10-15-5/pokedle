<template>
    <div class="card space-x-2 p-2 sm:space-x-[2px] sm:!p-1" :class="(isAllFieldsCorrect && !isGameWon) ? '!bg-red-300 dark:!bg-red-900' : ''">
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
import { guessState, TotalResultCardFlipDelay } from '../../constants.js';
import { useStore } from '../../stores/store.js';
import {onMounted, ref} from 'vue';
const store = useStore();

const props = defineProps({
    pokemonName: String,
    guessResult: Object,
    color: String,
    isGameWon: Boolean,
});

const isAllFieldsCorrect = ref(false)

onMounted(() => {
    console.log("computed");
    if(!props.guessResult) return false;
    console.log(props.guessResult);
    const arr = Object.values(props.guessResult);
    const arr2 = arr.shift();
    console.log(props.guessResult);
    console.log(arr);
    const t = arr.every((e) => e.guessState === guessState.CorrectGuess)
    console.log(t);

    setTimeout(() => {
        isAllFieldsCorrect.value = t;
    }, TotalResultCardFlipDelay)
});

</script>

<style scoped></style>
