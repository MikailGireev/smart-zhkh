<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';
import { registerUser } from '@/shared/api';

const router = useRouter();
const auth = useAuthStore();

const username = ref('');
const password = ref('');
const confirmPassword = ref('');
const errors = ref<string[]>([]);
const message = ref('');

async function handleRegister() {
  errors.value = [];
  message.value = '';

  if (!username.value.trim()) {
    errors.value.push('Введите логин');
  }
  if (password.value.length < 6) {
    errors.value.push('Пароль должен быть не менее 6 символов');
  }
  if (password.value !== confirmPassword.value) {
    errors.value.push('Пароли не совпадают');
  }
  if (errors.value.length) return;

  try {
    await registerUser(username.value, password.value);
    message.value = '✅ Регистрация прошла успешно';
    auth.login(username.value, 'fake-token', auth.userId);
    router.push('/dashboard');
  } catch (err: any) {
    message.value = err.message || 'Ошибка регистрации';
  }
}
</script>

<template>
  <div class="page-wrapper">
    <div class="form-card container">
      <!-- Иконка регистрации -->
      <svg viewBox="0 0 24 24" class="form-icon" aria-hidden="true">
        <path d="M12 12a5 5 0 1 0-5-5 5 5 0 0 0 5 5zm0 2c-4 0-7 3-7 5v3h14v-3c0-2-3-5-7-5z" />
      </svg>

      <h1 class="form-title">Регистрация</h1>

      <form @submit.prevent="handleRegister" novalidate>
        <div class="form-group">
          <label for="username">Логин</label>
          <input
            id="username"
            type="text"
            v-model="username"
            placeholder="Введите логин"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <input
            id="password"
            type="password"
            v-model="password"
            placeholder="Введите пароль"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="confirm">Повторите пароль</label>
          <input
            id="confirm"
            type="password"
            v-model="confirmPassword"
            placeholder="Повторите пароль"
            class="form-input"
          />
        </div>

        <button type="submit" class="submit-button btn-primary">📝 Зарегистрироваться</button>
      </form>

      <ul v-if="errors.length" class="errors">
        <li v-for="(err, i) in errors" :key="i">{{ err }}</li>
      </ul>

      <p v-if="message" :class="message.startsWith('✅') ? 'message-success' : 'message-error'">
        {{ message }}
      </p>
    </div>
  </div>
</template>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, var(--color-bg-light) 0%, var(--color-primary-light) 100%);
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

.message-success {
  margin-top: 1rem;
  color: var(--color-success);
  font-weight: 500;
}

.message-error {
  margin-top: 1rem;
  color: var(--color-error);
  font-weight: 500;
}

/* Анимации */
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
