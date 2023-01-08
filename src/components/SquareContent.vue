<template>
  <div class="square-content"
       :style="{ 'background-color': getColor }">
    <div v-if="guessResult">
      <p>{{ guessText }}</p>
    </div>
    <div v-else>
      <img class="mt-2"
           :src="'https://img.pokemondb.net/sprites/ruby-sapphire/normal/' + pokemon + '.png'"
           alt="" />
    </div>
  </div>
</template>

<script>
import { guessState } from '../constants.js';
const props = defineProps({
  guessResult: {
    type: String,
        },
  guessText: {
    type: String,
        },
  pokemon: {
    type: String
  }
});
  
const getColor = () => {
  switch (this.guessResult) {
    case guessState.CorrectGuess:
      return "greenyellow";
    case guessState.PartlyCorrectGuess:
      return "orange";
    case guessState.WrongGuess:
      return "red";
    default:
      return "black";
  }
}


</script>

<style scoped>
img {
  width: 100%;
}

.guess-text {
  display: flex;
  justify-content: center;
  align-items: center;
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
