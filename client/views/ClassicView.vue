<template>
    <div class="flex flex-col items-center justify-center gap-y-4 pt-4">
        <SearchField
            v-if="!store.isClassicGameWon"
            :pokemonNames="componentStore.pokemonNames"
            @submit-guess="submitGuess"
        />
        <GameWinContainer
            v-if="store.isClassicGameWon"
            :pokemon="componentStore.guesses[0]"
            :twitterText="classicTwitterText"
            :color="colors.at(-1)"
        />
        <HintContainer
            v-if="!store.isClassicGameWon && componentStore.guesses.length"
            :numberOfGuesses="componentStore.guesses.length"
            :stylingHintOne="'flex justify-center'"
            :stylingHintTwo="''"
            :stylingHintThree="'flex justify-center'"
        >
            <template #hint1>
                <div class="flex flex-col items-center gap-2 sm:gap-1">
                    <span class="card w-full justify-center py-1 sm:py-0">{{
                        hintTwo.header
                    }}</span>
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
                    <ResultSquare
                        :pokemon="secretPokemon.name"
                        :type="guessType.Blackout"
                        :color="'normal'"
                    />
                </div>
            </template>
        </HintContainer>
        <DailyGamesWonContainer
            v-if="!componentStore.guesses.length || store.isClassicGameWon"
            :dailyGamesWon="dailyGamesWon"
        />
        <div v-if="componentStore.guesses.length" class="mb-20 flex flex-col gap-y-2 sm:gap-y-1">
            <ResultHeader class="mb-n1 sm:!mb-0" />
            <ResultContainer
                v-for="(guess, i) in componentStore.guesses"
                :key="guess"
                :value="guess"
                :pokemonName="guess"
                :guessResult="
                    getGuessResults(
                        guess,
                        secretPokemon,
                        colors[componentStore.guesses.length - 1 - i]
                    )
                "
            />
        </div>
        <PreviousPokemonCard v-else :pokemonName="yesterdaysPokemon.name" />
    </div>
    <div v-if="isDevelopment" class="flex items-center justify-center">
        <button class="card p-2 text-xs hover:!bg-green-400" @click="revealPokemon">Reveal</button>
        <button class="card p-2 text-xs hover:!bg-purple-400" @click="getNewGame">
            Get New Game
        </button>
        <button class="card p-2 text-xs hover:!bg-blue-400" @click="setNewGame">
            Set New Game
        </button>
        <button class="card p-2 text-xs hover:!bg-pink-400" @click="lauchConfetti">Confetti</button>
        <button class="card p-2 text-xs hover:!bg-pink-400" @click="playWinnerSound">
            WIN SOUND
        </button>
    </div>
</template>

<script setup>
import ResultContainer from '../components/result/ResultContainer.vue';
import ResultHeader from '../components/result/ResultHeader.vue';
import GameWinContainer from '../components/GameWinContainer.vue';
import HintContainer from '../components/hints/HintContainer.vue';
import PreviousPokemonCard from '../components/PreviousPokemonCard.vue';
import SearchField from '../components/SearchField.vue';
import pokemonData from '../../server/data/pokemonData-v5-flavorText.json';
import HeaderContainer from '../components/headerIcons/HeaderIconContainer.vue';
import DailyGamesWonContainer from '../components/DailyGamesWonContainer.vue';
import ThemeButton from '../components/buttons/ThemeButton.vue';
import { getGuessResults } from '../services/guess';
import * as apiService from '../services/api/apiService.js';
import confetti from 'canvas-confetti';
import { useStore } from '../stores/store.js';
import ResultSquare from '../components/result/ResultSquare.vue';
import { reactive, ref, onBeforeMount, computed } from 'vue';
import {
    guessFieldTitles,
    guessState,
    guessType,
    TotalResultCardFlipDelay,
    ClassicGuessesNeededForHintOne,
} from '../constants.js';
import { getCurrentClassicPokemonNumber } from '../helpers.js';
import moment from 'moment-timezone';


const store = useStore();

const isDevelopment = computed(() => ENVIRONMENT === 'development');

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

const getRandomColor = () =>
    store.isShiny ? 'shiny' : Math.random() * 100 < 5 ? 'shiny' : 'normal';

const setDailyGamesWonCount = async () => {
    const date = moment().format('YYYY-MM-DD');
    const res = await apiService.getDailyStats(date);
    dailyGamesWon.value = res.data.gamesWon;
};

onBeforeMount(async () => {
    store.loadTheme();

    const [user] = await Promise.all([loadGameData(), setDailyGamesWonCount()]);

    console.log('Loaded at: ' + moment().toString());
    console.log('ENVIRONMENT: ' + ENVIRONMENT);

    if (user) {
        store.setUser(user);
    }
});

const playWinnerSound = () => {
    const audioEmeraldWinMusic = new Audio('/client/assets/audio/emerald-winner.mp3');
    const audioFireworks = new Audio('/client/assets/audio/fireworks.mp3');

    audioEmeraldWinMusic.volume = 0.2
    audioFireworks.volume = 0.15;

    audioEmeraldWinMusic.play();
    audioFireworks.play();
};

const removePokemonFromGuessPool = (guess) => {
    let guessRemovedFromList = false;
    let pokemonName = '';

    const updatedPokemonNames = componentStore.pokemonNames.filter((e) => {
        if (!guessRemovedFromList && e.startsWith(guess.toLowerCase())) {
            componentStore.guesses.unshift(e);
            pokemonName = e;
            guessRemovedFromList = true;
        } else return true;
    });

    return {
        pokemonName,
        updatedPokemonNames,
    };
};

const emojiResults = computed(() => {
    if (!store.isClassicGameWon) return [];

    const lastFiveGuesses = componentStore.guesses.slice(0, 4);

    const results = lastFiveGuesses.map((name) => getGuessResults(name, secretPokemon, 'normal'));

    const emojiResults = results.map((res) => {
        var emojis = '';
        for (const field in res.fields) {
            if (field === 'name') continue; //Skip name field

            var emoji = '';
            if (res.fields[`${field}`].guessState == guessState.CorrectGuess) emoji = '🟩';
            else if (res.fields[`${field}`].guessState == guessState.PartlyCorrectGuess)
                emoji = '🟧';
            else {
                emoji = '🟥';
            }
            emojis = emojis.concat(emoji);
        }
        return emojis;
    });


    return emojiResults;
});

const classicTwitterText = computed(() => {
    const sub1 =
        componentStore.guesses.length === 1
            ? 'FIRST TRY 🌟🥳🌠🏆'
            : `in ${componentStore.guesses.length} tries!🕵️🔎`;

    const header = `I guessed the #${getCurrentClassicPokemonNumber()} hidden #Pokedle Pokémon ${sub1}\n`;

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
});

const setHintTwo = () => {
    //TODO: make hint numbers constants
    if (
        componentStore.guesses.length < ClassicGuessesNeededForHintOne ||
        (hintTwo.header && componentStore.guesses.length > ClassicGuessesNeededForHintOne)
    )
        return;

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
            if (res[field] === guessState.CorrectGuess) {
                result[field] = false;
            }
        }
    });

    const hint = {
        guessResult: guessState.CorrectGuess,
    };

    if (result.type1) {
        hint.header = guessFieldTitles.Type1;
        hint.guessText = secretPokemon.type1;
        hint.type = guessType.Text;
    } else if (result.type2) {
        hint.header = guessFieldTitles.Type2;
        hint.guessText = secretPokemon.type2;
        hint.type = guessType.Text;
    } else if (result.color) {
        hint.header = guessFieldTitles.Color;
        hint.guessText = secretPokemon.color;
        hint.type = guessType.Text;
    } else {
        hint.header = guessFieldTitles.Habitat;
        hint.habitat = secretPokemon.habitat;
        hint.type = guessType.Habitat;
    }

    Object.assign(hintTwo, hint);
};

const incrementGamesWonCount = async () => {
    const response = await apiService.updateDailyGamesWonCount(componentStore.guesses.length);
    dailyGamesWon.value = response.data.dailyStats.gamesWon;
    dailyFirstTryWins.value = response.data.dailyStats.firstTryWins;
};

const updateUserWithGameWon = async () => {
    const userId = localStorage.getItem('userId');
    if (userId) {
        const response = await apiService.updateUserWithGameWon(
            userId,
            componentStore.guesses.length
        );
        store.setUser(response.data.user);
    }
};

const decideGame = (guess) => {
    if (guess === secretPokemon.name) {
        instantIsClassicGameWon.value = true;
        //Wait for all cards to flip
        setTimeout(() => {
            playWinnerSound();
        }, TotalResultCardFlipDelay - 500);
        setTimeout(() => {
            lauchConfetti();
            store.setIsClassicGameWon(true);
            console.log('🥳🎉🎊 Congrats! You guessed the secret pokemon: ' + guess);
            incrementGamesWonCount();
            updateUserWithGameWon();
        }, TotalResultCardFlipDelay);
    } else {
        console.log('❌❌❌ Wrong Guess. The secret pokemon was not ' + guess + ' ❌❌❌');
    }
};

const updateYesterdaysPokemon = async () => {
    yesterdaysPokemon.value = (await apiService.getPreviousSecretPokemon()).data;
};

const addGuessesToLocalStorage = () => {
    localStorage.setItem('guesses', JSON.stringify(componentStore.guesses));
};

const addColorsToLocalStorage = () => {
    localStorage.setItem('colors', JSON.stringify(colors));
};

const submitGuess = (guess) => {
    if (!guess || instantIsClassicGameWon.value) return;

    const { updatedPokemonNames, pokemonName } = removePokemonFromGuessPool(guess);

    if (updatedPokemonNames.length >= componentStore.pokemonNames.length) return;

    colors.push(getRandomColor());
    componentStore.pokemonNames = updatedPokemonNames;
    addGuessesToLocalStorage();
    addColorsToLocalStorage();
    setHintTwo();

    decideGame(pokemonName);
};

const loadSecretPokemon = () => {
    const loadedSecretPokemon = localStorage.getItem('secretPokemon');
    Object.assign(secretPokemon, JSON.parse(loadedSecretPokemon));
};

const loadGuesses = () => {
    const guesses = localStorage.getItem('guesses');
    if (guesses) componentStore.guesses = JSON.parse(guesses);
};

const loadColors = () => {
    const loadedColors = localStorage.getItem('colors');
    if (loadedColors) colors = JSON.parse(loadedColors);
};

const loadIsClassicGameWon = () => {
    const loadedIsClassicGameWon = localStorage.getItem('isClassicGameWon');
    if (loadedIsClassicGameWon === 'true') {
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

const setNewDate = () => localStorage.setItem('dayOfLastUpdate', moment().date());

const setNewSecretPokemon = async () => {
    const response = await apiService.newSecretPokemon();
    Object.assign(secretPokemon, response.data);
    localStorage.setItem('secretPokemon', JSON.stringify(secretPokemon));
};

const setSecretPokemon = async () => {
    const response = await apiService.getSecretPokemon();
    Object.assign(secretPokemon, response.data);
    localStorage.setItem('secretPokemon', JSON.stringify(secretPokemon));
};

const loadGameData = async () => {
    const dayOfLastUpdate = localStorage.getItem('dayOfLastUpdate');
    if (!dayOfLastUpdate) setNewDate();

    loadSecretPokemon();
    const currSecretPokemon = await (await apiService.getSecretPokemon()).data;

    if (
        parseInt(dayOfLastUpdate) == moment().date() &&
        secretPokemon &&
        secretPokemon?.name === currSecretPokemon?.name
    ) {
        loadColors();
        loadGuesses();
        loadIsClassicGameWon();
        removePokemonsFromGuessPool();
    } else {
        newGame();
        await setSecretPokemon();
    }
    setNewDate();
    setHintTwo();
    updateYesterdaysPokemon();
};

const revealPokemon = async () => {
    console.log(localStorage.getItem('secretPokemon'));
    const backendResponse = await apiService.getSecretPokemon();
    console.log(backendResponse.data);
};

const newGame = async () => {
    localStorage.removeItem('secretPokemon');
    localStorage.removeItem('guesses');
    localStorage.removeItem('colors');
    localStorage.removeItem('isClassicGameWon');
    colors = [];
    instantIsClassicGameWon.value = false;
    componentStore.guesses.splice(0);
    componentStore.pokemonNames = getSortedPokemonNames();
    setNewDate();
    store.setIsClassicGameWon(false);
};

const getNewGame = async () => {
    await newGame();
    await setSecretPokemon();
    await updateYesterdaysPokemon();
};

// Force update pokemon in DB
const setNewGame = async () => {
    await newGame();
    await setNewSecretPokemon();
    await updateYesterdaysPokemon();
};

const lauchConfetti = () => {
    var duration = 3 * 1000;
    var animationEnd = moment() + duration;
    var particleCount = 200;
    var defaults = {
        startVelocity: 25,
        spread: 360,
        ticks: 300,
        zIndex: 0,
        scalar: 1.4,
    };

    if (colors.at(-1) === 'shiny' || componentStore.guesses.length === 1) {
        defaults.shapes = ['star'];
        defaults.scalar = 1.1;
        particleCount = 90;
        if (colors.at(-1) === 'shiny' && componentStore.guesses.length === 1) {
            defaults.colors = ['9EFD38', '32CD32', 'A8E4A0', '98FB98', '7CFC00'];
        } else if (colors.at(-1) === 'shiny') {
            defaults.colors = ['FFE400', 'FFBD00', 'E89400', 'FFCA6C', 'FDFFB8'];
        } else if (componentStore.guesses.length === 1) {
            defaults.colors = ['00E5EE', '00FFFF', 'E0FFFF', '98F5FF'];
        }
    }

    const randomInRange = (min, max) => {
        return Math.random() * (max - min) + min;
    };

    const fire = () => {
        var timeLeft = animationEnd - moment();

        if (timeLeft <= 0) {
            return clearInterval(lauchConfetti);
        }

        var currParticleCount = particleCount * (timeLeft / duration);
        // since particles fall down, start a bit higher than random
        confetti(
            Object.assign({}, defaults, {
                currParticleCount,
                origin: {
                    x: randomInRange(0.1, 0.3),
                    y: Math.random() - 0.4,
                },
            })
        );
        confetti(
            Object.assign({}, defaults, {
                currParticleCount,
                origin: {
                    x: randomInRange(0.7, 0.9),
                    y: Math.random() - 0.4,
                },
            })
        );
    };

    fire();

    setInterval(() => {
        fire();
    }, 250);
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
