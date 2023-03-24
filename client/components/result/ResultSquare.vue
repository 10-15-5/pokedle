<template>
    <div class="relative">
        <div
            v-if="type === GuessType.Pokemon && color === 'shiny'"
            class="z-1 absolute mt-[-10px] ml-[46px] flex h-[22px] w-[22px] items-center justify-center sm:ml-[30px] sm:mt-[-10px]"
        >
            <SvgIcon
                type="mdi"
                :path="mdiStar"
                class="absolute z-10 h-[30px] w-[30px] text-black dark:!text-dark-border sm:h-[22px] sm:w-[22px]"
            />
            <SvgIcon
                type="mdi"
                :path="mdiStar"
                class="absolute z-10 h-[20px] w-[20px] text-yellow-500 sm:h-[14px] sm:w-[14px]"
            />
        </div>
        <VueFlip v-model="isCardFaceDown" transition="1s" :height="cardSize" :width="cardSize">
            <template v-slot:front>
                <div class="result-card rounded" :style="{ 'background-color': getColor }" variant="outlined">
                    <img
                        v-if="type === GuessType.Pokemon"
                        class="pokemon-bg bg-neutral-200 bg-no-repeat dark:!bg-neutral-700"
                        :src="'https://img.pokemondb.net/sprites/ruby-sapphire/' + color + '/' + pokemon + '.png'"
                        :title="pokemon"
                        alt="pokemon sprite"
                    />
                    <div
                        v-if="type === GuessType.Blackout"
                        class="w-full rounded bg-white bg-no-repeat dark:!bg-neutral-300"
                    >
                        <img
                            class="w-full brightness-0"
                            :src="'https://img.pokemondb.net/sprites/ruby-sapphire/' + color + '/' + pokemon + '.png'"
                            alt="pokemon sprite"
                        />
                    </div>
                    <div
                        v-if="type === GuessType.Habitat"
                        class="h-[39px] overflow-hidden rounded-[50%] border-2 border-light-border dark:!border-dark-border sm:h-[30px]"
                    >
                        <img :src="getHabitatImage(habitat)" :title="`${habitat}`" :alt="`${habitat}`" />
                    </div>
                    <p
                        v-if="type === GuessType.Text"
                        class="pt-1 pl-[2px] font-pkmEmerald text-[16px] capitalize dark:font-bold dark:text-dark-text sm:text-[12px] sm:font-bold"
                    >
                        {{ guessText }}
                    </p>
                </div>
            </template>
            <template v-slot:back>
                <div class="result-card bg-neutral-200 dark:!bg-neutral-800" variant="outlined">
                    <img :src="'./client/assets/result-cards/pokemon-cardback-pixel.png'" alt="pokemon sprite" />
                </div>
            </template>
        </VueFlip>
    </div>
</template>

<script setup>
import { GuessState, GuessType } from '../../constants.js';
import { computed, onBeforeMount, onMounted, ref, onUnmounted } from 'vue';
import { getHabitatImage } from '../../services/assets.js';
import { VueFlip } from 'vue-flip';
import { useStore } from '../../stores/store';
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiStar } from '@mdi/js';

const store = useStore();

const props = defineProps({
    guessResult: String,
    guessText: String,
    pokemon: String,
    habitat: String,
    color: {
        String,
        default: 'normal',
    },
    flipDelay: Number,
    type: String,
    isLarge: 
        Boolean,   
});

const windowWidth = ref(window.innerWidth);
const isCardFaceDown = ref(true);

const cardSize = computed(() => {
    if (!props.isLarge) {
        if (windowWidth.value > 576) return '60px';
        else if (windowWidth.value <= 576) return '43px';
    } else {
        if (windowWidth.value > 576) return '100px';
        else if (windowWidth.value <= 576) return '80px';
    }
});

onBeforeMount(() => {
    if (props.flipDelay === undefined || store.isClassicGameWon || props.type === GuessType.Pokemon)
        isCardFaceDown.value = false;
});
const onWidthChange = () => (windowWidth.value = window.innerWidth);

onMounted(() => {
    window.addEventListener('resize', onWidthChange);
    setTimeout(() => {
        isCardFaceDown.value = false;
    }, props.flipDelay * 350);
});

onUnmounted(() => {
    window.removeEventListener('resize', onWidthChange);
});

const getColor = computed(() => {
    switch (props.guessResult) {
        case GuessState.CorrectGuess:
            return store.isDark ? '#10b981' : 'rgb(70, 217, 48)';
        case GuessState.PartlyCorrectGuess:
            return store.isDark ? '#fb923c' : 'rgb(237, 159, 43)';
        case GuessState.WrongGuess:
            return store.isDark ? '#f87171' : 'rgb(237, 59, 43)';
        default:
            return 'transparent';
    }
});
</script>

<style scoped>
.pokemon-bg {
    width: 100%;
    background-image: url('/client/assets/result-cards/pokecenter-box-background.png');
    background-size: contain;
    border-radius: 2px;
}
</style>
