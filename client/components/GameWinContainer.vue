<template>
    <v-card class="card-container" variant="outlined">
        <v-card-title class="title">VICTORY!</v-card-title>
        <v-card-text class="card-text">You Guessed: {{ pokemon }}</v-card-text>
        <SquareContent :pokemon="removeSpecialCharactersExceptDashFromString(pokemon)" />
        <v-card-text class="card-smaller-text">Next Pokemon Will Appear In: {{
            " " +hoursRemaining + ":" + minRemaining + ":" +
                secRemaining
        }}</v-card-text>
    </v-card>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import SquareContent from './SquareContent.vue';
import { removeSpecialCharactersExceptDashFromString } from '../helpers';

const props = defineProps({
    pokemon: String
})

const secRemaining = ref("")
const minRemaining = ref("")
const hoursRemaining = ref("")
const secondsRemaining = ref(0)


const setSecRemaining = () => {
    const countDownToDate = new Date(new Date().setHours(0, 0, 0));
    countDownToDate.setDate(countDownToDate.getDate() + 1);
    secondsRemaining.value = Math.floor((countDownToDate.getTime() - Date.now()) / 1000);
}

onMounted(() => {
    setSecRemaining();
});

setInterval(() => {
    secondsRemaining.value--;
    const sec = Math.floor(secondsRemaining.value % 60);
    const min = Math.floor((secondsRemaining.value / 60) % 60);
    const hours = Math.floor((secondsRemaining.value / (60 * 60)) % 60);
    secRemaining.value = sec < 10 ? "0" + sec : "" + sec;
    minRemaining.value = min < 10 ? "0" + min : "" + min;
    hoursRemaining.value = hours < 10 ? "0" + hours : "" + hours;
}, 1000)

</script>

<style scoped>
.card-container {
    background-color: white;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 8px;
}

.card-text, .title, .card-smaller-text {
    font-family: pkmEmerald;
    text-transform: capitalize;
}

.card-text {
    font-size: 18px;
}

.card-smaller-text {
    font-size: 14px;
}
.title {
    font-size: 32px;
    color: rgb(34, 153, 0);
}
</style>