<template>
  <div class="add-account-container">
    <h2>Добавить счёт</h2>

    <form @submit.prevent="onSubmit" class="account-form">
      <label>
        Номер счёта:
        <input v-model="form.account_number" type="text" required />
      </label>

      <label>
        ФИО:
        <input v-model="form.full_name" type="text" required />
      </label>

      <label>
        Адрес:
        <input v-model="form.address" type="text" required />
      </label>

      <label>
        Площадь (м²):
        <input v-model.number="form.area" type="number" min="1" required />
      </label>

      <button type="submit">Сохранить</button>
    </form>
  </div>
</template>

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
    alert('Счёт успешно создан');
    router.push('/profile');
  } catch (e) {
    const error = e as Error;
    alert('Ошибка: ' + error.message);
  }
}
</script>

<style scoped>
.add-account-container {
  max-width: 600px;
  margin: 2rem auto;
  padding: 2rem;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 0 15px rgba(0, 0, 0, 0.05);
}

.account-form {
  display: grid;
  gap: 1rem;
}

input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
}

button {
  padding: 0.6rem 1.2rem;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
}

button:hover {
  background-color: #1d4ed8;
}
</style>
