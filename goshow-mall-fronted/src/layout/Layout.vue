<script setup lang="ts">
import { Layout } from 'ant-design-vue'
import GlobalHeader from './GlobalHeader.vue'
import GlobalFooter from './GlobalFooter.vue'
</script>

<template>
  <Layout class="global-layout">
    <GlobalHeader />
    <Layout.Content class="main-content">
      <slot />
    </Layout.Content>
    <GlobalFooter />
  </Layout>
</template>

<style scoped>
.global-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;
  padding: 24px 16px;
}

/* 响应式 */
@media (max-width: 768px) {
  .main-content {
    padding: 16px 12px;
  }
}
</style>
