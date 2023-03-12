<template>
    <div
        class="card flex min-w-[80px] max-w-[400px] flex-col font-pkmEmerald sm:max-w-[330px] sm:text-sm"
    >
        <p v-if="guessesRemainingForHint > 0" class="gap-2 p-2 px-3 sm:p-1 sm:px-2">
            Guess <b class="text-light-orange dark:text-dark-orange"> {{ guessesRemainingForHint }} </b> more times for a HINT
        </p>
        <div v-else class="flex w-full flex-row items-center justify-between gap-2 p-2 px-3 pb-1 sm:pb-1 sm:pt-2">
            <div class="flex flex-row text-base">
                <p>HINT</p>
            </div>
            <div>
                <button v-if="isShowHint" class="pi pi-minus" @click="isShowHint = false"></button>
                <button v-else class="pi pi-plus" @click="isShowHint = true"></button>
            </div>
        </div>
        <div v-if="isShowHint" class="w-full bg-gray-100 p-4 pt-3 text-justify dark:bg-zinc-700">
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
    },
    forceEnableHint: Boolean,
});

const isShowHint = ref(false);

const guessesRemainingForHint = computed(() => 5 - props.numberOfGuesses);
</script>
