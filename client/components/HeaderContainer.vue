<template>
    <div class="flex items-center justify-center">
        <HeaderIcon>
            <p class="text-[30px]">🔥</p>
            <p class="streak-count mt-2 font-pkmEmerald text-light-text dark:text-gray-600">
                {{ store.user.currentStreak }}
            </p>
        </HeaderIcon>
        <HeaderIcon @click="dialogSettings = true">
            <v-icon class="text-light-text dark:text-dark-text" icon="mdi-cog" size="small" />
        </HeaderIcon>
        <HeaderIcon @click="dialogStatistics = true">
            <v-icon class="text-light-text dark:text-dark-text" icon="mdi-poll" size="small" />
        </HeaderIcon>
        <HeaderIcon @click="dialogInfo = true">
            <v-icon
                class="text-light-text dark:text-dark-text"
                icon="mdi-information-outline"
                size="small"
            >
            </v-icon>
        </HeaderIcon>
        <!-- TODO: this should be a view -->
        <v-dialog class="w-[720px]" v-model="dialogSettings">
            <div class="dialog-card">
                <h1 class="font-pkmEmerald text-[32px] uppercase">SETTINGS</h1>
                <p class="font-pkmEmerald text-[20px]">Coming Soon....</p>
                <button class="dialog-close" block @click="dialogSettings = false">Close</button>
            </div>
        </v-dialog>
        <!-- TODO: this should be a view -->
        <v-dialog class="w-[720px]" v-model="dialogStatistics">
            <div class="dialog-card">
                <h1 class="font-pkmEmerald text-[32px] uppercase">STATISTICS</h1>
                <p class="font-pkmEmerald text-[20px]">Coming Soon....</p>
                <button class="dialog-close" block @click="dialogStatistics = false">Close</button>
            </div>
        </v-dialog>
        <!-- TODO: this should be a view -->
        <v-dialog class="w-[800px]" v-model="dialogInfo">
            <div class="dialog-card">
                <h1 class="font-pkmEmerald text-[32px] uppercase">About</h1>
                <p class="font-pkmEmerald text-[20px]">
                    Welcome to the world of POKEMéN! <br />
                    Every day a wild secret POKéMON will be hiding and its your job as a POKéMON
                    TRAINER to catch (guess) it!
                </p>
                <hr class="divider" />
                <h1 class="font-pkmEmerald text-[32px] uppercase">How To Play</h1>
                <div class="font-pkmEmerald text-[20px]">
                    <p>
                        <b>1. Guess:</b> Guess any random POKEMéN <b>from gen 1-3</b>, since we
                        currently have no clues. We will guess <b>SQUIRTLE</b> here.
                    </p>
                    <p>
                        <b>2. Guess:</b> We now know that we are looking for a <b>BLUE</b> POKEMéN.
                        Furthermore the POKEMéN has to be in its <b>final evolution stage</b>.
                        Lastly, the POKEMéN is <b>not</b> from gen 1. <br />
                        Let's guess <b> HERACROSS </b> since it matches these characteristics.
                    </p>
                    <p>
                        <b>3. Guess:</b> Now we know the hidden POKEMéN <b>must</b> be from gen 3.
                        Since the hidden POKEMéN cannot be a <b>WATER TYPE</b>, this leaves only a
                        couple of possible options. <br />
                        Let's guess
                        <b>SALAMANCE.</b>
                    </p>
                    <p>
                        <b>4. Guess:</b> We now know that the POKEMéN lives in
                        <b>ROUGH TERRAIN</b> and is in <b>evolution stage 3</b> <br />
                        This leaves only one POKEMéN; <b>METAGROSS.</b>
                    </p>
                    <v-img
                        class="card mx-auto w-[600px]"
                        :src="'./client/assets/game-example.png'"
                    ></v-img>
                </div>
                <hr class="divider" />
                <h1 class="font-pkmEmerald text-[32px] uppercase">Credit</h1>
                <p class="font-pkmEmerald text-[20px]">
                    Pokédlé is heavily inspired by loldle.net and wordle. Original artist/creator of
                    images, please DM for removal or credit.
                </p>
                <hr class="divider" />
                <h1 class="font-pkmEmerald text-[32px] uppercase">Contact</h1>
                <p class="font-pkmEmerald text-[20px]">
                    If you have any questions/feedback/bug reports feel free to contact me on the
                    email below. <br />
                    <a
                        class="text-[24px] font-bold text-blue-500 hover:text-blue-400 dark:text-blue-300 dark:hover:text-blue-400"
                        href="mailto: pokedle.contact@gmail.com"
                    >
                        pokedle.contact@gmail.com
                    </a>
                </p>
                <button class="dialog-close" block @click="dialogInfo = false">Close</button>
            </div>
        </v-dialog>
    </div>
</template>

<script setup>
import { useStore } from '../stores/store';
import { ref } from 'vue';
import HeaderIcon from './HeaderIcon.vue';
const store = useStore();
let dialogInfo = ref(false);
let dialogStatistics = ref(false);
let dialogSettings = ref(false);
</script>

<style scoped>
p {
    padding: 4px;
}

.streak-count {
    position: absolute;
    font-weight: bold;
    margin-left: 1px;
    text-shadow: -0.5px 0.5px 1px #fffcfc, 0.5px 0.5px 1px #ffffff, 0.5px -0.5px 0 #ffffff,
        -0.5px -0.5px 0 #ffffff;
}
</style>
