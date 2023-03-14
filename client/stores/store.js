import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'

export const useStore = defineStore('store', () => {
    const user = reactive({});
    const isGameWon = ref(false);
    const isShiny = ref(false);
    const isDifficultyInsane = ref(false);

    const setShiny = (isShinyParam) => {
        isShiny.value = isShinyParam;
        localStorage.setItem('isShiny', `${isShiny.value}`)
    }
    
    const toggleShiny = () => {
        setShiny(!isShiny.value);
    }
    
    const setDifficultyInsane = (isDifficultyInsaneParam) => {
        isDifficultyInsane.value = isDifficultyInsaneParam;
        localStorage.setItem('isDifficultyInsane', `${isDifficultyInsane.value}`)
    }


    const toggleDifficultyInsane = () => {
        setDifficultyInsane(!isDifficultyInsane.value);
    }

    const setUser = (newUser) => {
        Object.assign(user, newUser);
    }

    const setIsGameWon = (isWon) => {
        isGameWon.value = isWon;
        localStorage.setItem('isGameWon', `${isWon}`);
    }

    

    return {
        user,
        isGameWon,
        isShiny,
        isDifficultyInsane,
        setUser,
        setIsGameWon,
        toggleShiny,
        setShiny,
        setDifficultyInsane,
        toggleDifficultyInsane,
    }
});