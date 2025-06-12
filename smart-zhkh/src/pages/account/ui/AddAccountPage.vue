<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';
import { createAccount } from '@/shared/api/accounts';

const auth = useAuthStore();
const router = useRouter();

const form = reactive({
  account_number: '',
  full_name: '',
  address: '',
  area: 0,
});

async function onSubmit() {
  try {
    await createAccount({
      user_id: auth.userId,
      account_number: form.account_number,
      full_name: form.full_name,
      address: form.address,
      area: form.area,
    });
    alert('✅ Счёт успешно создан');
    router.push('/profile');
  } catch (e: any) {
    alert('❌ Ошибка: ' + (e.message || 'Не удалось добавить счёт'));
  }
}
</script>

<template>
  <div class="add-account-container container">
    <h2 class="form-title">
      <svg viewBox="0 0 24 24" class="form-icon" aria-hidden="true">
        <path d="M12 12a5 5 0 1 0-5-5 5 5 0 0 0 5 5zm0 2c-4 0-7 3-7 5v2h14v-2c0-2-3-5-7-5z" />
      </svg>
      Добавить счёт
    </h2>

    <form @submit.prevent="onSubmit" class="account-form">
      <div class="form-group">
        <label for="account_number">Номер счёта</label>
        <input
          id="account_number"
          v-model="form.account_number"
          type="text"
          required
          class="form-input"
          placeholder="Введите номер счёта"
        />
      </div>

      <div class="form-group">
        <label for="full_name">ФИО</label>
        <input
          id="full_name"
          v-model="form.full_name"
          type="text"
          required
          class="form-input"
          placeholder="Введите ФИО"
        />
      </div>

      <div class="form-group">
        <label for="address">Адрес</label>
        <input
          id="address"
          v-model="form.address"
          type="text"
          required
          class="form-input"
          placeholder="Введите адрес"
        />
      </div>

      <div class="form-group">
        <label for="area">Площадь (м²)</label>
        <input
          id="area"
          v-model.number="form.area"
          type="number"
          min="1"
          required
          class="form-input"
          placeholder="Введите площадь"
        />
      </div>

      <button type="submit" class="btn btn-primary submit-button">Сохранить</button>
    </form>
  </div>
</template>

<style scoped>
.add-account-container {
  margin: 3rem auto;
  padding: 2.5rem 1.5rem;
  background: var(--color-bg-light);
  border-radius: 1.5rem;
  box-shadow: var(--shadow-lg);
  animation: fadeIn 0.4s ease;
}

.form-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.75rem;
  color: var(--color-primary-dark);
  margin-bottom: 1.5rem;
}

.form-icon {
  width: 2rem;
  height: 2rem;
  fill: var(--color-primary);
  transition: var(--transition-default);
}
.form-icon:hover {
  transform: scale(1.1);
}

.account-form {
  display: grid;
  gap: 1.25rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
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
  justify-self: end;
  padding: 0.75rem 1.5rem;
}

.btn-primary {
  background-color: var(--color-primary);
  color: var(--color-text-light);
  transition: var(--transition-default);
}
.btn-primary:hover {
  background-color: var(--color-primary-dark);
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

/* Адаптивность */
@media (max-width: 768px) {
  .add-account-container {
    padding: 2rem 1rem;
  }
  .form-title {
    font-size: 1.5rem;
  }
  .submit-button {
    justify-self: stretch;
  }
}

@media (max-width: 480px) {
  .form-input {
    font-size: 0.95rem;
    padding: 0.6rem 0.8rem;
  }
  .form-title {
    font-size: 1.25rem;
  }
}
</style>
