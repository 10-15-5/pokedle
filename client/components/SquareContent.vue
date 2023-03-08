<template>
    <div>
        <div
            v-if="pokemon && color === 'shiny'"
            class="z-1 absolute mt-[-8px] ml-[46px] flex h-[22px] w-[22px] items-center justify-center"
        >
            <v-icon
                class="absolute z-10 text-3xl text-black dark:!text-dark-border"
                icon="mdi-star"
            />
            <v-icon class="absolute z-10 text-xl text-yellow-500" icon="mdi-star" />
        </div>
        <VueFlip
            v-model="isCardFaceDown"
            transition="1s"
            class="h-[60px] w-[60px] sm:h-[45px] sm:w-[45px]"
        >
            <template v-slot:front>
                <div
                    class="result-card rounded"
                    :style="{ 'background-color': getColor }"
                    variant="outlined"
                >
                    <img
                        v-if="pokemon"
                        class="pokemon-bg bg-neutral-200 dark:!bg-neutral-700"
                        :src="
                            'https://img.pokemondb.net/sprites/ruby-sapphire/' +
                            color +
                            '/' +
                            pokemon +
                            '.png'
                        "
                        alt="pokemon sprite"
                    />
                    <div
                        v-if="habitat"
                        class="h-[39px] overflow-hidden rounded-[50%] border-2 border-light-border dark:!border-dark-border sm:h-[30px]"
                    >
                        <img
                            :src="getHabitatImage(habitat)"
                            :title="`${habitat}`"
                            :alt="`${habitat}`"
                        />
                    </div>
                    <p
                        v-else-if="guessResult"
                        class="pt-1 pl-[2px] font-pkmEmerald text-[16px] capitalize dark:font-bold dark:text-dark-text sm:text-[12px] sm:font-bold"
                    >
                        {{ guessText }}
                    </p>
                </div>
            </template>
            <template v-slot:back>
                <div class="result-card bg-neutral-200 dark:!bg-neutral-800" variant="outlined">
                    <img :src="'./client/assets/pokemon-cardback-pixel.png'" alt="pokemon sprite" />
                </div>
            </template>
        </VueFlip>
    </div>
</template>

<script setup>
import { guessState } from '../constants.js';
import { computed, onBeforeMount, onMounted, ref } from 'vue';
import { getHabitatImage } from '../services/assets.js';
import { VueFlip } from 'vue-flip';
import { useDark } from '@vueuse/core';
import { useStore } from '../stores/store';
const isDark = useDark();
const store = useStore();

const props = defineProps({
    guessResult: String,
    guessText: String,
    pokemon: String,
    habitat: String,
    color: String,
    flipDelay: Number,
});

const isCardFaceDown = ref(true);

onBeforeMount(() => {
    if (props.flipDelay === undefined || store.isGameWon) isCardFaceDown.value = false;
});

onMounted(() => {
    setTimeout(() => {
        isCardFaceDown.value = false;
    }, props.flipDelay * 350);
});

const getColor = computed(() => {
    switch (props.guessResult) {
        case guessState.CorrectGuess:
            return isDark.value ? '#10b981' : 'rgb(70, 217, 48)';
        case guessState.PartlyCorrectGuess:
            return isDark.value ? '#fb923c' : 'rgb(237, 159, 43)';
        case guessState.WrongGuess:
            return isDark.value ? '#f87171' : 'rgb(237, 59, 43)';
        default:
            return 'transparent';
    }
});
</script>

<style scoped>
.pokemon-bg {
    width: 100%;
    background-image: url('../assets/pokecenter-box-background.png');
    background-size: contain;
}
</style>
