<script setup lang="ts">
import { ref } from 'vue'
import { Layout, Menu, Button } from 'ant-design-vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

interface MenuItem {
  label: string
  key: string
  path: string
}

const menuItems = ref<MenuItem[]>([
  {
    label: '首页',
    key: 'home',
    path: '/',
  },
  {
    label: '关于',
    key: 'about',
    path: '/about',
  },
])

const currentKey = ref(['home'])

const handleMenuClick = (item: MenuItem) => {
  currentKey.value = [item.key]
  router.push(item.path)
}

const handleLogin = () => {
  authStore.openLoginModal()
}
</script>

<template>
  <Layout.Header class="global-header">
    <div class="header-container">
      <div class="logo-section">
        <img alt="Logo" src="@/assets/logo.svg" class="logo" />
        <span class="title">够秀Mall</span>
      </div>

      <Menu
        v-model:selectedKeys="currentKey"
        mode="horizontal"
        class="menu"
        :items="
          menuItems.map((item) => ({
            label: item.label,
            key: item.key,
          }))
        "
        @click="(e: any) => {
          const item = menuItems.find(m => m.key === e.key)
          if (item) handleMenuClick(item)
        }"
      />

      <div class="auth-section">
        <Button type="primary" @click="handleLogin">登录</Button>
      </div>
    </div>
  </Layout.Header>
</template>

<style scoped>
.global-header {
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 0 !important;
}

.header-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 24px;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 12px;
  min-width: 180px;
}

.logo {
  width: 40px;
  height: 40px;
  display: block;
}

.title {
  font-size: 18px;
  font-weight: 600;
  color: #000;
  white-space: nowrap;
}

.menu {
  flex: 1;
  border-bottom: none;
  background: transparent;
  margin: 0 24px;
}

.auth-section {
  display: flex;
  align-items: center;
  gap: 16px;
}

/* 响应式 */
@media (max-width: 768px) {
  .header-container {
    padding: 0 16px;
    flex-wrap: wrap;
    gap: 16px;
  }

  .logo-section {
    min-width: 100%;
  }

  .menu {
    flex: 1;
    margin: 0;
    font-size: 14px;
  }

  .logo {
    width: 32px;
    height: 32px;
  }

  .title {
    font-size: 16px;
  }
}
</style>
