<template>
    <div class="relative">
        <input
            class="card w-[200px] p-4 font-pkmEmerald focus:outline-none"
            label="Type pokemon name"
            type="text"
            placeholder="Type Pokemon Name..."
            v-model="searchTerm"
            @keypress.enter="submitGuess(searchPokemonNames[idx === -1 ? 0 : idx])"
            :onFocus="() => {isSearchFieldActive = true}"
            v-click-outside="onClickOutsideSearchField"
            @keydown.arrow-down="scrollDown"
            @keydown.arrow-up="scrollUp"
        />
        <div class="absolute z-10">
            <VirtualScroller
                v-if="searchPokemonNames.length && isSearchFieldActive"
                class="card absolute z-10 mt-[-2px] w-[200px]"
                :class="getSearchSuggestionsHeight"
                :items="searchPokemonNames"
                :itemSize="itemSize"
                ref="virtualScroller"
            >
                <template v-slot:item="{ item, options }">
                    <div
                        :style="{ height: `${itemSize}px` }"
                        @click="submitGuess(item)"
                        class="justify-left hover flex cursor-pointer items-center border-light-border bg-light-bg text-black hover:!bg-neutral-100 hover:!text-light-green dark:border-dark-border dark:!bg-dark-bg dark:!text-dark-text hover:dark:!bg-neutral-600 hover:dark:!text-dark-green"
                        :class="getHover(options.index)"
                    >
                        <img
                            class="sprite-img ml-5"
                            :src="
                                'https://img.pokemondb.net/sprites/ruby-sapphire/normal/' +
                                removeSpecialCharactersExceptDashFromString(item) +
                                '.png'
                            "
                            alt=""
                        />
                        <p class="text-4 ml-5 font-pkmEmerald capitalize">
                            {{ item }}
                        </p>
                    </div>
                </template>
            </VirtualScroller>
        </div>
    </div>
</template>

<script setup>
import { computed, ref, watch } from 'vue';
import { removeSpecialCharactersExceptDashFromString } from '../helpers.js';

const itemSize = 52;

const searchFieldHeights = [
    '!h-[0px]',
    '!h-[52px]',
    '!h-[104px]',
    '!h-[156px]',
    '!h-[208px]',
    '!h-[260px]',
];

const idx = ref(0);

let suggestionViewportStart = 0;
const suggestionViewportSize = 4;

const props = defineProps({
    pokemonNames: Object,
});

const getHover = (index) => {
    if (idx.value === index)
        return '!bg-neutral-100 !text-light-green dark:!bg-neutral-600 dark:!text-dark-green';
};

const isSearchFieldActive = ref(false);
const searchTerm = ref('');
const virtualScroller = ref(null);

const scrollDown = () => {
    idx.value += searchPokemonNames.value.length - 1 > idx.value ? 1 : 0;
    if (idx.value > suggestionViewportStart + suggestionViewportSize) {
        suggestionViewportStart++;
        virtualScroller.value?.scrollTo({ top: suggestionViewportStart * itemSize });
    }
};

const scrollUp = () => {
    idx.value -= idx.value > 0 ? 1 : 0;
    if (idx.value < suggestionViewportStart) {
        suggestionViewportStart--;
        virtualScroller.value?.scrollTo({ top: suggestionViewportStart * itemSize });
    }
};

const onClickOutsideSearchField = () => {
    isSearchFieldActive.value = false;
    resetSuggestionsView();
};

const emit = defineEmits(['submitGuess']);

watch(searchTerm, () => {
    resetSuggestionsView();
});

const resetSuggestionsView = () => {
    idx.value = -1;
    suggestionViewportStart = 0;
    virtualScroller.value?.scrollToIndex(0);
};

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
    searchPokemonNames.value.length >= 5
        ? searchFieldHeights[5]
        : searchFieldHeights[searchPokemonNames.value.length]
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
