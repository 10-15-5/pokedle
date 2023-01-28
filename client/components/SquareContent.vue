<template>
    <v-card class="square-content elevation-0"
            variant="outlined"
            :style="{ 'background-color': getColor }">
        <div v-if="guessResult" class="text-content">
            <p>{{ guessText }}</p>
        </div>
        <div v-else>
            <img class="mt-2"
                 :src="'https://img.pokemondb.net/sprites/ruby-sapphire/' + pokemonColor + '/' + pokemon + '.png'"
                 alt="" />
        </div>
    </v-card>
</template>

<script setup>
import { guessState } from '../constants.js';
import { computed } from 'vue';

const props = defineProps({
    guessResult: String,
    guessText: String,
    pokemon: String
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
img {
    width: 100%;
    background-image: url("../assets/pokecenter-box-background.png");
    background-size: contain;
    background-color: white;
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

.text-content{
    padding: 4px 0px 0px 2px;
}
p {
    font-family: pkmEmerald;
    font-size: 20px;
}
</style>
