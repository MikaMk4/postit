import { defineStore } from "pinia";
import { ref } from "vue";

export const useAppStore = defineStore("appStore", {
    persist: true,
    state: () => ({
        historyCount: 0,
        darkMode: ref(false),
        animationsEnabled: ref(true),
    }),
    actions: {
        setHistoryCount(count) {
            this.historyCount = count;
        }
    },
});