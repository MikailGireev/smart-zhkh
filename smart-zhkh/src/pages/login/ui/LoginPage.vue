<template>
  <div class="page-wrapper">
    <div class="form-card">
      <h1>Вход в систему</h1>

      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">Логин</label>
          <input v-model="username" id="username" type="text" placeholder="Введите логин" />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <input v-model="password" id="password" type="password" placeholder="Введите пароль" />
        </div>

        <button type="submit" class="submit-button">➡️ Войти</button>
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

  if (username.value.trim() === '') errors.value.push('Введите логин');
  if (password.value.trim() === '') errors.value.push('Введите пароль');
  else if (password.value.length < 6) errors.value.push('Пароль должен быть от 6 символов');

  if (errors.value.length > 0) return;

  try {
    const res = await loginUser(username.value, password.value);
    auth.login(res.username, res.token || 'mock-token', res.userId); // обязательно userId
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
  background: linear-gradient(135deg, #e0f2fe, #f0f9ff);
  animation: fadeIn 0.6s ease;
}

.form-card {
  background: white;
  padding: 2.5rem;
  border-radius: 16px;
  width: 100%;
  max-width: 420px;
  box-shadow: 0 12px 30px rgba(0, 0, 0, 0.1);
  animation: slideUp 0.5s ease;
}

h1 {
  margin-bottom: 2rem;
  text-align: center;
  font-size: 28px;
  color: #1e40af;
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
  border-color: #2563eb;
  outline: none;
  background-color: #f8fafc;
}

.submit-button {
  width: 100%;
  padding: 0.75rem;
  background: #2563eb;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s;
}

.submit-button:hover {
  background: #1e40af;
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
