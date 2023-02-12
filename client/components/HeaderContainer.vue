<template>
    <div class="header-container">
        <div class="header-icon">
            <p class="fire-emoji">🔥</p>
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
        <v-dialog class="dialog-content"
                  v-model="dialogSettings">
            <v-card>
                <v-card-title class="dialog-text">SETTINGS</v-card-title>
                <v-card-text class="dialog-text">Coming Soon...</v-card-text>
                <v-card-actions>
                    <v-btn color="primary"
                           block
                           @click="dialogSettings = false">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-dialog class="dialog-content"
                  v-model="dialogStatistics">
            <v-card>
                <v-card-title class="dialog-text">STATISTICS</v-card-title>
                <v-card-text class="dialog-text">Coming Soon...</v-card-text>
                <v-card-actions>
                    <v-btn color="primary"
                           block
                           @click="dialogStatistics = false">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-dialog class="dialog-content"
                  v-model="dialogInfo">
            <v-card>
                <v-card-title class="dialog-text">About</v-card-title>
                <v-card-text class="dialog-text">Welcome to the world of POKEMéN! <br> Every day a wild secret POKéMON
                    will be hiding and its your job as a POKéMON TRAINER to catch (guess) it! </v-card-text>
                <v-card-text class="dialog-text">• Reset time <br> • Inspired by<br /> • Contact for removal/credit of
                    background image<br /> • Suggestion<br /> • Bugs<br /> • Feedback<br />
                </v-card-text>
                <v-card-actions>
                    <v-btn color="primary"
                           block
                           @click="dialogInfo = false">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const dialogInfo = ref(false);
const dialogStatistics = ref(false);
const dialogSettings = ref(false);

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

.dialog-text {
    font-family: pkmEmerald;
}

.dialog-content {
    width: 500px;
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

.fire-emoji {
    font-size: x-large;
}
</style>