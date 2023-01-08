<template>
	<div class="content">
        <SquareContent :pokemon="removeSpecialCharactersFromString(pokemonData.Name).toLowerCase()" />
		<SquareContent :guessResult="'CorrectGuess'" :guessText="'fire'" />
		<SquareContent :guessResult="'WrongGuess'" :guessText="'water'" />
		<SquareContent :guessResult="'PartlyCorrectGuess'" :guessText="'Fully evolved'" />
		<SquareContent :guessResult="'WrongGuess'" :guessText="'gen 1'" />
	</div>
</template>

<script setup>
import SquareContent from './SquareContent.vue';

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