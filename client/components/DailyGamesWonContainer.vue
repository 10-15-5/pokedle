<template>
    <v-card class="games-won-card"
    variant="outlined">
        {{ dailyGamesWon + " Trainers Already Found Out!" }}
    </v-card>
</template>

<script setup>
import { onBeforeMount, ref } from 'vue';
import {getDailyStats} from '../services/service.js';
const dailyGamesWon = ref(0);

const updateDailyGamesWonCount = async () => {
    var date = (new Date()).toISOString().split('T')[0];
    const res = await getDailyStats(date);
    dailyGamesWon.value = res.data.gamesWon;
}

onBeforeMount(async () =>{ 
    await updateDailyGamesWonCount()
});

</script>

<style scoped>
.games-won-card{
    padding: 5px 8px 3px 8px;
    background-color: rgb(255, 255, 255);
    color: rgb(0, 0, 0);
    font-size: small;
    font-family: pkmEmerald;
    border-width: 2px;
}
</style>