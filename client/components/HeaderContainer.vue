<template>
    <div class="header-container">
        <div class="header-icon">
            <!-- <v-img :src="'./client/assets/fire-emoji.png'"></v-img> -->
            <p class="fire-emoji">🔥</p>
            <p class="streak-count mt-2">{{ store.user.currentStreak }}</p>
        </div>
        <div class="header-icon"
             @click="dialogSettings = true">
            <v-icon icon="mdi-cog"
                    size="small" />
        </div>
        <div class="header-icon"
             @click="dialogStatistics = true">
            <v-icon icon="mdi-poll"
                    size="small" />
        </div>
        <div class="header-icon"
             @click="dialogInfo = true">
            <v-icon icon="mdi-information-outline"
                    size="small">
            </v-icon>
        </div>
        <!-- TODO: this should be a view -->
        <v-dialog class="dialog-content"
                  v-model="dialogSettings">
            <v-card class="card-content" variant="outlined">
                <v-card-title class="card-title">SETTINGS</v-card-title>
                <v-card-text class="card-text"> Coming Soon....</v-card-text>
                <v-card-actions>
                    <v-btn class="dialog-btn"
                           block
                           @click="dialogSettings = false">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <!-- TODO: this should be a view -->
        <v-dialog class="dialog-content"
                  v-model="dialogStatistics">
            <v-card class="card-content" variant="outlined">
                <v-card-title class="card-title">STATISTICS</v-card-title>
                <v-card-text class="card-text"> Coming Soon....</v-card-text>
                <v-card-actions>
                    <v-btn class="dialog-btn"
                           block
                           @click="dialogStatistics = false">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <!-- TODO: this should be a view -->
        <v-dialog class="dialog-content"
                  v-model="dialogInfo">
            <v-card class="card-content" variant="outlined">
                <v-card-title class="card-title">About</v-card-title>
                <v-card-text class="card-text">Welcome to the world of POKEMéN! <br> Every day a wild secret POKéMON
                    will be hiding and its your job as a POKéMON TRAINER to catch (guess) it! </v-card-text>
                <v-card-title class="card-title">Credit</v-card-title>
                <v-card-text class="card-text"> Pokédlé is heavily inspired by loldle.net and wordle. Original
                    artist/creator of images, please DM for removal or credit. </v-card-text>
                <v-card-title class="card-title">Contact</v-card-title>
                <v-card-text class="card-text"> If you have any questions/feedback/bug reports feel free to contact me
                    on the email below. <br> > <a
                       href="mailto: pokedle.contact@gmail.com">pokedle.contact@gmail.com</a></v-card-text>
                <v-card-actions>
                    <v-btn class="dialog-btn"
                           block
                           @click="dialogInfo = false">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script setup>
import { useStore } from '../stores/store';
import { ref } from 'vue';
const store = useStore();
let dialogInfo = ref(false);
let dialogStatistics = ref(false);
let dialogSettings = ref(false);

</script>

<style scoped>
.header-icon {
    width: 32px;
    height: 32px;
    margin: 4px;
    background-color: rgb(255, 255, 255);
    border-style: solid;
    border-width: 2px;
    border-radius: 4px;
    justify-content: center;
    align-items: center;
    display: flex;
    transition: all 120ms ease-in-out;
}

.fire-emoji {
    font-size: 30px;
}

.streak-count {
    position: absolute;
    font-weight: bold;
    margin-left: 1px;
    text-shadow: -0.5px 0.5px 1px #fffcfc,
				  0.5px 0.5px 1px #ffffff,
				  0.5px -0.5px 0 #ffffff,
				  -0.5px -0.5px 0 #ffffff;
}

.card-text,
.card-title,
.dialog-btn, .streak-count {
    font-family: pkmEmerald;
}
.card-title {
    font-size: larger;
    font-weight: bold;
}

.dialog-content {
    width: 500px;
}

.card-content {
    padding: 8px;
    border-style: solid;
    border-width: 2px;
    background-color: white;
}

.dialog-btn {
    font-size: large;
    padding: 4px;
    font-weight: bold;
    color: red;
    margin-top: 8px;
}

.header-icon:hover {
    transform: scale(1.1);
    cursor: pointer;
}

.header-container {
    display: flex;
    flex-direction: row;
    justify-content: center;
}

</style>