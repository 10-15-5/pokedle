<template>
    <v-card class="card">
        <v-card-title class="card-title">Victory!</v-card-title>
        <v-card-text class="card-text">You guessed: {{ pokemon }}</v-card-text>
        <SquareContent :pokemon="removeSpecialCharactersExceptDashFromString(pokemon)" />
        <v-card-text class="card-text">Next pokemon will appear in: {{
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
.card {
    width: 280px;
    height: 200px;
    background-color: white;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 12px;
}

.card-text,
.card-title {
    font-family: pkmEmerald;
}

.card-text {}

.card-title {
    color: rgb(34, 153, 0);
}
</style>