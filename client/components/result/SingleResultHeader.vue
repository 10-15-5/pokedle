<template>
    <div class="single-result-container mb-1 items-center justify-center">
        <span
            class="border-b-2 border-light-border text-center font-pkmEmerald text-lg font-bold dark:border-dark-border"
            >{{ getLanguage().guessType.pokemon }}
        </span>
    </div>
</template>

<script setup>
import { getLanguage } from '../../services/language';
</script>
