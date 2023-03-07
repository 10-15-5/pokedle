<template>
    <div class="card p-[3px] px-2 pt-[4px] font-pkmEmerald text-sm">
      <p> <b class=" text-light-orange dark:text-dark-orange"> {{ dailyGamesWon }}</b>   Trainers Already Found Out!</p>
    </div>
</template>

<script setup>
const props = defineProps({
    dailyGamesWon: Number,
});
</script>

<style scoped></style>
