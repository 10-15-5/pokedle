<template>
    <div id="app">
        <div
            class="min-h-screen w-full"
            :class="{
                'background-white': !(isDark && store.isDifficultyInsane),
                'background-black': isDark && !store.isDifficultyInsane,
                'background-white-difficulty-insane': !isDark && store.isDifficultyInsane,
                'background-black-difficulty-insane': isDark && store.isDifficultyInsane,
            }"
        >
            <ThemeButton class="absolute right-0" />
            <div class="flex flex-col items-center justify-center gap-y-4">
                <div>
                    <img
                        src="/client/assets/logo/pokedle-logo.png"
                        class="title-img mt-5 mb-3 w-[280px] transition-all duration-100 ease-linear hover:scale-110 sm:w-[220px]"
                    />
                    <HeaderContainer />
                </div>
                <SearchField
                    v-if="!store.isGameWon"
                    :pokemonNames="componentStore.pokemonNames"
                    @submit-guess="submitGuess"
                />
                <GameWinContainer
                    v-else
                    :pokemon="componentStore.guesses[0]"
                    :color="colors.at(-1)"
                />
                <HintContainer
                    v-if="!store.isGameWon && componentStore.guesses.length"
                    :text="secretPokemon.flavorText"
                    :numberOfGuesses="componentStore.guesses.length"
                />
                <DailyGamesWonContainer
                    v-if="!componentStore.guesses.length || store.isGameWon"
                    :dailyGamesWon="dailyGamesWon"
                />
                <div
                    v-if="componentStore.guesses.length"
                    class="mb-20 flex flex-col gap-y-2 sm:gap-y-1"
                >
                    <ResultHeader class="mb-n1 sm:!mb-0" />
                    <ResultContainer
                        v-for="(guess, i) in componentStore.guesses"
                        :key="guess"
                        :value="guess"
                        :pokemonName="guess"
                        :isGameWon="instantIsGameWon"
                        :guessResult="getGuessResults(guess, secretPokemon)"
                        :color="colors.at(componentStore.guesses.length - 1 - i)"
                    />
                </div>
                <PreviousPokemonCard :pokemonName="yesterdaysPokemon.name" v-else />
            </div>
        </div>
        <div class="flex items-center justify-center">
            <button class="card p-2 text-xs hover:!bg-green-400" @click="revealPokemon">
                Reveal
            </button>
            <button class="card p-2 text-xs hover:!bg-purple-400" @click="getNewGame">
                Get New Game
            </button>
            <button class="card p-2 text-xs hover:!bg-blue-400" @click="setNewGame">
                Set New Game
            </button>
            <button class="card p-2 text-xs hover:!bg-pink-400" @click="lauchConfetti">
                Confetti
            </button>
        </div>
    </div>
</template>

<script setup>
import ResultContainer from './components/result/ResultContainer.vue';
import ResultHeader from './components/result/ResultHeader.vue';
import GameWinContainer from './components/GameWinContainer.vue';
import HintContainer from './components/hints/HintContainer.vue';
import PreviousPokemonCard from './components/PreviousPokemonCard.vue';
import SearchField from './components/SearchField.vue';
import pokemonData from '../server/data/pokemonData-v5-flavorText.json';
import HeaderContainer from './components/headerIcons/HeaderIconContainer.vue';
import DailyGamesWonContainer from './components/DailyGamesWonContainer.vue';
import ThemeButton from './components/buttons/ThemeButton.vue';
import { reactive, ref, onBeforeMount } from 'vue';
import { getGuessResults } from './services/guess';
import * as apiService from './services/api/apiService.js';
import confetti from 'canvas-confetti';
import { useStore } from './stores/store';
import { useDark } from '@vueuse/core';
import { TotalResultCardFlipDelay } from './constants';

const isDark = useDark();

const getSortedPokemonNames = () => pokemonData.map((pokemonInfo) => pokemonInfo.name).sort();

const componentStore = reactive({
    pokemonNames: reactive(getSortedPokemonNames()),
    guesses: reactive([]),
});

const dailyGamesWon = ref(0);
const dailyFirstTryWins = ref(0);
const store = useStore();
const yesterdaysPokemon = ref('');
//Updated as soon an correct pokemon is guessed, contrary to store.isGameWon Which is only updated after TotalResultCardFlipDelay
const instantIsGameWon = ref(false); 
let colors = [];
const secretPokemon = reactive({});

const getRandomColor = () =>
    store.isShiny ? 'shiny' : Math.random() * 100 < 5 ? 'shiny' : 'normal';

const setDailyGamesWonCount = async () => {
    var date = new Date().toISOString().split('T')[0]; //Get current date in the format YYYY-MM-DD
    const res = await apiService.getDailyStats(date);
    dailyGamesWon.value = res.data.gamesWon;
};

onBeforeMount(async () => {
    const [user] = await Promise.all([getOrCreateUser(), loadGameData(), setDailyGamesWonCount()]);

    console.log(user);

    if (user) {
        store.setUser(user);
    }
});

const getOrCreateUser = async () => {
    const userId = localStorage.getItem('userId');

    console.log(userId);
    if (userId) {
        await apiService.updateUserStreak(userId);
        const response = await apiService.getUser(userId);
        return response.data.user;
    }

    const response = await apiService.createUser();
    const { user } = response.data;
    localStorage.setItem('userId', user._id);
    return user;
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
        instantIsGameWon.value = true;
        //Wait for all cards to flip
        setTimeout(() => {
            lauchConfetti();
            store.setIsGameWon(true);
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
    if (!guess || instantIsGameWon.value) return;

    const { updatedPokemonNames, pokemonName } = removePokemonFromGuessPool(guess);

    if (updatedPokemonNames.length >= componentStore.pokemonNames.length) return;

    colors.push(getRandomColor());
    componentStore.pokemonNames = updatedPokemonNames;
    addGuessesToLocalStorage();
    addColorsToLocalStorage();

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

const loadIsGameWon = () => {
    const loadedIsGameWon = localStorage.getItem('isGameWon');
    if (loadedIsGameWon === 'true') {
        store.setIsGameWon(true);
    } else {
        store.setIsGameWon(false);
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

const setNewDate = () =>
    localStorage.setItem('dayOfLastUpdate', new Date().getUTCDate().toString());

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

const loadIsShiny = () => {
    const isShinyString = localStorage.getItem('isShiny');
    if (isShinyString && isShinyString === 'true') {
        store.setShiny(true);
        return;
    }
    store.setShiny(false);
};

const loadGameData = async () => {
    const dayOfLastUpdate = localStorage.getItem('dayOfLastUpdate');
    if (!dayOfLastUpdate) setNewDate();

    loadSecretPokemon();
    const currSecretPokemon = await (await apiService.getSecretPokemon()).data;

    if (
        parseInt(dayOfLastUpdate) == new Date().getUTCDate() &&
        secretPokemon &&
        secretPokemon?.name === currSecretPokemon?.name
    ) {
        loadColors();
        loadGuesses();
        loadIsGameWon();
        removePokemonsFromGuessPool();
    } else {
        newGame();
        await setSecretPokemon();
    }
    setNewDate();
    loadIsShiny();
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
    localStorage.removeItem('isGameWon');
    colors = [];
    instantIsGameWon.value = false;
    componentStore.guesses.splice(0);
    componentStore.pokemonNames = getSortedPokemonNames();
    setNewDate();
    store.setIsGameWon(false);
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
    var animationEnd = Date.now() + duration;
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
            defaults.colors = ['63C5DA', '48AAD', '52B2BF', '3944BC'];
        }
    }

    const randomInRange = (min, max) => {
        return Math.random() * (max - min) + min;
    };

    const fire = () => {
        var timeLeft = animationEnd - Date.now();

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
    background-image: url('./client/assets/backgrounds/background-white.jpg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black {
    background: linear-gradient(rgba(0, 0, 0, 0.25), rgba(0, 0, 0, 0.25)),
        url('./client/assets/backgrounds/background-black.jpg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-white-difficulty-insane {
    background-image: url('./client/assets/backgrounds/background-white-trainer-red.jpeg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black-difficulty-insane {
    background-image: url('./client/assets/backgrounds/background-black-cubone.jpeg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}


</style>
