<template>
  <div class="add-charge-container container">
    <h2 class="add-charge-title">
      <svg viewBox="0 0 24 24" class="add-icon" aria-hidden="true">
        <line x1="12" y1="5" x2="12" y2="19" />
        <line x1="5" y1="12" x2="19" y2="12" />
      </svg>
      Добавить новое начисление
    </h2>

    <form @submit.prevent="submitForm" class="add-charge-form">
      <div class="form-group">
        <label for="category">📂 Категория</label>
        <input
          v-model="category"
          id="category"
          type="text"
          required
          placeholder="Например: Вода"
          class="form-input"
        />
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
          class="form-input"
        />
      </div>

      <div class="form-group">
        <label for="date">📅 Дата</label>
        <input v-model="date" id="date" type="date" required class="form-input" />
      </div>

      <button type="submit" class="btn btn-primary submit-button">
        <span>Добавить</span>
      </button>
    </form>

    <p v-if="message" class="message">{{ message }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';
import { createCharge } from '@/shared/api/charges';

const router = useRouter();
const auth = useAuthStore();

const category = ref('');
const amount = ref(0);
const date = ref('');
const message = ref('');

async function submitForm() {
  try {
    await createCharge({
      user_id: Number(auth.userId),
      category: category.value,
      amount: amount.value,
      date: date.value,
    });
    message.value = '✅ Начисление добавлено!';
    setTimeout(() => router.push('/charges'), 800);
  } catch (e) {
    message.value = '❌ Ошибка при добавлении начисления';
  }
}
</script>

<style scoped>
.add-charge-container {
  margin: 3rem auto;
  padding: 2.5rem 1.5rem;
  background: var(--color-bg-light);
  border-radius: 1.5rem;
  box-shadow: var(--shadow-lg);
  animation: fadeIn 0.4s ease;
}

.add-charge-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.5rem;
  color: var(--color-primary-dark);
  margin-bottom: 1.5rem;
}

.add-icon {
  width: 1.75rem;
  height: 1.75rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: var(--transition-default);
}
.add-icon:hover {
  transform: scale(1.1);
}

.add-charge-form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: var(--color-text-dark);
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
  background: var(--color-bg-light);
  outline: none;
}

.submit-button {
  align-self: flex-end;
  padding: 0.75rem 1.5rem;
}

.message {
  margin-top: 1.5rem;
  font-weight: 500;
  color: var(--color-success);
  text-align: center;
}

/* Animations */
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

/* Responsive */
@media (max-width: 768px) {
  .add-charge-container {
    padding: 2rem 1rem;
  }
  .add-charge-title {
    font-size: 1.25rem;
  }
  .submit-button {
    width: 100%;
  }
}

@media (max-width: 480px) {
  .form-input {
    font-size: 0.95rem;
    padding: 0.6rem 0.8rem;
  }
  .add-charge-title {
    font-size: 1.125rem;
  }
}
</style>
