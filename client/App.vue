<template>
    <v-app id="app">
        <v-main
            :class="{
                'background-white': !isDark,
                'background-black': isDark,
            }"
        >
            <div class="flex flex-col items-center justify-center gap-y-4">
                <div>
                    <v-img
                        src="/client/assets/pokedle-logo.png"
                        class="title-img mt-5 mb-3 w-[280px] transition-all duration-100 ease-linear hover:scale-110"
                    />
                    <HeaderContainer />
                </div>
                <SearchField
                    v-if="!isGameWon"
                    :pokemonNames="state.pokemonNames"
                    @submit-guess="submitGuess"
                />
                <GameWinContainer v-else :pokemon="state.guesses[0]" :color="colors.at(-1)" />
                <DailyGamesWonContainer :dailyGamesWon="dailyGamesWon" />
                <div v-if="state.guesses.length" class="mb-20 flex flex-col gap-y-2">
                    <SquareContentHeader class="mb-n1" />
                    <div v-for="(guess, i) in state.guesses" :key="guess" :value="guess">
                        <SquareContainer
                            :pokemonName="guess"
                            :guessResult="getGuessResults(guess, secretPokemon)"
                            :color="colors.at(state.guesses.length - 1 - i)"
                        />
                    </div>
                </div>
            </div>
        </v-main>
        <div class="flex items-center justify-center">
            <v-btn class="text-xs" @click="revealPokemon">Reveal</v-btn>
            <v-btn class="text-xs" @click="newGame">New Game</v-btn>
            <v-btn class="text-xs" @click="lauchConfetti">Confetti</v-btn>
            <v-btn class="text-xs" @click="toggleDark()">Theme</v-btn>
            <p>IsDark: {{ isDark }}</p>
        </div>
    </v-app>
</template>

<script setup>
import SquareContainer from './components/SquareContainer.vue';
import SquareContentHeader from './components/SquareContentHeader.vue';
import GameWinContainer from './components/GameWinContainer.vue';
import SearchField from './components/SearchField.vue';
import pokemonData from '../server/data/pokemonData-v4.json';
import HeaderContainer from './components/HeaderContainer.vue';
import DailyGamesWonContainer from './components/DailyGamesWonContainer.vue';
import { onMounted, reactive, ref, onBeforeMount } from 'vue';
import { getGuessResults } from './services/guess';
import * as service from './services/apiService.js';
import * as helpers from './helpers.js';
import confetti from 'canvas-confetti';
import { useStore } from './stores/store';
import { useDark, useToggle } from '@vueuse/core';

const isDark = useDark();
const toggleDark = useToggle(isDark);

const state = reactive({
    pokemonNames: pokemonData.map((pokemonInfo) => pokemonInfo.name).sort(),
    guesses: [],
});
const isGameWon = ref(false);
const isSearchFieldDisabled = ref(false);
const dailyGamesWon = ref(0);
const dailyFirstTryWins = ref(0);
const store = useStore();

let colors = [];
let secretPokemon;

const getRandomColor = () => (Math.random() * 100 < 5 ? 'shiny' : 'normal');

const setDailyGamesWonCount = async () => {
    var date = new Date().toISOString().split('T')[0]; //Get current date in the format YYYY-MM-DD
    const res = await service.getDailyStats(date);
    dailyGamesWon.value = res.data.gamesWon;
};

onBeforeMount(async () => {
    await setDailyGamesWonCount();
});

const removePokemonFromGuessPool = (guess) => {
    let guessRemovedFromList = false;
    let pokemonName = '';
    const updatedPokemonNames = state.pokemonNames.filter((e) => {
        if (!guessRemovedFromList && e.startsWith(guess.toLowerCase())) {
            state.guesses.unshift(e);
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
    const response = await service.updateDailyGamesWonCount(state.guesses.length);
    dailyGamesWon.value = response.data.dailyStats.gamesWon;
    dailyFirstTryWins.value = response.data.dailyStats.firstTryWins;
};

const updateUserWithGameWon = async () => {
    const { userId } = helpers.getCookie(document);
    const response = await service.updateUserWithGameWon(userId, state.guesses.length);
    store.setUser(response.data.user);
};

const decideGame = (guess) => {
    if (guess === secretPokemon.name) {
        isSearchFieldDisabled.value = true;

        //Wait for all cards to flip
        setTimeout(() => {
            lauchConfetti();
            isGameWon.value = true;
            localStorage.setItem('isGameWon', 'true');
            console.log('🥳🎉🎊 Congrats! You guessed the secret pokemon: ' + guess);
            incrementGamesWonCount();
            updateUserWithGameWon();
        }, 2750);
    } else {
        console.log('❌❌❌ Wrong Guess. The secret pokemon was not ' + guess + ' ❌❌❌');
    }
};

const addGuessesToLocalStorage = () => {
    localStorage.setItem('guesses', JSON.stringify(state.guesses));
};

const addColorsToLocalStorage = () => {
    localStorage.setItem('colors', JSON.stringify(colors));
};

const submitGuess = (guess) => {
    if (!guess || isSearchFieldDisabled.value) return;

    const { updatedPokemonNames, pokemonName } = removePokemonFromGuessPool(guess);

    if (updatedPokemonNames.length >= state.pokemonNames.length) return;

    colors.push(getRandomColor());
    state.pokemonNames = updatedPokemonNames;
    addGuessesToLocalStorage();
    addColorsToLocalStorage();

    decideGame(pokemonName);
};

const loadSecretPokemon = () => {
    const loadedSecretPokemon = localStorage.getItem('secretPokemon');
    secretPokemon = JSON.parse(loadedSecretPokemon);
};

const loadGuesses = () => {
    const guesses = localStorage.getItem('guesses');
    if (guesses) state.guesses = JSON.parse(guesses);
};

const loadColors = () => {
    const loadedColors = localStorage.getItem('colors');
    if (loadedColors) colors = JSON.parse(loadedColors);
};

const loadIsGameWon = () => {
    const loadedIsGameWon = localStorage.getItem('isGameWon');
    if (loadedIsGameWon === 'true') {
        isGameWon.value = true;
    } else {
        localStorage.setItem('isGameWon', 'false');
        isGameWon.value = false;
    }
};

const removePokemonsFromGuessPool = () => {
    state.pokemonNames = state.pokemonNames.filter((pokemon) => {
        for (const guessedPokemon of state.guesses) {
            if (guessedPokemon === pokemon) return false;
        }
        return true;
    });
};

const setNewDate = () =>
    localStorage.setItem('dayOfLastUpdate', new Date().getUTCDate().toString());

const setNewSecretPokemon = async () => {
    //TODO: currently we set new secret pokemon every day at 00:00.
    // Users will see a random pokemon each, since we rely on this instead of job/CDN
    const response = await service.newSecretPokemon();
    secretPokemon = response.data;
    localStorage.setItem('secretPokemon', JSON.stringify(secretPokemon));
};

const setSecretPokemon = async () => {
    const response = await service.getSecretPokemon();
    secretPokemon = response.data;
    localStorage.setItem('secretPokemon', JSON.stringify(secretPokemon));
};

const loadGameData = async () => {
    const dayOfLastUpdate = localStorage.getItem('dayOfLastUpdate');
    if (!dayOfLastUpdate) setNewDate();

    if (parseInt(dayOfLastUpdate) == new Date().getUTCDate()) {
        loadSecretPokemon();

        const currSecretPokemon = await (await service.getSecretPokemon()).data;

        if (secretPokemon && secretPokemon?.name === currSecretPokemon?.name) {
            loadColors();
            loadGuesses();
            loadIsGameWon();
            removePokemonsFromGuessPool();
        } else {
            localStorage.clear();
            setNewDate();
            await setSecretPokemon();
        }
    } else {
        localStorage.clear();
        setNewDate();
        //TODO: this should be replaced with CDN
        await setNewSecretPokemon();
    }
};

const getOrCreateUser = async () => {
    const { userId } = helpers.getCookie(document);

    if (!userId) {
        await service.createUser();
        return undefined;
    }

    const response = await service.getUser(userId);
    return response.data.user;
};

onMounted(async () => {
    const [_, user] = await Promise.all([loadGameData(), getOrCreateUser()]);

    console.log(user);

    if (user) {
        store.setUser(user);
        console.log(store.user);
    }
});

const revealPokemon = async () => {
    console.log(localStorage.getItem('secretPokemon'));
    const backendResponse = await service.getSecretPokemon();
    console.log(backendResponse.data);
};

const newGame = async () => {
    localStorage.clear();
    location.reload();
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

    if (colors.at(-1) === 'shiny') {
        (defaults.shapes = ['star']),
            (defaults.colors = ['FFE400', 'FFBD00', 'E89400', 'FFCA6C', 'FDFFB8']);
        defaults.scalar = 1.1;
        particleCount = 90;
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
    background-image: url('./client/assets/background-white.jpg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black {
    background: linear-gradient(rgba(0, 0, 0, 0.25), rgba(0, 0, 0, 0.25)), url('./client/assets/background-black.jpg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}
</style>
