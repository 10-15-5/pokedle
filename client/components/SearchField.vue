<template>
    <div>
        <v-text-field label="Type pokemon name"
                      class="search-field border-black"
                      variant="solo"
                      hide-details
                      v-model="searchTerm"
                      @keypress.enter="submitGuess(searchTerm)"
                      v-click-outside="onClickOutsideSearchField"
                      :active="isSearchFieldActive"
                      @click="isSearchFieldActive = true" />
        <v-card class="search-suggestion-dropdown border-black bg-white"
                variant="outlined"
                v-if="searchPokemonNames.length && isSearchFieldActive">
            <v-virtual-scroll :items="searchPokemonNames"
                              max-height="260"
                              item-height="52">
                <template v-slot:default="{ item, index }">
                    <v-list-item @click="submitGuess(item)"
                                 class="search-suggestions">
                        <template v-slot:prepend>
                            <img class="sprite-img ml-2"
                                 :src="'https://img.pokemondb.net/sprites/ruby-sapphire/normal/' + removeSpecialCharactersExceptDashFromString(item) + '.png'"
                                 alt="" />
                        </template>
                        <template v-slot:default>
                            <p class="ml-1">{{ item }}</p>
                        </template>
                    </v-list-item>
                    <v-divider v-if="index + 1 < searchPokemonNames.length"
                               class="suggestion-divider border-black"
                               thickness="2px"></v-divider>
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

const isSearchFieldActive = ref(false);

const onClickOutsideSearchField = () => {
    isSearchFieldActive.value = false;
}

const emit = defineEmits(['submitGuess'])
let searchTerm = ref('');

const searchPokemonNames = computed(() => {
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
.sprite-img {
    width: 35px;
    height: 35px;
}

.search-field {
    font-weight: 600;
    border-style: solid;
    border-radius: 5px;
    border-width: 2px;
}

.search-suggestions {
    display: flex;
    gap: 14%;
    justify-content: left;
    padding: 8px;
}

p {
    font-family: pkmEmerald;
    font-size: 16px;
    text-transform: capitalize;
}

.suggestion-divider {
    opacity: 100%;
}

.search-suggestions,
.search-field {
    font-family: pkmEmerald;
    text-transform: capitalize;
}

.search-suggestion-dropdown {
    position: absolute;
    z-index: 2;
    margin-top: -1px;
    border-width: 2px;

}

.search-field,
.search-suggestion-dropdown {
    width: 200px;
}
</style>