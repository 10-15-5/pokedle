<template>
    <div class="relative">
        <input
            class="card w-[200px] p-4 font-pkmEmerald focus:outline-none"
            label="Type pokemon name"
            type="text"
            placeholder="Type Pokemon Name..."
            v-model="searchTerm"
            @keypress.enter="submitGuess(searchTerm)"
            v-click-outside="onClickOutsideSearchField"
            :active="isSearchFieldActive"
            @click="isSearchFieldActive = true"
        />
        <div class="absolute z-10">
            <VirtualScroller
                v-if="searchPokemonNames.length && isSearchFieldActive"
                class="card absolute z-10 mt-[-2px] w-[200px]"
                :class="getSearchSuggestionsHeight"
                :items="searchPokemonNames"
                :itemSize="itemSize"
            >
                <template v-slot:item="{ item, options }">
                    <div
                        :style="{ height: `${itemSize}px` }"
                        @click="submitGuess(item)"
                        class="justify-left flex cursor-pointer items-center bg-light-bg text-black hover:!bg-neutral-200 hover:!text-green-500 dark:!bg-dark-bg dark:!text-dark-text hover:dark:!bg-neutral-700 hover:dark:!text-green-300"
                    >
                        <img
                            class="sprite-img ml-3"
                            :src="
                                'https://img.pokemondb.net/sprites/ruby-sapphire/normal/' +
                                removeSpecialCharactersExceptDashFromString(item) +
                                '.png'
                            "
                            alt=""
                        />
                        <p class="text-4 ml-4 font-pkmEmerald capitalize">
                            {{ item }}
                        </p>
                    </div>
                    <hr
                        class="border-t-2 border-light-border dark:border-dark-border"
                        v-if="options.index + 1 < searchPokemonNames.length"
                    />
                </template>
            </VirtualScroller>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';
import { removeSpecialCharactersExceptDashFromString } from '../helpers.js';

const itemSize = 52;
//Multiplications of 56 since; itemSize: 52 + borderWidth: 2*2 = 56
const searchFieldHeights = [
    'h-[0px]',
    'h-[56px]',
    'h-[112px]',
    'h-[168px]',
    'h-[224px]',
    'h-[270px]',
]

const props = defineProps({
    pokemonNames: Object,
});

const isSearchFieldActive = ref(false);

const onClickOutsideSearchField = () => {
    isSearchFieldActive.value = false;
};

const emit = defineEmits(['submitGuess']);
let searchTerm = ref('');

const searchPokemonNames = computed(() => {
    if (searchTerm.value === '') {
        return [];
    }

    const test = props.pokemonNames.filter((name) =>
        name.startsWith(searchTerm.value.toLowerCase())
    );
    return test;
});

const getSearchSuggestionsHeight = computed(() => 
     searchPokemonNames.value.length >= 5 ? searchFieldHeights[5] : searchFieldHeights[searchPokemonNames.value.length]
);

const submitGuess = (pokemonName) => {
    searchTerm.value = '';
    emit('submitGuess', pokemonName);
};
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
