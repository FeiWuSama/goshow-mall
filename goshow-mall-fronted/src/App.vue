<script setup lang="ts">
import { RouterView } from 'vue-router'
import { useAuthStore } from './stores/auth'
import Layout from './layout/Layout.vue'
import LoginModal from './components/LoginModal.vue'

const authStore = useAuthStore()

const handleLoginSuccess = (token: string) => {
  console.log('用户已登录，token:', token)
  authStore.closeLoginModal()
}
</script>

<template>
  <Layout>
    <RouterView />
  </Layout>
  <!-- 全局登录弹窗 -->
  <LoginModal
    :visible="authStore.isLoginModalVisible"
    @update:visible="authStore.closeLoginModal"
    @login-success="handleLoginSuccess"
  />
</template>

<style>
#app {
  width: 100%;
  height: 100%;
}
</style>
