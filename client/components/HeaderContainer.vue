<template>
    <div class="header-container">
        <v-icon class="header-icon"
                icon="mdi-fire-circle"
                size="small"></v-icon>
        <v-icon class="header-icon"
                icon="mdi-cog"
                size="small"
                @click="dialogSettings = true" />
        <v-icon class="header-icon"
                icon="mdi-poll"
                size="small"
                @click="dialogStatistics = true" />
        <v-icon class="header-icon"
                icon="mdi-information-outline"
                size="small"
                @click="dialogInfo = true">
        </v-icon>
        <v-dialog v-model="dialogSettings">
            <v-card>
                <v-card-text> SETTINGS DIALOG </v-card-text>
                <v-card-actions>
                    <v-btn color="primary"
                           block
                           @click="dialogSettings = false">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-dialog v-model="dialogStatistics">
            <v-card>
                <v-card-text> STATISTICS DIALOG </v-card-text>
                <v-card-actions>
                    <v-btn color="primary"
                           block
                           @click="dialogStatistics = false">Close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-dialog v-model="dialogInfo">
            <v-card>
                <v-card-text>INFO DIALOG</v-card-text>
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
    border-width: 1px;
    border-radius: 4px;
}

.header-icon:hover {
    transform: scale(1.1);
    cursor: pointer;
}

.header-container {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}
</style>