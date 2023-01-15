<template>
  <v-card class="content" variant="outlined">
    <v-card v-for="(result, i) in guessResults"
            :key="i"
            :value="result"
            >
      <SquareContent v-if="result.name"
      :pokemon="removeSpecialCharactersExceptDashFromString(result.name)"/>
      <SquareContent v-else
      :guessResult="result.guessState"
                     :guessText="result.text" />
    </v-card>
</v-card>
</template>

<script setup>
import SquareContent from './SquareContent.vue';
import SquareContentHeader from './SquareContentHeader.vue';
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
  gap: 6px;
  width: 100%;
  background-color: rgb(255, 255, 255, .9);
  padding: 8px;
}
</style>