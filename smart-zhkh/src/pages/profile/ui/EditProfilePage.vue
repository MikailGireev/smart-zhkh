<template>
  <div class="edit-account-container">
    <h2>Редактировать счёт</h2>

    <form @submit.prevent="submitForm" v-if="account">
      <div class="form-group">
        <label>ФИО</label>
        <input v-model="account.full_name" type="text" required />
      </div>

      <div class="form-group">
        <label>Адрес</label>
        <input v-model="account.address" type="text" required />
      </div>

      <div class="form-group">
        <label>Площадь (м²)</label>
        <input v-model.number="account.area" type="number" min="1" required />
      </div>

      <button type="submit">Сохранить изменения</button>
    </form>

    <p v-else>Загрузка счёта...</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { fetchAccountById, updateAccount } from '@/shared/api/accounts';

const route = useRoute();
const router = useRouter();
const account = ref<any | null>(null);

onMounted(async () => {
  const id = Number(route.params.id);
  try {
    account.value = await fetchAccountById(id);
  } catch (e) {
    console.error('Ошибка загрузки:', e);
  }
});

async function submitForm() {
  try {
    await updateAccount(account.value);
    router.push('/profile');
  } catch (e) {
    alert('Ошибка при сохранении');
  }
}
</script>

<style scoped>
.edit-account-container {
  max-width: 600px;
  margin: 2rem auto;
  padding: 2rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 0 12px rgba(0, 0, 0, 0.05);
}
.form-group {
  margin-bottom: 1rem;
  display: flex;
  flex-direction: column;
}
label {
  font-weight: 500;
  margin-bottom: 0.4rem;
}
input {
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 6px;
}
button {
  margin-top: 1rem;
  padding: 0.7rem 1.2rem;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 6px;
}
button:hover {
  background-color: #1e40af;
}
</style>
