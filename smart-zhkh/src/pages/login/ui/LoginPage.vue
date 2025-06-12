<script setup lang="ts">
// скрипт оставлен пустым — здесь можно добавить логику позже
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

<template>
  <div class="page-wrapper container">
    <div class="form-card">
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

        <button type="submit" class="submit-button">🔒 Войти</button>
      </form>

      <ul v-if="errors.length" class="errors">
        <li v-for="(err, i) in errors" :key="i">{{ err }}</li>
      </ul>

      <p v-if="message" class="message">{{ message }}</p>
    </div>
  </div>
</template>

<style scoped>
.page-wrapper {
  position: relative;
  overflow: hidden;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem 1rem;
  background: linear-gradient(145deg, rgba(37, 99, 235, 0.95) 0%, rgba(29, 78, 216, 0.95) 100%);
  animation: fadeInUp 0.8s ease;
}
.page-wrapper::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle at center, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  animation: rotate 20s linear infinite;
  z-index: 0;
}

.form-card {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 400px;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(8px);
  border-radius: 1.5rem;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow:
    0 10px 20px -5px rgba(0, 0, 0, 0.2),
    0 6px 10px -4px rgba(0, 0, 0, 0.1),
    inset 0 1px 2px rgba(255, 255, 255, 0.5);
  text-align: center;
  animation: cardEntrance 0.8s cubic-bezier(0.22, 1, 0.36, 1);
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
  font-weight: 800;
  color: var(--color-primary-dark);
  margin-bottom: 1rem;
  position: relative;
}
.form-title::after {
  content: '';
  position: absolute;
  bottom: -6px;
  left: 50%;
  transform: translateX(-50%);
  width: 40px;
  height: 3px;
  background: var(--color-primary-light);
  border-radius: 3px;
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.6);
}

.form-group {
  margin-bottom: 1.5rem;
  text-align: left;
}

label {
  display: block;
  margin-bottom: 0.4rem;
  color: var(--color-text-dark);
  font-weight: 600;
}

.form-input {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 1px solid var(--color-primary-light);
  border-radius: 0.75rem;
  font-size: 1rem;
  background: var(--color-text-light);
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
  background: var(--color-primary);
  color: var(--color-text-light);
  border: none;
  border-radius: 0.75rem;
  font-size: 1rem;
  font-weight: 600;
  transition: var(--transition-default);
}
.submit-button:hover {
  background: var(--color-primary-dark);
  transform: translateY(-2px);
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
  font-weight: 600;
}

@keyframes cardEntrance {
  from {
    opacity: 0;
    transform: translateY(20px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

/* Responsive */
@media (max-width: 480px) {
  .form-card {
    padding: 2.5rem 1.5rem;
  }
  .form-title {
    font-size: 1.5rem;
  }
  .form-input {
    padding: 0.6rem 0.8rem;
    font-size: 0.95rem;
  }
  .submit-button {
    font-size: 0.95rem;
  }
}
</style>
