<template>
    <div class="card mt-[-8px] p-1 font-pkmEmerald text-xs capitalize sm:text-[10px]">
        <p>
            Yesterdays Pokemon Was
            <b class="text-light-emerald dark:!text-dark-emerald">{{ pokemonName }}</b>
        </p>
    </div>
</template>

<script setup>

const props = defineProps({
    pokemonName: String,
});
</script>
