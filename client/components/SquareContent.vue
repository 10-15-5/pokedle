<template>
    <div>
        <div v-if="pokemon && color === 'shiny'"
             class="star-container">
            <v-icon class="star-icon-background"
                    icon="mdi-star" />
            <v-icon class="star-icon"
                    icon="mdi-star" />
        </div>
        <div class="square-content">
            <div class="square-content-inner" :style="{'animation-duration': getDelay }">
                <v-card class="square-content-front" :style="{ 'background-color': getColor }" variant="outlined">
                    <img v-if="pokemon"
                         class="pokemon-bg"
                         :src="'https://img.pokemondb.net/sprites/ruby-sapphire/' + color + '/' + pokemon + '.png'"
                         alt="pokemon sprite" />
                    <img v-if="habitat"
                         class="pokemon-habitat"
                         :src="getHabitatImage(habitat)"
                         :title="`${habitat}`"
                         :alt="`${habitat}`" />
                    <p v-else-if="guessResult"
                       class="text-content">{{ guessText }}</p>
                </v-card>
                <v-card class="square-content-back" variant="outlined">
                    <p>BACK</p>
                </v-card>
            </div>
        </div>
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
    color: String,
    flipDelay: String,
});

const getDelay = computed(() => {
    console.log(props.flipDelay)
    return props.flipDelay + 's';
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

/* The flip card container - set the width and height to
 whatever you want. We have added the border property to 
 demonstrate that the flip itself goes out of the box on
  hover (remove perspective if you don't want the 3D effect */
.square-content {
    width: 60px;
    height: 60px;
    min-height: 60px;
    min-width: 60px;
    background-color: transparent;
    perspective: 1000px;
    /* Remove this if you don't want the 3D effect */
}

/* This container is needed to position the front and back side */
.square-content-inner {
    position: relative;
    width: 100%;
    height: 100%;
    text-align: center;
    animation: fadein;
    transform-style: preserve-3d;
}

@keyframes fadein {
    from { transform: rotateY(180deg); }
    to   { transform: 0; }
}

/* Do an horizontal flip when you move the mouse over the flip box container */
/* .square-content:hover .square-content-inner {
    transform: rotateY(180deg);
} */

/* Position the front and back side */
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

/* Style the front side (fallback if image is missing) */
.square-content-front {
    color: black;
}

/* Style the back side */
.square-content-back {
    background-color: dodgerblue;
    color: black;
    transform: rotateY(180deg);
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
</style>
