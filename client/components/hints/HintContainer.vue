<template>
    <div
        class="card flex min-w-[190px] max-w-[400px] flex-col font-pkmEmerald sm:max-w-[330px] sm:text-sm"
    >
        <div class="flex w-full flex-row items-center justify-between p-2 px-3 pb-1">
            <div class="text-base flex flex-row">
                <p>HINT</p>
            </div>
            <p v-if="guessesRemainingForHint > 0" class=""> {{ guessesRemainingForHint }} More Guesses</p>
            <div v-else-if="guessesRemainingForHint <= 0">
                <button v-if="isShowHint" class="pi pi-minus" @click="isShowHint = false"></button>
                <button v-else class="pi pi-plus" @click="isShowHint = true"></button>
            </div>
        </div>
        <div v-if="isShowHint" class="w-full bg-gray-10 dark:bg-zinc-700 p-4 pt-3 text-justify">
            {{ text }}
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';

const props = defineProps({
    text: String,
    numberOfGuesses: {
        type: Number,
        default: 0,
    }
});

const isShowHint = ref(false)

const guessesRemainingForHint = computed(() =>  5 - props.numberOfGuesses);

</script>
