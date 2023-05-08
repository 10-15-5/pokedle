<template>
    <div>
        <div id="nav" class="flex flex-col items-center gap-4">
            <RouterLink to="/classic"> <GameModeButton :gameMode="GameModes.Classic" /> </RouterLink>
            <RouterLink to="/silhouette"> <GameModeButton :gameMode="GameModes.Silhouette" /> </RouterLink>
            <RouterLink to="/flavortext"> <GameModeButton :gameMode="GameModes.Flavortext" /> </RouterLink>
        </div>
    </div>
</template>

<script setup>
import { onMounted } from 'vue';
import GameModeButton from '../components/buttons/GameModeButton.vue';
import { GameModes } from '../constants';
import { updateCurrentUserStreakDisplay } from '../services/game';

onMounted(() => {
    updateCurrentUserStreakDisplay(GameModes.Classic);
});
</script>
