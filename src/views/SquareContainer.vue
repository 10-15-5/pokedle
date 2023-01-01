<template>
	<div class="content">
		<SquareContent :guessResult="'CorrectGuess'" :guessText="'fire'" />
		<SquareContent
			:guessResult="'PartlyCorrectGuess'"
			:guessText="'water'"
		/>
        <SquareContent
			:guessResult="'WrongGuess'"
			:guessText="'Fully evolved'"
		/>
	</div>
</template>

<script>
import SquareContent from './SquareContent.vue';
import {guessState} from '../constants.js';

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