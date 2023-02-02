<template>
    <div>
        <v-text-field label="Type pokemon name"
                      class="search-field"
                      variant="outlined"
                      hide-details
                      v-model="searchTerm"
                      @keypress.enter="submitGuess(searchTerm)" />
        <v-card class="search-suggestion-dropdown"
                v-if="searchPokemonNames.length">
            <v-virtual-scroll :items="searchPokemonNames"
                              max-height="300"
                              item-height="48">
                <template v-slot:default="{ item }">
                    <v-list-item :title="`${item}`"
                                 @click="submitGuess(item)"
                                 class="search-suggestions">
                        <template v-slot:prepend>
                            <img class="mt-2"
                                 :src="'https://img.pokemondb.net/sprites/ruby-sapphire/normal/' + removeSpecialCharactersExceptDashFromString(item) + '.png'"
                                 alt="" />
                        </template>
                    </v-list-item>
                </template>
            </v-virtual-scroll>
        </v-card>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';
import { removeSpecialCharactersExceptDashFromString } from '../helpers.js';

const props = defineProps({
    pokemonNames: Object,
});
const emit = defineEmits(['submitGuess'])
let searchTerm = ref('');

const searchPokemonNames = computed(() => {
    console.log("test")
    if (searchTerm.value === '') {
        return [];
    }

    const test = props.pokemonNames.filter(name =>
        name.startsWith(searchTerm.value.toLowerCase()));
    return test;
});

const submitGuess = (pokemonName) => {
    searchTerm.value = "";
    emit('submitGuess', pokemonName);
}

</script>

<style scoped>
img {
    width: 35px;
}

.search-field {
    background: white;
    font-family: pkmEmerald;
    font-weight: 600;
    border-radius: 4px;
}

.search-suggestions {
    display: flex;
    gap: 10%;
    justify-content: left;
    align-items: center;
}

.search-suggestion-dropdown {
    position: absolute;
    z-index: 2;
}

.search-field,
.search-suggestion-dropdown {
    width: 200px;
}
</style>