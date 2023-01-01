<template>
	<div class="square-content" :style="{ background: getColor }">
		<p>{{ guessText }}</p>
	</div>
</template>

<script>
import {guessState} from '../constants.js';
export default {
    props: {
        guessResult: {
            required: true,
            type: String,
        },
        guessText: {
            required: true,
            type: String,
        }
    },
    computed: {
        getColor () {
            switch(this.guessResult) {
                case guessState.CorrectGuess:
                   return "greenyellow";
                case guessState.PartlyCorrectGuess:
                    return "orange";
                default:
                    return "red";
            }
        }
    }
}

</script>

<style>
.square-content {
	width: 60px;
	height: 60px;
	border-style: solid;
	border-color: black;
}
</style>
