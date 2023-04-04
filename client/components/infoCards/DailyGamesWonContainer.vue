<template>
    <div class="card p-[3px] px-2 pt-1 font-pkmEmerald text-sm sm:pb-0 sm:text-[12px]">
        <p>
            <b class="text-light-orange dark:text-dark-orange"> {{ dailyGamesWon }}</b>{{ getLanguage().dailyGamesWon.trainersAlreadyFoundOut }}
        </p>
    </div>
</template>

<script setup>
import { getLanguage } from '../../services/language';
const props = defineProps({
    dailyGamesWon: Number,
});
</script>

<style scoped></style>
