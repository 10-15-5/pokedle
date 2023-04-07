<template>
    <div v-if="isLoaded" class="flex flex-col items-center justify-center gap-y-4 pb-20">
        <SearchField
            v-if="!store.isClassicGameWon"
            :pokemonNames="componentStore.pokemonNames"
            @submit-guess="submitGuess"
        />
        <GameWinContainer
            v-if="store.isClassicGameWon && secretPokemon.name"
            :pokemon="secretPokemon.name"
            :twitterText="classicTwitterText()"
            :color="colors.at(-1)"
        />
        <RouterLink v-if="store.isClassicGameWon" to="/silhouette">
            <GameModeButton :gameMode="GameModes.Silhouette" :isAttention="true" />
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
                    <span class="card w-full justify-center py-1 sm:py-0">{{ text().guessFieldTitles[hintOne.header]
                     }}</span>
                    <ResultSquare
                        :color="hintOne.color"
                        :guessResult="hintOne.guessResult"
                        :guessText="hintOne.guessText"
                        :type="hintOne.type"
                        :habitat="hintOne.habitat"
                    />
                </div>
            </template>
            <template #hint2>
                {{ getFlavortextLanguageFrom(secretPokemon) }}
            </template>
            <template #hint3>
                <div class="flex flex-col items-center gap-2 sm:gap-1">
                    <span class="card w-full justify-center py-1 sm:py-0">{{ text().guessFieldTitles.shape
                     }}</span>
                    <ResultSquare :pokemon="secretPokemon.name" :type="GuessType.Blackout" />
                </div>
            </template>
        </HintContainer>
        <DailyGamesWonContainer v-if="!componentStore.guesses.length" :dailyGamesWon="store.dailyClassicGamesWon" />
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
import pokemonData from '../../server/data/pokemonData-v6-translations.json';
import DailyGamesWonContainer from '../components/infoCards/DailyGamesWonContainer.vue';
import { getGuessResults } from '../services/guess';
import { useStore } from '../stores/store.js';
import ResultSquare from '../components/result/ResultSquare.vue';
import { reactive, ref, onBeforeMount, computed,onMounted } from 'vue';
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
import * as game from '../services/game';
import * as localStorageService from '../services/localStorage';
import moment from 'moment-timezone';
import { text,getFlavortextLanguageFrom } from '../services/language';

const store = useStore();
const GAME_MODE = GameModes.Classic;
const getSortedPokemonNames = () => pokemonData.map((pokemonInfo) => pokemonInfo.name).sort();

const componentStore = reactive({
    pokemonNames: getSortedPokemonNames(),
    guesses: [],
});
const isLoaded = ref(false);

const secretPokemon = reactive({});
const yesterdaysPokemon = ref('');

const hintOne = reactive({});

let colors = [];

//Updated as soon an correct pokemon is guessed, contrary to store.isGameWon Which is only updated after TotalResultCardFlipDelay
const instantIsClassicGameWon = ref(false);

onBeforeMount(async () => {
    await Promise.all([game.loadGameData(GAME_MODE, setHintOne, secretPokemon, componentStore), game.setDailyGamesWonCount(GAME_MODE), updateYesterdaysPokemon()]);
    colors = localStorageService.getColorsFromLocalStorage(GAME_MODE);
    isLoaded.value=true;
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
    const initialHeader =
        componentStore.guesses.length === 1 ? text().twitterText.classic.headerFirstTry : text().twitterText.classic.headerXTries;

    
    const headerWithPokemonNumber = initialHeader.replace("§1§", game.getCurrentPokemonNumber(GAME_MODE, moment()));

    const headerWithNumberOfGuesses = headerWithPokemonNumber.replace("§2§", componentStore.guesses.length)

    var emojiBody = '';

    emojiResults.value.forEach((res) => {
        emojiBody = emojiBody.concat(res + '\n');
    });

    let moreGuesses = "";
    if(componentStore.guesses.length>4){
        moreGuesses = componentStore.guesses.length === 5 ? text().twitterText.classic.plusOneMoreGuess : text().twitterText.classic.plusXMoreGuesses;
        moreGuesses = moreGuesses + "\n";
    }

    const moreGuessesWithNumber = moreGuesses.replace("§1§", componentStore.guesses.length - 4);

    const footer = text().twitterText.playAt;

    return headerWithNumberOfGuesses + "\n\n" + emojiBody + "\n"  + moreGuessesWithNumber + footer;
};

const setHintOne = () => {
    if (
        componentStore.guesses.length < ClassicGuessesNeededForHintOne ||
        (hintOne.header && componentStore.guesses.length > ClassicGuessesNeededForHintOne)
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

    Object.assign(hintOne, hint);
    console.log(hintOne);
    console.log(text().guessFieldTitles);
};

const updateYesterdaysPokemon = async () => {
    yesterdaysPokemon.value = await game.getYesterdaysPokemon(GAME_MODE);
};

const submitGuess = async (guess) => {
    if (instantIsClassicGameWon.value) return;

    const pokemonName = game.submitGuess(GAME_MODE, guess, componentStore, colors);
    if (!pokemonName) return;
    await game.decideGame(
        GAME_MODE,
        pokemonName,
        secretPokemon.name,
        colors.at(-1),
        componentStore.guesses.length,
        TotalResultCardFlipDelay - 500,
        TotalResultCardFlipDelay,
        () => {
            instantIsClassicGameWon.value = true;
        }
    );
    setHintOne();
};
</script>
