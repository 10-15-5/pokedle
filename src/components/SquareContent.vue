<template>
  <div class="square-content"
       :style="{ 'background-color': getColor }">
    <div v-if="guessResult">
      <p>{{ guessText }}</p>
    </div>
    <div v-else>
      <img class="mt-2"
           :src="'https://img.pokemondb.net/sprites/ruby-sapphire/' + pokemonColor + '/' + pokemon + '.png'"
           alt="" />
    </div>
  </div>
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
  const random = Math.random()*100;
  return random <8 ? 'shiny' : 'normal';
}

const pokemonColor = getShiny();

const getColor = computed(() => {
  switch (props.guessResult) {
    case guessState.CorrectGuess:
      return "greenyellow";
    case guessState.PartlyCorrectGuess:
      return "orange";
    case guessState.WrongGuess:
      return "red";
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
}

.square-content {
  width: 60px;
  height: 60px;
  border-style: solid;
  border-color: black;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
