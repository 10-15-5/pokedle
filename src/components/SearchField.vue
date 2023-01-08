<template>
  <div>
    <v-text-field label="Type pokemon name"
                  variant="solo"
                  v-model="searchTerm"
                  @keypress.enter="[$emit('someEvent', searchTerm), clearSearchField()]" />
    <v-card class="mt-n6"
            v-if="searchPokemonNames.length">
      <v-list-item v-for="pokemonName in searchPokemonNames"
                   :key="pokemonName"
                   :value="pokemonName"
                   @click="selectPokemonName(pokemonName)"
                   >
        <div class="search-suggestion-dropdown">
          <img class="mt-2"
               :src="'https://img.pokemondb.net/sprites/ruby-sapphire/normal/' + removeSpecialCharactersExceptDashFromString(pokemonName) + '.png'"
               alt="" />
          {{ pokemonName }}
        </div>
      </v-list-item>
  </v-card>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue';
import {removeSpecialCharactersExceptDashFromString} from '../helpers.js';

const props = defineProps({
  pokemonNames: Object,
});

let searchTerm = ref('');

const searchPokemonNames = computed(() => {
  if (searchTerm.value === '') {
    return [];
  }

  let matches = 0

  return props.pokemonNames.filter(name => {
    if (
      name.startsWith(searchTerm.value.toLowerCase())
      && matches < 5
    ) {
      matches++;
      return name
    }
  });
});

const selectPokemonName = (pokemonName) => {
  searchTerm.value = pokemonName;
}

const clearSearchField = () => {
  console.log(searchPokemonNames.value)
  if (searchPokemonNames.value.length>0) {
    searchTerm.value = "";
  }
}
</script>

<style>
img {
  width: 35px;
}

.search-suggestion-dropdown {
  display: flex;
  gap: 10%;
  justify-content: left;
  align-items: center;
}
</style>