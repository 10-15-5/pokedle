<template>
    <div class="flex flex-col items-center justify-center gap-y-4 pb-20">
        <GameWinContainer
            v-if="store.isSilhouetteGameWon"
            :pokemon="componentStore.guesses[0]"
            :color="colors.at(-1)"
            :twitterText="silhouetteTwitterText()"
        />
        <div
            v-else-if="secretPokemon.name"
            class="card relative flex flex-col items-center justify-center gap-3 py-4 px-8 sm:px-4 sm:text-sm"
        >
            <span class="card py-1 px-2 font-pkmEmerald">Silhouette, Rotated</span>
            <ResultSquare
                class="md:transform-none"
                :class="getRotation"
                :pokemon="secretPokemon.name"
                :type="GuessType.Blackout"
                :isLarge="true"
            />
        </div>
        <RouterLink v-if="store.isSilhouetteGameWon" to="/flavortext">
            <GameModeButton :gameMode="GameModes.Flavortext" />
        </RouterLink>
        <SearchField
            v-if="!store.isSilhouetteGameWon"
            :pokemonNames="componentStore.pokemonNames"
            @submit-guess="submitGuess"
        />
        <HintContainer
            v-if="!store.isSilhouetteGameWon && componentStore.guesses.length"
            :numberOfGuesses="componentStore.guesses.length"
            :stylingHintThree="'flex justify-center'"
            :guessesRequiredForHintOne="SilhouetteGuessesNeededForHintOne"
            :guessesRequiredForHintTwo="SilhouetteGuessesNeededForHintTwo"
        >
            <template #hint1
                ><div class="flex flex-row items-center justify-center gap-6 sm:gap-2">
                    <div v-for="(field, i) in hintOne" :key="i" :value="field" class="flex flex-col gap-1">
                        <span class="card items-center justify-center py-1 sm:py-0">{{ field.title }}</span>
                        <ResultSquare :guessResult="field.guessState" :guessText="field.text" :type="field.type" />
                    </div>
                </div>
            </template>
            <template #hint2>
                <div class="flex flex-row items-center justify-center gap-6 sm:gap-2">
                    <p>
                        Secret Pokemon Starts With The Letter:
                        <b class="text-lg text-light-orange dark:!text-dark-orange">
                            {{ secretPokemon.name.charAt(0).toUpperCase() }}</b
                        >
                    </p>
                </div>
            </template>
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
import GameModeButton from '../components/buttons/GameModeButton.vue';
import SingleResultHeader from '../components/result/SingleResultHeader.vue';
import ResultSquare from '../components/result/ResultSquare.vue';
import SingeResultContainer from '../components/result/SingeResultContainer.vue';
import DailyGamesWonContainer from '../components/infoCards/DailyGamesWonContainer.vue';
import PreviousPokemonCard from '../components/infoCards/PreviousPokemonCard.vue';
import { reactive, ref, onBeforeMount, computed } from 'vue';
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
    SilhouetteGuessesNeededForHintOne,
    SilhouetteGuessesNeededForHintTwo,
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
import { getCurrentSilhouettePokemonNumber } from '../helpers.js';

const store = useStore();

const getSortedPokemonNames = () => pokemonData.map((pokemonInfo) => pokemonInfo.name).sort();

const componentStore = reactive({
    pokemonNames: getSortedPokemonNames(),
    guesses: [],
});

const hintOne = reactive([]);

const dailyGamesWon = ref(0);
const dailyFirstTryWins = ref(0);
const yesterdaysPokemon = ref('');
//Updated as soon an correct pokemon is guessed, contrary to store.isGameWon Which is only updated after TotalResultCardFlipDelay
let colors = [];
const secretPokemon = reactive({});

const rotate = ['!rotate-90', '!rotate-180', '!-rotate-90'];

const getRotation = computed(() => rotate[Math.floor(Math.random() * rotate.length)]);

const setDailyGamesWonCount = async () => {
    dailyGamesWon.value = await getDailyGamesWonCount(GameModes.Silhouette);
};

onBeforeMount(async () => {
    await Promise.all([loadSilhouetteGameData(), setDailyGamesWonCount(), updateYesterdaysPokemon()]);
});

const silhouetteTwitterText = () => {
    const sub1 =
        componentStore.guesses.length === 1
            ? 'FIRST TRY 🎰🍀🥳🤩'
            : `in ${componentStore.guesses.length} tries!🍇🥭🍒💖`;

    const header = `I guessed the #${getCurrentSilhouettePokemonNumber()} silhouette #Pokedle Pokémon ${sub1}\n\n`;

    const footer = `Play at pokedle.gg 🎮!`;

    return header + footer;
};

const setHintOne = () => {
    if (hintOne.length) {
        return;
    }

    const { fields: correctFields } = getGuessResults(secretPokemon.name, secretPokemon);

    const titles = Object.keys(GuessFieldTitles);
    const correctFieldsWithTitles = Object.keys(correctFields).map((f, i) => ({
        ...correctFields[f],
        title: GuessFieldTitles[titles.at(i)],
    }));
    correctFieldsWithTitles[7].title = 'Gen';
    hintOne.push(correctFieldsWithTitles[1], correctFieldsWithTitles[2], correctFieldsWithTitles[5]); //Type 1, Type 2
};

//TODO: can refactor?
const incrementGamesWonCount = async () => {
    const response = await apiService.updateStatsSilhouetteWins(componentStore.guesses.length);
    dailyGamesWon.value = response.data.dailyStats.silhouetteGamesWon;
    dailyFirstTryWins.value = response.data.dailyStats.silhouetteFirstTryWins;
};

const decideGame = async (guess) => {
    if (guess === secretPokemon.name) {
        //Wait for all cards to flip
        playWinnerSound();

        launchConfetti(colors.at(-1) === 'shiny', componentStore.guesses.length === 1);
        store.setIsSilhouetteGameWon(true);
        console.log('🥳🎉🎊 Congrats! You guessed the secret pokemon: ' + guess);
        const [user] = await Promise.all([
            updateUserWithGameWon(GameModes.Silhouette, componentStore.guesses.length),
            incrementGamesWonCount(),
        ]);
        store.setUser(user);
    } else {
        console.log('❌❌❌ Wrong Guess. The secret pokemon was not ' + guess + ' ❌❌❌');
    }
};

const updateYesterdaysPokemon = async () => {
    yesterdaysPokemon.value = (await apiService.getSilhouettePreviousSecretPokemon()).data;
};

const submitGuess = async (guess) => {
    if (!guess) return;

    const { removedName, updatedNames } = removePokemonNameFromArray(guess, componentStore.pokemonNames);

    if (updatedNames.length >= componentStore.pokemonNames.length) return;

    colors.push(getRandomColor());
    componentStore.guesses.unshift(removedName);
    componentStore.pokemonNames = updatedNames;
    addGuessesToLocalStorage(GameModes.Silhouette, componentStore.guesses);
    addColorsToLocalStorage(GameModes.Silhouette, colors);
    await decideGame(removedName);
};

//TODO: can refactor?
const loadSecretPokemon = () => {
    if (!localStorage.silhouetteSecretPokemon) return;
    Object.assign(secretPokemon, JSON.parse(localStorage.silhouetteSecretPokemon));
};

//TODO: can refactor?
const loadGuesses = () => {
    const guesses = localStorage.silhouetteGuesses;
    if (guesses) componentStore.guesses = JSON.parse(guesses);
};

//TODO: can refactor?
const loadColors = () => {
    const loadedColors = localStorage.silhouetteColors;
    if (loadedColors) colors = JSON.parse(loadedColors);
};

//TODO: can refactor?
const loadIsSilhouetteGameWon = () => {
    const loadedIsSilhouetteGameWon = localStorage.isSilhouetteGameWon;
    if (loadedIsSilhouetteGameWon && loadedIsSilhouetteGameWon === 'true') {
        store.setIsSilhouetteGameWon(true);
    } else {
        store.setIsSilhouetteGameWon(false);
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
const loadSilhouetteGameData = async () => {
    const dayOfLastUpdate = localStorage.dayOfLastUpdate;
    if (!dayOfLastUpdate) setNewDate();

    loadSecretPokemon();
    const currSecretPokemon = await (await apiService.getSilhouetteSecretPokemon()).data;

    if (
        parseInt(dayOfLastUpdate) == moment().date() &&
        secretPokemon &&
        secretPokemon?.name === currSecretPokemon?.name
    ) {
        loadColors();
        loadGuesses();
        loadIsSilhouetteGameWon();
        removePokemonsFromGuessPool();
    } else {
        clearLocalStorageGameMode(GameModes.Silhouette);
        store.setIsSilhouetteGameWon(false);
        setSecretPokemon(GameModes.Silhouette, currSecretPokemon);
        loadSecretPokemon();
        setNewDate();
    }
    setHintOne();
    updateCurrentUserStreakDisplay(GameModes.Silhouette);
};
</script>
