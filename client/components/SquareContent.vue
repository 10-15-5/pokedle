<template>
    <v-card class="square-content elevation-0"
            variant="outlined"
            :style="{ 'background-color': getColor }">
        <div v-if="pokemon">
            <img class="mt-2 pokemon-bg"
                 :src="'https://img.pokemondb.net/sprites/ruby-sapphire/' + pokemonColor + '/' + pokemon + '.png'"
                 alt="pokemon sprite" />
        </div>
        <div v-if="habitat">
            <img class="mt-2 pokemon-habitat"
                 :src="getHabitatImage(habitat)"
                 alt="" />
        </div>
        <div v-else-if="guessResult"
             class="text-content">
            <p>{{ guessText }}</p>
        </div>
    </v-card>
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
});

const getShiny = () => {
    const random = Math.random() * 100;
    return random < 8 ? 'shiny' : 'normal';
}

const pokemonColor = getShiny();

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
    width: 75px;
    overflow: hidden;
    background-size: cover;
}
.square-content {
    width: 60px;
    height: 60px;
    min-height: 60px;
    min-width: 60px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.text-content {
    padding: 4px 0px 0px 2px;
}

p {
    font-family: pkmEmerald;
    font-size: 16px;
    text-transform: capitalize;
}
</style>
