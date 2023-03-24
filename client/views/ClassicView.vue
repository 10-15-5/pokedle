<template>
    <div class="flex flex-col items-center justify-center gap-y-4 pb-20">
        <SearchField
            v-if="!store.isClassicGameWon"
            :pokemonNames="componentStore.pokemonNames"
            @submit-guess="submitGuess"
        />
        <GameWinContainer
            v-if="store.isClassicGameWon"
            :pokemon="componentStore.guesses[0]"
            :twitterText="classicTwitterText()"
            :color="colors.at(-1)"
        />
        <RouterLink v-if="store.isClassicGameWon" to="/silhouette">
            <GameModeButton :gameMode="GameModes.Silhouette" />
        </RouterLink>
        <HintContainer
            v-if="!store.isClassicGameWon && componentStore.guesses.length"
            :numberOfGuesses="componentStore.guesses.length"
            :stylingHintOne="'flex justify-center'"
            :stylingHintTwo="''"
            :stylingHintThree="'flex justify-center'"
            :guessesRequiredForHintOne="ClassicGuessesNeededForHintOne"
            :guessesRequiredForHintTwo="ClassicGuessesNeededForHintTwo"
            :guessesRequiredForHintThree="ClassicGuessesNeededForHintThree"
        >
            <template #hint1>
                <div class="flex flex-col items-center gap-2 sm:gap-1">
                    <span class="card w-full justify-center py-1 sm:py-0">{{ hintTwo.header }}</span>
                    <ResultSquare
                        :color="hintTwo.color"
                        :guessResult="hintTwo.guessResult"
                        :guessText="hintTwo.guessText"
                        :type="hintTwo.type"
                        :habitat="hintTwo.habitat"
                    />
                </div>
            </template>
            <template #hint2>
                {{ secretPokemon.flavorText }}
            </template>
            <template #hint3>
                <div class="flex flex-col items-center gap-2 sm:gap-1">
                    <span class="card w-full justify-center py-1 sm:py-0">Shape</span>
                    <ResultSquare :pokemon="secretPokemon.name" :type="GuessType.Blackout" />
                </div>
            </template>
        </HintContainer>
        <DailyGamesWonContainer v-if="!componentStore.guesses.length" :dailyGamesWon="dailyGamesWon" />
        <div v-if="componentStore.guesses.length" class="flex flex-col gap-y-2 sm:gap-y-1">
            <ResultsHeader class="mb-n1 sm:!mb-0" />
            <ResultsContainer
                v-for="(guess, i) in componentStore.guesses"
                :key="guess"
                :value="guess"
                :guessResult="getGuessResults(guess, secretPokemon, colors[componentStore.guesses.length - 1 - i])"
            />
        </div>
        <PreviousPokemonCard v-else-if="yesterdaysPokemon.name" :pokemonName="yesterdaysPokemon.name" />
    </div>
</template>

<script setup>
import ResultsContainer from '../components/result/ResultsContainer.vue';
import ResultsHeader from '../components/result/ResultsHeader.vue';
import GameWinContainer from '../components/GameWinContainer.vue';
import HintContainer from '../components/hints/HintContainer.vue';
import PreviousPokemonCard from '../components/infoCards/PreviousPokemonCard.vue';
import SearchField from '../components/SearchField.vue';
import GameModeButton from '../components/buttons/GameModeButton.vue';
import pokemonData from '../../server/data/pokemonData-v5-flavorText.json';
import DailyGamesWonContainer from '../components/infoCards/DailyGamesWonContainer.vue';
import { getGuessResults, removePokemonNameFromArray, getRandomColor } from '../services/guess';
import * as apiService from '../services/api/apiService.js';
import { launchConfetti } from '../services/confetti.js';
import { useStore } from '../stores/store.js';
import ResultSquare from '../components/result/ResultSquare.vue';
import { reactive, ref, onBeforeMount, computed } from 'vue';
import {
    GuessFieldTitles,
    GuessState,
    GuessType,
    TotalResultCardFlipDelay,
    ClassicGuessesNeededForHintOne,
    GameModes,
    ClassicGuessesNeededForHintTwo,
    ClassicGuessesNeededForHintThree,
} from '../constants.js';
import { getCurrentClassicPokemonNumber } from '../helpers.js';
import moment from 'moment-timezone';
import {
    setSecretPokemon,
    getDailyGamesWonCount,
    updateUserWithGameWon,
    updateCurrentUserStreakDisplay,
} from '../services/game';
import { playWinnerSound } from '../services/sound';
import {
    clearLocalStorageGameMode,
    setNewDate,
    addColorsToLocalStorage,
    addGuessesToLocalStorage,
} from '../services/localStorage';

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
const hintTwo = reactive({});
const instantIsClassicGameWon = ref(false);

const setDailyGamesWonCount = async () => {
    dailyGamesWon.value = await getDailyGamesWonCount(GameModes.Classic);
};

onBeforeMount(async () => {
    await Promise.all([loadClassicGameData(), setDailyGamesWonCount(), updateYesterdaysPokemon()]);
});

const emojiResults = computed(() => {
    if (!store.isClassicGameWon) return [];

    const lastFiveGuesses = componentStore.guesses.slice(0, 4);

    const results = lastFiveGuesses.map((name) => getGuessResults(name, secretPokemon, 'normal'));

    const emojiResults = results.map((res) => {
        var emojis = '';
        for (const field in res.fields) {
            if (field === 'name') continue; //Skip name field

            var emoji = '';
            if (res.fields[`${field}`].guessState == GuessState.CorrectGuess) emoji = '🟩';
            else if (res.fields[`${field}`].guessState == GuessState.PartlyCorrectGuess) emoji = '🟧';
            else {
                emoji = '🟥';
            }
            emojis = emojis.concat(emoji);
        }
        return emojis;
    });

    return emojiResults;
});

const classicTwitterText = () => {
    const sub1 =
        componentStore.guesses.length === 1 ? 'FIRST TRY 🌟🥳🌠🏆' : `in ${componentStore.guesses.length} tries!🕵️🔎`;

    const header = `I guessed the #${getCurrentClassicPokemonNumber()} classic hidden #Pokedle Pokémon ${sub1}\n`;

    var emojiBody = '';

    emojiResults.value.forEach((res) => {
        emojiBody = emojiBody.concat(res + '\n');
    });

    const moreGuesses =
        componentStore.guesses.length > 4
            ? `＋ ${componentStore.guesses.length - 4} more ${
                  componentStore.guesses.length === 5 ? 'guess.' : 'guesses...'
              }\n\n`
            : '\n';

    const footer = `Play at pokedle.gg 🎮!`;

    return header + emojiBody + moreGuesses + footer;
};

const setHintOne = () => {
    if (
        componentStore.guesses.length < ClassicGuessesNeededForHintOne ||
        (hintTwo.header && componentStore.guesses.length > ClassicGuessesNeededForHintOne)
    ) {
        return;
    }

    const firstFiveGuesses = componentStore.guesses.slice(
        Math.max(componentStore.guesses.length - ClassicGuessesNeededForHintOne, 0)
    );

    const results = firstFiveGuesses
        .map((name) => getGuessResults(name, secretPokemon, 'normal'))
        .map((res) => ({
            type1: res.fields.type1.guessState,
            type2: res.fields.type2.guessState,
            color: res.fields.color.guessState,
            habitat: res.fields.habitat.guessState,
        }));

    const result = {
        type1: true,
        type2: true,
        color: true,
        habitat: true,
    };

    results.forEach((res) => {
        for (const field in res) {
            if (res[field] === GuessState.CorrectGuess) {
                result[field] = false;
            }
        }
    });

    const hint = {
        guessResult: GuessState.CorrectGuess,
    };

    if (result.type1) {
        hint.header = GuessFieldTitles.Type1;
        hint.guessText = secretPokemon.type1;
        hint.type = GuessType.Text;
    } else if (result.type2) {
        hint.header = GuessFieldTitles.Type2;
        hint.guessText = secretPokemon.type2;
        hint.type = GuessType.Text;
    } else if (result.color) {
        hint.header = GuessFieldTitles.Color;
        hint.guessText = secretPokemon.color;
        hint.type = GuessType.Text;
    } else {
        hint.header = GuessFieldTitles.Habitat;
        hint.habitat = secretPokemon.habitat;
        hint.type = GuessType.Habitat;
    }

    Object.assign(hintTwo, hint);
};

const incrementGamesWonCount = async () => {
    const response = await apiService.updateStatsClassicWins(componentStore.guesses.length);
    dailyGamesWon.value = response.data.dailyStats.classicGamesWon;
    dailyFirstTryWins.value = response.data.dailyStats.classicFirstTryWins;
};

const decideGame = (guess) => {
    if (guess === secretPokemon.name) {
        instantIsClassicGameWon.value = true;
        //Wait for all cards to flip
        setTimeout(() => {
            playWinnerSound();
        }, TotalResultCardFlipDelay - 500);
        setTimeout(async () => {
            launchConfetti(colors.at(-1) === 'shiny', componentStore.guesses.length === 1);
            store.setIsClassicGameWon(true);
            console.log('🥳🎉🎊 Congrats! You guessed the secret pokemon: ' + guess);
            incrementGamesWonCount();
            const user = await updateUserWithGameWon(GameModes.Classic, componentStore.guesses.length);
            store.setUser(user);
        }, TotalResultCardFlipDelay);
    } else {
        console.log('❌❌❌ Wrong Guess. The secret pokemon was not ' + guess + ' ❌❌❌');
    }
};

const updateYesterdaysPokemon = async () => {
    yesterdaysPokemon.value = (await apiService.getClassicPreviousSecretPokemon()).data;
};

const submitGuess = (guess) => {
    if (!guess || instantIsClassicGameWon.value) return;

    const { removedName, updatedNames } = removePokemonNameFromArray(guess, componentStore.pokemonNames);

    if (updatedNames.length >= componentStore.pokemonNames.length) return;

    colors.push(getRandomColor());
    componentStore.guesses.unshift(removedName);
    componentStore.pokemonNames = updatedNames;
    addGuessesToLocalStorage(GameModes.Classic, componentStore.guesses);
    addColorsToLocalStorage(GameModes.Classic, colors);
    setHintOne();
    decideGame(removedName);
};

const loadSecretPokemon = () => {
    if (!localStorage.classicSecretPokemon) return;
    Object.assign(secretPokemon, JSON.parse(localStorage.classicSecretPokemon));
};

const loadGuesses = () => {
    const guesses = localStorage.classicGuesses;
    if (guesses) componentStore.guesses = JSON.parse(guesses);
};

const loadColors = () => {
    const loadedColors = localStorage.classicColors;
    if (loadedColors) colors = JSON.parse(loadedColors);
};

const loadIsClassicGameWon = () => {
    const loadedIsClassicGameWon = localStorage.isClassicGameWon;
    if (loadedIsClassicGameWon && loadedIsClassicGameWon === 'true') {
        store.setIsClassicGameWon(true);
    } else {
        store.setIsClassicGameWon(false);
    }
};

const removePokemonsFromGuessPool = () => {
    componentStore.pokemonNames = componentStore.pokemonNames.filter((pokemon) => {
        for (const guessedPokemon of componentStore.guesses) {
            if (guessedPokemon === pokemon) return false;
        }
        return true;
    });
};

const loadClassicGameData = async () => {
    const dayOfLastUpdate = localStorage.dayOfLastUpdate;
    if (!dayOfLastUpdate) setNewDate();

    loadSecretPokemon();
    const currSecretPokemon = await (await apiService.getClassicSecretPokemon()).data;

    if (
        parseInt(dayOfLastUpdate) == moment().date() &&
        secretPokemon &&
        secretPokemon?.name === currSecretPokemon?.name
    ) {
        //Load
        loadColors();
        loadGuesses();
        loadIsClassicGameWon();
        removePokemonsFromGuessPool();
    } else {
        //Fresh game
        clearLocalStorageGameMode(GameModes.Classic);
        store.setIsClassicGameWon(false);
        setSecretPokemon(GameModes.Classic, currSecretPokemon);
        loadSecretPokemon();
        setNewDate();
    }
    setHintOne();
    updateCurrentUserStreakDisplay(GameModes.Classic);
};
</script>

<style scoped>
html.dark {
    color-scheme: dark;
}

.background-white {
    background-image: url('../client/assets/backgrounds/background-white.png');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black {
    background: linear-gradient(rgba(0, 0, 0, 0.25), rgba(0, 0, 0, 0.25)),
        url('../client/assets/backgrounds/background-black.jpg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-white-difficulty-insane {
    background-image: url('../client/assets/backgrounds/background-white-trainer-red.jpeg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black-difficulty-insane {
    background-image: url('../client/assets/backgrounds/background-black-cubone.jpeg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}
</style>
