<template>
    <div class="relative">
        <input
            ref="inputField"
            class="card relative w-[200px] p-4 font-pkmEmerald focus:outline-none"
            type="text"
            :class="filteredPokemons.length && isSearchFieldActive ? ' z-40' : ''"
            :placeholder="text().searchField.typePokemon"
            v-model="searchTerm"
            @keypress.enter="submitGuess(filteredPokemons[idx === -1 ? 0 : idx])"
            :onFocus="
                () => {
                    isSearchFieldActive = true;
                }
            "
            @keydown.arrow-down="scrollDown"
            @keydown.arrow-up="scrollUp"
            autocomplete="off"
        />
        <div class="absolute z-20">
            <VirtualScroller
                v-if="filteredPokemons.length && isSearchFieldActive"
                class="card mt-[-2px] w-[200px]"
                :class="getSearchSuggestionsHeight"
                :items="filteredPokemons"
                :itemSize="itemSize"
                ref="virtualScroller"
            >
                <template v-slot:item="{ item, options }">
                    <div
                        :style="{ height: `${itemSize}px` }"
                        @click="submitGuess(item)"
                        class="justify-left hover flex cursor-pointer items-center border-light-border bg-light-bg text-black hover:!bg-neutral-100 hover:!text-light-emerald dark:border-dark-border dark:!bg-dark-bg dark:!text-dark-text hover:dark:!bg-neutral-600 hover:dark:!text-dark-emerald"
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
import { removeSpecialCharactersExceptDashFromString } from '../services/helpers.js';
import { useClickOutside } from '../composables/useClickOutside.js';
import { text } from '../services/language';

const itemSize = 52;

const searchFieldHeights = ['!h-[0px]', '!h-[52px]', '!h-[104px]', '!h-[156px]', '!h-[208px]', '!h-[260px]'];

const isSearchFieldActive = ref(false);
const searchTerm = ref('');
const inputField = ref(null);

useClickOutside(inputField, () => {
    isSearchFieldActive.value = false;
    resetSuggestionsView();
});

const idx = ref(0);

let suggestionViewportStart = 0;
const suggestionViewportSize = 4;

const props = defineProps({
    pokemonNames: Object,
});

const getHover = (index) => {
    if (idx.value === index) return '!bg-neutral-100 !text-light-emerald dark:!bg-neutral-600 dark:!text-green-300';
};

const virtualScroller = ref(null);

const scrollDown = () => {
    idx.value += filteredPokemons.value.length - 1 > idx.value ? 1 : 0;
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

const emit = defineEmits(['submitGuess']);

watch(searchTerm, () => {
    resetSuggestionsView();
});

const resetSuggestionsView = () => {
    idx.value = -1;
    suggestionViewportStart = 0;
    virtualScroller.value?.scrollToIndex(0);
};

const filteredPokemons = computed(() => {
    if (searchTerm.value === '') {
        return [];
    }

    const test = props.pokemonNames.filter((n) => n.startsWith(searchTerm.value.toLowerCase()));
    return test;
});

const getSearchSuggestionsHeight = computed(() =>
    filteredPokemons.value.length >= 5 ? searchFieldHeights[5] : searchFieldHeights[filteredPokemons.value.length]
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

</style>
