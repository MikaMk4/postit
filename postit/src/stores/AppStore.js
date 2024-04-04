import { defineStore } from "pinia";

export const useAppStore = defineStore("appStore", {
    state: () => ({
        historyCount: 0,
    }),
    actions: {
        setHistoryCount(count) {
            this.historyCount = count;
        }
    },
});