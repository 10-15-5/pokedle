<template>
  <div class="content">
    <v-card v-for="(result, i) in guessResults"
            :key="i"
            :value="result">
      <SquareContent v-if="result.name"
      :pokemon="removeSpecialCharactersFromString(result.name)"/>
      <SquareContent v-else
      :guessResult="result.guessState"
                     :guessText="result.text" />
    </v-card>
  </div>
</template>

<script setup>
import SquareContent from './SquareContent.vue';

const props = defineProps({
  pokemonName: String,
  guessResults: Object
});
const removeSpecialCharactersFromString = (text) => {
  return text.replace(/[^\w\s]/gi, '');
}

</script>

<style>
.content {
  display: flex;
  flex-direction: row;
  gap: 10px;
  width: 100%;
}
</style>