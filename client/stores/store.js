import { defineStore } from 'pinia'
import { reactive } from 'vue'

export const useStore = defineStore('store', () => {
    const user = reactive({})

    const setUser = (newUser) => {
        Object.assign(user, newUser);
    }

    return {
        user,
        setUser,
    }
});