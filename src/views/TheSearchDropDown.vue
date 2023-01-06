<template>
  <div class="guess-field">
          <v-text-field label="Type pokemon name"
                        variant="solo"
                        v-model="searchTerm"
                        @keypress.enter="$emit('someEvent', selectedPokemonName)"
                        />
          <v-card class="search-dropdown"
                  v-if="searchPokemonNames.length">
            <v-list-item v-for="pokemonName in searchPokemonNames"
                         :key="pokemonName"
                         @click="selectPokemonName(pokemonName)">{{ pokemonName }}</v-list-item>
          </v-card>
        </div>
</template>

<script>
import { computed, ref } from 'vue';

export default {
  props: {
    pokemonNames: Object,
  },
  setup(props) {
    let searchTerm = ref('');
    let selectedPokemonName = ref('');
    let showDropdown = ref(true);

    const searchPokemonNames = computed(() => {
      if (searchTerm.value === '') {
        return [];
      }
      
      console.log("call to searchPokemonNames")

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
      showDropdown.value=false;
      console.log("select: " + selectedPokemonName.value)
    }

    return {
      searchTerm,
      searchPokemonNames,
      selectPokemonName,
      selectedPokemonName,
    }
  },
}

</script>

<style>

.search-dropdown {
  margin-top: -22px;
  position: relative;
}
</style>