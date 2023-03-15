import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'

export const useStore = defineStore('store', () => {
    const user = reactive({});
    const isClassicGameWon = ref(false);
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

    const setIsClassicGameWon = (isWon) => {
        isClassicGameWon.value = isWon;
        localStorage.setItem('isClassicGameWon', `${isWon}`);
    }

    

    return {
        user,
        isClassicGameWon,
        isShiny,
        isDifficultyInsane,
        setUser,
        setIsClassicGameWon,
        toggleShiny,
        setShiny,
        setDifficultyInsane,
        toggleDifficultyInsane,
    }
});