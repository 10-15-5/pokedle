<template>
  <div>
    <v-text-field label="Type pokemon name"
                  variant="solo"
                  v-model="searchTerm"
                  @keypress.enter="$emit('someEvent', searchTerm)" />
    <v-card class="mt-n6"
            v-if="searchPokemonNames.length">
      <v-list-item v-for="pokemonName in searchPokemonNames"
                   :key="pokemonName"
                   @click="selectPokemonName(pokemonName)">{{ pokemonName }}</v-list-item>
    </v-card>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue';

const props = defineProps({
  pokemonNames: Object,
});

let searchTerm = ref('');
let selectedPokemonName = ref('');

const searchPokemonNames = computed(() => {
  if (searchTerm.value === '') {
    return [];
  }

  console.log("call to searchPokemonNames")
  console.log(props.pokemonNames.length)
  let matches = 0

  return props.pokemonNames.filter(name => {
    if (
      name.toLowerCase().startsWith(searchTerm.value.toLowerCase())
      && matches < 5
    ) {
      matches++;
      return name
    }
  });
});

const selectPokemonName = (pokemonName) => {
  selectedPokemonName.value = pokemonName;
  searchTerm.value = pokemonName;
  console.log("select: " + selectedPokemonName.value)
}
</script>

<style>

</style>