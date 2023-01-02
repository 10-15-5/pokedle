<template>
	<div class="content">
        <SquareContent :guessPokemon="'charizard'" />
		<SquareContent :guessResult="'CorrectGuess'" :guessText="'fire'" />
		<SquareContent :guessResult="'WrongGuess'" :guessText="'water'" />
		<SquareContent :guessResult="'PartlyCorrectGuess'" :guessText="'Fully evolved'" />
		<SquareContent :guessResult="'WrongGuess'" :guessText="'gen 1'" />
	</div>
</template>

<script>
import SquareContent from './SquareContent.vue';

export default {
    components:{
        SquareContent,
    }
}

</script>

<style>
.content {
	display: flex;
	gap: 10px;
}
</style>