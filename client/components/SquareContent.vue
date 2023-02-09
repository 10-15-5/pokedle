<template>
    <div>
        <div v-if="pokemon && color === 'shiny'"
             class="star-container">
            <v-icon class="star-icon-background"
                    icon="mdi-star" />
            <v-icon class="star-icon"
                    icon="mdi-star" />
        </div>
        <v-card class="square-content elevation-0"
                variant="outlined"
                :style="{ 'background-color': getColor }">
            <div v-if="pokemon">
                <img class="mt-2 pokemon-bg"
                     :src="'https://img.pokemondb.net/sprites/ruby-sapphire/' + color + '/' + pokemon + '.png'"
                     alt="pokemon sprite" />
            </div>
            <div v-if="habitat">
                <img class="mt-2 pokemon-habitat"
                     :src="getHabitatImage(habitat)"
                     :title="`${habitat}`"
                     :alt="`${habitat}`" />
            </div>
            <div v-else-if="guessResult"
                 class="text-content">
                <p>{{ guessText }}</p>
            </div>
        </v-card>
    </div>
</template>

<script setup>
//getHabitatImage(habitat)
import { guessState } from '../constants.js';
import { computed } from 'vue';
import { getHabitatImage } from '../services/assets.js';

const props = defineProps({
    guessResult: String,
    guessText: String,
    pokemon: String,
    habitat: String,
    color: String
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
    background-color: white;
}

.pokemon-habitat {
    width: 100%;
    border-style: solid;
    border-radius: 60%;
    border-width: 2px;
}

.square-content {
    width: 60px;
    height: 60px;
    min-height: 60px;
    min-width: 60px;
    display: flex;
    justify-content: center;
    align-items: center;
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
    color: black;
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

img:hover {}
</style>
