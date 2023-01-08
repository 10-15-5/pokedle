<template>
	<div class="content">
        <SquareContent :guessPokemon="pokemonData.Name.toLowerCase()" />
		<SquareContent :guessResult="'CorrectGuess'" :guessText="'fire'" />
		<SquareContent :guessResult="'WrongGuess'" :guessText="'water'" />
		<SquareContent :guessResult="'PartlyCorrectGuess'" :guessText="'Fully evolved'" />
		<SquareContent :guessResult="'WrongGuess'" :guessText="'gen 1'" />
	</div>
</template>

<script setup>
import SquareContent from './SquareContent.vue';

const props = defineProps({
  pokemonData: String
});

</script>

<style>
.content {
	display: flex;
	gap: 10px;
}
</style>