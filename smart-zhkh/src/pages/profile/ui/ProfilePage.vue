<template>
  <div class="profile-container">
    <h2>Профиль</h2>

    <div class="user-info">
      <p><strong>Пользователь:</strong> {{ auth.username }}</p>
    </div>

    <h3>Ваши счета</h3>

    <div v-if="isLoading">Загрузка...</div>

    <div v-else-if="filteredAccounts.length === 0">
      <p>У вас пока нет заполненного профиля.</p>
      <RouterLink to="/accounts/add" class="add-profile-button">📝 Заполнить профиль</RouterLink>
    </div>

    <div v-else class="account-list">
      <div v-for="acc in filteredAccounts" :key="acc.id" class="account-card">
        <p><strong>Номер счёта:</strong> {{ acc.account_num }}</p>
        <p><strong>ФИО:</strong> {{ acc.full_name }}</p>
        <p><strong>Адрес:</strong> {{ acc.address }}</p>
        <p><strong>Площадь:</strong> {{ acc.area }} м²</p>
        <RouterLink :to="`/accounts/${acc.id}/edit`" class="edit-link">✏️ Редактировать</RouterLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { fetchAccounts } from '@/shared/api/accounts';

const auth = useAuthStore();
const accounts = ref<any[]>([]);
const isLoading = ref(true);

onMounted(async () => {
  try {
    const data = await fetchAccounts();
    accounts.value = data;
  } catch (e) {
    console.error('Ошибка загрузки счетов:', e);
  } finally {
    isLoading.value = false;
  }
});

const filteredAccounts = computed(() =>
  accounts.value.filter((acc) => Number(acc.user_id) === Number(auth.userId)),
);
</script>

<style scoped>
.profile-container {
  max-width: 700px;
  margin: 2rem auto;
  padding: 2rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 0 15px rgba(0, 0, 0, 0.05);
}

.user-info {
  margin-bottom: 2rem;
}

.account-list {
  display: grid;
  gap: 1rem;
}

.account-card {
  background: #f8fafc;
  border-radius: 8px;
  padding: 1rem;
  border: 1px solid #e2e8f0;
}

.edit-link {
  display: inline-block;
  margin-top: 0.5rem;
  color: #2563eb;
  font-size: 0.9rem;
}

.edit-link:hover {
  text-decoration: underline;
}

.add-profile-button {
  display: inline-block;
  padding: 0.5rem 1rem;
  margin-top: 1rem;
  background: #10b981;
  color: white;
  border-radius: 6px;
  text-decoration: none;
  font-weight: 500;
}
.add-profile-button:hover {
  background: #059669;
}
</style>
