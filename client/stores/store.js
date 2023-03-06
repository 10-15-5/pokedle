import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'

export const useStore = defineStore('store', () => {
    const user = reactive({})
    const isGameWon = ref(false)

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
        setUser,
        setIsGameWon,
    }
});