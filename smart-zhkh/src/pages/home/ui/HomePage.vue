<script setup lang="ts">
// пока скрипт не используется, но оставлен по стандарту
</script>

<template>
  <main class="hello-page">
    <div class="content-box">
      <h1>Привет</h1>
      <p>Добро пожаловать в ваше IoT-приложение!</p>
    </div>
  </main>
</template>

<style scoped>
.hello-page {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #f0f9ff, #e0f2fe);
  animation: fadeIn 0.6s ease;
}

.content-box {
  background: white;
  padding: 3rem 4rem;
  border-radius: 20px;
  text-align: center;
  box-shadow: 0 12px 30px rgba(0, 0, 0, 0.08);
  animation: slideUp 0.5s ease-out;
}

h1 {
  font-size: 36px;
  margin-bottom: 1rem;
  color: #1e40af;
}

p {
  font-size: 18px;
  color: #475569;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    background-color: #f8fafc;
  }
  to {
    opacity: 1;
  }
}

@keyframes slideUp {
  from {
    transform: translateY(16px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
</style>
