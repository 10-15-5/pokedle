<template>
    <div class="flex flex-col items-center justify-center gap-y-4 pb-20">
        <GameWinContainer
            v-if="store.isFlavortextGameWon && secretPokemon.name"
            :pokemon="secretPokemon.name"
            :color="colors.at(-1)"
            :twitterText="flavortextTwitterText()"
        />
        <div v-else-if="!store.isFlavortextGameWon && secretPokemon.name" class="card w-[450px] p-4 text-justify font-pkmEmerald text-xl italic sm:w-[350px] sm:text-base">
            "{{ secretPokemon.flavorText }}"
        </div>
        <SearchField
            v-if="!store.isFlavortextGameWon"
            :pokemonNames="componentStore.pokemonNames"
            @submitGuess="submitGuess"
        />
        <HintContainer
            v-if="!store.isFlavortextGameWon && componentStore.guesses.length"
            :numberOfGuesses="componentStore.guesses.length"
            :stylingHintThree="'flex justify-center'"
            :guessesRequiredForHintOne="FlavortextGuessesNeededForHintOne"
            :guessesRequiredForHintTwo="FlavortextGuessesNeededForHintTwo"
            :guessesRequiredForHintThree="FlavortextGuessesNeededForHintThree"
        >
            <template #hint1
                ><div class="flex flex-row items-center justify-center gap-6 sm:gap-2">
                    <div v-for="(field, i) in hintOne" :key="i" :value="field" class="flex flex-col gap-1">
                        <span class="card items-center justify-center py-1 sm:py-0">{{ field.title }}</span>
                        <ResultSquare :guessResult="field.guessState" :guessText="field.text" :type="field.type" :habitat="field.habitat" />
                    </div>
                </div>
            </template>
            <template #hint2>
                <div class="flex flex-row items-center justify-center gap-6 sm:gap-2">
                    <div v-for="(field, i) in hintTwo" :key="i" :value="field" class="flex flex-col gap-1">
                        <span class="card items-center justify-center py-1 sm:py-0" :class="getTextSizeShortenedTitle(field.title)">{{ field.title }}</span>
                        <ResultSquare
                            :guessResult="field.guessState"
                            :guessText="field.text"
                            :type="field.type"
                            :habitat="field.habitat"
                        />
                    </div>
                </div>
            </template>
            <template #hint3
                ><div class="flex flex-row items-center justify-center gap-6 sm:gap-2">
                    <div class="flex flex-col gap-1">
                        <span class="card items-center justify-center py-1 sm:py-0">Shape</span>
                        <ResultSquare :pokemon="secretPokemon.name" :type="GuessType.Blackout" />
                    </div></div
            ></template>
        </HintContainer>
        <DailyGamesWonContainer v-if="!componentStore.guesses.length" :dailyGamesWon="dailyGamesWon" />
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
        <PreviousPokemonCard v-else-if="yesterdaysPokemon.name" :pokemonName="yesterdaysPokemon.name" />
    </div>
</template>

<script setup>
import SearchField from '../components/SearchField.vue';
import GameWinContainer from '../components/GameWinContainer.vue';
import SingleResultHeader from '../components/result/SingleResultHeader.vue';
import ResultSquare from '../components/result/ResultSquare.vue';
import SingeResultContainer from '../components/result/SingeResultContainer.vue';
import DailyGamesWonContainer from '../components/infoCards/DailyGamesWonContainer.vue';
import PreviousPokemonCard from '../components/infoCards/PreviousPokemonCard.vue';
import { reactive, ref, onBeforeMount } from 'vue';
import pokemonData from '../../server/data/pokemonData-v5-flavorText.json';
import { useStore } from '../stores/store.js';
import HintContainer from '../components/hints/HintContainer.vue';
import { removePokemonNameFromArray, getRandomColor, isCorrectGuess, getGuessResults } from '../services/guess';
import { playWinnerSound } from '../services/sound';
import { launchConfetti } from '../services/confetti.js';
import {
    GameModes,
    GuessType,
    GuessFieldTitles,
    FlavortextGuessesNeededForHintOne,
    FlavortextGuessesNeededForHintTwo,
    FlavortextGuessesNeededForHintThree,
} from '../constants';
import * as apiService from '../services/api/apiService.js';
import moment from 'moment';
import {
    clearLocalStorageGameMode,
    setNewDate,
    addColorsToLocalStorage,
    addGuessesToLocalStorage,
} from '../services/localStorage';
import {
    setSecretPokemon,
    getDailyGamesWonCount,
    updateUserWithGameWon,
    updateCurrentUserStreakDisplay,
} from '../services/game';
import { getCurrentFlavortextPokemonNumber } from '../helpers.js';

const store = useStore();

const getSortedPokemonNames = () => pokemonData.map((pokemonInfo) => pokemonInfo.name).sort();

const ShortenedFieldTitles = {
    EvolutionLevel: 'Evol. Lvl.',
    Evolutions: 'Fully Evol.',
    Generation: 'Gen.',
}

const componentStore = reactive({
    pokemonNames: getSortedPokemonNames(),
    guesses: [],
});

const hintOne = reactive([]);
const hintTwo = reactive([]);

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
    await Promise.all([loadFlavortextGameData(), setDailyGamesWonCount(), updateYesterdaysPokemon()]);
});

const flavortextTwitterText = () => {
    const sub1 =
        componentStore.guesses.length === 1 ? 'FIRST TRY 🤯🤩⚡️✨' : `in ${componentStore.guesses.length} tries!🍉🍓🫧`;

    const header = `I guessed the #${getCurrentFlavortextPokemonNumber()} flavortext #Pokedle Pokémon ${sub1}\n\n`;

    const footer = `Play at pokedle.gg 🎮!`;

    return header + footer;
};

const getTextSizeShortenedTitle = (shortenedTitle) => {
    switch(shortenedTitle){
        case ShortenedFieldTitles.EvolutionLevel:
            return 'text-[14px] sm:text-[10px]'
        case ShortenedFieldTitles.Evolutions:
            return 'text-[13px] sm:text-[10px]'
        default:
            return '';
    }
}

const setHints = () => {
    if (hintOne.length || hintTwo.length) {
        return;
    }

    const { fields: correctFields } = getGuessResults(secretPokemon.name, secretPokemon);

    const titles = Object.keys(GuessFieldTitles);
    const correctFieldsWithTitles = Object.keys(correctFields).map((f, i) => ({
        ...correctFields[f],
        title: GuessFieldTitles[titles.at(i)],
    }));

    correctFieldsWithTitles[3].title = ShortenedFieldTitles.EvolutionLevel;
    correctFieldsWithTitles[4].title = ShortenedFieldTitles.Evolutions;
    correctFieldsWithTitles[7].title = ShortenedFieldTitles.Generation;

    hintOne.push(correctFieldsWithTitles[1], correctFieldsWithTitles[2], correctFieldsWithTitles[6]); //Type 1, Type 2, Habitat
    hintTwo.push(correctFieldsWithTitles[3], correctFieldsWithTitles[4], correctFieldsWithTitles[7]); //Evolution lvl, isFullyEvolved, Gen
};

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
        const [user] = await Promise.all([
            updateUserWithGameWon(GameModes.Flavortext, componentStore.guesses.length),
            incrementGamesWonCount(),
        ]);
        store.setUser(user);
    }
};

const updateYesterdaysPokemon = async () => {
    yesterdaysPokemon.value = (await apiService.getFlavortextPreviousSecretPokemon()).data;
};

const submitGuess = async (guess) => {
    if (!guess) return;

    const { removedName, updatedNames } = removePokemonNameFromArray(guess, componentStore.pokemonNames);

    if (updatedNames.length >= componentStore.pokemonNames.length) return;

    colors.push(getRandomColor());
    componentStore.guesses.unshift(removedName);
    componentStore.pokemonNames = updatedNames;
    addGuessesToLocalStorage(GameModes.Flavortext, componentStore.guesses);
    addColorsToLocalStorage(GameModes.Flavortext, colors);
    await decideGame(removedName);
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
        loadColors();
        loadGuesses();
        loadIsFlavortextGameWon();
        removePokemonsFromGuessPool();
    } else {
        clearLocalStorageGameMode(GameModes.Flavortext);
        store.setIsFlavortextGameWon(false);
        setSecretPokemon(GameModes.Flavortext, currSecretPokemon);
        loadSecretPokemon();
        setNewDate();
    }
    setHints();
    updateCurrentUserStreakDisplay(GameModes.Flavortext);
};
</script>
