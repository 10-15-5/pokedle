<template>
    <div>
        <div v-if="pokemon && color === 'shiny'"
             class="star-container">
            <v-icon class="star-icon-background"
                    icon="mdi-star" />
            <v-icon class="star-icon"
                    icon="mdi-star" />
        </div>
        <VueFlip class="square-content"
                 v-model="isCardFaceDown"
                 transition="1s"
                 width="60px"
                 height="60px">
            <template v-slot:front>
                <v-card class="square-content-front border-black"
                        :style="{ 'background-color': getColor }"
                        variant="outlined">
                    <img v-if="pokemon"
                         class="pokemon-bg bg-white"
                         :src="'https://img.pokemondb.net/sprites/ruby-sapphire/' + color + '/' + pokemon + '.png'"
                         alt="pokemon sprite" />
                    <img v-if="habitat"
                         class="pokemon-habitat border-black"
                         :src="getHabitatImage(habitat)"
                         :title="`${habitat}`"
                         :alt="`${habitat}`" />
                    <p v-else-if="guessResult"
                       class="text-content">{{ guessText }}</p>
                </v-card>
            </template>
            <template v-slot:back>
                <v-card class="square-content-back border-black"
                        variant="outlined">
                        <img
                        class="pokemon-bg"
                         :src="'./client/assets/pokemon-cardback-pixel.png'"
                         alt="pokemon sprite" />
                </v-card>
            </template>
        </VueFlip>
    </div>
</template>

<script setup>

import { guessState } from '../constants.js';
import { computed, onBeforeMount, onMounted, ref } from 'vue';
import { getHabitatImage } from '../services/assets.js';
import { VueFlip } from 'vue-flip';

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
    if(props.flipDelay===undefined) isCardFaceDown.value = false;
});

onMounted(() => {
    setTimeout(() => {
        isCardFaceDown.value = false;
    }, props.flipDelay * 350);
});

const getColor = computed(() => {
    switch (props.guessResult) {
        case guessState.CorrectGuess:
            return "rgb(70, 217, 48)";
        case guessState.PartlyCorrectGuess:
            return "rgb(237, 159, 43)";
        case guessState.WrongGuess:
            return "rgb(237, 59, 43)";
        default:
            return "transparent";
    }
});
</script>

<style scoped>
.pokemon-bg {
    width: 100%;
    background-image: url("../assets/pokecenter-box-background.png");
    background-size: contain;
}

.pokemon-habitat {
    width: 101%;
    border-style: solid;
    border-radius: 60%;
    border-width: 2px;
}

.square-content-front,
.square-content-back {
    border-width: 2px;
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    width: 100%;
    height: 100%;
    -webkit-backface-visibility: hidden;
    /* Safari */
    backface-visibility: hidden;
}

/* Style the back side */
.square-content-back {
    border-width: 2px;
}

.text-content {
    padding: 4px 0px 0px 2px;
}

.star-container {
    width: 22px;
    height: 22px;
    z-index: 1;
    position: absolute;
    margin: -8px 0px 0px 46px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.star-icon {
    color: rgb(249, 194, 53);
    z-index: 1;
    position: absolute;
    font-size: 17px;
}

.star-icon-background {
    z-index: 1;
    position: absolute;
    font-size: 25px;
}

p {
    font-family: pkmEmerald;
    font-size: 16px;
    text-transform: capitalize;
}

/* capitalize alt text*/
img {
    text-transform: capitalize;
}
</style>
