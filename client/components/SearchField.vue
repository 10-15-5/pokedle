<template>
    <div class="relative">
        <input
            class="card w-[200px] p-4 font-pkmEmerald focus:outline-none"
            label="Type pokemon name"
            type="text"
            placeholder="Type Pokemon Name..."
            v-model="searchTerm"
            @keypress.enter="submitGuess(searchPokemonNames[idx]), resetScroll"
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
                        class="justify-left hover flex cursor-pointer items-center border-b-2 border-light-border bg-light-bg text-black hover:!bg-neutral-200 hover:!text-green-500 dark:border-dark-border dark:!bg-dark-bg dark:!text-dark-text hover:dark:!bg-neutral-700 hover:dark:!text-green-300"
                        :class="getHover(options.index)"
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
    '!h-[54px]',
    '!h-[104px]',
    '!h-[156px]',
    '!h-[208px]',
    '!h-[260px]',
];

const idx = ref(0);

const props = defineProps({
    pokemonNames: Object,
});

const getHover = (index) => {
    if (idx.value === index)
        return '!bg-neutral-200 !text-green-500 dark:!bg-neutral-700 dark:!text-green-300';
};

const isSearchFieldActive = ref(false);
const searchTerm = ref('');
const virtualScroller = ref(null);

const scrollDown = () => {
    idx.value += searchPokemonNames.value.length - 1 > idx.value ? 1 : 0;
    console.log(idx.value);
    virtualScroller.value?.scrollTo({ top: idx.value * itemSize });
};

const scrollUp = () => {
    idx.value -= idx.value > 0 ? 1 : 0;
    console.log(idx.value);
    virtualScroller.value?.scrollTo({ top: idx.value * itemSize });
};

const onClickOutsideSearchField = () => {
    isSearchFieldActive.value = false;
};

const emit = defineEmits(['submitGuess']);

watch(searchTerm, () => {
    idx.value = 0;
    virtualScroller.value?.scrollToIndex(0);
});

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
