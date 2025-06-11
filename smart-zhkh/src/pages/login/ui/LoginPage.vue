<template>
  <div class="page-wrapper">
    <div class="form-card">
      <h1>Вход</h1>

      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">Логин</label>
          <input v-model="username" id="username" type="text" required />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <input v-model="password" id="password" type="password" required />
        </div>

        <button type="submit" class="submit-button">Войти</button>
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

  if (username.value.trim() === '') {
    errors.value.push('Введите логин');
  }
  if (password.value.trim() === '') {
    errors.value.push('Введите пароль');
  } else if (password.value.length < 6) {
    errors.value.push('Пароль должен быть от 6 символов');
  }

  if (errors.value.length > 0) return;

  try {
    const res = await loginUser(username.value, password.value);

    // 👇 обязательно передаём userId
    auth.login(res.username, res.token || 'mock-token', res.userId);

    router.push('/dashboard');
  } catch (err: any) {
    message.value = '❌ ' + (err.message || 'Ошибка входа');
  }
}
</script>

<style scoped>
.page-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: #f5f7fa;
}
.form-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  width: 100%;
  max-width: 400px;
}
h1 {
  margin-bottom: 1.5rem;
  text-align: center;
  font-size: 24px;
  color: #333;
}
.form-group {
  margin-bottom: 1rem;
}
label {
  display: block;
  margin-bottom: 0.25rem;
  color: #555;
  font-weight: 500;
}
input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 6px;
  font-size: 16px;
}
input:focus {
  border-color: #3b82f6;
  outline: none;
}
.submit-button {
  width: 100%;
  padding: 0.6rem;
  background-color: #3b82f6;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
}
.submit-button:hover {
  background-color: #2563eb;
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
</style>
