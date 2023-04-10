<template>
    <div v-if="isLoaded" class="flex flex-col items-center justify-center gap-y-4 pb-20">
        <GameWinContainer
            v-if="store.isFlavortextGameWon && secretPokemon.name"
            :pokemon="secretPokemon.name"
            :color="colors.at(-1)"
            :twitterText="flavortextTwitterText()"
        />
        <div
            v-else-if="!store.isFlavortextGameWon && secretPokemon.name"
            class="card w-[450px] p-4 text-justify font-pkmEmerald text-xl italic sm:w-[350px] sm:text-base"
        >
            "{{ getFlavortextLanguageFrom(secretPokemon) }}"
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
                        <span class="card items-center justify-center py-1 sm:py-0">{{ text().guessFieldTitles[field.title] }}</span>
                        <ResultSquare
                            :guessResult="field.guessState"
                            :guessText="field.text"
                            :type="field.type"
                            :habitat="field.habitat"
                        />
                    </div>
                </div>
            </template>
            <template #hint2>
                <div class="flex flex-row items-center justify-center gap-6 sm:gap-2">
                    <div v-for="(field, i) in hintTwo" :key="i" :value="field" class="flex flex-col gap-1">
                        <span
                            class="card items-center justify-center py-1 sm:py-0"
                            :class="getTextSizeShortenedTitle(field.title)"
                            >{{ text().guessFieldTitles[field.title] }}</span
                        >
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
                        <span class="card items-center justify-center py-1 sm:py-0">{{ text().guessFieldTitles.shape }}</span>
                        <ResultSquare :pokemon="secretPokemon.name" :type="GuessType.Blackout" />
                    </div></div
            ></template>
        </HintContainer>
        <DailyGamesWonContainer v-if="!componentStore.guesses.length" :dailyGamesWon="store.dailyFlavortextGamesWon" />
        <div v-if="componentStore.guesses.length" class="flex flex-col gap-1">
            <SingleResultHeader />
            <SingeResultContainer
                v-for="(guess, i) in componentStore.guesses"
                :key="guess"
                :value="guess"
                :pokemonName="guess"
                :color="colors[componentStore.guesses.length - 1 - i]"
                :isCorrect="isCorrectGuess(guess, secretPokemon.name)"
            />
        </div>
        <PreviousPokemonCard v-else-if="yesterdaysPokemon.name" :pokemonName="yesterdaysPokemon.names" />
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
import pokemonData from '../../server/data/pokemonData-v6-translations.json';
import { useStore } from '../stores/store.js';
import HintContainer from '../components/hints/HintContainer.vue';
import { isCorrectGuess, getGuessResults } from '../services/guess';
import {
    GameModes,
    GuessType,
    GuessFieldTitles,
    FlavortextGuessesNeededForHintOne,
    FlavortextGuessesNeededForHintTwo,
    FlavortextGuessesNeededForHintThree,
} from '../constants';
import * as game from '../services/game';
import moment from 'moment-timezone';
import { text,getFlavortextLanguageFrom } from '../services/language';

const store = useStore();
const GAME_MODE = GameModes.Flavortext;
const getSortedPokemonNames = () => pokemonData.map((pokemonInfo) => pokemonInfo.name).sort();

const componentStore = reactive({
    pokemonNames: getSortedPokemonNames(),
    guesses: [],
});

const isLoaded = ref(false);

const secretPokemon = reactive({});
const yesterdaysPokemon = ref('');

const hintOne = reactive([]);
const hintTwo = reactive([]);

let colors = [];

onBeforeMount(async () => {
    await Promise.all([game.loadGameData(GAME_MODE,setHints,secretPokemon, componentStore), game.setDailyGamesWonCount(GAME_MODE), updateYesterdaysPokemon()]);
    isLoaded.value=true;
});

const getTextSizeShortenedTitle = (shortenedTitle) => {
    switch (shortenedTitle) {
        case text().guessFieldTitles.evolutionLevel:
            return 'text-[14px] sm:text-[10px]';
        case text().guessFieldTitles.evolutions:
            return 'text-[13px] sm:text-[10px]';
        default:
            return '';
    }
};

const flavortextTwitterText = () => {
    const initialHeader =
        componentStore.guesses.length === 1 ? text().twitterText.flavortext.headerFirstTry : text().twitterText.flavortext.headerXTries;

    
    const headerWithPokemonNumber = initialHeader.replace("§1§", game.getCurrentPokemonNumber(GAME_MODE, moment()));

    const headerWithNumberOfGuesses = headerWithPokemonNumber.replace("§2§", componentStore.guesses.length)

    const footer = text().twitterText.playAt;

    return headerWithNumberOfGuesses + "\n\n" + footer;
};


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

    correctFieldsWithTitles[3].title = text().guessFieldTitles.evolutionLevel;
    correctFieldsWithTitles[4].title =  text().guessFieldTitles.evolutions;
    correctFieldsWithTitles[7].title =  text().guessFieldTitles.generation;

    hintOne.push(correctFieldsWithTitles[1], correctFieldsWithTitles[2], correctFieldsWithTitles[6]); //Type 1, Type 2, Habitat
    hintTwo.push(correctFieldsWithTitles[3], correctFieldsWithTitles[4], correctFieldsWithTitles[7]); //Evolution lvl, isFullyEvolved, Gen

    console.log(hintTwo);
};

const updateYesterdaysPokemon = async () => {
    yesterdaysPokemon.value = await game.getYesterdaysPokemon(GAME_MODE);
};

const submitGuess = async (guess) => {
    const pokemonName = game.submitGuess(GAME_MODE, guess, componentStore, colors);
    if (!pokemonName) return;
    await game.decideGame(GAME_MODE, pokemonName, secretPokemon.name, colors.at(-1), componentStore.guesses.length);
    setHints()
};
</script>
