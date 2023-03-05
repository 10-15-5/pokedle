<template>
    <div class="card p-[3px] px-2 pt-[4px] font-pkmEmerald text-sm">
        {{ dailyGamesWon + ' Trainers Already Found Out!' }}
    </div>
</template>

<script setup>
const props = defineProps({
    dailyGamesWon: Number,
});
</script>

<style scoped></style>
