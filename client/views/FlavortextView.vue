<template>
    <div class="flex flex-col items-center justify-center gap-y-4 pb-20">
        <GameWinContainer
            v-if="store.isFlavortextGameWon"
            :pokemon="componentStore.guesses[0]"
            :color="colors.at(-1)"
        />
        <div v-else class="card w-[450px] p-4 text-justify font-pkmEmerald text-xl italic">
            "{{ secretPokemon.flavorText }}"
        </div>
        <SearchField
            v-if="!store.isFlavortextGameWon"
            :pokemonNames="componentStore.pokemonNames"
            @submit-guess="submitGuess"
        />
        <HintContainer
            v-if="!store.isFlavortextGameWon && componentStore.guesses.length"
            :numberOfGuesses="componentStore.guesses.length"
            :stylingHintOne="''"
            :stylingHintTwo="''"
            :stylingHintThree="''"
        >
            <template #hint1>HINT 1 </template>
            <template #hint2> HINT 2 </template>
            <template #hint3> HINT 3 </template>
        </HintContainer>
        <DailyGamesWonContainer
            v-if="!componentStore.guesses.length"
            :dailyGamesWon="dailyGamesWon"
        />
        <div v-if="componentStore.guesses.length" class="flex flex-col gap-1">
            <SingleResultHeader :headerText="GuessType.Pokemon" />
            <SingeResultContainer
                v-for="(guess, i) in componentStore.guesses"
                :key="guess"
                :value="guess"
                :pokemonName="guess"
                :color="colors[componentStore.guesses.length - 1 - i]"
                :isCorrect="isCorrectGuess(guess, secretPokemon.name)"
            />
        </div>
        <PreviousPokemonCard v-else :pokemonName="yesterdaysPokemon.name" />
    </div>
</template>

<script setup>
import SearchField from '../components/SearchField.vue';
import GameWinContainer from '../components/GameWinContainer.vue';
import SingleResultHeader from '../components/result/SingleResultHeader.vue';
import SingeResultContainer from '../components/result/SingeResultContainer.vue';
import DailyGamesWonContainer from '../components/infoCards/DailyGamesWonContainer.vue';
import PreviousPokemonCard from '../components/infoCards/PreviousPokemonCard.vue';
import { reactive, ref, onBeforeMount } from 'vue';
import pokemonData from '../../server/data/pokemonData-v5-flavorText.json';
import { useStore } from '../stores/store.js';
import HintContainer from '../components/hints/HintContainer.vue';
import { removePokemonNameFromArray, getRandomColor, isCorrectGuess } from '../services/guess';
import { playWinnerSound } from '../services/sound';
import { launchConfetti } from '../services/confetti.js';
import { GameModes, GuessType } from '../constants';
import * as apiService from '../services/api/apiService.js';
import moment from 'moment';
import {
    clearLocalStorageGameMode,
    setNewDate,
    addColorsToLocalStorage,
    addGuessesToLocalStorage,
} from '../services/localStorage';
import { setSecretPokemon, getDailyGamesWonCount, updateUserWithGameWon } from '../services/game';

const store = useStore();

const getSortedPokemonNames = () => pokemonData.map((pokemonInfo) => pokemonInfo.name).sort();

const componentStore = reactive({
    pokemonNames: getSortedPokemonNames(),
    guesses: [],
});

const dailyGamesWon = ref(0);
const dailyFirstTryWins = ref(0);
const yesterdaysPokemon = ref('');
//Updated as soon an correct pokemon is guessed, contrary to store.isGameWon Which is only updated after TotalResultCardFlipDelay
let colors = [];
const secretPokemon = reactive({});

const setDailyGamesWonCount = async () => {
    dailyGamesWon.value = await getDailyGamesWonCount(GameModes.Flavortext);
};

onBeforeMount(async () => {
    await Promise.all([loadFlavortextGameData(), setDailyGamesWonCount()]);
});

//TODO: can refactor?
const incrementGamesWonCount = async () => {
    const response = await apiService.updateStatsFlavortextWins(componentStore.guesses.length);
    dailyGamesWon.value = response.data.dailyStats.flavortextGamesWon;
    dailyFirstTryWins.value = response.data.dailyStats.flavortextFirstTryWins;
};

const decideGame = async (guess) => {
    if (guess === secretPokemon.name) {
        //Wait for all cards to flip
        playWinnerSound();

        launchConfetti(colors.at(-1) === 'shiny', componentStore.guesses.length === 1);
        store.setIsFlavortextGameWon(true);
        console.log('🥳🎉🎊 Congrats! You guessed the secret pokemon: ' + guess);
        incrementGamesWonCount();
        const user = await updateUserWithGameWon(GameModes.Flavortext, componentStore.guesses.length);
        store.setUser(user);
    } else {
        console.log('❌❌❌ Wrong Guess. The secret pokemon was not ' + guess + ' ❌❌❌');
    }
};

const updateYesterdaysPokemon = async () => {
    yesterdaysPokemon.value = (await apiService.getFlavortextPreviousSecretPokemon()).data;
};

const submitGuess = (guess) => {
    if (!guess) return;

    const { removedName, updatedNames } = removePokemonNameFromArray(guess, componentStore.pokemonNames);

    if (updatedNames.length >= componentStore.pokemonNames.length) return;

    colors.push(getRandomColor());
    componentStore.guesses.unshift(removedName);
    componentStore.pokemonNames = updatedNames;
    addGuessesToLocalStorage(GameModes.Flavortext, componentStore.guesses);
    addColorsToLocalStorage(GameModes.Flavortext, colors);
    decideGame(removedName);
};

//TODO: can refactor?
const loadSecretPokemon = () => {
    if (!localStorage.flavortextSecretPokemon) return;
    Object.assign(secretPokemon, JSON.parse(localStorage.flavortextSecretPokemon));
};

//TODO: can refactor?
const loadGuesses = () => {
    const guesses = localStorage.flavortextGuesses;
    if (guesses) componentStore.guesses = JSON.parse(guesses);
};

//TODO: can refactor?
const loadColors = () => {
    const loadedColors = localStorage.flavortextColors;
    if (loadedColors) colors = JSON.parse(loadedColors);
};

//TODO: can refactor?
const loadIsFlavortextGameWon = () => {
    const loadedIsFlavortextGameWon = localStorage.isFlavortextGameWon;
    if (loadedIsFlavortextGameWon && loadedIsFlavortextGameWon === 'true') {
        store.setIsFlavortextGameWon(true);
    } else {
        store.setIsFlavortextGameWon(false);
    }
};

//TODO: can refactor?
const removePokemonsFromGuessPool = () => {
    componentStore.pokemonNames = componentStore.pokemonNames.filter((pokemon) => {
        for (const guessedPokemon of componentStore.guesses) {
            if (guessedPokemon === pokemon) return false;
        }
        return true;
    });
};

//TODO: can refactor?
const loadFlavortextGameData = async () => {
    const dayOfLastUpdate = localStorage.dayOfLastUpdate;
    if (!dayOfLastUpdate) setNewDate();

    loadSecretPokemon();
    const currSecretPokemon = await (await apiService.getFlavortextSecretPokemon()).data;

    if (
        parseInt(dayOfLastUpdate) == moment().date() &&
        secretPokemon &&
        secretPokemon?.name === currSecretPokemon?.name
    ) {
        //Load TODO: Implement
        loadColors();
        loadGuesses();
        loadIsFlavortextGameWon();
        removePokemonsFromGuessPool();
    } else {
        //Fresh game TODO: Implement
        clearLocalStorageGameMode(GameModes.Flavortext);
        store.setIsFlavortextGameWon(false);
        await setSecretPokemon(GameModes.Flavortext);
        setNewDate();
    }
    updateYesterdaysPokemon();
};
</script>
