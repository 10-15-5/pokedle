<template>
  <v-card class="content" variant="outlined">
    <div v-for="(result, i) in guessResults"
            :key="i"
            :value="result"
            >
      <SquareContent v-if="result.name"
      :pokemon="removeSpecialCharactersExceptDashFromString(result.name)"/>
      <SquareContent v-else
      :guessResult="result.guessState"
                     :guessText="result.text" />
    </div>
</v-card>
</template>

<script setup>
import SquareContent from './SquareContent.vue';
import {removeSpecialCharactersExceptDashFromString} from '../helpers.js';
const props = defineProps({
  pokemonName: String,
  guessResults: Object
});
</script>

<style scoped>
.content {
  display: flex;
  flex-direction: row;
  justify-content: center;
  gap: 8px;
  padding: 8px 0 8px 8px;
  background-color: rgb(255, 255, 255, .9);
}
</style>