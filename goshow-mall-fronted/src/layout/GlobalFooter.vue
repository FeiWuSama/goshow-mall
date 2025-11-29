<script setup lang="ts"></script>

<template>
  <div class="global-footer">
    <div class="footer-content">
      <span>够秀Mall by
        <a href="https://github.com/FeiWuSama/yikou-ai-feiwu" target="_blank" rel="noopener noreferrer">绯雾Sama</a>
      </span>
    </div>
  </div>
</template>

<style scoped>
.global-footer {
  background: #fafafa;
  border-top: 1px solid #f0f0f0;
  padding: 24px 16px;
  text-align: center;
  width: 100%;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  font-size: 14px;
  color: #666;
}

a {
  color: #1890ff;
  text-decoration: none;
  transition: color 0.3s ease;
}

a:hover {
  color: #40a9ff;
  text-decoration: underline;
}
</style>
