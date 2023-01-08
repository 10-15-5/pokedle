<template>
  <div class="content">
    <SquareContent :pokemon="removeSpecialCharactersFromString(pokemonData.name).toLowerCase()" />
    <SquareContent :guessResult="'CorrectGuess'"
                   :guessText="'fire'" />
    <SquareContent :guessResult="'WrongGuess'"
                   :guessText="'water'" />
    <SquareContent :guessResult="'PartlyCorrectGuess'"
                   :guessText="'Fully evolved'" />
    <SquareContent :guessResult="'WrongGuess'"
                   :guessText="'gen 1'" />
  </div>
</template>

<script setup>
import SquareContent from './SquareContent.vue';
const correctFields = {
  "number": 6,
    "name": "charizard",
    "type 1": "fire",
    "type 2": "flying",
    "total": 534,
    "hp": 78,
    "attack": 84,
    "defense": 78,
    "sp. atk": 109,
    "sp. def": 85,
    "speed": 100,
    "generation": 1,
    "legendary": "false"
}
const props = defineProps({
  pokemonData: Object
});
const removeSpecialCharactersFromString = (text) => {
  return text.replace(/[^\w\s]/gi, '');
}

</script>

<style>
.content {
  display: flex;
  gap: 10px;
}
</style>