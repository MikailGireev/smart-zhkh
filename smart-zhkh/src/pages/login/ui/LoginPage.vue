<template>
  <div class="page-wrapper">
    <div class="form-card container">
      <!-- Иконка замка -->
      <svg viewBox="0 0 24 24" class="form-icon" aria-hidden="true">
        <rect x="5" y="11" width="14" height="10" rx="2" ry="2" />
        <path d="M12 11V7a4 4 0 0 0-8 0v4" />
        <path d="M12 11V7a4 4 0 0 1 8 0v4" />
      </svg>

      <h1 class="form-title">Вход в систему</h1>

      <form @submit.prevent="handleLogin" novalidate>
        <div class="form-group">
          <label for="username">Логин</label>
          <input
            v-model="username"
            id="username"
            type="text"
            placeholder="Введите логин"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <input
            v-model="password"
            id="password"
            type="password"
            placeholder="Введите пароль"
            class="form-input"
          />
        </div>

        <button type="submit" class="submit-button btn-primary">🔒 Войти</button>
      </form>

      <ul v-if="errors.length" class="errors">
        <li v-for="(err, i) in errors" :key="i">{{ err }}</li>
      </ul>

      <p v-if="message" class="message">{{ message }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
// пока скрипт не используется, но оставлен по стандарту
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';
import { loginUser } from '@/shared/api';

const username = ref('');
const password = ref('');
const errors = ref<string[]>([]);
const message = ref('');

const auth = useAuthStore();
const router = useRouter();

async function handleLogin() {
  errors.value = [];
  message.value = '';

  if (!username.value.trim()) errors.value.push('Введите логин');
  if (!password.value.trim()) errors.value.push('Введите пароль');
  else if (password.value.length < 6) errors.value.push('Пароль должен быть не менее 6 символов');

  if (errors.value.length) return;

  try {
    const res = await loginUser(username.value, password.value);
    auth.login(res.username, res.token || 'mock-token', res.userId);
    router.push('/dashboard');
  } catch (err: any) {
    message.value = '❌ ' + (err.message || 'Ошибка входа');
  }
}
</script>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: var(--color-bg-light);
  animation: fadeIn 0.6s ease;
  padding: 1rem;
}

.form-card {
  background: var(--color-text-light);
  padding: 2.5rem 3rem;
  border-radius: 1rem;
  width: 100%;
  max-width: 400px;
  box-shadow: var(--shadow-lg);
  animation: slideUp 0.5s ease;
  text-align: center;
  position: relative;
}

.form-icon {
  width: 3rem;
  height: 3rem;
  margin: 0 auto 1rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: var(--transition-default);
}
.form-icon:hover {
  transform: scale(1.1);
}

.form-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--color-primary-dark);
  margin-bottom: 1.5rem;
}

.form-group {
  margin-bottom: 1.25rem;
  text-align: left;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  color: var(--color-text-dark);
  font-weight: 600;
}

.form-input {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 1px solid var(--color-primary-light);
  border-radius: 0.5rem;
  font-size: 1rem;
  transition: var(--transition-default);
}
.form-input:focus {
  border-color: var(--color-primary);
  outline: none;
  background: var(--color-bg-light);
}

.submit-button {
  width: 100%;
  padding: 0.75rem;
  border-radius: 0.5rem;
  font-size: 1rem;
  font-weight: 600;
  transition: var(--transition-default);
}

.errors {
  margin-top: 1rem;
  list-style: none;
  color: var(--color-error);
  font-size: 0.875rem;
  text-align: left;
  padding-left: 1rem;
}

.message {
  margin-top: 1rem;
  color: var(--color-success);
  font-weight: 500;
}

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
    transform: translateY(20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

/* Адаптивность */
@media (max-width: 480px) {
  .form-card {
    padding: 2rem 1.5rem;
  }
  .form-title {
    font-size: 1.5rem;
  }
  .form-input {
    font-size: 0.95rem;
    padding: 0.6rem 0.8rem;
  }
  .submit-button {
    font-size: 0.95rem;
  }
}
</style>
