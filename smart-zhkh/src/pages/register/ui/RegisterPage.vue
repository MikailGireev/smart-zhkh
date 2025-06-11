<template>
  <div class="page-wrapper">
    <div class="form-card">
      <h1>Регистрация</h1>

      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label for="username">Логин</label>
          <input v-model="username" id="username" type="text" placeholder="Введите логин" />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <input v-model="password" id="password" type="password" placeholder="Введите пароль" />
        </div>

        <div class="form-group">
          <label for="confirm">Повторите пароль</label>
          <input
            v-model="confirmPassword"
            id="confirm"
            type="password"
            placeholder="Повторите пароль"
          />
        </div>

        <button type="submit" class="submit-button">Зарегистрироваться</button>
      </form>

      <ul v-if="errors.length" class="errors">
        <li v-for="(err, index) in errors" :key="index">{{ err }}</li>
      </ul>

      <p v-if="message" class="message">{{ message }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { registerUser } from '@/shared/api';
import { useAuthStore } from '@/shared/store/auth';

const router = useRouter();
const auth = useAuthStore();

const username = ref('');
const password = ref('');
const confirmPassword = ref('');
const errors = ref<string[]>([]);
const message = ref('');

async function handleRegister() {
  errors.value = [];

  if (username.value.trim() === '') errors.value.push('Введите логин');
  if (password.value.length < 6) errors.value.push('Пароль должен быть от 6 символов');
  if (password.value !== confirmPassword.value) errors.value.push('Пароли не совпадают');

  if (errors.value.length > 0) return;

  try {
    await registerUser(username.value, password.value);
    message.value = 'Регистрация успешна';
    auth.login(username.value, 'fake-token', auth.userId); // Заменить на реальный токен
    router.push('/dashboard');
  } catch (err: any) {
    message.value = err.message || 'Ошибка регистрации';
  }
}
</script>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #f0f9ff, #e0f2fe);
  animation: fadeIn 0.6s ease;
}

.form-card {
  background: white;
  padding: 2.5rem;
  border-radius: 16px;
  width: 100%;
  max-width: 420px;
  box-shadow: 0 12px 30px rgba(0, 0, 0, 0.1);
  animation: slideUp 0.4s ease;
}

h1 {
  margin-bottom: 2rem;
  text-align: center;
  font-size: 28px;
  color: #0f172a;
}

.form-group {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  margin-bottom: 0.4rem;
  color: #334155;
  font-weight: 600;
}

input {
  width: 100%;
  padding: 0.6rem 0.8rem;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 16px;
  transition: border 0.2s;
}

input:focus {
  border-color: #22c55e;
  outline: none;
  background-color: #f0fdf4;
}

.submit-button {
  width: 100%;
  padding: 0.75rem;
  background: #22c55e;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s;
}

.submit-button:hover {
  background: #15803d;
}

.errors {
  color: #dc2626;
  margin-top: 1rem;
  list-style: none;
  padding-left: 0;
  font-size: 14px;
}

.message {
  margin-top: 1rem;
  color: #16a34a;
  text-align: center;
  font-weight: 500;
}

@keyframes fadeIn {
  from {
    opacity: 0;
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
</style>
