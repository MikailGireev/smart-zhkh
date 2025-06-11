<template>
  <div class="add-charge-container">
    <h2>➕ Добавить новое начисление</h2>

    <form @submit.prevent="submitForm">
      <div class="form-group">
        <label for="category">📂 Категория</label>
        <input v-model="category" id="category" type="text" required placeholder="Например: Вода" />
      </div>

      <div class="form-group">
        <label for="amount">💰 Сумма (₽)</label>
        <input
          v-model.number="amount"
          id="amount"
          type="number"
          min="0"
          required
          placeholder="Введите сумму"
        />
      </div>

      <div class="form-group">
        <label for="date">📅 Дата</label>
        <input v-model="date" id="date" type="date" required />
      </div>

      <button type="submit" class="submit-button">Добавить</button>
    </form>

    <p v-if="message" class="message">{{ message }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { createCharge } from '@/shared/api/charges';
import { useRouter } from 'vue-router';

const router = useRouter();
const auth = useAuthStore();

const category = ref('');
const amount = ref(0);
const date = ref('');
const message = ref('');

async function submitForm() {
  try {
    await createCharge({
      user_id: auth.userId,
      category: category.value,
      amount: amount.value,
      date: date.value,
    });

    message.value = '✅ Начисление добавлено!';
    setTimeout(() => {
      router.push('/charges');
    }, 1000);
  } catch (e) {
    message.value = '❌ Ошибка при добавлении начисления';
  }
}
</script>

<style scoped>
.add-charge-container {
  max-width: 520px;
  margin: 3rem auto;
  padding: 2.5rem;
  background: linear-gradient(to bottom right, #ffffff, #f0f9ff);
  border-radius: 16px;
  box-shadow: 0 12px 35px rgba(0, 0, 0, 0.06);
  animation: fadeIn 0.3s ease;
}

h2 {
  font-size: 24px;
  color: #0f172a;
  margin-bottom: 1.5rem;
  text-align: center;
}

form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

label {
  font-weight: 600;
  margin-bottom: 0.5rem;
  color: #334155;
}

input {
  padding: 0.6rem 0.75rem;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 16px;
  background: white;
  transition: 0.2s;
}
input:focus {
  border-color: #3b82f6;
  outline: none;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15);
}

.submit-button {
  background-color: #10b981;
  color: white;
  padding: 0.75rem;
  font-size: 16px;
  font-weight: bold;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: 0.2s;
}
.submit-button:hover {
  background-color: #059669;
}

.message {
  margin-top: 1.5rem;
  text-align: center;
  font-weight: 600;
  color: #22c55e;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
