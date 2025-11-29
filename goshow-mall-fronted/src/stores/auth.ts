import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const isLoginModalVisible = ref(false)

  const openLoginModal = () => {
    isLoginModalVisible.value = true
  }

  const closeLoginModal = () => {
    isLoginModalVisible.value = false
  }

  return {
    isLoginModalVisible,
    openLoginModal,
    closeLoginModal,
  }
})
