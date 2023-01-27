<template>
    <v-card class="card">
        <v-card-title>Victory!</v-card-title>
        <v-card-text>You guessed: Baltoy</v-card-text>
        <v-card-text>Time left:{{ getMSRemaining() }}</v-card-text>
    </v-card>
</template>

<script setup>
    import { ref } from 'vue';
    const countdownTimer = ref("Time remaining");

    const getMSRemaining = () => {
       const date = new Date(new Date().setHours(0, 0, 0))
       date.setDate(date.getDate() + 1);
       return date.getTime()-Date.now();
    }

</script>

<style scoped>

.card {
    width: 360px;
    height: 200px;
    background-color: white;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

</style>