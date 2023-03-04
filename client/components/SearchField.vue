<template>
    <div class="relative">
        <input class="card p-4 font-pkmEmerald w-[200px] focus:outline-none"
               label="Type pokemon name"
               type="text"
               placeholder="Type Pokemon Name..."
               v-model="searchTerm"
               @keypress.enter="submitGuess(searchTerm)"
               v-click-outside="onClickOutsideSearchField"
               :active="isSearchFieldActive"
               @click="isSearchFieldActive = true" />
        <div class="absolute z-10">
            <VirtualScroller v-if="searchPokemonNames.length && isSearchFieldActive"
                             class="h-[260px] w-[200px] absolute z-10 card mt-[-2px]"
                             :items="searchPokemonNames"
                             :itemSize="52">
                <template v-slot:item="{ item, options }">
                    <div style="height: 52px;"
                         @click="submitGuess(item)"
                         class="flex justify-left items-center
                                 bg-light-bg text-black 
                                         dark:!text-dark-text
                                         dark:!bg-dark-bg cursor-pointer
                                         hover:!text-orange-500
                                         hover:!bg-neutral-200
                                         hover:dark:!text-orange-300
                                         hover:dark:!bg-neutral-700">
                        <img class="sprite-img ml-3"
                             :src="'https://img.pokemondb.net/sprites/ruby-sapphire/normal/' +
                                 removeSpecialCharactersExceptDashFromString(item) + '.png'"
                             alt="" />
                        <p class="ml-4 capitalize font-pkmEmerald text-4">{{ item }}</p>
                    </div>
                    <hr class="border-t-2 border-light-border dark:border-dark-border"
                        v-if="options.index + 1 < searchPokemonNames.length">
                </template>
            </VirtualScroller>
        </div>
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
</style>