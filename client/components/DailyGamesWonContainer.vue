<template>
    <v-card class="games-won-card bg-white border-black"
    variant="outlined">
        {{ dailyGamesWon + " Trainers Already Found Out!" }}
    </v-card>
</template>

<script setup>

const props = defineProps({
    dailyGamesWon: Number
});

</script>

<style scoped>
.games-won-card{
    padding: 5px 8px 3px 8px;
    font-size: small;
    font-family: pkmEmerald;
    border-width: 2px;
}
</style>