<script setup lang="ts">
// скрипт оставлен по стандарту, здесь можно добавить логику позже
</script>

<template>
  <main class="hello-page">
    <div class="content-box">
      <!-- SVG-иконка сети IoT -->
      <svg viewBox="0 0 64 64" class="hello-icon" aria-hidden="true">
        <circle cx="32" cy="8" r="6" />
        <circle cx="12" cy="48" r="6" />
        <circle cx="52" cy="48" r="6" />
        <line x1="32" y1="14" x2="12" y2="42" />
        <line x1="32" y1="14" x2="52" y2="42" />
        <line x1="12" y1="42" x2="52" y2="42" />
      </svg>

      <h1 class="hello-title">Привет</h1>
      <p class="hello-text">Добро пожаловать в ваше IoT-приложение!</p>
    </div>
  </main>
</template>

<style scoped>
.hello-page {
  /* Центрируем блок и задаём градиент фона через CSS-переменные */
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, var(--color-primary-light) 0%, var(--color-bg-light) 100%);
  animation: fadeIn 0.6s ease;
  padding: 1rem;
}

.content-box {
  /* Карточка с тенью и плавным «выездом» */
  background: var(--color-bg-light);
  padding: 3rem 4rem;
  border-radius: 1.5rem;
  text-align: center;
  box-shadow: var(--shadow-lg);
  animation: slideUp 0.5s ease-out;
  max-width: 400px;
  width: 100%;
  position: relative;
  overflow: hidden;
}

/* SVG-иконка */
.hello-icon {
  width: 4rem;
  height: 4rem;
  margin-bottom: 1rem;
  fill: var(--color-primary);
  transition: var(--transition-default);
}
.hello-icon:hover {
  transform: scale(1.1);
}

/* Заголовок */
.hello-title {
  font-size: 2rem;
  margin-bottom: 0.75rem;
  color: var(--color-primary-dark);
}

/* Подзаголовок */
.hello-text {
  font-size: 1.125rem;
  color: var(--color-text-dark);
}

/* Анимации появления */
@keyframes fadeIn {
  from {
    opacity: 0;
    background-color: var(--color-bg-light);
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

/* Responsive — планшеты и мобильные */
@media (max-width: 768px) {
  .content-box {
    padding: 2.5rem 3rem;
    max-width: 350px;
  }
  .hello-icon {
    width: 3.5rem;
    height: 3.5rem;
  }
  .hello-title {
    font-size: 1.75rem;
  }
  .hello-text {
    font-size: 1rem;
  }
}

@media (max-width: 480px) {
  .content-box {
    padding: 2rem 2rem;
    max-width: 90%;
  }
  .hello-icon {
    width: 3rem;
    height: 3rem;
  }
  .hello-title {
    font-size: 1.5rem;
  }
  .hello-text {
    font-size: 0.95rem;
  }
}
</style>
